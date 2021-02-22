package log

import (
	"base/pkg/app"
)

type OrmLoggerI struct {
}

func (o *OrmLoggerI) Print(v ...interface{}) {
	v = v[2:]  //只取sql语句
	app.Logger().Infof(" %v  ", v)
}

func NewGormLogger() *OrmLoggerI {
	return &OrmLoggerI{}
}
