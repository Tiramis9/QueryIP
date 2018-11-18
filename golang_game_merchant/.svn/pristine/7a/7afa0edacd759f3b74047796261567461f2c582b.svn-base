package action

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
	"time"
)

type UserGroupReq struct {
	Id               int     `json:"id"`
	GroupName        string  `json:"group_name" binding:"required"`    //等级
	FsSportRate      int     `json:"fs_sport_rate"`                    //体育返水
	FsLotteryRate    int     `json:"fs_lottery_rate"`                  //彩票返水
	FsPeopleRate     int     `json:"fs_people_rate"`                   //真人返水
	FsElectronicRate int     `json:"fs_electronic_rate"`               //电子游戏返水
	EffectiveBet     float64 `json:"effective_bet" binding:"required"` //有效投注额
	UpgradeReward    int     `json:"upgrade_reward"`                   //晋级彩金
	BirthdayReward   int     `json:"birthday_reward"`                  //生日彩金
}

type UserGroupDelReq struct {
	Id int `json:"id" binding:"required"`
}

type UserGroupInfo struct {
	Id               int     `json:"id"`
	MerchantId       int     `json:"merchant_id"`        //商户id
	GroupName        string  `json:"group_name"`         //等级
	FsSportRate      int     `json:"fs_sport_rate"`      //体育返水
	FsLotteryRate    int     `json:"fs_lottery_rate"`    //彩票返水
	FsPeopleRate     int     `json:"fs_people_rate"`     //真人返水
	FsElectronicRate int     `json:"fs_electronic_rate"` //电子游戏返水
	EffectiveBet     float64 `json:"effective_bet"`      //有效投注额
	UpgradeReward    int     `json:"upgrade_reward"`     //晋级彩金
	BirthdayReward   int     `json:"birthday_reward"`    //生日彩金
	MemberNum        int     `json:"member_num"`         //会员数
}

//等级详细信息
type UserGroupInfoListResp struct {
	List []UserGroupInfo
}

/*type UserGroupConfigReq struct {
	ExtraFsSwitch        string `json:"extra_fs_switch"`        //额外反水开关 0.关; 1.开
	UpgradeRewardSwitch  string `json:"upgrade_reward_switch"`  //晋级彩金开关 0.晋级彩金,自动全无;1.晋级彩金有,自动发送无; 2晋级彩金有,自动发送有
	BirthdayRewardSwitch string `json:"birthday_reward_switch"` //生日彩金开关 0.生日彩金,自动全无;1.生日彩金有,自动发送无; 2生日彩金有,自动发送有
}*/

type UserGroupConfigInfo struct {
	ExtraFsSwitch        int `json:"extra_fs_switch"`        //额外反水开关 0.关; 1.开
	UpgradeRewardSwitch  int `json:"upgrade_reward_switch"`  //晋级彩金开关 0.晋级彩金,自动全无;1.晋级彩金有,自动发送无; 2晋级彩金有,自动发送有
	BirthdayRewardSwitch int `json:"birthday_reward_switch"` //生日彩金开关 0.生日彩金,自动全无;1.生日彩金有,自动发送无; 2生日彩金有,自动发送有
}

/**
用户等级列表
*/
func GetMerchantUserGroupList(c *gin.Context) {
	var data UserGroupInfoListResp
	//todo 从token获取 merchantId
	merchantId := 1
	list, err := model.GetMerchantUserGroupList(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	for i := range list {
		temp := UserGroupInfo{
			Id:               list[i].Id,
			GroupName:        list[i].GroupName,
			FsSportRate:      list[i].FsSportRate,
			FsLotteryRate:    list[i].FsLotteryRate,
			FsPeopleRate:     list[i].FsPeopleRate,
			FsElectronicRate: list[i].FsElectronicRate,
			EffectiveBet:     list[i].EffectiveBet,
			BirthdayReward:   list[i].BirthdayReward,
			MemberNum:        list[i].MemberNum,
		}
		data.List = append(data.List, temp)
	}
	RespJson(c, status.OK, data)
}

/**
用户等级详情
*/
/*func GetUserGroupInfo(c *gin.Context) {
	var macReq UserClassEditReq
	c.BindJSON(&macReq)
	merchantId := 1
	info, err, count := model.GetMerchantUserGroupInfo(model.Db, merchantId, macReq.Id)
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
}*/

/**
用户等级操作(有 id 走编辑,无 id 走添加)
*/
func UserGroupOperate(c *gin.Context) {
	var macReq UserGroupReq
	var mac model.MerchantUserGroup
	fields := make(map[string]interface{})
	if err := c.BindJSON(&macReq); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo 从 token 获取 merchantId
	merchantId := 1

	if macReq.Id == 0 {
		//走添加等级
		timestamp := time.Now().Unix()
		mac.GroupName = macReq.GroupName
		mac.MerchantId = merchantId
		mac.FsSportRate = macReq.FsSportRate
		mac.FsLotteryRate = macReq.FsLotteryRate
		mac.FsPeopleRate = macReq.FsPeopleRate
		mac.FsElectronicRate = macReq.FsElectronicRate
		mac.UpgradeReward = macReq.UpgradeReward
		mac.BirthdayReward = macReq.BirthdayReward
		mac.CreateTime = timestamp

		_, err := mac.AddMerchantUserGroup(model.Db)
		if err != nil {
			RespServerErr(c)
			return
		}
		RespSuccess(c)

	} else {
		//有 id ,走编辑等级
		mac.Id = macReq.Id
		timestamp := time.Now().Unix()
		fields["group_name"] = macReq.GroupName
		fields["fs_sport_rate"] = macReq.FsSportRate           //体育返水
		fields["fs_lottery_rate"] = macReq.FsLotteryRate       //彩票返水
		fields["fs_people_rate"] = macReq.FsPeopleRate         //真人返水
		fields["fs_electronic_rate"] = macReq.FsElectronicRate //电子游戏返水
		fields["upgrade_reward"] = macReq.UpgradeReward        //晋级彩金
		fields["birthday_reward"] = macReq.BirthdayReward      //生日彩金
		fields["update_time"] = timestamp

		_, err := mac.UpdateMerchantUserGroup(model.Db, merchantId, fields)
		if err != nil {
			RespServerErr(c)
			return
		}
		RespSuccess(c)
	}

}

/**
用户等级删除
*/
func UserGroupDel(c *gin.Context) {
	var macReq UserGroupDelReq
	var mac model.MerchantUserGroup
	if err := c.BindJSON(&macReq); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	mac.Id = macReq.Id

	//todo 从 token 获取 merchantId
	merchantId := 1
	info, err := mac.DelMerchantUserGroup(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, info)
}

//获取用户等级配置
func GetUserGroupConfigInfo(c *gin.Context) {
	//var req UserGroupConfigInfo
	var data UserGroupConfigInfo
	/*if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}*/
	//todo 从 token 获取 merchantId
	merchantId := 1
	info, err := model.GetMerUserGroupConfigInfo(model.Db, merchantId)
	if err != nil {
		RespServerErr(c)
		return
	}
	if info == nil {
		RespSuccess(c)
		return
	}

	data.ExtraFsSwitch = info.ExtraFsSwitch
	data.UpgradeRewardSwitch = info.UpgradeRewardSwitch
	data.BirthdayRewardSwitch = info.BirthdayRewardSwitch

	RespJson(c, status.OK, data)
}

//编辑用户等级配置
func UserGroupConfigInfoEdit(c *gin.Context) {
	var req UserGroupConfigInfo
	var m model.MerchantUserGroupConfig
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	if req.ExtraFsSwitch < 0 || req.ExtraFsSwitch > 1 {
		RespParamErr(c)
		return
	}
	if req.BirthdayRewardSwitch < 0 || req.ExtraFsSwitch > 2 {
		RespParamErr(c)
		return
	}
	if req.UpgradeRewardSwitch < 0 || req.UpgradeRewardSwitch > 2 {
		RespParamErr(c)
		return
	}
	//todo 从 token 获取 merchantId
	merchantId := 1
	fields := make(map[string]interface{})
	//m.MerchantId = merchantId
	timestamp := time.Now().Unix()
	fields["extra_fs_switch"] = req.ExtraFsSwitch
	fields["upgrade_reward_switch"] = req.UpgradeRewardSwitch
	fields["birthday_reward_switch"] = req.BirthdayRewardSwitch
	fields["update_time"] = timestamp

	_, err := m.EditMerUserGroupConfigInfo(model.Db, merchantId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}

	RespSuccess(c)
}
