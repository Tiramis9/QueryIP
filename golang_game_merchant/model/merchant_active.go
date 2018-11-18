package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//表： merchant_warn_log
type MerchantActive struct {
	Id           int    `json:"id"`
	MerchantId   int    `json:"merchant_id"`
	ActTitle     string `json:"act_title"`
	ActType      int    `json:"act_type"`
	Status       int    `json:"status"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	ResourceWeb  string `json:"resource_web"`
	ResourceWap  string `json:"resource_wap"`
	Content      string `json:"content"`
	JoinLimit    string `json:"join_limit"`
	RewardConfig string `json:"reward_config"`
	RelatedGame  string `json:"related_game"`
	Description  string `json:"description"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}

func (m MerchantActive) AddMerchantActivePay(db *gorm.DB) (int, error) {
	if err := db.Create(&m).Error; err != nil {
		logrus.Error(err)
		return 0, err
	}
	return m.Id, nil
}

func (m MerchantActive) EditMerchantActivePay(db *gorm.DB) (int, error) {
	if err := db.Debug().Model(&m).Where("merchant_id=? AND id=?", m.MerchantId, m.Id).Updates(m).Error; err != nil {
		logrus.Error(err)
		return 0, err
	}
	return m.Id, nil
}

func GetMerchantActive(db *gorm.DB, id int, merchantId int) (*MerchantActive, error) {
	var m MerchantActive
	if err := db.Table("merchant_active").Where("id=? AND merchant_id=?", id, merchantId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}

	return &m, nil
}

// 获取数据库 (merchant_active) 的活动描述
func GetMerchantActiveListInfo(db *gorm.DB, msg map[string]interface{}) ([]MerchantActive, int, error) {
	var list []MerchantActive
	Pagination := make(map[string]int)
	for key, aInfo := range msg {
		switch value := aInfo.(type) {
		case int:
			Pagination[key] = value
		}
	}
	if err := db.Debug().Table("merchant_active").Select("id,description,create_time,status,end_time,content").Offset((Pagination["page"]-1)*Pagination["page_count"]).Limit(Pagination["page_count"]).Where("status = ? AND merchant_id = ?",
		Pagination["type"], Pagination["merchant_id"]).Order("status DESC,create_time DESC,id,description,end_time").Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, err
		}
		return nil, 0, err
	}
	var count int
	if err := db.Table(`merchant_active`).Where("status = ? AND merchant_id =?", Pagination["type"], Pagination["merchant_id"]).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
