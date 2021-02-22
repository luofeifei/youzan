package common

import (
	"base/plugin/pkg/database"
	"github.com/jinzhu/gorm"
)

type IEngine interface {
	Orm(string) *gorm.DB
	Mongo(string) *database.CollectionInfo
	Redis(string) *database.RedisClient
}

type ormDb struct {
	db   *gorm.DB
	hash string
}

type mongoDb struct {
	db   *database.CollectionInfo
	hash string
}

type redisDb struct {
	db   *database.RedisClient
	hash string
}

var (
	ormClient   map[string]ormDb
	mongoClient map[string]mongoDb
	redisClient map[string]redisDb
)

func init() {
	ormClient = make(map[string]ormDb)
	mongoClient = make(map[string]mongoDb)
	redisClient = make(map[string]redisDb)
}

// 通过 HASH 检测MySQL是否已连接
func HashOrm(dbname string, hash string) bool {
	return ormClient[dbname].hash != hash
}

// 通过 HASH 检测Mongo是否已连接
func HashMongo(dbname string, hash string) bool {
	return mongoClient[dbname].hash != hash
}

// 通过 HASH 检测Redis是否已连接
func HashRedis(dbname string, hash string) bool {
	return redisClient[dbname].hash != hash
}

func SetOrm(dbname string, hash string, db *gorm.DB) {
	ormClient[dbname] = ormDb{
		db:   db,
		hash: hash,
	}
}

func SetMongo(dbname string, hash string, db *database.CollectionInfo) {
	mongoClient[dbname] = mongoDb{
		db:   db,
		hash: hash,
	}
}

func SetRedis(dbname string, hash string, db *database.RedisClient) {
	redisClient[dbname] = redisDb{
		db:   db,
		hash: hash,
	}
}

type Engine struct {
}

func (e *Engine) Orm(dbname string) (res *gorm.DB) {
	if ormClient[dbname].db != nil {
		return ormClient[dbname].db
	}
	return res
}

func (e *Engine) Mongo(dbname string) (res *database.CollectionInfo) {
	if mongoClient[dbname].db != nil {
		return mongoClient[dbname].db
	}
	return res
}

func (e *Engine) Redis(dbname string) (res *database.RedisClient) {
	if redisClient[dbname].db != nil {
		return redisClient[dbname].db
	}
	return res
}
