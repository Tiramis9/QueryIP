package sb

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"game2/lib/game"
	"game2/lib/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type GameSB struct {
	Host     string
	Agent    string
	UserKey  string
	Password string
}

const (
	UserPrefix = "jkgsb"
)

type (
	//注册返回xml
	RegisterRec struct {
		Result string `xml:",chardata"`
	}

	//查询余额返回xml
	GetBalanceRec struct {
		Result string `xml:",chardata"`
	}

	//统一返回xml为 result
	CommonRec struct {
		Result string `xml:",chardata"`
	}
)

var langMap = map[string]string{"cn": "zh"}

//资金类别
var monneyMap = map[string]string{"en": "USD", "cn": "RMB", "zh": "RMB"}

func NewSBGame() game.Game {
	return &GameSB{
		Host:     GameSBHost,
		Agent:    GameSBAgent,
		UserKey:  GameSBUserKey,
		Password: GameSBPwd,
	}
}

func init() {
	game.Register("sb", NewSBGame)
}

// 注册新会员
func (s *GameSB) Register(m map[string]interface{}) (interface{}, error) {

	req := make(map[string]string)
	req["agent"] = s.Agent                                    //代理名
	req["username"] = m["game_user_name"].(string)            //账号
	req["moneysort"] = monneyMap[langMap[m["lang"].(string)]] //货币 美元(USD), 人民币(RMB),马来币Dollar(MYR),韩币(KRW), 新加坡币(SGD),香港币(HKD)
	req["password"] = s.Password
	req["method"] = "caca"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + s.UserKey)
	//拼接访问url
	url := s.Host + "?params=" + paramStrEncode + "&key=" + key
	data, err := utils.HttpGet(url)

	if err != nil {
		return nil, err
	}

	res := RegisterRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	switch res.Result {
	case "key_error":
		return nil, errors.New(ErrKeyWrong)
	case "0":
		return nil, errors.New(ERRFail)
	case "2":
		return nil, errors.New(ErrPassWrong)
	case "3":
		return nil, errors.New(ErrUserNameTooLong)
	case "10":
		return nil, errors.New(ErrAgentNotExist)
	case "1":
		return true, nil
	default:
		return nil, errors.New(ErrUnknow)
	}
}

//会员登入
func (s *GameSB) Login(m map[string]interface{}) (interface{}, error) {
	reqReg := make(map[string]string)
	reqReg["agent"] = s.Agent                                    //代理名
	reqReg["username"] = m["game_user_name"].(string)            //账号
	reqReg["moneysort"] = monneyMap[langMap[m["lang"].(string)]] //货币 美元(USD), 人民币(RMB),马来币Dollar(MYR),韩币(KRW), 新加坡币(SGD),香港币(HKD)
	reqReg["password"] = s.Password
	reqReg["method"] = "caca"

	//将数据拼接
	paramStrReg := httpBuildQuery(reqReg)
	//base64加密
	paramStrEncodeReg := base64.StdEncoding.EncodeToString([]byte(paramStrReg))
	//md5加密得到key
	keyReg := utils.Md5V(paramStrEncodeReg + s.UserKey)
	//拼接访问url
	urlReg := s.Host + "?params=" + paramStrEncodeReg + "&key=" + keyReg
	data, err := utils.HttpGet(urlReg)

	if err != nil {
		return nil, err
	}

	res := RegisterRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	switch res.Result {
	case "key_error":
		return nil, errors.New(ErrKeyWrong)
	case "0":
		return nil, errors.New(ERRFail)
	case "2":
		return nil, errors.New(ErrPassWrong)
	case "3":
		return nil, errors.New(ErrUserNameTooLong)
	case "10":
		return nil, errors.New(ErrAgentNotExist)
	case "1":
		gameType := "2"
		paltFormName := "IBC"
		if m["game_code"] == "SB_TY" {
			gameType = "2"
			paltFormName = "IBC"
		}
		req := make(map[string]string)
		req["agent"] = s.Agent                         //代理名
		req["username"] = m["game_user_name"].(string) //账号
		req["password"] = s.Password
		//req["domain"] = m["domain"].(string)//可选，退出游戏时的域名the value domain=www.Aviabet.net
		req["gametype"] = gameType //游戏类别 value: 1.视讯 2.体育 3.彩票 4.电子游戏 5.捕鱼游戏
		//TODO //暂时设置A
		req["oddtype"] = "A"                      //Ag盘口: A (20~50000) B (50~5000) C (20~10000) D (200~20000) E (300~30000) F (400~40000) G (500~50000) H (1000~100000) I (2000~200000)
		req["gamekind "] = "0"                    //默认为0 gametype=5 (捕鱼游戏)
		req["iframe"] = "0"                       //默认为 0,如果采用 iframe框架请设置为 1, https 非框架 iframe=2,框架 iframe=3
		req["platformname"] = paltFormName        //平台名称:bbin,ag,IBC, SBTech
		req["lang"] = langMap[m["lang"].(string)] //zh中文,en英文,jp日文,kr韩文,id印尼
		req["method"] = "tg"

		//将数据拼接
		paramStr := httpBuildQuery(req)
		//base64加密
		paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
		//md5加密得到key
		key := utils.Md5V(paramStrEncode + s.UserKey)
		//拼接访问url
		url := s.Host + "?params=" + paramStrEncode + "&key=" + key
		return url, nil
	default:
		return nil, errors.New(ErrUnknow)
	}

	/*gameType := "2"
	paltFormName := "IBC"
	if m["game_code"] == "SB_TY" {
		gameType = "2"
		paltFormName = "IBC"
	}
	req := make(map[string]string)
	req["agent"] = s.Agent                                 //代理名
	req["username"] = m["game_user_name"].(string) //账号
	req["password"] = s.Password
	//req["domain"] = m["domain"].(string)//可选，退出游戏时的域名the value domain=www.Aviabet.net
	req["gametype"] = gameType //游戏类别 value: 1.视讯 2.体育 3.彩票 4.电子游戏 5.捕鱼游戏
	//TODO //暂时设置A
	req["oddtype"] = "A"               //Ag盘口: A (20~50000) B (50~5000) C (20~10000) D (200~20000) E (300~30000) F (400~40000) G (500~50000) H (1000~100000) I (2000~200000)
	req["gamekind "] = "0"             //默认为0 gametype=5 (捕鱼游戏)
	req["iframe"] = "0"                //默认为 0,如果采用 iframe框架请设置为 1, https 非框架 iframe=2,框架 iframe=3
	req["platformname"] = paltFormName //平台名称:bbin,ag,IBC, SBTech
	req["lang"] = m["lang"].(string)   //zh中文,en英文,jp日文,kr韩文,id印尼
	req["method"] = "tg"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + s.UserKey)
	//拼接访问url
	url := s.Host + "?params=" + paramStrEncode + "&key=" + key
	return url, nil*/
}

//获取会员余额
func (s *GameSB) GetBalance(m map[string]interface{}) (interface{}, error) {

	paltFormName := "IBC"
	if m["game_code"] == "SB_TY" {
		paltFormName = "IBC"
	}
	req := make(map[string]string)
	req["agent"] = s.Agent                         //代理名
	req["username"] = m["game_user_name"].(string) //账号
	req["password"] = s.Password
	req["platformname"] = paltFormName //平台名称:bbin,ag,IBC, SBTech
	req["method"] = "gb"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + s.UserKey)
	//拼接访问url
	url := s.Host + "?params=" + paramStrEncode + "&key=" + key
	data, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}

	res := GetBalanceRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	if res.Result == "key_error" {
		return nil, errors.New(ErrKeyWrong)
	} else if res.Result == "Account_no_exist" {
		return nil, errors.New(ErrAccountNotExist)
	} else if res.Result == "10" {
		return nil, errors.New(ErrAgentNotExist)
	}
	balace, err := strconv.ParseFloat(res.Result, 64)
	if err != nil {
		return nil, err
	}
	return balace, nil
}

// 中心专户转到游戏账户
func (s *GameSB) Account2GameTransfer(m map[string]interface{}) (interface{}, error) {
	paltFormName := "IBC"
	if m["game_code"] == "SB_TY" {
		//gameType = "2"
		paltFormName = "IBC"
	}
	req := make(map[string]string)
	req["agent"] = s.Agent                         //代理名
	req["username"] = m["game_user_name"].(string) //账号
	req["password"] = s.Password
	req["billno"] = m["order_sn"].(string) //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	req["type"] = "IN"                     // “IN”, 表示存入 “OUT”.表示提出
	req["usertype"] = "0"                  // 1 正常;0 测试
	//req["usertype"] = m["usertype"].(string) // 1 正常;0 测试
	req["credit"] = m["amount"].(string) // 操作的额度，请使用整数，如:100
	req["platformname"] = paltFormName   //平台名称:bbin,ag,IBC, SBTech
	req["method"] = "ptc"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + s.UserKey)
	//拼接访问url
	url := s.Host + "?params=" + paramStrEncode + "&key=" + key
	data, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}

	res := CommonRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	switch res.Result {
	case "key_error":
		return nil, errors.New(ErrKeyWrong)
	case "account_no_exist":
		return nil, errors.New(ErrAccountNotExist)
	case "0":
		return nil, errors.New(ERRFail)
	case "1":
		return true, nil
	case "2":
		//不确定,比如.404 等等,需要用 ctc来确认是否成功
		checkParam := make(map[string]interface{})
		checkParam["username"] = req["username"]
		checkParam["billno"] = req["billno"]
		checkRes, err := ConfirmTransferCredit(checkParam)
		if err != nil {
			return nil, err
		}
		return checkRes, nil
	case "10":
		return nil, errors.New(ErrAgentNotExist)
	default:
		return nil, errors.New(ErrUnknow)
	}
	return true, nil

}

//游戏账户转到中心中户
func (s *GameSB) Game2AccountTransfer(m map[string]interface{}) (interface{}, error) {
	paltFormName := "IBC"
	if m["game_code"] == "SB_TY" {
		//gameType = "2"
		paltFormName = "IBC"
	}
	req := make(map[string]string)
	req["agent"] = s.Agent                         //代理名
	req["username"] = m["game_user_name"].(string) //账号
	req["password"] = s.Password
	req["billno"] = m["order_sn"].(string) //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	req["type"] = "OUT"                    // “IN”, 表示存入 “OUT”.表示提出
	req["usertype"] = "0"                  // 1 正常;0 测试
	req["credit"] = m["amount"].(string)   // 操作的额度，请使用整数，如:100
	req["platformname"] = paltFormName     //平台名称:bbin,ag,IBC, SBTech
	req["method"] = "ptc"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + s.UserKey)
	//拼接访问url
	url := s.Host + "?params=" + paramStrEncode + "&key=" + key
	data, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}

	res := CommonRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	switch res.Result {
	case "key_error":
		return nil, errors.New(ErrKeyWrong)
	case "account_no_exist":
		return nil, errors.New(ErrAccountNotExist)
	case "0":
		return nil, errors.New(ERRFail)
	case "1":
		return true, nil
	case "2":
		//不确定,比如.404 等等,需要用 ctc来确认是否成功
		checkParam := make(map[string]interface{})
		checkParam["username"] = req["username"]
		checkParam["billno"] = req["billno"]
		checkRes, err := ConfirmTransferCredit(checkParam)
		if err != nil {
			return nil, err
		}
		return checkRes, nil
		//return nil,errors.New("密钥错误")
	case "10":
		return nil, errors.New(ErrAgentNotExist)
	default:
		return nil, errors.New(ErrUnknow)
	}
	return true, nil
}

//查询转帐
func (s *GameSB) ConfirmTransferCredit(m map[string]interface{}) (interface{}, error) {
	req := make(map[string]string)
	req["agent"] = s.Agent                         //代理名
	req["username"] = m["game_user_name"].(string) //账号
	req["password"] = s.Password
	req["billno"] = m["order_sn"].(string) //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	req["method"] = "ctc"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + s.UserKey)
	//拼接访问url
	url := s.Host + "?params=" + paramStrEncode + "&key=" + key
	data, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}

	res := CommonRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	switch res.Result {
	case "key_error":
		return nil, errors.New(ErrKeyWrong)
	case "account_no_exist":
		return nil, errors.New(ErrAccountNotExist)
	case "0":
		return nil, errors.New(ERRFail)
	case "1":
		//成功
		return true, nil
	case "10":
		return nil, errors.New(ErrAgentNotExist)
	default:
		return nil, errors.New(ErrUnknow)
	}

	//return nil,nil
}

func ConfirmTransferCredit(m map[string]interface{}) (interface{}, error) {
	req := make(map[string]string)
	req["agent"] = GameSBAgent                     //代理名
	req["username"] = m["game_user_name"].(string) //账号
	req["password"] = GameSBPwd
	req["billno"] = m["order_sn"].(string) //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	req["method"] = "ctc"

	//将数据拼接
	paramStr := httpBuildQuery(req)
	//base64加密
	paramStrEncode := base64.StdEncoding.EncodeToString([]byte(paramStr))
	//md5加密得到key
	key := utils.Md5V(paramStrEncode + GameSBUserKey)
	//拼接访问url
	url := GameSBHost + "?params=" + paramStrEncode + "&key=" + key
	data, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}

	res := CommonRec{}
	err = xml.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, err
	}

	switch res.Result {
	case "key_error":
		return nil, errors.New(ErrKeyWrong)
	case "account_no_exist":
		return nil, errors.New(ErrAccountNotExist)
	case "0":
		return nil, errors.New(ERRFail)
	case "1":
		//成功
		return true, nil
	case "10":
		return nil, errors.New(ErrAgentNotExist)
	default:
		return nil, errors.New(ErrUnknow)
	}

	//return nil,nil
}

func (s *GameSB) QueryRecord(m map[string]interface{}) (interface{}, error) {
	return nil, nil
}

//类似php 的http_build_query
func httpBuildQuery(data map[string]string) string {
	var str string
	for k, v := range data {
		str = str + k + "=" + v + "$"
	}
	str2 := str[0 : len(str)-1]
	return string(str2)
}

//get 走代理方法 请求方法
func httpGetProxy(reqUrl string) (string, error) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}
	transport := &http.Transport{Proxy: proxy}
	c := &http.Client{Transport: transport}
	/*	query := url.Values{}
		for key, value := range params {
			query.Add(key, value)
		}*/
	resp, err := c.Get(reqUrl)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(utils.ERR_NET_ERROR)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(utils.ERR_NET_ERROR)
	}
	return string(body), nil
}

func (s *GameSB) GetPrefix() string {
	return UserPrefix
}
