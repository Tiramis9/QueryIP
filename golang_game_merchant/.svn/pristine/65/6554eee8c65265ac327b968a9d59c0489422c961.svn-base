package global

import (
	"github.com/spf13/viper"
)

type Config struct {
	// http服务器配置
	ServerHttpPort int

	// 数据库连接配置
	DbDriver   string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string

	// redis配置
	RedisType     string
	RedisHost     string
	RedisPort     int
	RedisDatabase string
	RedisPassword string
}

const (
	configFile = "app"
	configPath = "./config"
)

var (
	AppConfig *Config
)

func ConfigInit() {
	v := viper.New()
	v.SetConfigName(configFile)
	v.AddConfigPath(configPath)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	AppConfig = &Config{
		ServerHttpPort: v.GetInt("server.http.port"),
		DbDriver:       v.GetString("db.driver"),
		DbHost:         v.GetString("db.host"),
		DbPort:         v.GetInt("db.port"),
		DbUser:         v.GetString("db.user"),
		DbPassword:     v.GetString("db.password"),
		DbName:         v.GetString("db.name"),
		RedisType:      v.GetString("redis.type"),
		RedisHost:      v.GetString("redis.host"),
		RedisPort:      v.GetInt("redis.port"),
		RedisDatabase:  v.GetString("redis.database"),
		RedisPassword:  v.GetString("redis.password"),
	}
}
