package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Merchant struct {
		Id            int
		MerchantName  string
		Url           string
		ContactName   string
		ContactPhone  string
		ContactEmail  string
		Code          string
		Balance       float64
		CreateTime    int64
		EffectTime    int64
		ExpireTime    int64
		Bail          float64
		UpdateTime    int64
		WithdrawWin   float64
		WithdrawMax   float64
		DockingPeople string
		Appid         string
		Secret        string
	}

	MerchantPayConfig struct {
		Id            int
		SysPayTypeId  int
		ConfigContext string
		MerchantId    int
		Sort          int
		Code          string
		MerchNo       string
		Md5Key        string
		PublicKey     string
		SecretKey     string
		SimpleMin     float64
		SimpleMax     float64
		DayStopMax    float64
		FeeRate       int
		Status        int
		CreateTime    int64
		UpdateTime    int64
		Remark        string
		Account       string
		Qrcode        string
		Url           string
	}

	MerchantGame struct {
		Id             int
		MerchantId     int
		GameId         int
		Status         int
		CreateTime     int64
		UpdateTime     int64
		Url            string
		CommissionRate int
	}
)

func GetMerchantById(db *gorm.DB, id int) (*Merchant, error) {
	var merchant Merchant
	if err := db.Table(`merchant`).Where(`id=?`, id).Find(&merchant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &merchant, nil
}

func CheckMPCIdAndPayTagMatch(db *gorm.DB, id, payTag int) (bool, error) {
	var mpc MerchantPayConfig
	if err := db.Table(`merchant_pay_config`).Where(`id=? AND pay_tag=?`, id, payTag).Find(&mpc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (mpc *MerchantPayConfig) NewRecord(db *gorm.DB) error {
	return db.Table(`merchant_pay_config`).Create(mpc).Error
}

func (mpc *MerchantPayConfig) Update(db *gorm.DB, payTag int) error {
	m := make(map[string]interface{})
	m["sys_pay_type_id"] = mpc.SysPayTypeId
	m["sort"] = mpc.Sort
	m["code"] = mpc.Code
	m["single_min"] = mpc.SimpleMin
	m["single_max"] = mpc.SimpleMax
	m["day_stop_max"] = mpc.DayStopMax
	m["status"] = mpc.Status
	m["update_time"] = mpc.UpdateTime
	m["remark"] = mpc.Remark
	if payTag == 1 {
		m["merch_no"] = mpc.MerchNo
		m["md5_key"] = mpc.Md5Key
		m["public_key"] = mpc.PublicKey
		m["secret_key"] = mpc.SecretKey
		m["url"] = mpc.Url
	} else if payTag == 2 {
		m["account"] = mpc.Account
	}

	return db.Table(`merchant_pay_config`).Where(`id=? AND merchant_id=?`, mpc.Id, mpc.MerchantId).Updates(m).Error
}

func (mpc *MerchantPayConfig) Delete(db *gorm.DB) error {
	return db.Where(`merchant_id=? AND id=?`, mpc.MerchantId, mpc.Id).Delete(mpc).Error
}

func GetMGByMerchantIdAndGameId(db *gorm.DB, merchantId, gameId int) (*MerchantGame, error) {
	var mg MerchantGame
	if err := db.Table(`merchant_game`).Where(`
		merchant_id=? AND game_id=?
	`, merchantId, gameId).Find(&mg).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &mg, nil
}

func IsMidGidUidMatch(db *gorm.DB, merchantId, gameId, userId int) (bool, error) {
	var count int
	if err := db.Table(`merchant AS m`).Joins(`
		LEFT JOIN merchant_game AS mg ON mg.merchant_id=m.id
	`).Joins(`
		LEFT JOIN user AS u on u.merchant_id=m.merchant_id
	`).Where(`
		m.merchant_id=? AND mg.game_id=? AND u.id=? 
	`, merchantId, gameId, userId).Count(&count).Error; err != nil {
		return false, err
	}

	return count != 0, nil
}
