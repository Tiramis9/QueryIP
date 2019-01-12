package model

import "github.com/jinzhu/gorm"

type BlackIp struct {
	Id         int    `json:"id"`
	Ip         string `json:"ip"`
	CreateTime int64  `json:"create_time"`
	Status     int    `json:"status"`
	Area       string `json:"area"`
	MerchantId int    `json:"merchant_id"`
}

func GetIpInfo(db *gorm.DB, merchId int, ip string, area string) (*BlackIp, error) {
	var m BlackIp
	if err := db.Table("black_ip_list").
		Where("merchant_id=? AND ip=?", merchId, ip).Or("merchant_id=? AND area=?", merchId, area).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
