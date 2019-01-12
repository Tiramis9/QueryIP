package dg

import (
	"encoding/json"
	"fmt"
	"game2/lib/encrypt"
	"game2/lib/game"
	"game2/lib/utils"
	"strconv"
)

const (
	GameDGPlatform  = "24科技"
	GameDGCoinType  = "CNY/MYR/THB/TWD/KRW2/VND2/IDR2"
	GameDGHostHttp  = "http://api.dg99web.com"
	GameDGHostHttps = "https://api.dg99web.com"
	GameDGAgent     = "DGTE0101AE"
	GameDGPwd       = "abc123"
	GameDGKey       = "965de88b1bb745fa969cb09946612d86"
	GameDGAppSufix  = "HJK"
	UserPassword    = "xrbqn21j006cm75q" //密码

	RegisterUrl             = "/user/signup/%v/"
	LoginUrl                = "/user/login/%v/"
	TryPlayLoginUrl         = "/user/free/%v/"
	UpdateUserInfoUrl       = "/user/update/%v"
	GetUserBalanceUrl       = "/user/getBalance/%v"
	UserTransferUrl         = "/account/transfer/%v"
	IsDepositSuccessUrl     = "/account/checkTransfer/%v"
	UpdateUserLimitGroupUrl = "/game/updateLimit/%v"
	GetGameRecordUrl        = "/game/getReport/%v"
	TagGameRecordUrl        = "/game/markReport/%v"
	GetAgentUserListInfoUrl = "/user/onlineReport/%v"

	CodeOperateOk            = 0   //操作成功
	CodeParamsError          = 1   //参数错误
	CodeTokenError           = 2   //token验证失败
	CodeOperateFail          = 98  //操作失败
	CodeAccountNotFound      = 102 //账号不存在
	CodeAccountHasRegistered = 116 //账号已占用
	CodeNoEnoughMoney        = 120 //余额不足
)

var DGMoneyMap = map[string]string{
	"en": "USD",
	"cn": "CNY",
	"tw": "TWD",
	"kr": "KRW",
	"my": "MMK",
	"th": "THB",
}

type GameDG struct {
	Host         string
	Agent        string
	Key          string
	UserPassword string
}

type (
	RegisterMember struct {
		Username     string  `json:"username"`
		Password     string  `json:"password"`
		CurrencyName string  `json:"currencyName"`
		WinLimit     float64 `json:"winLimit"`
	}
	RegisterReq struct {
		Token  string         `json:"token"`
		Random string         `json:"random"`
		Data   string         `json:"data"`
		Member RegisterMember `json:"member"`
	}
	RegisterResp struct {
		CodeId int    `json:"codeId"`
		Token  string `json:"token"`
		Random string `json:"random"`
		Data   string `json:"data"`
	}

	LoginMember struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	LoginReq struct {
		Token  string      `json:"token"`
		Random string      `json:"random"`
		Lang   string      `json:"lang"`
		Member LoginMember `json:"member"`
	}
	LoginResp struct {
		CodeId int      `json:"codeId"`
		Token  string   `json:"token"`
		Random string   `json:"random"`
		List   []string `json:"list"`
	}

	GetUserBalanceMember struct {
		Username string `json:"username"`
	}
	GetUserBalanceReq struct {
		Token  string               `json:"token"`
		Random string               `json:"random"`
		Member GetUserBalanceMember `json:"member"`
	}
	GetUserBalanceResp struct {
		CodeId int    `json:"codeId"`
		Token  string `json:"token"`
		Random string `json:"random"`
		Member struct {
			UserName string  `json:"username"`
			Balance  float64 `json:"balance"`
		} `json:"member"`
	}

	UserTransferMember struct {
		UserName string  `json:"username"`
		Amount   float64 `json:"amount"`
	}
	UserTransferReq struct {
		Token  string             `json:"token"`
		Random string             `json:"random"`
		Data   string             `json:"data"`
		Member UserTransferMember `json:"member"`
	}
	UserTransferResp struct {
		CodeId int    `json:"codeId"`
		Token  string `json:"token"`
		Random string `json:"random"`
		Data   string `json:"data"`
		Member struct {
			UserName string  `json:"username"`
			Balance  float64 `json:"balance"`
		} `json:"member"`
	}

	IsDepositSuccessReq struct {
		Token  string `json:"token"`
		Random string `json:"random"`
		Data   string `json:"data"`
	}
	IsDepositSuccessResp struct {
		CodeId int    `json:"codeId"`
		Token  string `json:"token"`
		Random string `json:"random"`
	}

	GetGameRecordReq struct {
		Token  string `json:"token"`
		Random string `json:"random"`
	}
	GameRecordInfo struct {
		Id           int64   `json:"id"`
		LobbyId      int     `json:"lobbyId"`
		TableId      int     `json:"tableId"`
		ShoeId       int64   `json:"shoeId"`
		PlayId       int64   `json:"playId"`
		GameType     int     `json:"GameType"`
		GameId       int     `json:"GameId"`
		MemberId     int64   `json:"memberId"`
		BetTime      string  `json:"betTime"`
		CalTime      string  `json:"calTime"`
		WinOrLoss    float64 `json:"winOrLoss"`
		WinOrLossz   float64 `json:"winOrLossz"`
		BetPoints    float64 `json:"betPoints"`
		BetPointsz   float64 `json:"betPointsz"`
		AvailableBet float64 `json:"availableBet"`
		UserName     string  `json:"userName"`
		Result       string  `json:"result"`
		BetDetail    string  `json:"betDetail"`
		BetDetailz   string  `json:"betDetailz"`
		Ip           string  `json:"ip"`
		Ext          string  `json:"ext"`
		IsRevocation int     `json:"isRevocation"`
		ParentBetId  int64   `json:"parentBetId"`
		CurrencyId   int     `json:"currencyId"`
		DeviceType   int     `json:"deviceType"`
		Pluginid     int     `json:"pluginid"`
	}
	GetGameRecordResp struct {
		CodeId int              `json:"codeId"`
		Token  string           `json:"token"`
		Random string           `json:"random"`
		List   []GameRecordInfo `json:"list"`
	}

	TagGameRecordReq struct {
		Token  string `json:"token"`
		Random string `json:"random"`
		List   []int  `json:"list"`
	}
	TagGameRecordResp struct {
		CodeId int    `json:"codeId"`
		Token  string `json:"token"`
		Random string `json:"random"`
	}
)

// 产生token
func (g *GameDG) generateToken(random string) string {
	return encrypt.MD5(g.Agent + g.Key + random)
}

// 注册新会员
// 参数名			|类型	|说明
// --------------------------------
// game_user_name	|string	|用户名
// lang				|string	|语言
func (g *GameDG) Register(m map[string]interface{}) (interface{}, error) {
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := RegisterReq{
		Token:  g.generateToken(random),
		Random: random,
		Data:   "A",
		Member: RegisterMember{
			Username:     m["game_user_name"].(string),
			Password:     g.UserPassword,
			CurrencyName: DGMoneyMap[m["lang"].(string)],
			WinLimit:     0, //单日赢取不限额
		},
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(RegisterUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}
	var resp RegisterResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.CodeId != CodeOperateOk && resp.CodeId != CodeAccountHasRegistered {
		return nil, fmt.Errorf("register error, error code: %v", resp.CodeId)
	}

	return nil, nil
}

// 会员登录
// 参数名			|类型	|说明
// --------------------------------
// game_user_name	|string	|用户名
// lang				|string	|语言
func (g *GameDG) Login(m map[string]interface{}) (interface{}, error) {
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := LoginReq{
		Token:  g.generateToken(random),
		Random: random,
		Lang:   m["lang"].(string),
		Member: LoginMember{
			UserName: m["game_user_name"].(string),
			Password: g.UserPassword,
		},
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(LoginUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp LoginResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	//如果未注册，先注册
	if resp.CodeId == CodeAccountNotFound {
		_, err = g.Register(m)
		if err != nil {
			return nil, err
		}

		return g.Login(m)
	}

	if resp.CodeId != CodeOperateOk {
		return nil, fmt.Errorf("login error, error code: %v", resp.CodeId)
	}

	pc := 0
	return resp.List[pc] + resp.Token, nil
}

// 获取会员余额
// 参数名			|类型	|说明
// --------------------------------
// game_user_name	|string	|用户名
func (g *GameDG) GetBalance(m map[string]interface{}) (interface{}, error) {
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := GetUserBalanceReq{
		Token:  g.generateToken(random),
		Random: random,
		Member: GetUserBalanceMember{
			Username: m["game_user_name"].(string),
		},
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(GetUserBalanceUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp GetUserBalanceResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	if resp.CodeId != CodeOperateOk {
		return nil, fmt.Errorf("get balance error, error code %v", resp.CodeId)
	}

	return resp.Member.Balance, nil
}

// 检查存取款操作是否成功
// 参数名	|类型	|说明
// --------------------------------
// order_sn	|string	|流水号
func (g *GameDG) isDepositSuccess(m map[string]interface{}) (interface{}, error) {
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := IsDepositSuccessReq{
		Token:  g.generateToken(random),
		Random: random,
		Data:   m["order_sn"].(string),
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(IsDepositSuccessUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}
	var resp IsDepositSuccessResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.CodeId != CodeOperateOk {
		return nil, fmt.Errorf("check deposit error, error code: %v", resp.CodeId)
	}

	return nil, nil
}

// 中心专户转到游戏账户
// 参数名			|类型	|说明
// --------------------------------
// order_sn			|string	|流水号
// game_user_name	|string	|用户名
// amount			|string	|金额
func (g *GameDG) Account2GameTransfer(m map[string]interface{}) (interface{}, error) {
	amount, _ := strconv.ParseFloat(m["amount"].(string), 64)
	if amount < 0 {
		return nil, fmt.Errorf("account to game transfer error: amount < 0")
	}
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := UserTransferReq{
		Token:  g.generateToken(random),
		Random: random,
		Data:   m["order_sn"].(string),
		Member: UserTransferMember{
			UserName: m["game_user_name"].(string),
			Amount:   amount,
		},
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(UserTransferUrl, g.Agent), req)
	if err != nil {
		if _, err := g.isDepositSuccess(m); err != nil {
			return nil, err
		}
		return nil, nil
	}
	var resp UserTransferResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.CodeId != CodeOperateOk {
		return nil, fmt.Errorf("account to game transfer error, error code: %v", resp.CodeId)
	}
	return nil, nil
}

// 游戏账户转到中心中户
// 参数名			|类型	|说明
// --------------------------------
// order_sn			|string	|流水号
// game_user_name	|string	|用户名
// amount			|string	|金额
func (g *GameDG) Game2AccountTransfer(m map[string]interface{}) (interface{}, error) {
	amount, _ := strconv.ParseFloat(m["amount"].(string), 64)
	if amount < 0 {
		return nil, fmt.Errorf("game to account transfer error: amount < 0")
	}
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := UserTransferReq{
		Token:  g.generateToken(random),
		Random: random,
		Data:   m["order_sn"].(string),
		Member: UserTransferMember{
			UserName: m["game_user_name"].(string),
			Amount:   -amount, //减余额，取负数
		},
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(UserTransferUrl, g.Agent), req)
	if err != nil {
		if _, err := g.isDepositSuccess(m); err != nil {
			return nil, err
		}
		return nil, nil
	}
	var resp UserTransferResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.CodeId == CodeNoEnoughMoney {
		return nil, game.ErrNoEnoughMoney
	}
	if resp.CodeId != CodeOperateOk {
		return nil, fmt.Errorf("game to account transfer error, error code: %v", resp.CodeId)
	}
	return nil, nil
}

// 抓取注单报表
func (g *GameDG) QueryRecord(m map[string]interface{}) (interface{}, error) {
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := GetGameRecordReq{
		Token:  g.generateToken(random),
		Random: random,
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(GetGameRecordUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp GetGameRecordResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.List, nil
}

// 标记已抓取注单报表
// 参数名	|类型	|说明
// --------------------------------
// list		|[]int	|已入库的注单id集合
func (g *GameDG) TagGameRecord(list []int) (interface{}, error) {
	random := string(utils.Krand(32, utils.KC_RAND_KIND_ALL))
	req := TagGameRecordReq{
		Token:  g.generateToken(random),
		Random: random,
		List:   list,
	}
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(TagGameRecordUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp TagGameRecordResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.CodeId != CodeOperateOk {
		return nil, fmt.Errorf("tag game record error, error code: %v", resp.CodeId)
	}
	return nil, nil
}

func (g *GameDG) GetPrefix() string {
	return ""
}

func NewDGGame() game.Game {
	return &GameDG{
		Host:         GameDGHostHttp,
		Agent:        GameDGAgent,
		Key:          GameDGKey,
		UserPassword: UserPassword,
	}
}

func init() {
	game.Register("dg", NewDGGame)
}
