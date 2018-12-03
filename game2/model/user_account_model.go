package model

import (
	"github.com/jinzhu/gorm"
)

//INSERT INTO user_account_bill (account_name, user_id, money, ok, " + "old_balance, new_balance, bill_no, create_time, update_time) values(?,?,?,?,?,?,?,?,?)
type UserAccountBill struct {
	AccountName string  `json:"account_name"`
	UserId      int     `json:"user_id"`
	Money       float64 `json:"money"`
	Ok          int     `json:"ok"`
	OldBalance  float64 `json:"old_balance"`
	NewBalance  float64 `json:"new_balance"`
	BillNo      string  `json:"bill_no"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
	Type        int64   `json:"type"`
}

type UserAccountInterface interface {
	GetAccountByUserId()
}

type Account struct {
	Channel     string  `json:"channel"`
	GameName    string  `json:"game_name"`
	GameCode    string  `json:"game_code"`
	Type        int     `json:"type"`
	GameBalance float64 `json:"game_balance"`
}

type UserAccount struct {
	Type         int    `json:"type"`
	UserId       int    `json:"user_id"`
	GameUserName string `json:"game_user_name"`
	Status       int    `json:"status"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
	GameName     string `json:"game_name"`
	MerchantId   int    `json:"merchant_id"`
	MgId         int    `json:"mg_id"` //mg游戏 游戏账户id
}

func GetAccountListByUserId(db *gorm.DB,userId, merchantId int) ([]Account, error) {
	var m []Account
	if err := db.Table("user_account ua").
		Joins("LEFT JOIN sys_game sg ON ua.game_name=sg.game_code").
		Joins("LEFT JOIN merchant_game mg ON mg.game_id=sg.id").
		Select("sg.channel,sg.game_name,sg.game_code,sg.type").
		Where("ua.user_id=? and mg.merchant_id=? and sg.parent_id=0 ", userId, merchantId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

func GetAccountByGameName(db *gorm.DB, userId int, gameName string) (*UserAccount, error) {
	var m UserAccount
	if err := db.Table("user_account").Where("user_id=? AND game_name=?", userId, gameName).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ua UserAccount) InsertAccount(db *gorm.DB) (bool, error) {
	if err := db.Table("user_account").Create(&ua).Error; err != nil {
		return false, err
	}
	return true, nil
}

func UpdateAccount(db *gorm.DB, userId int, gameCode string, param map[string]interface{}) (bool, error) {
	if err := db.Table("user_account").Where("user_id=? and game_name=?", userId, gameCode).Updates(param).Error; err != nil {
		return false, err
	}
	return true, nil
}