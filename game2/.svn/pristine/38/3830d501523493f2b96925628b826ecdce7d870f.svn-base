package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type MerchantActive struct {
	Id          int    `json:"id"`
	MerchantId  int    `json:"merchant_id"`
	ActTitle    string `json:"act_title"`
	ActType     int    `json:"act_type"`
	Status      int    `json:"status"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	ResourceWeb string `json:"resource_web"`
	ResourceWap string `json:"resource_wap"`
	Content     string `json:"content"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

func GetMerchantActiveList(db *gorm.DB, merchantId int, nowTime int64) ([]MerchantActive, error) {
	var list []MerchantActive
	if err := db.Table("merchant_active").
		Select("id,resource_web,resource_wap,description,create_time,status,end_time").
		Where("merchant_id=? AND start_time<=? AND end_time>? AND status=1", merchantId, nowTime, nowTime).Order("create_time DESC").Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return list, nil
}

func GetMerchantActiveInfo(db *gorm.DB, id int, merchantId int) (*MerchantActive, error) {
	var m MerchantActive
	if err := db.Table("merchant_active").Where("id=? AND merchant_id=? AND status=1", id, merchantId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &m, nil
}
