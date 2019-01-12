package main

import (
	"game2/global"
	"game2/global/log"
	"game2/model"
	"game2/router"
)

func main() {
	global.ConfigInit()
	log.InitLog()

	model.DbInit()

	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
