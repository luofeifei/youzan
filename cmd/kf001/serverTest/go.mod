module server

go 1.15

require (
	base v0.0.0
	base/model v0.0.0
	base/server v0.0.0
	github.com/TarsCloud/TarsGo v1.1.5
	github.com/nsqio/go-nsq v1.0.8
	go.mongodb.org/mongo-driver v1.4.1
)

replace base => ../../../base/base

replace base/server => ../../../base/server

replace base/model => ../../../model
