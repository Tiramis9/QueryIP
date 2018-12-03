package agin

import (
	"encoding/xml"
	"fmt"
	"game2/lib/encrypt"
	"game2/lib/game"
	"game2/lib/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

/*
注意事项：
1.游戏账号必须少于20个字符，只能是数字，字母，下划线
2.密码必须少于20个字符，不支持以下字符：' , “ , \ , / , > , < , & , # , -- , % , ? , $, 空格, 双节字符(全角字), TAB键, NULL, 換行符(\N)
3.**BBIN平台的密码须为6~12码英文或数字且符合0~9及a~z字符
4.MG 平台的会员帐号(必须加上 cagent 前缀，如cagent=AAA_NMGE,会员账号必須加上前綴 AAA，賬號就是 AAAxxxx)
5.IPM 平台的会员帐号(不可超过 13 位元长度)
6.BBIN 平台在转入或转出额度时, 只能整数转入或转出,不能带小数
7.BBIN 平台的 billno，只可使用數字, 请使用 19 字符内, 例如:123456445676789098
8.MG 平台的 billno，只可使用數字, 请使用 8 字符内, 例如:12345678
9.OG 平台的 billno，序列为 13 位, 例如 1234567890953
*/

const (
	HostGI       = "https://gi.jkgsoft.com"
	HostGCI      = "https://gci.jkgsoft.com"
	CAgent       = "DJ9_AGIN"
	MD5Key       = "Rr6GhEHqjJ4s"
	DESKey       = "SzNwjutz"
	UserPassword = "wpf4x7e9o3c2"

	AcTypeTryPlay   = "0" //试玩账号
	AcTypeRealMoney = "1" //真钱账号

	DoBusinessUrl    = HostGI + "/doBusiness.do"
	DoForwardGameUrl = HostGCI + "/forwardGame.do"
)

var (
	Lang = map[string]string{
		"zh": "zh-cn",
		"cn": "zh-cn",
		"en": "en-us",
	}

	Currency = map[string]string{
		"zh": "CNY",
		"cn": "CNY",
		"en": "USD",
	}
)

type GameAGin struct {
	HostGI       string
	HostGCI      string
	Agent        string
	Md5Key       string
	DesKey       string
	UserPassword string
}

type BaseReq struct {
	Params string
	Key    string
}

func newGameAGin() *GameAGin {
	return &GameAGin{
		HostGI:       HostGI,
		HostGCI:      HostGCI,
		Agent:        CAgent,
		Md5Key:       MD5Key,
		DesKey:       DESKey,
		UserPassword: UserPassword,
	}
}

func generateUrlParam(params, key string) string {
	query := url.Values{}
	query.Add("params", params)
	query.Add("key", key)
	return query.Encode()
}

func (g *GameAGin) generateUrl(m map[string]string) string {
	var src string
	for k, v := range m {
		src += fmt.Sprintf(`%v=%v/\\\\/`, k, v)
	}
	if len(src) > 6 {
		src = src[:len(src)-6]
	}

	params := encrypt.DesEcbPkc5Encrypt(src, g.DesKey)
	key := encrypt.MD5(params + g.Md5Key)

	return generateUrlParam(params, key)
}

func (g *GameAGin) httpPost(postUrl string) ([]byte, error) {
	request, _ := http.NewRequest("POST", DoBusinessUrl+"?"+postUrl, nil)
	request.Header.Add("User-Agent", "WEB_LIB_GI_"+g.Agent)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func generateTransferSrcParams(agent, method, loginName, billNo, tp, credit, acType, password, cur string) string {
	return fmt.Sprintf(`cagent=%v/\\\\/method=%v/\\\\/loginname=%v/\\\\/billno=%v/\\\\/type=%v/\\\\/credit=%v/\\\\/actype=%v/\\\\/password=%v/\\\\/cur=%v`,
		agent, method, loginName, billNo, tp, credit, acType, password, cur)
}

func generateTransferConfirmSrcParams(agent, loginName, method, billNo, tp, credit, acType, flag, password, cur string) string {
	return fmt.Sprintf(`cagent=%v/\\\\/loginname=%v/\\\\/method=%v/\\\\/billno=%v/\\\\/type=%v/\\\\/credit=%v/\\\\/actype=%v/\\\\/flag=%v/\\\\/password=%v/\\\\/cur=%v`,
		agent, loginName, method, billNo, tp, credit, acType, flag, password, cur)
}

func generateQosSrcParams(agent, billNo, method, acType, cur string) string {
	return fmt.Sprintf(`cagent=%v/\\\\/billno=%v/\\\\/method=%v/\\\\/actype=%v/\\\\/cur=%v`,
		agent, billNo, method, acType, cur)
}

type (
	BaseResp struct {
		XMLName xml.Name `xml:"result"`
		Info    string   `xml:"info,attr"`
		Msg     string   `xml:"msg,attr"`
	}
)

// 跳转到游戏页面，进入游戏
// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// ac_type:			|string		|"0"-试玩账号，"1"-真钱账号
// game_user_name:	|string		|用户名
// lang				|string		|语言
// game_type		|string		|游戏类型【非必填】
func (g GameAGin) generateForwardGameUrl(m map[string]interface{}) (interface{}, error) {
	req := map[string]string{
		"cagent":    g.Agent,
		"loginname": m["game_user_name"].(string),
		"password":  g.UserPassword,
		"lang":      Lang[m["lang"].(string)],
		"cur":       Currency[m["lang"].(string)],
		"oddtype":   "A",                   //盘口，即可下注范围，默认A盘口，A盘口可下注范围：20~50000
		"actype":    m["ac_type"].(string), //0-试玩账号，1-真钱账号
		"sid":       g.Agent + string(utils.Krand(15, utils.KC_RAND_KIND_NUM)),
	}
	if v, ok := m["game_type"].(string); ok {
		req["gameType"] = v
	}

	return DoForwardGameUrl + "?" + g.generateUrl(req), nil
}

// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// ac_type:			|string		|"0"-试玩账号，"1"-真钱账号
// game_user_name:	|string		|用户名
// lang				|string		|语言,zh,cn,en
// game_type		|string		|游戏类型【非必填】
func (g *GameAGin) checkOrCreateGameAccount(m map[string]interface{}) (interface{}, error) {
	req := map[string]string{
		"cagent":    g.Agent,
		"loginname": m["game_user_name"].(string),
		"password":  g.UserPassword,
		"method":    "lg",                         //检测并创建游戏账号，login缩写
		"oddtype":   "A",                          //盘口，即可下注范围，默认A盘口，A盘口可下注范围：20~50000
		"cur":       Currency[m["lang"].(string)], //人民币
		"actype":    m["ac_type"].(string),        //0-试玩账号，1-真钱账号
	}
	body, err := g.httpPost(g.generateUrl(req))
	if err != nil {
		return nil, err
	}

	var resp BaseResp
	if err := xml.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if resp.Info != "0" {
		return nil, fmt.Errorf("check or create game account error, error inf: %v, error message: %v", resp.Info, resp.Msg)
	}

	return g.generateForwardGameUrl(m)
}

// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// ac_type:			|string		|"0"-试玩账号，"1"-真钱账号
// game_user_name:	|string		|用户名
// lang				|string		|语言,zh,cn,en
// game_type		|string		|游戏类型【非必填】
func (g *GameAGin) Register(m map[string]interface{}) (resp interface{}, err error) {
	return g.checkOrCreateGameAccount(m)
}

// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// ac_type:			|string		|"0"-试玩账号，"1"-真钱账号
// game_user_name:	|string		|用户名
// lang				|string		|语言,zh,cn,en
// game_type		|string		|游戏类型【非必填】
func (g *GameAGin) Login(m map[string]interface{}) (resp interface{}, err error) {
	return g.checkOrCreateGameAccount(m)
}

// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// ac_type:			|string		|"0"-试玩账号，"1"-真钱账号
// game_user_name:	|string		|用户名
// lang:			|string		|语言,zh,cn,en
func (g *GameAGin) GetBalance(m map[string]interface{}) (interface{}, error) {
	req := map[string]string{
		"cagent":    g.Agent,
		"loginname": m["game_user_name"].(string),
		"method":    "gb",                  //查询余额，GetBalance缩写
		"actype":    m["ac_type"].(string), //0-试玩账号，1-真钱账号
		"password":  g.UserPassword,
		"cur":       Currency[m["lang"].(string)],
	}
	body, err := g.httpPost(g.generateUrl(req))
	if err != nil {
		return nil, err
	}

	var resp BaseResp
	if err := xml.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	balance, err := strconv.ParseFloat(resp.Info, 64)
	if err != nil {
		return nil, fmt.Errorf("get balance error, error info: %v, error message: %v", resp.Info, resp.Msg)
	}

	return balance, nil
}

// 【试玩账号不接受转账】
// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// game_user_name	|string		|用户名
// order_sn:		|string		|流水号
// amount			|string		|金额
// lang				|string		|语言,zh,cn,en
func (g *GameAGin) Account2GameTransfer(m map[string]interface{}) (interface{}, error) {
	m["type"] = "IN" //IN: 从网站账号转款到游戏账号;OUT: 從遊戲账號转款到網站賬號
	//step1.预备转账
	err := g.prepareTransferCredit(m)
	if err != nil {
		return nil, err
	}

	//step2.转账确认
	return nil, g.transferCreditConfirm(m)
}

// 【试玩账号不接受转账】
// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// game_user_name	|string		|用户名
// order_sn:		|string		|流水号
// amount			|string		|金额
// lang				|string		|语言,zh,cn,en
func (g *GameAGin) Game2AccountTransfer(m map[string]interface{}) (interface{}, error) {
	m["type"] = "OUT" //IN: 从网站账号转款到游戏账号;OUT: 從遊戲账號转款到網站賬號
	//step1.预备转账
	err := g.prepareTransferCredit(m)
	if err != nil {
		return nil, err
	}

	//step2.转账确认
	return nil, g.transferCreditConfirm(m)
}

//billno=(cagent+序列),序列是唯一的13~16位数,例如:cagent=‘XXXXX’及序列=1234567890987,那么billno=XXXXX1234567890987
// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// game_user_name	|string		|用户名
// order_sn:		|string		|流水号
// type				|string		|IN: 从网站账号转款到游戏账号;OUT: 從遊戲账號转款到網站賬號
// amount			|string		|金额
// lang				|string		|语言,zh,cn,en
func (g *GameAGin) prepareTransferCredit(m map[string]interface{}) error {
	req := map[string]string{
		"cagent":    g.Agent,
		"loginname": m["game_user_name"].(string),
		"method":    "tc", //预备转账，PrepareTransferCredit
		"billno":    m["order_sn"].(string),
		"type":      m["type"].(string),   //IN: 从网站账号转款到游戏账号;OUT: 從遊戲账號转款到網站賬號
		"credit":    m["amount"].(string), //BBIN平台在转入或转出额度时,只能整数转入或转出,不能带小数
		"password":  g.UserPassword,
		"actype":    "1",
		"cur":       Currency[m["lang"].(string)],
	}
	body, err := g.httpPost(g.generateUrl(req))
	if err != nil {
		return err
	}

	var resp BaseResp
	if err := xml.Unmarshal(body, &resp); err != nil {
		return err
	}
	//重复转账，返回成功
	if resp.Info == "duplicate_transfer" {
		return nil
	}
	//余额不足
	if resp.Info == "not_enough_credit" {
		return game.ErrNoEnoughMoney
	}
	if resp.Info != "0" {
		return fmt.Errorf("prepare transfer credit error, error info: %v, error message: %v", resp.Info, resp.Msg)
	}

	return nil
}

//billno=(cagent+序列),序列是唯一的13~16位数,例如:cagent=‘XXXXX’及序列=1234567890987,那么billno=XXXXX1234567890987
// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// game_user_name	|string		|用户名
// order_sn:		|string		|流水号
// type				|string		|IN: 从网站账号转款到游戏账号;OUT: 從遊戲账號转款到網站賬號
// amount			|string		|金额
// lang				|string		|语言,zh,cn,en
func (g *GameAGin) transferCreditConfirm(m map[string]interface{}) error {
	req := map[string]string{
		"cagent":    g.Agent,
		"loginname": m["game_user_name"].(string),
		"method":    "tcc", //转账确认，TransferCreditConfirm
		"billno":    m["order_sn"].(string),
		"type":      m["type"].(string),   //IN: 从网站账号转款到游戏账号;OUT: 從遊戲账號转款到網站賬號
		"credit":    m["amount"].(string), //BBIN平台在转入或转出额度时,只能整数转入或转出,不能带小数
		"actype":    "1",
		"flag":      "1", //1-代表调用‘预备转账 PrepareTransferCredit’API成功；0-代表调用‘預備轉賬 PrepareTransferCredit’出错或 出现错误码
		"password":  g.UserPassword,
		"cur":       Currency[m["lang"].(string)],
	}
	body, err := g.httpPost(g.generateUrl(req))
	if err != nil {
		return err
	}

	var resp BaseResp
	if err := xml.Unmarshal(body, &resp); err != nil {
		return err
	}

	const (
		ConfirmTransferSuccess           = "0"
		ConfirmTransferNetworkError      = "network_error"
		ConfirmTransferDuplicateTransfer = "duplicate_transfer" //重复转账
	)
	if resp.Info == ConfirmTransferNetworkError {
		return g.queryOrderStatus(m)
	}
	if resp.Info != ConfirmTransferSuccess && resp.Info != ConfirmTransferDuplicateTransfer {
		return fmt.Errorf("confirm transfer error, error info: %v, error message: %v", resp.Info, resp.Msg)
	}
	return nil
}

// 参数列表：
// 参数名			|数据类型	|说明
// ---------------------------------------
// order_sn:		|string		|流水号
// lang				|string		|语言,zh,cn,en
func (g *GameAGin) queryOrderStatus(m map[string]interface{}) error {
	req := map[string]string{
		"cagent": g.Agent,
		"billno": m["order_sn"].(string),
		"method": "qos", //查询订单状态，QueryOrderStatus
		"actype": "1",
		"cur":    Currency[m["lang"].(string)],
	}
	body, err := g.httpPost(g.generateUrl(req))
	if err != nil {
		return err
	}

	var resp BaseResp
	if err := xml.Unmarshal(body, &resp); err != nil {
		return err
	}
	const ConfirmTransferSuccess = "0"
	if resp.Info != ConfirmTransferSuccess {
		return fmt.Errorf("query order status error, error info: %v, error message: %v", resp.Info, resp.Msg)
	}
	return nil
}

func (g *GameAGin) QueryRecord(req map[string]interface{}) (interface{}, error) {
	//todo
	return nil, nil
}

func (g *GameAGin) GetPrefix() string {
	return strings.Split(CAgent, "_")[0]
}

func NewAGinGame() game.Game {
	return newGameAGin()
}

func init() {
	game.Register("agin", NewAGinGame)
}
