package model

import "github.com/jinzhu/gorm"

type Agent struct {
	Id            int    `json:"id"`
	MerchId       int    `json:"merch_id"`
	UserName      string `json:"user_name"`
	TrueName      string `json:"true_name"`
	Phone         string `json:"phone"`
	Skype         string `json:"skype"`
	Email         string `json:"email"`
	Qq            string `json:"qq"`
	Salt          string `json:"salt"`
	Status        int    `json:"status"`
	RegIp         string `json:"reg_ip"`
	LastLoginTime int    `json:"last_login_time"`
	LoginIp       string `json:"login_ip"`
	ApplyMemo     string `json:"apply_memo"`
	Auditor       string `json:"auditor"`
	AuditTime     string `json:"audit_time"`
	ClassId       int    `json:"class_id"`
}

type SuAgent struct {
	Agent
	UserCount int
	ClassName string
}

/**
获取代理列表
*/
func GetAgentList(db *gorm.DB, merchId int, condition string, status int, page int, pageCount int) ([]SuAgent, error) {
	str := "="
	if condition == ">=" {
		str = ">="
	} else {
		str = "="
	}
	var agentList []SuAgent
	if err := db.Debug().Table(`agent AS ag`).Joins(`
		LEFT JOIN merchant_agent_class AS mac ON ag.class_id=mac.id`).Joins(`
		LEFT JOIN user AS us ON ag.id=us.parent_id`).Select(`
		ag.id,ag.user_name,ag.merch_id,ag.true_name,ag.phone,ag.skype,ag.email,ag.qq,ag.salt,ag.status,ag.reg_ip,ag.last_login_time,`+
		`ag.login_ip,ag.apply_memo,ag.class_id,ag.auditor,ag.audit_time,mac.class_name,count(us.id) as user_count
	`).Where(`ag.merch_id=? AND ag.status`+str+`?`, merchId, status).Group("ag.id").Offset((page - 1) * pageCount).Limit(pageCount).Find(&agentList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return agentList, nil
}
