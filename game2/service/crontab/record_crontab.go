package crontab

import (
	"errors"
	"game2/lib/game"
	_ "game2/lib/game/Imone"
	_ "game2/lib/game/allbet"
	_ "game2/lib/game/dg"
	"game2/lib/game/og"
	_ "game2/lib/game/sb"
	_ "game2/lib/game/vr"
	"game2/model"
	"strconv"
	"time"
)

type UserAccount2 struct {
	GameName    string  `json:"game_name,omitempty"`
	AccountName string  `json:"account_name,omitempty"`
	UserId      int     `json:"user_id,omitempty"`
	Money       float64 `json:"money,omitempty"`
}

const (
	OG    = "OG"
	DG    = "DG"
	IMONE = "IMONE"
)

//游戏类别
var gameMap = map[string]string{

	"OG":    "og",
	"DG":    "dg",
	"IMONE": "imone",
}

func GameRecord(gameCode string, info map[string]interface{}) (interface{}, error) {
	var str string
	//游戏代码,判断
	gameStr, ok := gameMap[gameCode]
	if !ok {
		return str, errors.New("game code error")
	}
	gameClass, err := game.NewGame(gameStr)
	if err != nil {
		return str, err
	}
	// 请求查询
	record, err := gameClass.QueryRecord(info)
	if err != nil {
		return str, err
	}
	// 每个游戏单独处理请求数据
	switch gameCode {
	case OG:
		timenow := time.Now().Unix()
		list := make([]model.GameRecordOg, 0)
		recordlist, ok := record.([]og.ResPutRecord)
		if !ok {
			return nil, errors.New("not recordlist ")
		}
		for i := range recordlist {
			bet, err := strconv.ParseFloat(recordlist[i].Bet, 64)
			if err != nil {
				return nil, err
			}
			lose, err := strconv.ParseFloat(recordlist[i].Bet, 64)
			if err != nil {
				return nil, err
			}
			betdate, err := MatchDate(recordlist[i].BettingDate)
			if err != nil {
				return nil, err
			}
			temp := model.GameRecordOg{
				GameProvider:   recordlist[i].GameProvider,
				MemberName:     recordlist[i].MemberName,
				GameName:       recordlist[i].GameName,
				BettingCode:    recordlist[i].BettingCode,
				BettingDate:    betdate,
				GameId:         recordlist[i].GameId,
				RoundNo:        recordlist[i].Roundno,
				Result:         recordlist[i].Result,
				Bet:            bet,
				Win_loseResult: lose,
				BettingAmount:  recordlist[i].BettinGamount,
				ValidBet:       recordlist[i].ValidBet,
				WinLoseAmount:  recordlist[i].Winloseamount,
				Balance:        recordlist[i].Balance,
				Currency:       recordlist[i].Currency,
				Handicap:       recordlist[i].Handicap,
				Status:         recordlist[i].Status,
				GameCategory:   recordlist[i].Gamecategory,
				SettleDate:     recordlist[i].Settledate,
				Remark:         recordlist[i].Remark,
				CreateTime:     timenow,
				UpdateTime:     timenow,
			}
			list = append(list, temp)
		}
		err = model.GameRecordOgAddInfo(model.Db, list)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case DG:
	}
	return nil, nil
}
