package action

import (
	"game2/global/status"
	"game2/model"
	"github.com/gin-gonic/gin"
)

type RechargeReq struct {
	Id int `json:"id"`
}

func PayTypetList(c *gin.Context) {
	var data []interface{}
	ch := make(chan int)
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
		return
	}
	merchId := int(merchantId.(float64))
	payTypeList, err := model.GetPayTypeList(model.Db, merchId)
	if err != nil {
		RespServerErr(c)
		return
	}
	for _, v := range payTypeList {
		go func(ty string) {
			res, err := model.GetPayConfig(model.Db, merchId, ty)
			if err != nil {
				RespServerErr(c)
				return
			}
			arr := map[string]interface{}{}
			arr["list"] = res
			arr["name"] = ty
			data = append(data, arr)
			ch <- 1
		}(v.PayType)
	}
	//等待通道数结束
	for i := 0; i < len(payTypeList); i++ {
		<-ch
	}
	RespJson(c, status.OK, data)
}

func Recharge(c *gin.Context) {
	var m RechargeReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
		return
	}
	merchId := int(merchantId.(float64))
	info, err:= model.GetPayInfo(model.Db, m.Id, merchId)
	if err!=nil{
		RespServerErr(c)
		return
	}
	data := make(map[string]interface{})
	data["pay_tag"] = info.PayTag
	data["redirect"] = "www.baidu.com"
	RespJson(c,status.OK,data)
}
