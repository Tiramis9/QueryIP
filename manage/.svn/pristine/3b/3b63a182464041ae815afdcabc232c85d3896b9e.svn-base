package model

import "github.com/jinzhu/gorm"

type (
	SysPayType struct {
		Id           int
		Platform     string
		PlatformCode string
		PayType      string
		Status       string
		CreateTime   int64
		UpdateTime   int64
		PayTag       string
	}
)

func GetSysPayTypeByPayType(db *gorm.DB, payType string) (*SysPayType, error) {
	var spt SysPayType
	if err := db.Table(`sys_pay_type`).Where(`pay_type=?`, payType).Find(&spt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &spt, nil
}
