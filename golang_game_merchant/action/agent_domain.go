package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
	"time"
)

//添加代理域名req
type AgentDomainAddReq struct {
	Domain  string `json:"domain" binding:"required"`
	AgentId int    `json:"agent_id" binding:"required"`
}

type AgentDomainReq struct {
	Id int `json:"id" binding:"required"`
}

type AgentDomainListReq struct {
	Id int `json:"id" binding:"required"`
}

type AgentDomainListResp struct {
	List []model.AgentDomain `json:"list"`
}

/**
代理域名列表
*/
func AgentDomainList(c *gin.Context) {
	//获取商户id
	var data AgentDomainListResp
	var alReq AgentDomainListReq
	if err := c.BindJSON(&alReq); err != nil {
		RespParamErr(c)
		return
	}
	merchantId := 1
	//获取列表
	list, err := model.GetAgentDomainList(model.Db, merchantId, alReq.Id)
	if err != nil {
		RespServerErr(c)
		return
	}
	data.List = list
	RespJson(c, status.OK, data)
}

/**
代理添加域名
*/
func AgentDomainAdd(c *gin.Context) {
	var adReq AgentDomainAddReq
	var ad model.AgentDomain
	if err := c.BindJSON(&adReq); err != nil {
		RespParamErr(c)
		return
	}
	//merchantId := 1
	timestamp := time.Now().Unix()
	ad.MerchantId = 1
	ad.AgentId = adReq.AgentId
	ad.Domain = adReq.Domain
	ad.CreateTime, ad.UpdateTime = timestamp, timestamp
	fmt.Println(ad)
	_, err := ad.AddAgentDomain(model.Db)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

/**
删除代理域名
*/
func AgentDomainDel(c *gin.Context) {
	var adReq AgentDomainReq
	var ad model.AgentDomain
	c.BindJSON(&adReq)
	ad.MerchantId = 1
	ad.Id = adReq.Id
	_, err := ad.DelAgentDomain(model.Db)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}
