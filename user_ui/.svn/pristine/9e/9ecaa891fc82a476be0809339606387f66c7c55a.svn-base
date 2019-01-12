package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	AnnouncementInfo struct {
		Content    string `json:"content"`
		Device     int    `json:"device"`
		StartTime  int64  `json:"start_time"`
		EndTime    int64  `json:"end_time"`
		Status     int    `json:"status"`
		Sort       int    `json:"sort"`
		Type       int    `json:"type"`
		Title      string `json:"title"`
		Url        string `json:"url"`
		MerchantID int    `json:"merchant_id"`
		CreateTime int64  `json:"create_time"`
		UpdateTime int64  `json:"update_time"`
		ID         int    `json:"id"`
	}
	MerchantAnnouncement struct {
		AnnouncementInfo
	}
)

func GetMerchantAnnouncementList(db *gorm.DB, msg map[string]interface{}) ([]MerchantAnnouncement, error) {
	var list []MerchantAnnouncement
	nowTime := time.Now().Unix()
	if err := db.Table("merchant_announcement").
		Select("content,title,url,start_time,end_time,device,sort,type,status,merchant_id").
		Where("merchant_id=? AND start_time<=? AND end_time>? AND type=1 AND status=1",
		msg["merchant_id"], nowTime, nowTime).Group("sort desc").Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil,err
		}
		return nil,err
	}
	return list, nil
}
