package model

import "github.com/jinzhu/gorm"

type MerchantWarn struct {
	Status    int `json:"status"`
	SysStatus int `json:"sys_status"`
}

func GetMerchantWarn(db *gorm.DB, merchId int, code string) (*MerchantWarn, error) {
	var m MerchantWarn
	if err := db.Table("merchant_warn as mw").Select("sw.status as sys_status, mw.status").
		Joins("LEFT JOIN sys_warning as sw ON mw.sys_warning_id=sw.id").
		Where("mw.merchant_id=? AND sw.code=?", merchId, code).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
