package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"public.info/config"
	"public.info/server"
)

//Task distribution handle, only support GET request
func HandleIPServer(connect http.ResponseWriter, response *http.Request) {
	log := config.Log.WithField("package", "tsak")
	log.Info("start running HandleIPServer ")

	robots, err := ioutil.ReadAll(response.Body)
	log.Info(response.Host + " Connection successful")
	response.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	server.DealQuery(connect, string(robots))
	log.Info("HandleIPServer running over")
}

//监听指定信号 ctrl+c kill
func Control() {
	log := config.Log.WithField("package", "tsak")
	channel := make(chan os.Signal)
	signal.Notify(channel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for sig := range channel {
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("退出", sig)
				ExitFunc()
			default:
				log.Println("other", sig)
			}
		}
	}()
}

//退出程序，并关闭文件描述符
func ExitFunc() {
	fmt.Println("Start out...")
	fmt.Println("IP database file being close down...")
	config.DbHander.Close()

	time.Sleep(time.Second * 1)
	fmt.Println("shutdown...")
	os.Exit(0)
}

//Task distribution handle, support GET and POST request
func HandleQueryServer(connect http.ResponseWriter, response *http.Request) {
	log := config.Log.WithField("package", "tsak")
	log.Info("start running HandleGETServer ")
	log.Info(response.Host + " Connection successful")

	getstr := response.URL.Query()
	getQuery, err := server.MapTOString(getstr)
	if err != nil {
		fmt.Println(getQuery, err)
		return
	}
	postQuery, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Info("start running switch")
	switch {
	case getQuery != "":
		log.Info("start running getQuery")
		server.DealQuery(connect, getQuery)
		break
	case postQuery != nil:
		log.Info("start running postQuery")
		server.DealQuery(connect, string(postQuery))
		break
	default:
		log.Info("start running other")
		log.Printf("%v,%v", postQuery, getstr)
		fmt.Fprintf(connect, "%v,%v", string(postQuery), getstr)
	}
	log.Info("HandleIPServer running over")
}
