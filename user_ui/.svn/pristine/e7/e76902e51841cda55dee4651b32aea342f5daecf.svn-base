package model

import "github.com/jinzhu/gorm"

type MerchantWebsite struct {
	Id                  int     `json:"id"`
	MerchantId          int     `json:"merchant_id"`
	Name                string  `json:"name,omitempty"`
	AgentPlatUrl        string  `json:"agent_plat_url"`
	AgentSpreadUrl      string  `json:"agent_spread_url"`
	AppDownloadUrl      string  `json:"app_download_url"`
	AppLogo             string  `json:"app_logo"`
	AllowIpMinute       int     `json:"allow_ip_minute"`
	RegStatus           int     `json:"reg_status"`
	ActiveStatus        int     `json:"active_status"`
	Code                string  `json:"code,omitempty"`
	WithdrawSingleMin   float64 `json:"withdraw_single_min"`
	WithdrawSingleMax   float64 `json:"withdraw_single_max"`
	Credit              float64 `json:"credit"`
	FsRate              int     `json:"fs_rate"`
	FyRate              int     `json:"fy_rate"`
	AgentEmail          string  `json:"agent_email"`
	ServiceEmail        string  `json:"service_email"`
	RiskEmail           string  `json:"risk_email"`
	WebsitePhone        string  `json:"website_phone"`
	RegPayPass          int     `json:"reg_pay_pass"`
	RegSecurityQuestion int     `json:"reg_security_question"`
	RegTrueName         int     `json:"reg_true_name"`
	RegPhone            int     `json:"reg_phone"`
	RegEmail            int     `json:"reg_email"`
	Appid               string  `json:"appid"`
	Secret              string  `json:"secret"`
	IpWhite             string  `json:"ip_white"`
	TempId              int     `json:"temp_id"`
	EffectTime          int     `json:"effect_time"`
	ExpireTime          int     `json:"expire_time"`
}

type MerchantWebsiteInterface interface {
	GetMerchantWebsiteREG()
}

func GetMerchantWebsiteReg(db *gorm.DB, merchantId int) (*MerchantWebsite, error) {
	/*stmt, errs := Db.Prepare("SELECT id,merchant_id,reg_pay_pass,reg_security_question,reg_true_name,reg_phone,reg_email FROM merchant_website" + " Where merchant_id=?")
	if errs != nil {
		fmt.Println(errs)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	conv := MerchantWebsite{}
	row.Scan(&conv.Id, &conv.Merchant_id, &conv.Reg_pay_pass, &conv.Reg_security_question, &conv.Reg_true_name, &conv.Reg_phone, &conv.Reg_email)

	return conv*/
	var m MerchantWebsite
	if err := db.Table("merchant_website").Where("merchant_id=?", merchantId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (m MerchantWebsite) GetMerchantWebsiteAppDownload(id int) (merchat_website MerchantWebsite) {
	/*stmt, errs := Db.Prepare("SELECT id,merchant_id,app_download_url FROM merchant_website" + " Where merchant_id=?")
	if errs != nil {
		fmt.Println(errs)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	conv := MerchantWebsite{}
	row.Scan(&conv.Id, &conv.Merchant_id, &conv.App_download_url)

	return conv*/ /*stmt, errs := Db.Prepare("SELECT id,merchant_id,app_download_url FROM merchant_website" + " Where merchant_id=?")
	if errs != nil {
		fmt.Println(errs)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	conv := MerchantWebsite{}
	row.Scan(&conv.Id, &conv.Merchant_id, &conv.App_download_url)

	return conv*/
	return MerchantWebsite{}
}

func GetDomainInitInfo(db *gorm.DB, domain string) (*MerchantWebsite, error) {
	var m MerchantWebsite
	if err := db.Table("merchant_domain as md").Joins("LEFT JOIN merchant m ON md.merchant_id=m.id").
		Joins("LEFT JOIN merchant_website mw ON md.merchant_id=mw.merchant_id").
		Select("m.id,m.appid,m.secret,md.ip_white,m.effect_time,m.expire_time,mw.name,mw.agent_plat_url,mw.agent_spread_url,"+
			"mw.app_download_url,mw.app_logo,mw.allow_ip_minute,mw.reg_status,mw.active_status,mw.withdraw_single_min,"+
			"mw.withdraw_single_max,mw.fs_rate,mw.fy_rate,mw.agent_email,mw.service_email,mw.risk_email,mw.website_phone,"+
			"mw.reg_pay_pass,mw.reg_security_question,mw.reg_true_name,mw.reg_phone,mw.reg_email").
		Where("md.domain=?", domain).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
