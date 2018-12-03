package allbet

import (
	"encoding/json"
	"errors"
	"game2/lib/game"
	"game2/lib/utils"
	"math/rand"
	"strconv"
)

const (
	DesKey         = "f0Q2638+aVam6t0+YXaAtwamFNXPME5V"
	Md5Key         = "pRyd5cngzAja4LMdzUeELeX2yqIs/V1DwHPwNuobAug="
	PropertyId     = "2143220"
	ApiUrl         = "https://api3.apidemo.net:8443/"
	AgentName      = "3iwmaa"
	UserPassword   = "123456"
	UserPrefix     = "jkg_"
	LoginReturnUrl = "http://api.jkgsoft.com/game/allbet"
)

type GameAllBet struct {

}

var errorCode = map[string]string{"OK": "ok", "INTERNAL_ERROR": "internet error",
	"ILLEGAL_ARGUMENT": "illegal argument", "SYSTEM_MATAINING": "system mataining",
	"AGENT_NOT_EXIST": "agent not exist", "CLIENT_EXIST": "client exist", "CLIENT_PASSWORD_INCORRECT": "client password incorrect",
	"TOO_FREQUENT_REQUEST": "too frequent request", "CLIENT_NOT_EXIST": "client not exist", "TRANS_EXISTED": "trans existed",
	"LACK_OF_MONEY": "lack of money", "DUPLICATE_CONFIRM": "duplicate confirm", "TRANS_NOT_EXIST": "trans not exist",
	"DECRYPTION_FAILURE": "decryption failure", "FORBIDDEN": "forbidden", "INCONSISTENT_WITH_PRE_TRANS": "inconsistent with pre tans",
	"INVALID_PROPERTYID": "invalid propertyid", "INVALID_SIGN": "invalid sign", "TRANS_FAILURE": "trans failure"}

var langMap = map[string]string{"cn": "zh_CN"}

func doRequest(apiUrl string, data map[string]string) map[string]interface{} {
	cryptEd, sign := utils.AbDesSign(data, DesKey, Md5Key)
	post := map[string]string{}
	post["data"] = cryptEd
	post["sign"] = sign
	post["propertyId"] = PropertyId
	str, _ := utils.HttpPost(ApiUrl+apiUrl, post)
	m := make(map[string]interface{})
	json.Unmarshal(str, &m)
	return m
}

//注册
func (g *GameAllBet) Register(info map[string]interface{}) (interface{}, error) {
	return nil, nil
}

//登录
func (g *GameAllBet) Login(info map[string]interface{}) (interface{}, error) {
	//根据用户传递的字符串类型
	if info["game_code"] == "ALLBET_DZ" {
		return loginEGame(info)
	} else {
		return login(info)
	}
}

func create(info map[string]interface{}) (bool, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["agent"] = AgentName
	data["random"] = randomStr
	data["client"] = info["game_user_name"].(string)
	data["nickName"] = info["game_user_name"].(string)
	data["password"] = UserPassword
	data["vipHandicaps"] = "12"
	data["orHandicaps"] = "1"
	data["orHallRebate"] = "0"
	mapRes := doRequest("/check_or_create", data)
	if mapRes["error_code"] != "OK" {
		return false, errors.New(mapRes["error_code"].(string))
	} else {
		return true, nil
	}
}

func login(info map[string]interface{}) (interface{}, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["random"] = randomStr
	data["password"] = UserPassword
	data["client"] = info["game_user_name"].(string)
	data["language"] = langMap[info["lang"].(string)]
	data["returnUrl"] = LoginReturnUrl
	mapRes := doRequest("/forward_game", data)
	//fmt.Println(mapRes)
	if mapRes["error_code"] != "OK" {
		if mapRes["error_code"] == "CLIENT_NOT_EXIST" {
			//客户端不存在，走注册流程
			_, err := create(info)
			if err != nil {
				return nil, err
			}
			//登录
			return login(info)
		}
		return nil, errors.New(mapRes["error_code"].(string))
	}
	return mapRes["gameLoginUrl"], nil
}

//获取余额
func (g *GameAllBet) GetBalance(info map[string]interface{}) (interface{}, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["random"] = randomStr
	data["client"] = info["game_user_name"].(string)
	data["password"] = UserPassword
	mapRes := doRequest("/get_balance", data)
	if mapRes["error_code"] != "OK" {
		return nil, errors.New(mapRes["error_code"].(string))
	} else {
		return mapRes["balance"], nil
	}
}

//账户转游戏
func (g *GameAllBet) Account2GameTransfer(info map[string]interface{}) (interface{}, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["random"] = randomStr
	data["agent"] = AgentName
	//待确认传递订单方式
	//TODO
	data["sn"] = utils.AbOrderSn(PropertyId)
	data["client"] = info["game_user_name"].(string)
	data["operFlag"] = "1" //"0"(提取)或者"1"(存入)
	amountStr := strconv.FormatFloat(info["amount"].(float64), 'E', -1, 64)
	data["credit"] = amountStr
	mapRes := doRequest("/agent_client_transfer", data)
	if mapRes["error_code"] != "OK" {
		return false, errors.New(mapRes["error_code"].(string))
	}else{
		return true,nil
	}
}

//游戏转账户
func (g *GameAllBet) Game2AccountTransfer(info map[string]interface{}) (interface{}, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["random"] = randomStr
	data["agent"] = AgentName
	//待确认传递订单方式
	//TODO
	data["sn"] = utils.AbOrderSn(PropertyId)
	data["client"] = info["game_user_name"].(string)
	data["operFlag"] = "0" //"0"(提取)或者"1"(存入)
	amountStr := strconv.FormatFloat(info["amount"].(float64), 'E', -1, 64)
	data["credit"] = amountStr
	mapRes := doRequest("/agent_client_transfer", data)
	if mapRes["error_code"] != "OK" {
		return false, errors.New(mapRes["error_code"].(string))
	}else{
		return true,nil
	}
}

func loginEGame(info map[string]interface{}) (interface{}, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["random"] = randomStr
	data["client"] = info["game_user_name"].(string)
	data["password"] = UserPassword
	data["egameType"] = "af"
	data["gameType"] = "1100"
	mapRes := doRequest("/forward_egame", data)
	if mapRes["error_code"] != "OK" {
		if mapRes["error_code"] == "CLIENT_NOT_EXIST" {
			//客户端不存在，走注册流程
			_, err := create(info)
			if err != nil {
				return nil, err
			}
			//登录
			return loginEGame(info)
		}
		return nil, errors.New(mapRes["error_code"].(string))
	}
	return mapRes["gameLoginUrl"], nil
}

func (g *GameAllBet) QueryRecord(info map[string]interface{}) (interface{}, error) {
	EgameBetlogHistories(info)
	return nil, nil
}

func EgameBetlogHistories(map[string]interface{}) (interface{}, error) {
	randomStr := strconv.Itoa(rand.Intn(1000))
	data := map[string]string{}
	data["random"] = randomStr
	data["egameType"] = "af"
	data["startTime"] = "2018-10-29 16:30:00"
	data["endTime"] = "2018-10-29 17:30:00"
	data["pageIndex"] = "1"
	data["pageSize"] = "1000"
	_ = doRequest("/egame_betlog_histories", data)
	return nil, nil
}

func (g *GameAllBet) GetPrefix() string {
	return UserPrefix
}

func NewAllBetGame() game.Game {
	return &GameAllBet{
	}
}

func init() {
	game.Register("allbet", NewAllBetGame)
}
