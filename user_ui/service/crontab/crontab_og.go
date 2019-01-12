package crontab

import (
	"errors"
	"fmt"
	"game2/lib/game"
	"game2/lib/game/og"
	"game2/model"
	"regexp"
	"strconv"
	"time"
)

type GameOGCrontab struct {
	name string
}

// 匹配时间 返回时间戳 1541030400
func MatchDate(Msg string) (string, error) {
	var str string
	myexp := regexp.MustCompile(`[0-9]+`)
	result := myexp.FindAllStringSubmatch(Msg, 1)
	if result == nil {
		return "", errors.New("Please enter valid IP information!")
	}
	for _, value := range result {
		if value != nil {
			for _, str = range value {
			}
		}
	}
	timeInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return "", err
	}
	timeInt /= 1000
	timeSter := strconv.FormatInt(timeInt, 10)
	//	timeSter := time.Unix(timeInt, 0)
	//	timeSter.Format("2006-01-02 03:04:05 PM")
	return timeSter, nil
}

// 处理时间字段 返回 格式2018-11-12 08:22:54
func getTimeStr(mtime string) string {
	var date string
	myexp := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})`)
	result := myexp.FindAllString(mtime, 1)
	for _, date = range result {
	}
	return date
}

//New time，返回 格式2018-11-12 08:22:54,间隔10分钟
func newOGTimeDate() (startTime, endTime string) {
	const base_format = "2006-01-02 15:04:05"
	timeInt := time.Now().Unix()
	timeSter := time.Unix(timeInt, 0)
	start, _ := time.ParseDuration("-10m")
	sDate := timeSter.Add(start)
	nowTime := getTimeStr(timeSter.String())
	begin := getTimeStr(sDate.String())
	startTime, endTime = begin, nowTime
	return
}

func (g *GameOGCrontab) QueryRecord() (interface{}, error) {
	gameName, err := game.NewGame(g.name)
	if err != nil {
		return nil, err
	}
	sDate, eDate := newOGTimeDate()
	//todo: 各个游戏设置自己的查询参数
	req := map[string]interface{}{
		"start_time": "2018-11-29 06:50:46",
		"end_time":   "2018-11-29 06:59:46",
	}
	fmt.Println("data:", sDate, eDate, req)
	return gameName.QueryRecord(req)
}

func (g *GameOGCrontab) RecordList2Db(src interface{}) error {
	list := make([]model.GameRecordOg, 0)
	recordlist, ok := src.([]og.ResPutRecord)
	if !ok {
		return errors.New("record list to db error: data type error!")
	}
	timenow := time.Now().Unix()
	for i := range recordlist {
		bet, err := strconv.ParseFloat(recordlist[i].Bet, 64)
		if err != nil {
			return err
		}
		lose, err := strconv.ParseFloat(recordlist[i].Bet, 64)
		if err != nil {
			return err
		}
		betdate, err := MatchDate(recordlist[i].BettingDate)
		if err != nil {
			return err
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
	err := model.GameRecordOgAddInfo(model.Db, list)
	if err != nil {
		return err
	}
	return nil
}

func NewOGTask() Task {
	return &GameOGCrontab{
		name: "og",
	}
}

func init() {
	Register("og", NewOGTask)
}
