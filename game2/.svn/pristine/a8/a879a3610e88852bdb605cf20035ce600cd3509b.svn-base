package action

import (
	"encoding/json"
	"game2/global/status"
	"game2/lib/redisclient"
	"game2/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

type WebsiteInfoReq struct {
	Domain string `json:"domain" binding:"required"`
}

//跟据站点的域名查看appid
func WebsiteInfo(c *gin.Context) {
	var w WebsiteInfoReq
	if err := c.BindJSON(&w); err != nil {
		RespParamErr(c)
		return
	}
	//php接口方地址
	ip := c.ClientIP()
	//获取列表
	data, err := model.GetDomainInitInfo(model.Db, w.Domain)
	if err != nil {
		RespServerErr(c)
		return
	}

	if data == nil || data.Appid == "" {

		RespParamErr(c)
		return
	}

	//判断ip是否在白名单中
	flag := strings.Contains(data.IpWhite, ip)
	if flag == false {
		logrus.Warn("ip banned")
		RespUnauthorized(c)
		return
	}

	//将网站信息与appid,secret存入redis
	input, _ := json.Marshal(data)
	conn := redisclient.Get()
	defer conn.Close()
	exTime := 1200000
	_, err = conn.Do("SET", "info_"+w.Domain, input, "EX", exTime)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	appIdInfo := make(map[string]interface{})
	appIdInfo["merchant_id"] = data.Id
	appIdInfo["secret"] = data.Secret
	jsonInfo, _ := json.Marshal(appIdInfo)
	_, err = conn.Do("SET", "appid_"+data.Appid, jsonInfo, "EX", exTime)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, data)
}
