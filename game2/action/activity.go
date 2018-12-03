package action

import (
	"fmt"
	"game2/model"
	"github.com/gin-gonic/gin"
	"game2/global/status"
	"github.com/sirupsen/logrus"
	"time"
)

type ActivityInfoReq struct {
	Id int `json:"id" binding:"required"`
}

type ActivityListResp struct {
	Id          int    `json:"id"`
	ActTitle    string `json:"act_title"`
	ActType     int    `json:"act_type"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	ResourceWeb string `json:"resource_web"`
	ResourceWap string `json:"resource_wap"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

func ActivityList(c *gin.Context) {
	var listResp []ActivityListResp
	data := make(map[string]interface{})
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
		return
	}
	merchId := int(merchantId.(float64))
	nowTime := time.Now().Unix()
	list, err := model.GetMerchantActiveList(model.Db, merchId, nowTime)
	if err != nil {
		RespServerErr(c)
		logrus.Error(err)
		return
	}
	for i :=range list{
		var m = ActivityListResp{
			Id:list[i].Id,
			ActTitle:list[i].ActTitle,
			ActType:list[i].ActType,
			StartTime:list[i].StartTime,
			EndTime:list[i].EndTime,
			ResourceWap:list[i].ResourceWap,
			ResourceWeb:list[i].ResourceWeb,
			Description:list[i].Description,
			CreateTime:list[i].CreateTime,
		}
		listResp = append(listResp, m)
	}
	data["list"] = listResp
	RespJson(c, status.OK, data)
}

func ActivityInfo(c *gin.Context) {
	var m ActivityInfoReq
	if err := c.BindJSON(&m); err != nil {
		fmt.Println(err)
		RespParamErr(c)
		return
	}
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
	}
	merchId := int(merchantId.(float64))
	info,err:= model.GetMerchantActiveInfo(model.Db,m.Id,merchId)
	if err!=nil{
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, info)
}
