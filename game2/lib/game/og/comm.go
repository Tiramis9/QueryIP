package og

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// OG
type (
	GameOG struct {
		//game.Server
		// Token  string `json:"token"`
		ServerHost string
		ReCordHost string
		Key        string
		Operator   string
		//////////
		ProviderId int
		Status     string                 `json:"status"`
		Data       map[string]interface{} `json:"data"`
		User       string
		GameCode   string
		/*
			Agent  string

			DesKey string
			Md5Key string
		*/
	}
	RespToken struct {
		Status string `json:"status"`
		Data   struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	/*
			200:
			[
		{
		"gameprovider":"og",
		"membername":"tzjk_myuser123",
		"gamename":"Baccarat",
		"bettingcode":"10882760273",
		"bettingdate":"\/Date(1543488845000)\/",
		"gameid":"C2",
		"roundno":"20-46",
		"result":null,
		"bet":"102",
		"winloseresult":"3",
		"bettingamount":20.000,
		"validbet":0.000,
		"winloseamount":0.000,
		"balance":7891.000,
		"currency":"RMB",
		"handicap":null,
		"status":"102^20.0^0.0^,",
		"gamecategory":"live",
		"settledate":null,
		"remark":null
		}
			]
	*/
	ResPutRecord struct {
		GameProvider  string  `json:"gameprovider"`
		MemberName    string  `json:"membername"`
		GameName      string  `json:"gamename"`
		BettingCode   string  `json:"bettingcode"`
		BettingDate   string  `json:"bettingdate"`
		GameId        string  `json:"gameid"`
		Roundno       string  `json:"roundno"`
		Result        string  `json:"result"`
		Bet           string  `json:"bet"`
		WinloseResult string  `json:"winloseresult"`
		BettinGamount float64 `json:"bettingamount"`
		ValidBet      float64 `json:"validbet"`
		Winloseamount float64 `json:"winloseamount"`
		Balance       float64 `json:"balance"`
		Currency      string  `json:"currency"`
		Handicap      string  `json:"handicap"`
		Status        string  `json:"status"`
		Gamecategory  string  `json:"gamecategory"`
		Settledate    string  `json:"settledate"`
		Remark        string  `json:"remark"`
	}

	ResPayRecord struct {
		Provider     string `json:"provider"`
		Id           string `json:"id"`
		UserName     string `json:"username"`
		Amount       string `json:"amount"`
		Currency     string `json:"currency"`
		Actions      string `json:"actions"`
		BaLance      string `json:"balance"`
		Transfercode string `json:"transfercode"`
		Createtime   string `json:"createtime"`
		Message      string `json:"Message"`
		Status       string `json:"status"`
		State        string `json:"State"`
	}
)

const (
	ERR_SYSTEM_ERROR       = "system error"
	ERR_NET_ERROR          = "network error"
	ERR_AUTH_FAILED        = "auth failed"
	ERR_SIGN_ERROR         = "sign error"
	ERR_APPID_ERROR        = "appid error"
	ERR_SECRET_ERROR       = "secret error"
	ERR_LACK_PARAMETERS    = "Lack of necessary parameters"
	ERR_BALANCE_NOT_ENOUGH = "balance is not enough"
	ERR_ILLEGAL_IP         = "illegal ip"
	KC_RAND_KIND_NUM       = 0    // 纯数字
	KC_RAND_KIND_LOWER     = 1    // 小写字母
	KC_RAND_KIND_UPPER     = 2    // 大写字母
	KC_RAND_KIND_ALL       = 3    // 数字、大小写字母
	DEFAULT_PAGE           = "1"  //默认第几页
	DEFAULT_PAGECOUNT      = "10" //默认一页多少数量
	LOGIN_EXPIRED_TIME     = 3600 //token过期时间
	//redis key值
	INFO_USER  = "info_user"
	INFO_MERCH = "info_merch"
	GAME_AG    = "1"
	GAME_BBIN  = "2"
	GAME_SB    = "3"
	GAME_AB    = "4"
)

// 代理本地服务器
func NewHttpClient() (*http.Client, error) {
	proxy, err := url.Parse(PROXYADDR)
	if err != nil {
		return nil, err
	}
	netTransport := &http.Transport{

		Proxy: http.ProxyURL(proxy),
	}
	return &http.Client{
		Transport: netTransport,
	}, nil
}

// 不走环回地址
func HttpPostQUERY(urlApi string, data interface{}) ([]byte, error) {
	client := &http.Client{}
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", urlApi, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//Post 请求方法查询表单
func HttpPOSTInquire(reqUrl string, msg map[string]interface{}) (body []byte, err error) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}
	query := url.Values{}
	for key, value := range msg {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	transport := &http.Transport{Proxy: proxy}
	c := &http.Client{Transport: transport}
	resp, err := c.PostForm(reqUrl, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 获取Token
func HttpGET(client *http.Client, url string) (body []byte, err error) {
	resp, err := http.NewRequest("GET", url, nil)
	resp.Header.Add("X-Operator", OPERATOR_NAME)
	resp.Header.Add("X-Key", OPERATOR_KEY)
	response, err := client.Do(resp)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK || err != nil {
		err = fmt.Errorf("HTTP GET Code=%v, URI=%v, err=%v", response.StatusCode, url, err)
		return
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

// 注册 update balance
func HttpPOST(client *http.Client, Url string, msg map[string]interface{}) (body []byte, err error) {
	data := make(map[string]interface{})
	for key, value := range msg {
		if key != "X-Token" {
			data[key] = fmt.Sprintf("%v", value)
		}
	}
	jsonstr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer([]byte(jsonstr)))
	req.Header.Add("X-Token", msg["X-Token"].(string))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	response, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("HTTP GET Code=%v, URI=%v, err=%v", response.StatusCode, Url, err)
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

// 获取游戏Url
func requestGetPlay(client *http.Client, url string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK || err != nil {
		err = fmt.Errorf("HTTP GET Code=%v, URI=%v, err=%v", response.StatusCode, url, err)
		return
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

// 获取key  获取余额
func HttpGETGlobal(client *http.Client, url string, msg map[string]interface{}) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Token", msg["X-Token"].(string))
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK || err != nil {
		err = fmt.Errorf("HTTP GET Code=%v, URI=%v, err=%v", response.StatusCode, url, err)
		return
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
