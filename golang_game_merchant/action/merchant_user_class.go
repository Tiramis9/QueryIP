package action

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
	"time"
)

type UserClassEditReq struct {
	Id                 int    `json:"id" binding:"required"`
	ClassName          string `json:"class_name" binding:"required"`           //层级
	SupportMerchantPay string `json:"support_merchant_pay" binding:"required"` //支持的商户支付通道,eg:1,4,6,8
	WithdrawLimitCount int    `json:"withdraw_limit_count" binding:"required"` //单日提款次数上限 -1为不限制提现次数,0将无法提现
	RuleId             int    `json:"rule_id" binding:"required"`              //反水规则id
	IsFs               int    `json:"is_fs" binding:"required"`                //0不支持反水;1支持反水
	//Remark             string `json:"remark"`               //备注
}

type UserClassAddReq struct {
	ClassName          string `json:"class_name" binding:"required"`           //层级
	SupportMerchantPay string `json:"support_merchant_pay" binding:"required"` //支持的商户支付通道,eg:1,4,6,8
	WithdrawLimitCount int    `json:"withdraw_limit_count" binding:"required"` //单日提款次数上限 -1为不限制提现次数,0将无法提现
	RuleId             int    `json:"rule_id" binding:"required"`              //反水规则id
	IsFs               int    `json:"is_fs" binding:"required"`                //0不支持反水;1支持反水
	//Remark             string `json:"remark"`                                  //备注
}

type MerchantUserClassList struct {
	Id        int    `json:"id"`
	ClassName string `json:"class_name"`
}

//层级列表返回
type MerchantUserClassListResp struct {
	List []MerchantUserClassList `json:"list"`
}

//层级详细信息
type UserClassInfoResp struct {
	Id                 int    `json:"id" binding:"required"`
	ClassName          string `json:"class_name"`           //层级
	SupportMerchantPay string `json:"support_merchant_pay"` //支持的商户支付通道,eg:1,4,6,8
	WithdrawLimitCount int    `json:"withdraw_limit_count"` //单日提款次数上限 -1为不限制提现次数,0将无法提现
	RuleId             int    `json:"rule_id"`              //反水规则id
	IsFs               int    `json:"is_fs"`                //0不支持反水;1支持反水
	MemberNum          int    `json:"member_num"`
	//Remark             string `json:"remark"`               //备注
}

/**
用户层级列表
*/
func UserClassList(c *gin.Context) {
	var data MerchantUserClassListResp
	//todo 从token获取 merchantId
	merchantId := 1
	list, err := model.GetMerchantUserClassList(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	for i := range list {
		temp := MerchantUserClassList{
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
func UserClassInfo(c *gin.Context) {
	var macReq UserClassEditReq
	if err := c.BindJSON(&macReq); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	merchantId := 1
	info, err, count := model.GetMerchantUserClassInfo(model.Db, merchantId, macReq.Id)
	if err != nil {
		RespServerErr(c)
		return
	}

	if info == nil {
		RespSuccess(c)
		return
	}

	res := UserClassInfoResp{
		Id:                 info.Id,
		ClassName:          info.ClassName,
		SupportMerchantPay: info.SupportMerchantPay,
		WithdrawLimitCount: info.WithdrawLimitCount,
		RuleId:             info.RuleId,
		IsFs:               info.IsFs,
		MemberNum:          count,
	}
	RespJson(c, status.OK, res)
}

/**
用户层级添加
*/
func UserClassAdd(c *gin.Context) {
	var macReq UserClassAddReq
	var mac model.MerchantUserClass

	if err := c.BindJSON(&macReq); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	//todo 从 token 获取 merchantId
	merchantId := 1
	timestamp := time.Now().Unix()
	mac.MerchantId = merchantId
	mac.ClassName = macReq.ClassName
	mac.SupportMerchantPay = macReq.SupportMerchantPay
	mac.WithdrawLimitCount = macReq.WithdrawLimitCount
	mac.RuleId = macReq.RuleId
	mac.IsFs = macReq.IsFs
	mac.CreateTime = timestamp
	_, err := mac.AddMerchantUserClass(model.Db)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

/**
用户层级编辑
*/
func UserClassEdit(c *gin.Context) {
	var macReq UserClassEditReq
	var mac model.MerchantUserClass
	fields := make(map[string]interface{})
	if err := c.BindJSON(&macReq); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	merchantId := 1
	mac.Id = macReq.Id
	timestamp := time.Now().Unix()

	fields["class_name"] = macReq.ClassName
	fields["support_merchant_pay"] = macReq.SupportMerchantPay
	fields["withdraw_limit_count"] = macReq.WithdrawLimitCount
	fields["rule_id"] = macReq.RuleId
	fields["is_fs"] = macReq.IsFs
	fields["update_time"] = timestamp

	_, err := mac.UpdateMerchantUserClass(model.Db, merchantId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

/**
会员层级删除
*/
func UserClassDel(c *gin.Context) {
	var macReq UserClassEditReq
	var mac model.MerchantUserClass
	if err := c.BindJSON(&macReq); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	mac.Id = macReq.Id
	//todo 从 token 获取merchantId
	merchantId := 1
	info, err := mac.DelMerchantUserClass(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, info)

}
