package log

import (
	"base/pkg/app"
	"strings"
)

type NsqLoggerI struct {
}

func (l *NsqLoggerI) Output(calldepth int, s string) error {
	level := strings.Split(s, " ")[0]
	switch level {
	case "INF":
		app.Logger().Infof(s)
	case "WRN":
		app.Logger().Warn(s)
	case "ERR":
		app.Logger().Error(s)
	default:
		app.Logger().Debug(s)
	}
	return nil
}

func NsqLogger() *NsqLoggerI {
	return &NsqLoggerI{}
}
