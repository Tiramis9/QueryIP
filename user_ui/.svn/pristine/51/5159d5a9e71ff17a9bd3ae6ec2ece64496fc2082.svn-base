package mg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"game2/lib/game"
	"game2/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type GameMG struct {
	Host        string
	ApiUserName string
	ApiPassWord string
	ParentId    string
	ClientId    string
	Secret      string
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
	Meta         Meta   `json:"meta"`
	Error2       Error2 `json:"error"`
}

type Json struct {
	Meta   Meta        `json:"meta"`
	Data   interface{} `json:"data"`
	Error2 Error2      `json:"error"`
}

type Meta struct {
	Currency       string `json:"currency"`
	TimeZone       string `json:"time_zone"`
	TransactionId  string `json:"transaction_id"`
	ProcessingTime int    `json:"processing_time"`
}

type Error2 struct {
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	//Fields  interface{} `json:"fields"`
}

func NewMGGame() game.Game {
	return &GameMG{
		Host:        GameMGHost,
		ApiUserName: GameMGApiUserName,
		ApiPassWord: GameMGApiPassWord,
		ParentId:    GameMGParentId,
		ClientId:    GameMGClientId,
		Secret:      GameMGClientSecret,
	}
}

func init() {
	game.Register("mg", NewMGGame)
}

var langMap = map[string]string{"cn": "zh"}

//资金类别
var monneyMap = map[string]string{"en": "USD", "cn": "CNY", "zh": "CNY"}

//timezone类别
var timeZoneMap = map[string]string{
	"-12": "UTC-12",
	"-11": "UTC-11",
	"-10": "UTC-10",
	"-9":  "UTC-9",
	"-8":  "UTC-8",
	"-7":  "UTC-7",
	"-6":  "UTC-6",
	"-5":  "UTC-5",
	"-4":  "UTC-4",
	"-3":  "UTC-3",
	"-2":  "UTC-2",
	"-1":  "UTC-1",
	"0":   "UTC",
	"1":   "UTC+1",
	"2":   "UTC+2",
	"3":   "UTC+3",
	"4":   "UTC+4",
	"5":   "UTC+5",
	"6":   "UTC+6",
	"7":   "UTC+7",
	"8":   "UTC+8",
	"9":   "UTC+9",
	"10":  "UTC+10",
	"11":  "UTC+11",
	"12":  "UTC+12",
	"13":  "UTC+13",
	"14":  "UTC+14",
}

func DoLogin(m map[string]interface{}) (interface{}, error) {
	sub_url := "/oauth/token"
	http_method := "POST"
	contentType := "query_string"
	timeZone := timeZoneMap[m["time_zone"].(string)]
	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	txId := GameMGParentId

	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(GameMGAuth + ":" + GameMGAuthSecret))
	auth := "Basic " + paramStrEncode
	//auth := paramStrEncode

	req := make(map[string]string)
	req["grant_type"] = "password"      //代理名
	req["username"] = GameMGApiUserName //应用账号
	req["password"] = GameMGApiPassWord //应用账号密码

	res, err := httpHandle(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	resMap := LoginResp{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}

	return resMap.AccessToken, err
}

//刷新token
func DoRefreshToken(m map[string]interface{}) (interface{}, error) {
	sub_url := "/oauth/token"

	http_method := "POST"
	contentType := "query_string"
	timeZone := timeZoneMap[m["time_zone"].(string)]
	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	txId := GameMGParentId

	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(GameMGAuth + ":" + GameMGAuthSecret))
	auth := "Basic " + paramStrEncode
	//auth := paramStrEncode

	req := make(map[string]string)
	req["grant_type"] = "refresh_token"                //代理名
	req["refresh_token"] = m["refresh_token"].(string) //应用账号密码

	res, err := httpHandle(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	resMap := LoginResp{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}

	return resMap.AccessToken, err
}

//会员注册
func (mg *GameMG) Register(m map[string]interface{}) (interface{}, error) {
	sub_url := "/v1/account/member"

	http_method := "POST"
	contentType := "json"
	access_token := m["access_token"].(string)
	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	timeZone := timeZoneMap[m["time_zone"].(string)]
	txId := GameMGParentId

	auth := "Bearer " + access_token

	req := make(map[string]string)
	req["parent_id"] = GameMGParentId                      //代理名
	req["username"] = UserPrefix + m["user_name"].(string) //注册用户名
	req["password"] = GameMGPwd                            //注册密码

	res, err := httpHandle(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	resMap := Json{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}
	accountId := resMap.Data.(map[string]interface{})["id"] //注册的用户游戏id
	accountId = int(accountId.(float64))                    //将id转换格式

	return accountId, nil

}

//登录
func (mg *GameMG) Login(m map[string]interface{}) (interface{}, error) {
	//先获取access_token
	accessToken, err := DoLogin(m)
	if err != nil {
		return nil, err
	}

	//查询会员是否有account_id,无则注册
	userId := m["user_id"].(int)
	gameCode := m["game_code"].(string)
	flag := false

	var accountId interface{}
	var errs error
	account, err := model.GetAccountByGameName(model.Db, userId, gameCode)
	if err != nil {
		return nil, err
	}
	if account.MgId != 0 { //用户注册了游戏
		accountId = account.MgId
		flag = true
	}
	//账号不存在,走注册
	if flag == false {

		//不报错则注册成功
		m["access_token"] = accessToken
		accountId, errs = mg.Register(m)

		if errs != nil {
			return nil, errs
		}
		//更新游戏用户表的 mg_id

		field := make(map[string]interface{})
		field["mg_id"] = accountId
		_, err = model.UpdateAccount(model.Db, userId, gameCode, field)
		if err != nil {
			return nil, err
		}

	}
	//登录
	sub_url := "/v1/launcher/item"

	http_method := "POST"
	contentType := "json"
	//access_token := m["access_token"].(string)
	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	timeZone := timeZoneMap[m["time_zone"].(string)]
	txId := GameMGParentId

	auth := "Bearer " + accessToken.(string)

	req := make(map[string]interface{})
	reqMap := make(map[string]interface{})

	reqMap["lang"] = language //会员游戏id

	req["account_id"] = accountId.(int)      //会员游戏id
	req["item_id"] = m["game_type"].(string) //游戏id
	req["app_id"] = m["app_id"].(string)     //游戏id
	req["login_context"] = reqMap            //游戏id

	res, err := httpHandle2(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	resMap := Json{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}
	return resMap.Data, nil
}

//查询余额
func (mg *GameMG) GetBalance(m map[string]interface{}) (interface{}, error) {
	//先获取access_token
	accessToken, err := DoLogin(m)
	if err != nil {
		return nil, err
	}
	//查询用户的 mg 游戏id(account_id)
	userId := m["user_id"].(int)
	gameCode := m["game_code"].(string)

	account, err := model.GetAccountByGameName(model.Db, userId, gameCode)
	if err != nil {
		return nil, err
	}
	if account.MgId == 0 {
		return nil, errors.New(ErrAccountNotExist)
	}

	accountId := strconv.Itoa((account.MgId))

	sub_url := "/v1/wallet?account_id=" + accountId
	http_method := "GET"
	contentType := "nil"

	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	timeZone := timeZoneMap[m["time_zone"].(string)]
	txId := GameMGParentId

	auth := "Bearer " + accessToken.(string)

	req := make(map[string]string)

	res, err := httpHandle(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	resMap := Json{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}

	return resMap.Data.([]interface{})[0].(map[string]interface{})["credit_balance"], nil
}

//中心金额转游戏金额
func (mg *GameMG) Account2GameTransfer(m map[string]interface{}) (interface{}, error) {
	//先获取access_token
	accessToken, err := DoLogin(m)
	if err != nil {
		return nil, err
	}

	//查询用户的 mg 游戏id(account_id)
	userId := m["user_id"].(int)
	gameCode := m["game_code"].(string)

	account, err := model.GetAccountByGameName(model.Db, userId, gameCode)
	if err != nil {
		return nil, err
	}
	if account.MgId == 0 {
		return nil, errors.New(ErrAccountNotExist)
	}

	accountId := strconv.Itoa((account.MgId))

	sub_url := "/v1/transaction"

	http_method := "POST"
	contentType := "json"

	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	timeZone := timeZoneMap[m["time_zone"].(string)]
	txId := GameMGParentId

	auth := "Bearer " + accessToken.(string)

	var req []interface{}
	reqMap := make(map[string]interface{})

	reqMap["account_id"] = accountId          //会员游戏账户id
	reqMap["type"] = "CREDIT"                 //中心金额转游戏金额
	reqMap["balance_type"] = "CREDIT_BALANCE" //中心金额转游戏金额
	reqMap["category"] = "TRANSFER"           //金额
	reqMap["amount"] = m["amount"].(string)   //金额
	req = append(req, reqMap)

	res, err := httpHandle2(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	resMap := Json{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}

	return true, nil
}

//游戏金额转中心金额
func (mg *GameMG) Game2AccountTransfer(m map[string]interface{}) (interface{}, error) {
	//先获取access_token
	accessToken, err := DoLogin(m)
	if err != nil {
		return nil, err
	}

	//查询用户的 mg 游戏id(account_id)
	userId := m["user_id"].(int)
	gameCode := m["game_code"].(string)

	account, err := model.GetAccountByGameName(model.Db, userId, gameCode)
	if err != nil {
		return nil, err
	}
	if account.MgId == 0 {
		return nil, errors.New(ErrAccountNotExist)
	}

	accountId := strconv.Itoa((account.MgId))

	sub_url := "/v1/transaction"

	http_method := "POST"
	contentType := "json"

	currency := monneyMap[langMap[m["lang"].(string)]]
	language := langMap[m["lang"].(string)]
	timeZone := timeZoneMap[m["time_zone"].(string)]
	txId := GameMGParentId

	auth := "Bearer " + accessToken.(string)

	var req []interface{}
	reqMap := make(map[string]interface{})
	reqMap["account_id"] = accountId          //会员游戏账户id
	reqMap["type"] = "DEBIT"                  //游戏金额转中心金额
	reqMap["balance_type"] = "CREDIT_BALANCE" //游戏金额转中心金额
	reqMap["category"] = "TRANSFER"           //转账
	reqMap["amount"] = m["amount"].(string)   //金额
	req = append(req, reqMap)

	res, err := httpHandle2(sub_url, http_method, auth, contentType, timeZone, currency, txId, language, req)

	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	resMap := Json{}
	json.Unmarshal(res, &resMap)

	if resMap.Error2 != (Error2{}) {
		return nil, errors.New(resMap.Error2.Message)
	}

	return true, nil
}

//查询记录
func (mg *GameMG) QueryRecord(m map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func httpHandle(reqUrl, method, auth, contentType, timeZone, currency, txId, language string, params map[string]string) ([]byte, error) {

	client := &http.Client{}
	var req *http.Request
	var err error

	apiUrl := GameMGHost + reqUrl
	//dataStr := params

	if contentType == "json" {

		dataStr, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		//fmt.Println(apiUrl)
		//fmt.Println(params)
		req, err = http.NewRequest(method, apiUrl, bytes.NewBuffer(dataStr))
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	} else if contentType == "query_string" {

		dataStr := http_build_query(params)
		//fmt.Println(dataStr)
		//fmt.Println(apiUrl)

		req, err = http.NewRequest(method, apiUrl+"?"+dataStr, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	} else {
		req, err = http.NewRequest(method, apiUrl, nil)
	}
	//设置header
	req.Header.Add("Authorization", auth)
	req.Header.Add("X-DAS-TZ", timeZone)
	req.Header.Add("X-DAS-CURRENCY", currency)
	req.Header.Add("X-DAS-TX-ID", txId)
	req.Header.Add("X-DAS-LANG", language)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

//转账用
func httpHandle2(reqUrl, method, auth, contentType, timeZone, currency, txId, language string, params interface{}) ([]byte, error) {

	client := &http.Client{}
	var req *http.Request
	var err error

	apiUrl := GameMGHost + reqUrl
	//dataStr := params

	dataStr, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	//fmt.Println(apiUrl)
	//fmt.Println(params)
	req, err = http.NewRequest(method, apiUrl, bytes.NewBuffer(dataStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	//设置header
	req.Header.Add("Authorization", auth)
	req.Header.Add("X-DAS-TZ", timeZone)
	req.Header.Add("X-DAS-CURRENCY", currency)
	req.Header.Add("X-DAS-TX-ID", txId)
	req.Header.Add("X-DAS-LANG", language)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (mg *GameMG) GetPrefix() string {
	return UserPrefix
}

//类似php 的http_build_query
func http_build_query(data map[string]string) string {
	query := url.Values{}
	for k, v := range data {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	return query.Encode()
}
