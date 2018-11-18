package model

import "github.com/jinzhu/gorm"

//涉及表：merchant（读）、 merchant_website
type BasicInfo struct {
	AgentPlatUrl      string `json:"agent_plat_url"`
	AgentSpreadUrl    string `json:"agent_spread_url"`
	AllowIpMinute     int    `json:"allow_ip_minute"`
	AppDownloadUrl    string `json:"app_download_url"`
	AppLogo           string `json:"app_logo"`
	ActiveStatus      int    `json:"active_status"`
	Code              string    `json:"code"`
	RegStatus         int    `json:"reg_status"`
	ServiceOnlineUrl  string `json:"service_online_url"`
	WithdrawSimpleMax float64 `json:"withdraw_single_max"`
	WithdrawSimpleMin float64 `json:"withdraw_single_min"`
	Name              string `json:"name"`
	Id                string `json:"id"`   //在另一个表 merchant
	Bail              int    `json:"bail"` //在另一个表 merchant
}

func QueryMerchantBasicInfo(db *gorm.DB, mid int) (BasicInfo, error) {
	var biList BasicInfo
	if err := db.Table(`merchant AS m`).Select(`
	mw.active_status,mw.agent_plat_url,mw.agent_spread_url,mw.allow_ip_minute,
	mw.app_download_url,mw.app_logo,mw.code,mw.name,mw.reg_status,
    mw.service_online_url,mw.withdraw_single_max,mw.withdraw_single_min,m.id
	`).Where("mw.merchant_id = ?", mid).Joins(`LEFT JOIN merchant_website as mw ON m.id = mw.merchant_id`).Find(&biList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return BasicInfo{}, nil
		}
		return BasicInfo{}, err
	}
	return biList, nil
}

// {
// "withdraw_single_max":"withdraw_max",
// "withdraw_single_min":"withdraw_min",
// "agent_plat_url":"example.com/plat",
// "agent_spread_url":"example.com/spread",
// "app_download_url":"example.com/app_download",
// "service_online_url":"example.com/online_service",
// "app_logo":"logoX",
// "allow_ip_minute":10,
// "active_status":1,
// "reg_status":1
// }
func ModifyMerchantBasicInfo(db *gorm.DB, bi *BasicInfo, mid int) error {
	fields := map[string]interface{}{
		"withdraw_single_max": bi.WithdrawSimpleMax,
		"withdraw_single_min": bi.WithdrawSimpleMin,
		"agent_plat_url":      bi.AgentPlatUrl,
		"agent_spread_url":    bi.AgentSpreadUrl,
		"app_download_url":    bi.AppDownloadUrl,
		"service_online_url":  bi.ServiceOnlineUrl,
		"app_logo":            bi.AppLogo,
		"allow_ip_minute":     bi.AllowIpMinute,
		"active_status":       bi.ActiveStatus,
		"reg_status":          bi.RegStatus,
	}

	return db.Table(`merchant_website`).Where(`merchant_id = ?`, mid).Updates(fields).Error
}
