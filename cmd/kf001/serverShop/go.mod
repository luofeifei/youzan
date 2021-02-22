module server

go 1.15

require (
	base v0.0.0
	base/model v0.0.0
	base/server v0.0.0
	github.com/TarsCloud/TarsGo v1.1.5
	go.mongodb.org/mongo-driver v1.4.1
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0 // indirect
)

replace base => ../../../base/base

replace base/server => ../../../base/server

replace base/model => ../../../model
