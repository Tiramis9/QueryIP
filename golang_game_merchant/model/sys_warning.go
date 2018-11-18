package model

import (
	"github.com/jinzhu/gorm"
)

//表： merchant_warn_log
type WarnLog struct {
	UserName   string `json:"user_name"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	IP         string `json:"ip"`
	Device     int    `json:"device"`
}

//表 merchant_warn 和 sys_warning
type WarningPolicy struct {
	Id           int    `json:"id"`             //out
	SysWarningId int    `json:"sys_warning_id"` //out
	Status       int    `json:"status"`         //out
	Value        string `json:"value"`          //out
	Name         string `json:"name"`           //out,本字段在sys_warning中
	Code         string `json:"code"`           //out,本字段在sys_warning中
}

type SysWarningPolicy struct {
	Id int `json:"sys_warning_id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

//SELECT mwl.content,mwl.create_time,mwl.ip,mwl.device,mwl.merchant_id,u.user_name,u.id FROM merchant_warn_log AS mwl LEFT JOIN user AS u ON mwl.user_id = u.id  WHERE mwl.merchant_id = 1;
func QueryWarningLog(db *gorm.DB, mid int, page, pageCount int) ([]WarnLog, int, error) {

	filter := db.Table(`merchant_warn_log AS mwl
     `).Joins(`LEFT JOIN user AS u ON mwl.user_id = u.id
     `).Select(`u.user_name,mwl.content,mwl.create_time,mwl.ip,mwl.device`).Where

	var count int
	if err := filter(`mwl.merchant_id = ?`, mid).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var m []WarnLog
	if err := filter(`mwl.merchant_id = ?`, mid).Offset((page - 1) * pageCount).Limit(pageCount).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return m, count, nil
}

// SQL类似: SELECT mw.sys_warning_id, mw.status, mw.value, sw.name FROM merchant_warn AS mw RIGHT JOIN sys_warning AS sw ON mw.sys_warning_id = sw.id WHERE mw.merchant_id = 1 AND sw.status = 1;
func QueryWarningPolicy(db *gorm.DB, mid int) ([]WarningPolicy, error) {
	var wplist []WarningPolicy
	err := db.Table(`merchant_warn AS mw`).Select(`mw.id, mw.sys_warning_id, mw.status, mw.value, sw.name, sw.code
    `).Joins(`RIGHT JOIN sys_warning AS sw ON mw.sys_warning_id = sw.id`).Where(`mw.merchant_id = ? AND sw.status = ?`, mid, 1).Find(&wplist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return wplist, nil
}

// 表sys_warning
func QuerySysWarningPolicy(db *gorm.DB, ) ([]SysWarningPolicy, error) {
	var swplist []SysWarningPolicy
	err := db.Table(`sys_warning`).Select(`id, name, code
    `).Where(`status = ?`,  1).Find(&swplist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return swplist, nil
}

//表 merchant_warn 和 sys_warning
func ModifyWarningPolicy(db *gorm.DB, mid int, wp *WarningPolicy) error {
	fields := map[string]interface{}{"status": wp.Status, "value": wp.Value}
	err := db.Table(`merchant_warn`).Where(`merchant_id = ? AND id = ?`, mid, wp.Id).Updates(fields).Error
	if err != nil {
		return nil
	}

	return nil
}

// add 预警策略到merchant_warn
func AddWarningPolicy(db *gorm.DB, mid, sysWarningId int, value string) error {
	err := db.Table("merchant_warn").Create(struct {
		SysWarningId int
		Value        string
	}{sysWarningId, value}).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除预警策略到merchant_warn
func DelWarningPolicy(db *gorm.DB, mid, warnId int) error {
	err := db.Table("merchant_warn").Where(`merchant_id = ? AND id = ?`, mid, warnId).Delete(WarningPolicy{}).Error
	if err != nil {
		return err
	}
	return nil
}
