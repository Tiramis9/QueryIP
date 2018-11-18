package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type MerchantAgentClass struct {
	Id            int     `json:"id"`
	MerchantId int `json:"merchant_id"`
	ClassName     string  `json:"class_name"`
	Mode          int     `json:"mode"`
	FdConfig      string  `json:"fd_config"`
	FcConfig      string  `json:"fc_config"`
	BonusCutRate  int     `json:"bonus_cut_rate"`
	RebateCutRate int     `json:"rebate_cut_rate"`
	SpreadAward   float64 `json:"spread_award"`
	CreateTime    int     `json:"create_time"`
	UpdateTime    int     `json:"update_time"`
}

/**
获取代理层级列表
*/
func GetMerchantAgentClassList(db *gorm.DB, merchId int) ([]MerchantAgentClass, error) {
	var agentClassList []MerchantAgentClass
	if err := db.Table("merchant_agent_class").Select("id, class_name").Find(&agentClassList, map[string]interface{}{"merchant_id": merchId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return agentClassList, nil
}

/**
获取代理层级详情
*/
func GetMerchantAgentClassInfo(db *gorm.DB, classId int, merchId int) (*MerchantAgentClass, error) {
	var agentClass MerchantAgentClass
	if err := db.First(&agentClass, map[string]interface{}{"id": classId, "merchant_id": merchId}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &agentClass, nil
}

/**
编辑代理层级信息
*/
func (m MerchantAgentClass) UpdateMerchantAgentClass(db *gorm.DB, merchId int, fields map[string]interface{}) (bool, error) {
	if err := db.Model(&m).Debug().Where("merchant_id=?", merchId).Updates(fields).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

/**
删除代理层级信息
*/
func (m MerchantAgentClass) DelMerchantAgentClass(db *gorm.DB, merchId int) (bool, error) {
	if err := db.Delete(&m).Where("merchant_id=?", merchId).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

/**
新增代理层级信息
*/
func (m MerchantAgentClass) AddMerchantAgentClass(db *gorm.DB) (int, error) {
	fmt.Println(m)
	if err := db.Create(&m).Error; err != nil {
		logrus.Error(err)
		return 0, err
	}
	return m.Id, nil
}
