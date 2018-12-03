package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Agent struct {
	Id            int     `json:"id"`
	MerchantId    int     `json:"merchant_id"`
	UserName      string  `json:"user_name"`
	AgentCode     string  `json:"agent_code"`
}

func GetAgentByAgentCode(db *gorm.DB, agentCode string) (*Agent, error) {
	var m Agent
	if err := db.Table("agent").Where("agent_code=?", agentCode).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &m, nil
}