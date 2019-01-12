package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type MerchantUserGroup struct {
	Id               int     `json:"id"`
	MerchantId       int     `json:"merchant_id"`        //商户id
	GroupName        string  `json:"group_name"`         //等级
	CreateTime       int64   `json:"create_time"`        //创建时间
	UpdateTime       int64   `json:"update_time"`        //更新时间
	Remark           string  `json:"remark"`             //备注
	FsSportRate      int     `json:"fs_sport_rate"`      //体育返水
	FsLotteryRate    int     `json:"fs_lottery_rate"`    //彩票返水
	FsPeopleRate     int     `json:"fs_people_rate"`     //真人返水
	FsElectronicRate int     `json:"fs_electronic_rate"` //电子游戏返水
	EffectiveBet     float64 `json:"effective_bet"`      //有效投注额
	UpgradeReward    int     `json:"upgrade_reward"`     //晋级彩金
	BirthdayReward   int     `json:"birthday_reward"`    //生日彩金
}

type MerchantUserGroupList struct {
	MerchantUserGroup
	MemberNum int `json:"member_num"` //会员数
}

type MerchantUserGroupSimpleList struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"` //等级
}

type MerchantUserGroupConfig struct {
	Id                   int    `json:"id"`
	MerchantId           int    `json:"merchant_id"`            //商户id
	CreateTime           int64  `json:"create_time"`            //创建时间
	UpdateTime           int64  `json:"update_time"`            //更新时间
	Remark               string `json:"remark"`                 //备注
	ExtraFsSwitch        int    `json:"extra_fs_switch"`        //额外反水开关 0.关; 1.开
	UpgradeRewardSwitch  int    `json:"upgrade_reward_switch"`  //晋级彩金开关 0.晋级彩金,自动全无;1.晋级彩金有,自动发送无; 2晋级彩金有,自动发送有
	BirthdayRewardSwitch int    `json:"birthday_reward_switch"` //生日彩金开关 0.生日彩金,自动全无;1.生日彩金有,自动发送无; 2生日彩金有,自动发送有

}

/**
获取会员等级列表
*/
func GetMerchantUserGroupList(db *gorm.DB, merchantId int) ([]MerchantUserGroupList, error) {
	var userGroupList []MerchantUserGroupList
	if err := db.Table("merchant_user_group").Select(`
		id, 
		group_name,
		fs_sport_rate,
		fs_lottery_rate,
		fs_people_rate,
		fs_electronic_rate,
		effective_bet,
		upgrade_reward,
		birthday_reward
	`).Find(&userGroupList, map[string]interface{}{"merchant_id": merchantId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	var count int
	for k, value := range userGroupList {
		if err := db.Table(`user AS u`).Where("u.group_id = ? and merchant_id = ?", value.Id, merchantId).Count(&count).Error; err != nil {
			return nil, err
		}
		userGroupList[k].MemberNum = count
	}
	return userGroupList, nil
}

//获取简易会员等级列表
func GetMerchantUserSimpleGroupList(db *gorm.DB, merchantId int) ([]MerchantUserGroupSimpleList, error) {
	var userGroupList []MerchantUserGroupSimpleList
	if err := db.Table("merchant_user_group").Select(`
		id, 
		group_name
	`).Find(&userGroupList, map[string]interface{}{"merchant_id": merchantId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return userGroupList, nil
}

/**
获取用户等级详情
*/
/*func GetMerchantUserGroupInfo(db *gorm.DB, merchantId int, classId int) (*MerchantUserGroup, error, int) {
	var userGroup MerchantUserGroup
	if err := db.First(&userGroup, map[string]interface{}{"id": classId, "merchant_id": merchantId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, 0
		}
		return nil, err, 0
	}
	var count int
	if err := db.Table(`user AS u`).Where(
		"u.merchant_id = ? and u.class_id = ?", merchantId, classId).Count(&count).Error; err != nil {
		return nil, err, 0
	}
	return &userGroup, nil, count
}*/

/**
编辑会员等级信息
*/
func (m MerchantUserGroup) AddMerchantUserGroup(db *gorm.DB) (bool, error) {
	if err := db.Debug().Create(&m).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

/**
编辑会员等级信息
*/
func (m MerchantUserGroup) UpdateMerchantUserGroup(db *gorm.DB, merchantId int, fields map[string]interface{}) (bool, error) {
	if err := db.Model(&m).Where("merchant_id=?", merchantId).Updates(fields).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

/**
删除用户等级信息
*/
func (m MerchantUserGroup) DelMerchantUserGroup(db *gorm.DB, merchId int) (bool, error) {
	if err := db.Delete(&m).Where("merchant_id=?", merchId).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

//获取会员等级配置信息
func GetMerUserGroupConfigInfo(db *gorm.DB, merchantId int) (*MerchantUserGroupConfig, error) {
	var userGroupInfo MerchantUserGroupConfig
	if err := db.Table("merchant_user_group_config").Select(`
		extra_fs_switch,
		upgrade_reward_switch,
		birthday_reward_switch
	`).Find(&userGroupInfo, map[string]interface{}{"merchant_id": merchantId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &userGroupInfo, nil
}

/**
编辑会员等级信息
*/
func (m MerchantUserGroupConfig) EditMerUserGroupConfigInfo(db *gorm.DB, merchantId int, fields map[string]interface{}) (bool, error) {
	if err := db.Model(&m).Where("merchant_id=?", merchantId).Updates(fields).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}
