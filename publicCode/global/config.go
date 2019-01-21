package global

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
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

	//oplog 配置
	LogConf lumberjack.Logger

	//分页默认配置
	Page      int
	PageCount int
}

const (
	devFile    = "app"
	liveFile   = "live"
	configPath = "./config"
)

var (
	AppConfig *Config
)

func ConfigInit() {
	v := viper.New()
	configFile := devFile
	if !LOCAL {
		configFile = liveFile
	}
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
		LogConf: lumberjack.Logger{
			Filename:   v.GetString("log.filename"),
			MaxSize:    v.GetInt("log.maxsize"),
			MaxAge:     v.GetInt("log.maxage"),
			MaxBackups: v.GetInt("log.maxbackups"),
			LocalTime:  v.GetBool("log.localtime"),
			Compress:   v.GetBool("log.compress"),
		},
		Page:      v.GetInt("page.page"),
		PageCount: v.GetInt("page.pagecount"),
	}
}
