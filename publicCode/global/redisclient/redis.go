package redisclient

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"golang_game_merchant/global"
	"encoding/json"
	"github.com/sirupsen/logrus"
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

func HsetToRedis(key string, field string, value string) error {
	conn := Get()

	//记得销毁本次链连接
	defer conn.Close()

	_, err := conn.Do("HSET", key, field, value)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

//Hget获取redis中的信息
func HgetFromRedis(key string, field string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	conn := Get()

	//记得销毁本次链连接
	defer conn.Close()

	info, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		logrus.Error(err)
		return m, err
	}
	err = json.Unmarshal([]byte(info), &m)
	if err != nil {
		logrus.Error(err)
		return m, err
	}
	return m, nil
}

func RedisSet(key string, value string, ex_time int) error {
	conn := Get()

	//记得销毁本次链连接
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "EX", ex_time)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func RedisGetString(key string) (string, error) {
	conn := Get()

	//记得销毁本次链连接
	defer conn.Close()

	info, err := redis.String(conn.Do("GET", key))
	if err != nil {
		logrus.Error(err)
		return info, err
	}
	return info, nil
}

func RedisGetMap(key string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	conn := Get()

	//记得销毁本次链连接
	defer conn.Close()

	info, err := redis.String(conn.Do("GET", key))
	if err != nil {
		logrus.Error("RedisGetMap:"+key,err)
		return m, err
	}
	err = json.Unmarshal([]byte(info), &m)
	if err != nil {
		logrus.Error("RedisGetMap"+key,err)
		return m, err
	}
	return m, nil
}

func UserInfoByRedis(token string) map[string]interface{} {
	//用户信息
	mapInfo, err := RedisGetMap(token)
	if err != nil {
		return nil
	}
	return mapInfo
}


func init() {
	redisPool = redis.NewPool(NewRedisConn, poolSize)
}
