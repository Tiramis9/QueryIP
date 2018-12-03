package crontab

import (
	"fmt"
	"game2/lib/game"
	"game2/lib/game/ky"
	"game2/model"
	"time"
)

type TaskKy struct {
	name string
}

func (t *TaskKy) QueryRecord() (interface{}, error) {
	gameName, _ := game.NewGame(t.name)

	now := time.Now()
	req := map[string]interface{}{
		"start_time": fmt.Sprint(now.Add(-5*time.Second).UnixNano() / 1e6),
		"end_time":   fmt.Sprint(now.UnixNano() / 1e6),
	}
	return gameName.QueryRecord(req)
}

func (t *TaskKy) RecordList2Db(recordList interface{}) error {
	list, ok := recordList.([]ky.RecordInfo)
	if !ok {
		return fmt.Errorf("record list to db error: data type error")
	}

	now := time.Now().Unix()
	for i := range list {
		record := &model.GameRecordKy{
			GameId:        list[i].GameId,
			Account:       list[i].Accounts,
			ServerId:      list[i].ServerId,
			KindId:        list[i].KindId,
			TableId:       list[i].TableId,
			ChairId:       list[i].ChairId,
			UserCount:     list[i].UserCount,
			CardValue:     list[i].CardValue,
			CellScore:     list[i].CellScore,
			AllBet:        list[i].AllBet,
			Profit:        list[i].Profit,
			Revenue:       list[i].Revenue,
			GameStartTime: list[i].GameStartTime,
			GameEndTime:   list[i].GameEndTime,
			ChannelId:     list[i].ChannelId,
			LineCode:      list[i].LineCode,
			CreateTime:    now,
			UpdateTime:    now,
		}

		// record.NewRecord中已做去重处理
		if err := record.NewRecord(model.Db); err != nil {
			return err
		}
	}

	return nil
}

func NewKyTask() Task {
	return &TaskKy{
		name: "ky",
	}
}

func init() {
	Register("ky", NewKyTask)
}
