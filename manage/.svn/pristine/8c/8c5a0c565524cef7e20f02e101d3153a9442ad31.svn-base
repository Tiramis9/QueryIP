package action

import (
	"encoding/json"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MerchantAgentClassReq struct {
	Id            int           `json:"id" binding:"required"`
	ClassName     string        `json:"class_name"`
	Mode          int           `json:"mode"`
	FdSport       []interface{} `json:"fd_sport"`
	FdLottery     []interface{} `json:"fd_lottery"`
	FdPeople      []interface{} `json:"fd_people"`
	FdChess       []interface{} `json:"fd_chess"`
	FdElectornic  []interface{} `json:"fd_electornic"`
	FdBetUser     []interface{} `json:"fd_bet_user"`
	FdBetBill     []interface{} `json:"fd_bet_bill"`
	BonusCutRate  int           `json:"bonus_cut_rate"`
	RebateCutRate int           `json:"rebate_cut_rate"`
	FcTeamProfit  []interface{} `json:"fc_team_profit"`
	FcSport       []interface{} `json:"fc_sport"`
	FcLottery     []interface{} `json:"fc_lottery"`
	FcPeople      []interface{} `json:"fc_people"`
	FcChess       []interface{} `json:"fc_chess"`
	FcElectornic  []interface{} `json:"fc_electornic"`
	FcBetUser     []interface{} `json:"fc_bet_user"`
	FcBetBill     []interface{} `json:"fc_bet_bill"`
	SpreadAward   float64       `json:"spread_award"`
	CreateTime    int           `json:"create_time"`
	UpdateTime    int           `json:"update_time"`
}

type MerchantAgentClassAddReq struct {
	ClassName string `json:"class_name" binding:"required"`
}

type MerchantAgentClass struct {
	Id        int    `json:"id"`
	ClassName string `json:"class_name"`
}

type MerchantAgentClassListResp struct {
	List []MerchantAgentClass `json:"list"`
}

/**
代理层级列表
*/
func AgentClassList(c *gin.Context) {
	var data MerchantAgentClassListResp
	merchantId := 1
	list, err := model.GetMerchantAgentClassList(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		logrus.Error(err)
		return
	}
	for i := range list {
		temp := MerchantAgentClass{
			Id:        list[i].Id,
			ClassName: list[i].ClassName,
		}
		data.List = append(data.List, temp)
	}
	RespJson(c, status.OK, data)
}

/**
代理层级详情
*/
func AgentClassInfo(c *gin.Context) {
	var macReq MerchantAgentClassReq
	if err := c.BindJSON(&macReq); err != nil {
		RespParamErr(c)
		return
	}
	merchantId := 1
	info, err := model.GetMerchantAgentClassInfo(model.Db, macReq.Id, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, info)
}

/**
代理层级新增
*/
func AgentClassAdd(c *gin.Context) {
	var m MerchantAgentClassAddReq
	var ma model.MerchantAgentClass
	if err:= c.BindJSON(&m);err!=nil{
		RespParamErr(c)
		return
	}
	merchantId := 1
	ma.ClassName = m.ClassName
	ma.MerchantId = merchantId
	id, err:= ma.AddMerchantAgentClass(model.Db)
	if err!=nil{
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, id)
}

/**
代理层级编辑
*/
func AgentClassEdit(c *gin.Context) {
	var macReq MerchantAgentClassReq
	var mac model.MerchantAgentClass
	fields := make(map[string]interface{})
	if err := c.BindJSON(&macReq); err != nil {
		RespParamErr(c)
		return
	}
	merchantId := 1
	mac.Id = macReq.Id
	timestamp := time.Now().Unix()
	fields["class_name"] = macReq.ClassName
	fields["mode"] = macReq.Mode
	fields["bonus_cut_rate"] = macReq.BonusCutRate
	fields["rebate_cut_rate"] = macReq.RebateCutRate

	count := len(macReq.FdSport)
	var fdConfig []interface{}
	i := 0
	m := make(map[string]interface{})
	for i < count {
		m["fd_sport"] = macReq.FdSport[i]
		m["fd_lottery"] = macReq.FdLottery[i]
		m["fd_people"] = macReq.FdPeople[i]
		m["fd_chess"] = macReq.FdChess[i]
		m["fd_electornic"] = macReq.FdElectornic[i]
		m["fd_bet_user"] = macReq.FdBetUser[i]
		m["fd_bet_bill"] = macReq.FdBetBill[i]
		fdConfig = append(fdConfig, m)
		i++
	}
	count = len(macReq.FcSport)
	var fcConfig []interface{}
	i = 0
	m2 := make(map[string]interface{})
	for i < count {
		m2["fc_team_profit"] = macReq.FcTeamProfit[i]
		m2["fc_sport"] = macReq.FcSport[i]
		m2["fc_lottery"] = macReq.FcLottery[i]
		m2["fc_people"] = macReq.FcPeople[i]
		m2["fc_chess"] = macReq.FcChess[i]
		m2["fc_bet_user"] = macReq.FcBetUser[i]
		m2["fc_bet_bill"] = macReq.FcBetBill[i]
		fcConfig = append(fcConfig, m2)
		i++
	}

	fdConfigJson, err := json.Marshal(fdConfig)
	if err != nil {
		RespServerErr(c)
		return
	}
	fields["fd_config"] = string(fdConfigJson)

	fcConfigJson, err := json.Marshal(fcConfig)
	fields["fc_config"] = string(fcConfigJson)
	if err != nil {
		RespServerErr(c)
		return
	}
	fields["spread_award"] = macReq.SpreadAward
	fields["create_time"], fields["update_time"] = timestamp, timestamp
	_, err = mac.UpdateMerchantAgentClass(model.Db, merchantId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

/**
代理层级删除
*/
func AgentClassDel(c *gin.Context) {
	var macReq MerchantAgentClassReq
	var mac model.MerchantAgentClass
	if err := c.BindJSON(&macReq); err != nil {
		RespParamErr(c)
		return
	}
	mac.Id = macReq.Id
	merchantId := 1
	info, err := mac.DelMerchantAgentClass(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, info)
}
