package crontab

import (
	"fmt"
	"game2/lib/game"
	"game2/lib/game/dg"
	"game2/model"
	"time"
)

type TaskDg struct {
	name string
}

func (t *TaskDg) QueryRecord() (interface{}, error) {
	gameName, _ := game.NewGame(t.name)
	req := make(map[string]interface{})
	return gameName.QueryRecord(req)
}

func (t *TaskDg) RecordList2Db(recordList interface{}) error {
	list, ok := recordList.([]dg.GameRecordInfo)
	if !ok {
		return fmt.Errorf("record list to db error: data type error")
	}

	now := time.Now().Unix()
	for i := range list {
		record := &model.GameRecordDg{
			LobbyId:      list[i].LobbyId,
			TableId:      list[i].TableId,
			DgId:         list[i].Id,
			PlayId:       list[i].PlayId,
			GameType:     list[i].GameType,
			GameId:       list[i].GameId,
			ShoeId:       list[i].ShoeId,
			MemberId:     list[i].MemberId,
			WinOrLoss:    list[i].WinOrLoss,
			WinOrLossz:   list[i].WinOrLossz,
			BetPoints:    list[i].BetPoints,
			BetPointsz:   list[i].BetPointsz,
			AvailableBet: list[i].AvailableBet,
			UserName:     list[i].UserName,
			Result:       list[i].Result,
			BetDetail:    list[i].BetDetail,
			BetDetailz:   list[i].BetDetailz,
			Ip:           list[i].Ip,
			Ext:          list[i].Ext,
			IsRevocation: list[i].IsRevocation,
			ParentBetId:  list[i].ParentBetId,
			CurrencyId:   list[i].CurrencyId,
			DeviceType:   list[i].DeviceType,
			PluginId:     list[i].Pluginid,
			CreateTime:   now,
			UpdateTime:   now,
		}
		record.BetTime, _ = time.ParseInLocation("2006-01-02 15:04:05", list[i].BetTime, time.Local)
		record.CalTime, _ = time.ParseInLocation("2006-01-02 15:04:05", list[i].CalTime, time.Local)

		// record.NewRecord中已做去重处理
		if err := record.NewRecord(model.Db); err != nil {
			return err
		}
	}

	return nil
}

func NewDgTask() Task {
	return &TaskDg{
		name: "dg",
	}
}

func init() {
	Register("dg", NewDgTask)
}
