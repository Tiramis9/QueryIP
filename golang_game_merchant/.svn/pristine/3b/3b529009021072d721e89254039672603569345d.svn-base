package model

import (
	"github.com/jinzhu/gorm"
)

type MerchantRebateLog struct {
	Id           int `json:"id"`
	MerchantId   int `json:"merchant_id"`   //商户id
	EffectiveBet int `json:"effective_bet"` //有效投注
	Rebate       int `json:"rebate"`        //反水金额
	Uid          int `json:"uid"`           //用户id
	CreateTime   int `json:"create_time"`   //创建时间
	UpdateTime   int `json:"update_time"`   //更新时间
	Status       int `json:"status"`        //0.未操作；1.已反水; 2.反水驳回
	Operator     int `json:"operator"`      //操作员
}

type RebateLogInfo struct {
	MerchantRebateLog
	UserName string `json:"user_name"` //用户名
}
type Sum struct {
	RebateTotal int `json:"rebate_total"` //反水总数
}

//商户反水生成记录
func RebateLogList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]RebateLogInfo, int, *Sum, error) {
	whereStr := "rl.merchant_id=?"
	condition := []interface{}{merchantId}

	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=rl.create_time AND rl.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=rl.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND rl.create_time<=?"
			condition = append(condition, v2)
		}
	}

	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var rebateLog []RebateLogInfo
	if err := db.Table("merchant_rebate_log as rl").Joins(`
	left join user as u on rl.uid = u.id`).Select(`
		rl.id, 
		rl.effective_bet,
		rl.rebate,
		rl.status,
		rl.create_time,
		u.user_name
	`).Where(whereStr, condition...).Offset((page - 1) * pageCount).Limit(pageCount).Find(&rebateLog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil, nil
		}
		return nil, 0, nil, err
	}
	//反水总数
	var rebateTotal Sum
	if err := db.Table("merchant_rebate_log as rl").Joins(`
	left join user as u on rl.uid = u.id`).Select(`
		sum(rebate) as rebate_total
	`).Where(whereStr, condition...).Offset((page - 1) * pageCount).Limit(pageCount).Find(&rebateTotal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil, nil
		}
		return nil, 0, nil, err
	}
	//
	var count int
	if err := db.Table("merchant_rebate_log as rl").Joins(`
	left join user as u on rl.uid = u.id`).Select(`
		sum(rebate) as rebate_total
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, nil, err
	}

	return rebateLog, count, &rebateTotal, nil
}

//商户反水生成记录
/*func RebateSuccessLogList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]RebateLogInfo, int, *Sum, error) {
	whereStr := "rl.merchant_id=?"
	condition := []interface{}{merchantId}

	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=rl.create_time AND rl.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=rl.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND rl.create_time<=?"
			condition = append(condition, v2)
		}
	}

	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var rebateLog []RebateLogInfo
	if err := db.Table("merchant_rebate_log as rl").Joins(`
	left join user as u on rl.uid = u.id`).Select(`
		rl.id,
		rl.effective_bet,
		rl.rebate,
		rl.status,
		rl.create_time,
		u.user_name
	`).Where(whereStr, condition...).Offset((page - 1) * pageCount).Limit(pageCount).Find(&rebateLog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil, nil
		}
		return nil, 0, nil, err
	}
	//反水总数
	var rebateTotal Sum
	if err := db.Table("merchant_rebate_log as rl").Joins(`
	left join user as u on rl.uid = u.id`).Select(`
		sum(rebate) as rebate_total
	`).Where(whereStr, condition...).Offset((page - 1) * pageCount).Limit(pageCount).Find(&rebateTotal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil, nil
		}
		return nil, 0, nil, err
	}
	//
	var count int
	if err := db.Table("merchant_rebate_log as rl").Joins(`
	left join user as u on rl.uid = u.id`).Select(`
		sum(rebate) as rebate_total
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, nil, err
	}

	return rebateLog, count, &rebateTotal, nil
}*/
