package action

import (
	"fmt"
	"game2/model"
	"github.com/gin-gonic/gin"
	"game2/global/status"
)

type AdvertisementReq struct {
	Type     int `json:"type"`
	Location int `json:"location"`
}

func AdvertisementList(c *gin.Context) {
	var m AdvertisementReq
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
	where := make(map[string]interface{})
	if m.Location == 1 || m.Location == 2 {
		where["location"] = m.Location
	}
	if m.Type == 1 || m.Type == 2 || m.Type == 3 {
		where["type"] = m.Type
	}
	data, err := model.GetAdvertisementList(model.Db, merchId, where)
	if err != nil {
		//处理
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, data)
}
