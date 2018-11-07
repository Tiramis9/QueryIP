package main

import (
	"log"
	"net/http"

	geoip2 "github.com/oschwald/geoip2-golang"
	"public.info/cmd/task"
	"public.info/config"
)

//初始化IP库
func init() {
	config.ConfigLocalFilesystemLogger(config.Path, config.File, config.MaxAge, config.Interval)
	db, err := geoip2.Open(config.MmdbFile)
	if err != nil {
		log.Fatal("Open IP Database File failed..")
	}
	config.DbHander = db
}

///主动连接服务器
func main() {
	log := config.Log.WithField("package", "main")
	log.Info("welcome to main")
	go task.Control()
	http.HandleFunc("/", task.HandleQueryServer)
	log.Fatal(http.ListenAndServe(config.ADDRPORT, nil))
}
