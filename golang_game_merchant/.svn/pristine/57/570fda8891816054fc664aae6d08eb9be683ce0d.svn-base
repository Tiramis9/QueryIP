package model

import (
	"github.com/jinzhu/gorm"
)

type MerchantWebsite struct {
	Id                  int     `json:"id"`
	MerchantId          int     `json:"merchant_id"`
	Name                string  `json:"name,omitempty"`
	AgentPlatUrl        string  `json:"agent_plat_url,omitempty"`
	AgentSpreadUrl      string  `json:"agent_spread_url,omitempty"`
	AppDownloadUrl      string  `json:"app_download_url,omitempty"`
	AppLogo             string  `json:"app_logo,omitempty"`
	AllowIpMinute       int     `json:"allow_ip_minute,omitempty"`
	RegStatus           int     `json:"reg_status,omitempty"`
	ActiveStatus        int     `json:"active_status,omitempty"`
	Code                string  `json:"code,omitempty"`
	WithdrawSimpleMin   float64 `json:"withdraw_simple_min,omitempty"`
	WithdrawSimpleMax   float64 `json:"withdraw_simple_max,omitempty"`
	Credit              float64 `json:"credit,omitempty"`
	FsRate              int     `json:"fs_rate,omitempty"`
	FyRate              int     `json:"fy_rate,omitempty"`
	AgentEmail          string  `json:"agent_email,omitempty"`
	ServiceEmail        string  `json:"service_email,omitempty"`
	RiskEmail           string  `json:"risk_email,omitempty"`
	RegPayPass          int     `json:"reg_pay_pass,omitempty"`
	RegSecurityQuestion int     `json:"reg_security_question,omitempty"`
	RegTrueName         int     `json:"reg_true_name,omitempty"`
	RegPhone            int     `json:"reg_phone,omitempty"`
	RegEmail            int     `json:"reg_email,omitempty"`
}

// SELECT * FROM `merchant_website`  WHERE (merchant_id='{id}')
func GetMerchantWebsiteReg(db *gorm.DB, id int) (*MerchantWebsite, error) {
	var mw MerchantWebsite
	if err := db.Table(`merchant_website`).Where("merchant_id=?", id).Find(&mw).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}

		return nil, err
	}

	return &mw, nil
}

func GetMerchantWebsiteAppDownload(db *gorm.DB, id int) (*MerchantWebsite, error) {
	return GetMerchantWebsiteReg(db, id)
}
