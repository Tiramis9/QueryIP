package service

import (
	"errors"
	"game2/lib/game"
	_ "game2/lib/game/Imone"
	_ "game2/lib/game/agin"
	_ "game2/lib/game/allbet"
	_ "game2/lib/game/dg"
	_ "game2/lib/game/ky"
	_ "game2/lib/game/mg"
	_ "game2/lib/game/og"
	_ "game2/lib/game/sb"
	_ "game2/lib/game/vr"
	"game2/lib/utils"
	"game2/model"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserAccount2 struct {
	GameName    string  `json:"game_name,omitempty"`
	AccountName string  `json:"account_name,omitempty"`
	UserId      int     `json:"user_id,omitempty"`
	Money       float64 `json:"money,omitempty"`
}

const (
	AllBet   = "ALLBET"
	AllBetDz = "ALLBET_DZ"
	OG       = "OG"
	DG       = "DG"
	IMONE    = "IMONE"
	AGIN     = "AGIN"
	KY       = "KY"
)

//游戏类别
var gameMap = map[string]string{
	AllBet:    "allbet",
	AllBetDz:  "allbet",
	"SB_TY":   "sb",
	"VR_CP":   "vr",
	"BBIN_DZ": "bbin",
	"BBIN_CP": "bbin",
	"BBIN_TY": "bbin",
	"BBIN_SX": "bbin",
	OG:        "og",
	DG:        "dg",
	IMONE:     "imone",
	"MG_DZ":   "mg",
	AGIN:      "agin",
	KY:        "ky",
}

func GameLogin(gameCode string, userInfo map[string]interface{}, GameType string, AppId string) (string, error) {
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
	//每个游戏单独处理输入参数
	switch gameCode {
	case AllBet:
	case AllBetDz:
	case OG:
	case DG:
	case AGIN:
	case KY:
	}
	//判断游戏账户是否存在,不存在则创建
	Registered, err := dealRegister(gameCode, gameClass, userInfo)
	userInfo["registered"] = Registered //bool
	if err != nil {
		return str, err
	}
	userInfo["game_code"] = gameCode
	userInfo["game_type"] = GameType
	userInfo["app_id"] = AppId
	userInfo["game_user_name"] = getGameUserName(gameClass, userInfo)
	loginUrl, err := gameClass.Login(userInfo)
	if err != nil {
		return str, err
	}
	return loginUrl.(string), nil
}

func GameGetBalance(gameCode string, userInfo map[string]interface{}) (interface{}, error) {
	//游戏代码,判断
	gameStr := gameMap[gameCode]
	gameClass, err := game.NewGame(gameStr)
	userInfo["game_user_name"] = getGameUserName(gameClass, userInfo)
	if err != nil {
		return 0, nil
	}
	balance, err := gameClass.GetBalance(userInfo)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (u UserAccount2) TransAccount(db *gorm.DB, userId int, direction int, accountCode string, amount float64) (int, error) {
	var user model.User
	//从什么账户转至什么账户(开启事务)
	tx := db.Begin()
	if err := tx.Table("user").Select("id,balance,user_name,lang,time_zone").Where("id=?", userId).First(&user).Error; err != nil {
		tx.Callback()
		return 0, err
	}

	//将结构体转为json
	userInfo := make(map[string]interface{})
	userInfo["user_name"] = user.UserName
	userInfo["user_id"] = user.Id
	userInfo["game_code"] = accountCode
	userInfo["time_zone"] = strconv.Itoa(user.TimeZone)
	userInfo["lang"] = user.Lang

	var uab model.UserAccountBill
	uab.OldBalance = user.Balance
	uab.AccountName = u.AccountName
	uab.UserId = u.UserId
	uab.Money = u.Money
	uab.BillNo = utils.CreateOrderNo(u.UserId)
	//user_id, type, sett_amt, about, " + "balance, old_balance, order_sn, code, code_sn, create_time, update_time) values(?,?,?,?,?,?,?,?,?,?,?)"
	var ub model.UserBill
	ub.UserId = u.UserId
	ub.SettAmt = u.Money
	ub.OldBalance = user.Balance
	ub.OrderSn = utils.CreateOrderNo(u.UserId)
	ub.About = u.AccountName
	ub.Code = 300
	ub.CodeSn = uab.BillNo
	userInfo["order_sn"] = uab.BillNo
	//转出账户是中心账户
	if direction == 1 {
		if user.Balance < amount { //余额不足
			tx.Rollback()
			return 101, nil
		}
		uab.Type = 1
		ub.Type = -1
		//先减中心账户余额,再调用第三方接口
		if err := tx.Table("user").Where("id=?", userId).UpdateColumn("balance", gorm.Expr("balance-?", amount)).Error; err != nil {
			logrus.Error(err)
			tx.Rollback()
			return 0, err
		}
		//调用第三方接口
		_, err := gameTrans(accountCode, 1, amount, userInfo)
		if err != nil { //第三方接口调用失败
			logrus.Error(err)
			tx.Rollback()
			return 500, nil
		}
		uab.NewBalance = user.Balance - amount
		ub.Balance = user.Balance - amount
	} else {

		//查询第三方接口余额是否充足
		balance, err := GameGetBalance(accountCode, userInfo)
		if err != nil {
			tx.Rollback()
			return 500, nil
		}
		if balance.(float64) < amount {
			tx.Rollback()
			return 101, nil
		}
		//调用第三方接口
		_, err = gameTrans(accountCode, 0, amount, userInfo)
		if err != nil {
			logrus.Error(err)
			tx.Rollback()
			return 500, nil
		}
		uab.Type = 2
		ub.Type = 1
		//加中心账户余额
		if err := tx.Table("user").Where("id=?", userId).UpdateColumn("balance", gorm.Expr("balance+?", amount)).Error; err != nil {
			logrus.Error(err)
			tx.Rollback()
			return 0, err
		}
		uab.NewBalance = user.Balance + amount
		ub.Balance = user.Balance + amount
	}
	timestamp := time.Now().Unix()
	uab.CreateTime = timestamp
	uab.UpdateTime = timestamp
	//增加转账记录
	if err := tx.Create(&uab).Error; err != nil {
		logrus.Error(err)
		tx.Rollback()
		return 0, err
	}
	//增加交易明细
	if err := tx.Create(&ub).Error; err != nil {
		logrus.Error(err)
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return 200, nil
}

func gameTrans(gameCode string, direction int, amount float64, userInfo map[string]interface{}) (interface{}, error) {
	//处理转账数据库
	var info = make(map[string]interface{})
	//游戏代码,判断
	gameStr := gameMap[gameCode]
	gameClass, err := game.NewGame(gameStr)
	if err != nil {
		return 0, nil
	}
	for k, v := range userInfo {
		info[k] = v
	}
	info["amount"] = strconv.FormatFloat(amount, 'f', 0, 64)
	info["game_code"] = gameCode
	info["game_user_name"] = getGameUserName(gameClass, userInfo)
	info["user_name"] = userInfo["user_name"]
	info["user_id"] = userInfo["user_id"]
	info["time_zone"] = userInfo["time_zone"]
	info["lang"] = userInfo["lang"]
	if direction == 1 { //账户转游戏
		_, err = gameClass.Account2GameTransfer(info)
		if err != nil {
			return nil, err
		}
		return true, nil
	}
	//游戏转账户
	_, err = gameClass.Game2AccountTransfer(info)
	if err != nil {
		return nil, err
	}
	return true, nil
}

//检查游戏账户是否存在
//bool代表是否曾经注册
func dealRegister(gameCode string, gameClass game.Game, userInfo map[string]interface{}) (bool, error) {
	userId := userInfo["user_id"].(int)
	gameUserName := getGameUserName(gameClass, userInfo)
	account, err := model.GetAccountByGameName(model.Db, userId, gameCode)
	if err != nil {
		return false, err
	}
	if account != nil { //账号存在
		return true, nil
	}
	timestamp := time.Now().Unix()
	//账号不存在->创建账号
	ua := model.UserAccount{
		UserId:       userId,
		MerchantId:   userInfo["merchant_id"].(int),
		Type:         1,
		GameName:     gameCode,
		Status:       1,
		GameUserName: gameUserName,
		CreateTime:   timestamp,
		UpdateTime:   timestamp,
	}
	_, err = ua.InsertAccount(model.Db)
	if err != nil {
		return false, err
	}
	//账号创建成功
	return false, nil
}

//获取用户名
func getGameUserName(gameClass game.Game, userInfo map[string]interface{}) string {
	return gameClass.GetPrefix() + userInfo["user_name"].(string)
}
