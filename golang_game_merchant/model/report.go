package model

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
)

type BizReport struct {
	SummaryUserBill // 一览表
	FeatureList     // `json:"data"`             // 游戏列表
	TipsList        // `json:"tips_list"`        // 打赏小费列表
	GameActiveList  // `json:"game_active_list"` // 游戏活动列表
}
type FeatureList struct {
	Chess `json:"chess"`
	Game  `json:"game"`
}
type Chess struct {
}
type Game struct {
}
type TipsList struct {
}
type GameActiveList struct {
	pt_Game `json:"pt_game"`
}
type pt_Game struct {
}
type SummaryUserBill struct {
	NewRegister    int     // `json:"new_register"`    // 新注册会员 .
	BonusAmount    float64 // `json:"bonus_amount"`    // 红利 ..
	Rebate         float64 // `json:"rebate"`          // 反水 .
	RechargeAmount float64 // `json:"rechargeAmount"` // 充值金额 .
	RechargeMember int     // `json:"recharge_member"` // 充值会员 .
	WithdrawAmount float64 // `json:"withdraw_amount"` // 提现 .
	EffectiveBet   float64 // `json:"effective_bet"`   // 有效投注
	Win            float64 // `json:"win"`             // 输赢结果
}

func StringtoInt(message string) int {
	date, err := strconv.Atoi(message)
	if err != nil {
		panic(err)
	}
	return date
}

func GetMerchantAnnouncement(db *gorm.DB, msg map[string]interface{}) (*BizReport, error) {
	var (
		start_time int
		end_time   int
	)
	for point, data := range msg {
		if point == "start_time" {
			switch m_begin := data.(type) {
			case string:
				start_time = StringtoInt(m_begin)
			case int:
				start_time = m_begin
			}
		}
		if point == "end_time" {
			switch m_end := data.(type) {
			case string:
				end_time = StringtoInt(m_end)
			case int:
				end_time = m_end
			}
		}
	}
	if start_time == 0 || end_time == 0 {
		return nil, errors.New("time invaild because is zero")
	}
	//var bizReport []BizReport
	bizReport := new(BizReport)
	if err := db.Table(`user`).Where("reg_time>= ? AND reg_time <=? AND status=1 ",
		start_time, end_time).Count(&(bizReport.NewRegister)).Error; err != nil {
		return nil, err
	}
	if err := db.Table(`user_bill`).Where("create_time>= ? AND create_time <=? AND status=1 AND code=100",
		start_time, end_time).Count(&(bizReport.RechargeMember)).Error; err != nil {
		return nil, err
	}
	if err := db.Debug().Table("user_bill").Select(`SUM(sett_amt) AS rebate `).Where("create_time>= ? AND create_time <=? AND status=1  AND code=800",
		start_time, end_time).Find(bizReport).Error; err != nil {
		return nil, err
	}
	if err := db.Debug().Table("user_bill").Select(`SUM(sett_amt) AS withdraw_amount `).Where("create_time>= ? AND create_time <=? AND status=1   AND code=200",
		start_time, end_time).Find(bizReport).Error; err != nil {
		return nil, err
	}
	if err := db.Debug().Table("user_bill").Select(`SUM(sett_amt) AS bonus_amount `).Where("create_time>= ? AND create_time <=? AND status=1  AND code=400",
		start_time, end_time).Find(bizReport).Error; err != nil {
		return nil, err
	}
	if err := db.Debug().Table("user_bill").Select(`SUM(sett_amt) AS recharge_amount `).Where("create_time>= ? AND create_time <=? AND status=1   AND code=100",
		start_time, end_time).Find(bizReport).Error; err != nil {
		return nil, err
	}
	return bizReport, nil
}
