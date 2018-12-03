package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

/*
`game_provider` varchar(10) DEFAULT NULL COMMENT '游戏提供',
  `member_name` varchar(20) DEFAULT NULL COMMENT '玩家帐号',
  `game_name` varchar(20) DEFAULT NULL COMMENT '游戏名',
  `betting_code` varchar(30) DEFAULT NULL COMMENT '下注代码',
  `betting_date` varchar(30) DEFAULT NULL COMMENT '下注日期',
  `game_id` varchar(20) DEFAULT NULL COMMENT '游戏id',
  `round_no` varchar(30) DEFAULT NULL,
  `result` varchar(30) DEFAULT NULL,
  `bet` decimal(15,2) DEFAULT NULL,
  `win_lose_result` decimal(15,2) DEFAULT NULL COMMENT '输赢结果',
  `betting_amount` decimal(15,2) DEFAULT NULL COMMENT '下注金额',
  `valid_bet` decimal(15,2) DEFAULT NULL COMMENT '有效下注额',
  `win_lose_amount` decimal(15,2) DEFAULT NULL COMMENT '输赢金额',
  `balance` decimal(20,2) DEFAULT NULL COMMENT '余额',
  `currency` varchar(5) DEFAULT NULL COMMENT '币种',
  `handicap` varchar(50) DEFAULT NULL,
  `status` varchar(30) DEFAULT NULL,
  `game_category` varchar(10) DEFAULT NULL COMMENT '游戏类别',
  `settle_date` date DEFAULT NULL COMMENT '清算日期',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
*/
type (
	GameRecordOg struct {
		GameProvider   string  `json:"game_provider"`
		MemberName     string  `json:"member_name"`
		GameName       string  `json:"game_name"`
		BettingCode    string  `json:"betting_code"`
		BettingDate    string  `json:"betting_date"`
		GameId         string  `json:"game_id"`
		RoundNo        string  `json:"round_no"`
		Result         string  `json:"result"`
		Bet            float64 `json:"bet"`
		Win_loseResult float64 `json:"win_lose_result"`
		BettingAmount  float64 `json:"betting_amount"`
		ValidBet       float64 `json:"valid_bet"`
		WinLoseAmount  float64 `json:"win_lose_amount"`
		Balance        float64 `json:"balance"`
		Currency       string  `json:"currency"`
		Handicap       string  `json:"handicap"`
		Status         string  `json:"status"`
		GameCategory   string  `json:"game_category"`
		SettleDate     string  `json:"settle_date"`
		Remark         string  `json:"remark"`
		CreateTime     int64   `json:"create_time"`
		UpdateTime     int64   `json:"update_time"`
	}
)

// 插入SQL (sys_message 与UserMessage)信息
func (in GameRecordOg) GameRecordOgAddInfo(db *gorm.DB) error {
	return db.Create(&in).Error
}

// 插入SQL (sys_message 与UserMessage)信息
func GameRecordOgAddInfo(db *gorm.DB, in []GameRecordOg) error {
	insertPost := fmt.Sprintf("%v", "INSERT INTO  game_record_og  (game_provider,member_name,game_name,betting_code,betting_date,game_id,round_no,result,bet,win_lose_result,betting_amount,valid_bet,win_lose_amount,balance,currency,handicap,status,game_category,settle_date,remark,create_time,update_time) VALUES ")
	const rowSQL = "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	var inserts []string
	vals := []interface{}{}
	for i := range in {
		var count int
		if err := db.Model(&GameRecordOg{}).Where("betting_code=?", in[i].BettingCode).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			inserts = append(inserts, rowSQL)
			vals = append(vals, in[i].GameProvider, in[i].MemberName, in[i].GameName, in[i].BettingCode, in[i].BettingDate, in[i].GameId, in[i].RoundNo, in[i].Result,
				in[i].Bet, in[i].Win_loseResult, in[i].BettingAmount, in[i].ValidBet, in[i].WinLoseAmount, in[i].Balance, in[i].Currency, in[i].Handicap, in[i].Status, in[i].GameCategory,
				in[i].SettleDate, in[i].Remark, in[i].CreateTime, in[i].UpdateTime)
		}
	}
	insertPost += strings.Join(inserts, ",")
	if len(vals) > 0 {
		err := db.Exec(insertPost, vals...).Error
		if err != nil {
			return err
		}
	}
	return nil
}
