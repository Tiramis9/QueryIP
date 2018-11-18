package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type OperationLogReq struct {
	Token     int    `json:"token"`
	UserName  string `json:"user_name"`
	MainMenu  string `json:"main_menu"`
	SubMenu   string `json:"sub_menu"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
	Page      int    `json:"page"`
	PageCount int    `json:"page_count"`
}

type OperationLog struct {
	UserName   string `json:"user_name"` //本字段在merchant_admin表，其他字段在merchant_admin_log表
	MainMenu   string `json:"main_menu"`
	SubMenu    string `json:"sub_menu"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	OperatorId int    `json:"id"`
	Ip         string `json:"ip"`
}

func constructWhere(req *OperationLogReq, id int) (string, []interface{}, error) {
	whereStr := make([]string, 1)
	cond := make([]interface{}, 1)
	whereStr[0] = " mal.operator_id = ? "
	cond[0] = id

	if req.UserName != "" {
		whereStr = append(whereStr, " ma.user_name = ? ")
		cond = append(cond, req.UserName)
	}

	if req.MainMenu != "" {
		whereStr = append(whereStr, " mal.main_menu = ? ")
		cond = append(cond, req.MainMenu)
	}

	if req.SubMenu != "" {
		whereStr = append(whereStr, " mal.sub_menu = ? ")
		cond = append(cond, req.SubMenu)
	}

	if req.StartTime != 0 {
		if req.EndTime < 0 {
			return "", nil, errors.New("invalid end_time")
		}
		if req.StartTime >= req.EndTime {
			return "", nil, errors.New("start_time > end_time")
		}
		whereStr = append(whereStr, " mal.create_time BETWEEN ? AND ? ")
		cond = append(cond, req.StartTime)
		cond = append(cond, req.EndTime)
	}
	return strings.Join(whereStr, " AND "), cond, nil
}

//类似 SELECT mal.main_menu,mal.content ,ma.user_name FROM merchant_admin_log AS al LEFT JOIN merchant_admin AS a ON mal.operator_id = ma.merchant_id where mal.main_menu='admin' AND mal.operator_id = 1;
func QueryOperationLog(db *gorm.DB, req *OperationLogReq, merchant_id int) ([]OperationLog, int, error) {
	whereStr, cond, err := constructWhere(req, merchant_id)
	if err != nil {
		return nil, 0, err
	}

	filter := db.Table(`merchant_admin_log AS mal`).Joins(`
	    LEFT JOIN merchant_admin AS ma ON mal.operator_id = ma.merchant_id
	`).Select(`
	    mal.main_menu,mal.sub_menu,mal.content,mal.create_time,mal.operator_id,mal.ip,ma.user_name
    `).Where

	var logs []OperationLog
	if err := filter(whereStr, cond...).Offset((req.Page - 1) * req.PageCount).Limit(req.PageCount).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	var count int
	if err := filter(whereStr, cond...).Count(&count).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return logs, count, nil
}
