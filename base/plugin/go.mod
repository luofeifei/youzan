module base/plugin

go 1.15

require (
	base v0.0.0
    github.com/jinzhu/gorm v1.9.16 // indirect
    go.mongodb.org/mongo-driver v1.4.1 // indirect
    github.com/go-redis/redis v6.15.9+incompatible // indirect
)

replace base => ../base
