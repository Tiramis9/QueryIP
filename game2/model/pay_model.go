package model

import "github.com/jinzhu/gorm"

type PayType struct {
	Id           int     `json:"id,omitempty"`
	Channel      string  `json:"channel"`
	PlatformCode string  `json:"platform_code,omitempty"`
	PayType      string  `json:"pay_type,omitempty"`
	MerchantId   int     `json:"merchant_id"`
	Account      string  `json:"account"`
	DayStopMax   float64 `json:"day_stop_max"`
	FeeRate      float64 `json:"fee_rate"`
	PayTag       float64 `json:"pay_tag"`
	Qrcode       string  `json:"qrcode"`
	SimpleMax    float64 `json:"simple_max"`
	SimpleMin    float64 `json:"simple_min"`
	MerchNo      string  `json:"merch_no"`
	Md5Key       string  `json:"md5_key"`
	PublicKey    string  `json:"public_key,omitempty"`
	SecretKey    string  `json:"secret_key,omitempty"`
}

//查询商户支持哪几种类型的支付方式
func GetPayTypeList(db *gorm.DB, merchantId int) ([]PayType, error) {
	var payTypeList []PayType
	if err := db.Table("merchant_pay_config mpc").Joins("LEFT JOIN sys_pay_type spt ON mpc.sys_pay_type_id = spt.id").
		Select("spt.pay_type").Where("mpc.merchant_id=?", merchantId).Group("spt.pay_type").Find(&payTypeList).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return payTypeList, nil
}

//根据商户支持的支付方式查询出具体支持的通道
func GetPayConfig(db *gorm.DB, merchantId int, payType string) ([]PayType, error) {
	var payTypeList []PayType
	if err := db.Table("merchant_pay_config mpc").Joins("LEFT JOIN sys_pay_type spt ON mpc.sys_pay_type_id = spt.id").
		Select("mpc.id,mpc.account,mpc.day_stop_max,mpc.fee_rate,spt.pay_tag,mpc.qrcode,mpc.simple_max,mpc.simple_min").
		Where("mpc.merchant_id=? AND spt.pay_type=?", merchantId, payType).Find(&payTypeList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return payTypeList, nil
}

//查看单独的一条支付类型的详情
func GetPayInfo(db *gorm.DB, id int, merchantId int) (*PayType,error) {
	var m PayType
	if err := db.Table("merchant_pay_config as mpc").Joins("LEFT JOIN sys_pay_type as spt ON mpc.sys_pay_type_id = spt.id").
		Select("mpc.id,mpc.account,mpc.day_stop_max,mpc.fee_rate,spt.pay_tag,mpc.qrcode,mpc.simple_max,mpc.simple_min," +
		"mpc.merch_no,mpc.md5_key,mpc.public_key,mpc.secret_key").Where("mpc.id=? AND mpc.merchant_id=? AND mpc.status=1", id, merchantId).Find(&m).Error; err != nil {
		if err==gorm.ErrRecordNotFound{
			return nil,nil
		}
		return nil, err
	}
	return &m,nil
}
