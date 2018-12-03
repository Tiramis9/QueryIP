package action

import (
	"fmt"
	"game2/global/status"
	"game2/logic"
	"game2/model"
	"game2/service"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type GameListReq struct {
	Type int `json:"type"`
}

type GameLoginReq struct {
	GameCode string `json:"game_code" binding:"required"`
	GameType string `json:"game_type"`
	AppId    string `json:"app_id"`
}

type GameSubListReq struct {
	Id int `json:"id"`
}

//游戏列表
func GameTypeList(c *gin.Context) {
	var data []interface{}
	ch := make(chan int)
	defer close(ch)
	//获取类型
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		logrus.Error("merchantId获取错误")
		RespServerErr(c)
		return
	}
	merchId := int(merchantId.(float64))
	gameTypeList, err := model.GetGameTypeList(model.Db, merchId)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	fmt.Println(gameTypeList)
	//获取列表
	for _, v := range gameTypeList {
		go func(ty int) {
			res, err := model.GetGameList(model.Db, ty)
			if err != nil {
				RespServerErr(c)
				return
			}
			arr := map[string]interface{}{}
			arr["list"] = res
			arr["type"] = ty
			data = append(data, arr)
			ch <- 1
		}(v.Type)
	}
	//等待通道数结束
	for i := 0; i < len(gameTypeList); i++ {
		<-ch
	}
	RespJson(c, status.OK, data)
}

func GameList(c *gin.Context) {
	var m GameListReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	info, err := model.GetGameList(model.Db, m.Type)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, info)
}

func GameSubList(c *gin.Context) {
	//要查看父游戏是否开通
	//TODO

	var data = make(map[string]interface{})
	//根据父id查询子游戏列表
	var m GameSubListReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	list, err := model.GetGameSubList(model.Db, m.Id)
	if err != nil {
		RespServerErr(c)
		return
	}
	data["list"] = list
	RespJson(c, status.OK, data)
}

func GameLogin(c *gin.Context) {
	var data = make(map[string]interface{})
	var m GameLoginReq

	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}

	//根据游戏代码选择不同包
	gameCode := m.GameCode
	tokenS, ok := c.Request.Header["Token"]
	if !ok {
		RespServerErr(c)
		return
	}
	token := tokenS[0]
	//用户信息
	userInfo := make(map[string]interface{})
	mapInfo := logic.UserInfoByRedis(token)

	userInfo["user_id"] = int(mapInfo["id"].(float64))
	userInfo["user_name"] = mapInfo["user_name"].(string)
	userInfo["lang"] = mapInfo["lang"].(string)
	userInfo["time_zone"] = strconv.Itoa(int(mapInfo["time_zone"].(float64)))
	userInfo["merchant_id"] = int(mapInfo["merchant_id"].(float64))
	userInfo["login_ip"] = mapInfo["login_ip"].(string)
	//增添GameType表示电子游戏的具体类型
	loginUrl, err := service.GameLogin(gameCode, userInfo, m.GameType, m.AppId)

	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	data["url"] = loginUrl
	RespJson(c, status.OK, data)
}
