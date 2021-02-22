package subscriber

import (
	"base/server/pkg/queue"
	json "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
)

func RegisterSubscriber() {
	//可监听多个主题
	err := queue.Subscribe("mall.server.co", "channel", new(NsqHandler))
	if err != nil {
		return
	}
}

type NsqHandler struct{}

func (s *NsqHandler) HandleMessage(message *nsq.Message) error {
	var msg queue.Message
	err := json.Unmarshal(message.Body, &msg)
	if err != nil {
		return err
	}
	switch msg.Name {
	case "coInit":
		// 企业创建初始化
		_ = CoInit(msg.Body)
	case "staff":
		// 企业角色
		_ = UserStaffData(msg.Body)
	}
	return nil
}
