package queue

import (
	"base/pkg/app"
	"base/pkg/config"
	"base/pkg/log"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type configs struct {
	Address     string `yaml:"address"`
	Lookup      string `yaml:"lookup"`
	MaxInFlight int    `yaml:"max_in_flight"`
}

type Message struct {
	Name string
	Body interface{}
}

var (
	producer       *nsq.Producer
	addrNsqLookups []string
	logLevel       nsq.LogLevel
	consumers      []*nsq.Consumer
	nsqConfig      *nsq.Config
	conf           configs
)

func Start() bool {
	err := config.Config().Bind(app.CfgName, "nsq", &conf, func() {
		if config.Config().Bind(app.CfgName, "nsq", &conf, nil) == nil {
			initStart()
		}
	})
	if err != nil {
		app.Logger().Debug(fmt.Sprintf("nsq 配置错误 err: %s", err.Error()))
		return false
	}
	return initStart()
}

func initStart() bool {
	if conf.Address == "" {
		conf.Address = "127.0.0.1:4150"
	}
	addrNsqLookups = strings.Split(conf.Lookup, ",")
	nsqConfig = nsq.NewConfig()
	nsqConfig.MaxInFlight = conf.MaxInFlight
	p, err := nsq.NewProducer(conf.Address, nsqConfig)
	if err != nil {
		app.Logger().Warn("nsq connect error: ")
		app.Logger().Warn(err)
		return false
	}
	logLevel = nsq.LogLevelWarning
	p.SetLogger(log.NsqLogger(), logLevel)
	producer = p
	if err = p.Ping(); err != nil {
		app.Logger().Warn(err)
		return false
	}
	app.Logger().Info("nsq[" + conf.Address + "] connect success")
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		gracefulStop()
	}()
	return true
}

func Publish(topic string, name string, body interface{}, delay ...time.Duration) (err error) {
	msg, err := json.Marshal(&Message{Name: name, Body: body})
	if err != nil {
		return err
	}
	if len(delay) == 0 {
		err = producer.Publish(topic, msg)
	} else {
		err = producer.DeferredPublish(topic, delay[0], msg)
	}
	return
}

func Subscribe(topic string, channel string, handler nsq.Handler) (err error) {
	c, errs := nsq.NewConsumer(topic, channel, nsqConfig)
	if errs != nil {
		app.Logger().Warn(err)
		return errs
	}
	c.AddHandler(handler)
	c.SetLogger(log.NsqLogger(), logLevel)
	err = c.ConnectToNSQLookupds(addrNsqLookups)
	if err != nil {
		app.Logger().Warn(err)
		return
	}
	consumers = append(consumers, c)
	return
}

func gracefulStop() {
	producer.Stop()
	var wg sync.WaitGroup
	for _, c := range consumers {
		wg.Add(1)
		go func(c *nsq.Consumer) {
			c.Stop()
			// disconnect from all lookupd
			for _, addr := range addrNsqLookups {
				err := c.DisconnectFromNSQLookupd(addr)
				if err != nil {
					app.Logger().Warn(err)
				}
			}
			wg.Done()
		}(c)
	}
	wg.Wait()
}
