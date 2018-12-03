package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Advertisement struct {
	Type     int    `json:"type,omitempty"`
	Location int    `json:"location"`
	Sort     int    `json:"srot"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Url      string `json:"url"`
}

func GetAdvertisementList(db *gorm.DB, merchantId int, where map[string]interface{}) ([]Advertisement, error) {
	var adsList []Advertisement
	var params []interface{}
	params = append(params, merchantId)
	str := "merchant_id=? "
	if v, ok := where["type"]; ok {
		str += " AND type=?"
		params = append(params, v)
	}
	if v2, ok := where["location"]; ok {
		str += " AND location=?"
		params = append(params, v2)
	}
	if err := db.Table("merchant_ads").Where(str, params...).Find(&adsList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return adsList, nil
}
