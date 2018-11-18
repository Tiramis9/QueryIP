package model

import "github.com/jinzhu/gorm"

type AgentDomain struct {
	Id         int    `json:"id"`
	AgentId    int    `json:"agent_id"`
	MerchantId int    `json:"merchant_id"`
	Domain     string `json:"domain"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

/**
获取代理列表
*/
func (ad AgentDomain) AddAgentDomain(db *gorm.DB) (int, error) {
	if err := db.Create(&ad).Error; err != nil {
		return 0, err
	}
	return ad.Id, nil
}

/**
获取代理域名列表
*/
func GetAgentDomainList(db *gorm.DB, merchId int, agentId int) ([]AgentDomain, error) {
	var agentdomainList []AgentDomain
	if err := db.Table(`agent_domain`).Select(`id,domain`).
		Where("agent_id=? AND merchant_id=?", agentId, merchId).Find(&agentdomainList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return agentdomainList, nil
}

/**
删除代理域名信息
*/
func (ad AgentDomain) DelAgentDomain(db *gorm.DB) (bool, error) {
	if err := db.Where("merchant_id=?", ad.MerchantId).Delete(&ad).Error; err != nil {
		return false, err
	}
	return true, nil
}