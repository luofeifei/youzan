module client

go 1.15

require (
	base v0.0.0
	base/client v0.0.0
	base/model v0.0.0
	github.com/TarsCloud/TarsGo v1.1.5
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/gin-gonic/gin v1.6.3
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
)

replace base => ../../../base/base

replace base/client => ../../../base/client

replace base/model => ../../../model
