package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//表： merchant_warn_log
type SysActive struct {
	Id        int    `json:"id"`
	Type      int    `json:"type"`
	Condition string `json:"condition"`
	Desc      string `json:"desc"`
}

type PayType struct {
	PayType string `json:"pay_type"`
}

type Group struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
}

type Class struct {
	Id        int    `json:"id"`
	ClassName string `json:"class_name"`
}

//Group 等级； Class 层级 PayType
type PayActiveInfo struct {
	GroupList     []Group   `json:"group_list"`
	ClassList     []Class   `json:"class_list"`
	PayOnlineList []PayType `json:"pay_online_list"`
	PayTransList  []PayType `json:"pay_trans_list"`
}

func GetSysActiveList(db *gorm.DB) ([]SysActive, error) {
	var saList []SysActive
	if err := db.Table("sys_active").Find(&saList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return saList, nil
}

func GetSysActiveInfo(db *gorm.DB, aType int) (*SysActive,error){
	var a SysActive
	if err := db.Table("sys_active").Where("type=?",aType).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &a, nil
}

func GetPayActiveInfo(db *gorm.DB, merchId int) (*PayActiveInfo, error) {
	//获取系统活动的信息
	var res PayActiveInfo
	var classList []Class
	var groupList []Group
	var payTypeOnline []PayType
	var payTypeTrans []PayType
	//1.获取等级
	merchantUserGroupList, err := GetMerchantUserGroupList(db, merchId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Error(err)
			return nil, err
		}
	}
	for i := range merchantUserGroupList {
		var group Group
		group.Id = merchantUserGroupList[i].Id
		group.GroupName = merchantUserGroupList[i].GroupName
		groupList = append(groupList, group)
	}
	res.GroupList = groupList
	//2.获取层级
	merchantUserClassList, err := GetMerchantUserClassList(db, merchId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Error(err)
			return nil, err
		}
	}
	for i := range merchantUserClassList {
		var class Class
		class.Id = merchantUserClassList[i].Id
		class.ClassName = merchantUserClassList[i].ClassName
		classList = append(classList, class)
	}
	res.ClassList = classList
	//3.获取充值方式
	//在线支付
	if err := db.Table("sys_pay_type").Where("pay_tag=?", 1).Find(&payTypeOnline).Group("pay_type").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	res.PayOnlineList = payTypeOnline
	//转账汇款
	if err := db.Table("sys_pay_type").Where("pay_tag=?", 2).Find(&payTypeTrans).Group("pay_type").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	res.PayTransList = payTypeTrans
	return &res, nil
}

//查询所有的支付方式
func GetPayTypeList(db *gorm.DB) ([]PayType, error){
	var list []PayType
	if err := db.Table("sys_pay_type").Find(&list).Group("pay_type").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return list, nil
}
