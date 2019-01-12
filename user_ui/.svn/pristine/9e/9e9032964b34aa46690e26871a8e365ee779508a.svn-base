package model

import "github.com/jinzhu/gorm"

type GameType struct {
	Type int `json:"type"`
}

type Game struct {
	Id       int    `json:"id"`
	Type     int    `json:"type,omitempty"`
	Channel  string `json:"channel"`
	GameName string `json:"game_name"`
	GameCode string `json:"game_code"`
}

type GameInfo struct {
	Id       int    `json:"id"`
	Type     int    `json:"type,omitempty"`
	Channel  string `json:"channel"`
	GameName string `json:"game_name"`
	GameCode string `json:"game_code"`
	Memo     string `json:"memo"`
	ParentId int    `json:"parent_id"`
	AppId    int    `json:"app_id"` //mg 游戏的app_id
}

func GetGameTypeList(db *gorm.DB, merchantId int) ([]GameType, error) {
	var gameTypeList []GameType
	if err := db.Table("merchant_game as mg").Select("sg.type").Joins("LEFT JOIN sys_game sg ON mg.game_id=sg.id").
		Where("mg.merchant_id=?", merchantId).Group("sg.type").Find(&gameTypeList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return gameTypeList, nil
}

func GetGameList(db *gorm.DB, ty int) ([]Game, error) {
	var gameList []Game
	if err := db.Table("sys_game").Where("type=? AND parent_id=0", ty).Find(&gameList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return gameList, nil
}

func GetGameSubList(db *gorm.DB, id int) ([]GameInfo, error) {
	var gameList []GameInfo
	if err := db.Table("sys_game").Where("parent_id=?", id).Find(&gameList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return gameList, nil
}

func GetGameSubCount(db *gorm.DB, id int) (int, error) {
	var total int
	if err := db.Table("sys_game").Where("parent_id=?", id).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
