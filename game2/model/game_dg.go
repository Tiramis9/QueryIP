package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GameRecordDg struct {
	Id           int
	LobbyId      int
	TableId      int
	DgId         int64
	PlayId       int64
	GameType     int
	GameId       int
	ShoeId       int64
	MemberId     int64
	BetTime      time.Time
	CalTime      time.Time
	WinOrLoss    float64
	WinOrLossz   float64
	BetPoints    float64
	BetPointsz   float64
	AvailableBet float64
	UserName     string
	Result       string
	BetDetail    string
	BetDetailz   string
	Ip           string
	Ext          string
	IsRevocation int
	ParentBetId  int64
	CurrencyId   int
	DeviceType   int
	PluginId     int
	CreateTime   int64
	UpdateTime   int64
}

func (dg *GameRecordDg) NewRecord(db *gorm.DB) error {
	return db.Where(`game_id=?`, dg.DgId).FirstOrCreate(dg).Error
}
