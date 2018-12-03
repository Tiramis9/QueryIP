package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GameRecordKy struct {
	Id            int
	GameId        string //具有唯一性
	Account       string
	ServerId      int
	KindId        int
	TableId       int
	ChairId       int
	UserCount     int
	CardValue     string
	CellScore     float64
	AllBet        float64
	Profit        float64
	Revenue       float64
	GameStartTime time.Time
	GameEndTime   time.Time
	ChannelId     int
	LineCode      string
	CreateTime    int64
	UpdateTime    int64
}

func (ky *GameRecordKy) NewRecord(db *gorm.DB) error {
	return db.Where(`game_id=?`, ky.GameId).FirstOrCreate(ky).Error
}
