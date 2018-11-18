package model

import (
	"github.com/jinzhu/gorm"
)

type GameDescription struct {
	GameId   int    `json:"game_id"`  //out
	GameName string `json:"game_name"` //out
	Status   int    `json:"status"` //out
}
type classfiedGameDescription struct {
	GameDescription
	Type int `json:"type"`
}

//ex: SELECT sg.game_name,sg.type,mg.game_id,mg.status FROM sys_game AS sg INNER JOIN merchant_game AS mg ON sg.id = mg.game_id WHERE mg.merchant_id =1;
func QueryMerchantGameStatus(db *gorm.DB, mid int) (map[int][]GameDescription, error) {
	var gds []classfiedGameDescription
	err := db.Table("sys_game AS sg").Joins(`
        INNER JOIN merchant_game AS mg ON sg.id = mg.game_id
	`).Where(`mg.merchant_id = ?`, mid).Select(`
        sg.game_name,sg.type,mg.game_id,mg.status`).Find(&gds).Error
	if err != nil {
		return nil, err
	}

	result := make(map[int][]GameDescription)
	for _, gd := range gds {
		if _, ok := result[gd.Type]; !ok {
			gdlist := []GameDescription{}
			result[gd.Type] = gdlist
		}
		result[gd.Type] = append(result[gd.Type], gd.GameDescription)
	}

	return result, nil
}

func ModifyMerchantGameStatus(db *gorm.DB, gd *GameDescription, mid int) error {
	err := db.Table("merchant_game AS mg").Where("mg.merchant_id = ?", mid).Updates(map[string]interface{}{"game_id": gd.GameId, "status": gd.Status}).Error
	if err != nil {
		return err
	}
	return nil
}
