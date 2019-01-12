package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserBill struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id,omitempty"`
	SettAmt    float64 `json:"sett_amt"`
	Type       int     `json:"type"`
	Memo       string  `json:"memo"`
	OrderSn    string  `json:"order_no"`
	Status     int     `json:"status"`
	CreateTime int64   `json:"create_time"`
	Code       int     `json:"code,omitempty"`
	UpdateTime int64   `json:"update_time"`
	CodeSn     string  `json:"code_sn"`
	OldBalance float64 `json:"old_balance"`
	About      string  `json:"about"`
	Balance    float64 `json:"balance"`
}

type UserWithdraw struct {
	UserId     int     `json:"user_id"`
	MerchantId int     `json:"merchant_id"`
	CardNo     string  `json:"card_no"`
	Money      float64 `json:"money"`
	Status     int     `json:"status"`
	Memo       string  `json:"memo"`
	Addition   string  `json:"addition"`
	CreateTime int64   `json:"create_time"`
	OrderSn    string  `json:"order_sn"`
	Fee        float64 `json:fee`
}

type UserBillTotal struct {
	Code  int     `json:"code"`
	Total float64 `json:"total"`
}

type UserBillInterface interface {
	GetUserBillListByUserId()
	GetUserBillCount()
}

//获取资金列表
func GetUserBillListByUserId(db *gorm.DB, userId int, page int, pageCount int, startTime int64, endTime int64) ([]UserBill, error) {
	var billList []UserBill
	if err := db.Table("user_bill").Select("id,sett_amt,memo,type,order_sn,status,create_time").
		Where("user_id=? AND create_time>=? AND create_time<=?", userId, startTime, endTime).
		Order("id desc").Offset((page - 1) * pageCount).Limit(pageCount).Find(&billList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return billList, nil
}

//获取资金列表记录数
func GetUserBillCount(db *gorm.DB, userId int, startTime int64, endTime int64) (int, error) {
	var total int
	if err := db.Table("user_bill").Where("user_id=? AND create_time>=? AND create_time<=?", userId, startTime, endTime).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (userWithdraw UserWithdraw) ApplayWithdraw(db *gorm.DB) (error, int) {
	var user User
	//事务开始
	tx := db.Begin()
	//判断中心余额是否足够
	if err := tx.Table("user").Select("balance").Where("id=?", userWithdraw.UserId).First(&user).Error; err != nil {
		logrus.Error(err)
		return err, 0
	}
	balance := user.Balance
	if balance < userWithdraw.Money {
		tx.Callback()
		return nil, 101 //中心余额不足
	}
	//较少中心账户余额
	if err := tx.Table("user").Where("id=?", userWithdraw.UserId).UpdateColumn("balance", gorm.Expr("balance-?", userWithdraw.Money)).Error; err != nil {
		logrus.Error(err)
		tx.Callback()
		return err, 0
	}
	//插入提现记录
	if err := db.Create(&userWithdraw).Error; err != nil {
		logrus.Error(err)
		tx.Callback()
		return err, 0
	}
	tx.Commit()
	return nil, 200
}

//获取资金列表
func GetUserBillTotalByUserId(db *gorm.DB, userId int) ([]UserBillTotal, error) {
	var ubTotal []UserBillTotal
	if err := db.Table("user_bill").Select("code, sum(sett_amt) as total").Where("user_id=?", userId).Group("code").Find(&ubTotal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return ubTotal, nil
}
