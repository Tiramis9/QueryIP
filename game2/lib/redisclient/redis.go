package redisclient

import (
	"fmt"
	"game2/global"
	"game2/lib/utils"
	"github.com/gomodule/redigo/redis"
)

const poolSize = 20

var redisPool *redis.Pool

func NewRedisConn() (redis.Conn, error) {
	conf := global.AppConfig
	c, err := redis.Dial(conf.RedisType, fmt.Sprintf("%v:%v", conf.RedisHost, conf.RedisPort))
	if err != nil {
		utils.Error(err)
		return nil, err
	}
	_, err2 := c.Do("AUTH", conf.RedisPassword)
	if err2 != nil {
		utils.Error(err2)
		return nil, err2
	}
	return c, nil
}

func Get() redis.Conn {
	return redisPool.Get()
}

func init() {
	redisPool = redis.NewPool(NewRedisConn, poolSize)
}
