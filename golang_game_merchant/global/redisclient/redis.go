package redisclient

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"golang_game_merchant/global"
)

const poolSize = 20

var redisPool *redis.Pool

func NewRedisConn() (redis.Conn, error) {
	conf := global.AppConfig
	c, err := redis.Dial(conf.RedisType, fmt.Sprintf("%v:%v", conf.RedisHost, conf.RedisPort))
	if err != nil {
		return nil, err
	}
	_, err2 := c.Do("AUTH", conf.RedisPassword)
	if err2 != nil {
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
