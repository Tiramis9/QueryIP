package bbin

import (
	"encoding/json"
	"errors"
	"fmt"
	"game2/lib/game"
	"game2/lib/utils"
	"game2/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"time"
)

//用户名仅可输入英文字母以及数字的组合
const (
	Str8                         = "djfkdkfk"
	Str1                         = "d"
	Str9                         = "qewrtyqaz"
	Str4                         = "tyeu"
	Str7                         = "qwertyu"
	WebSite                      = "avia"
	UpperName                    = "djksoft"
	UserPrefix                   = "bbin"
	API888                       = "http://888.tcy789.com/app/WebService/JSON/display.php"
	APIlink                      = "http://linkapi.tcy789.com/app/WebService/JSON/display.php"
	CreateMemberKey              = "3QcgFxyY0"
	LoginKey                     = "fV98jAu"
	Login2Key                    = "fV98jAu"
	LogoutKey                    = "x2b7x"
	CheckUsrBalanceKey           = "7pxyd9c0a"
	TransferKey                  = "10WyHdOdZ"
	CheckTransferKey             = "5Jr57Ya8c7"
	TransferRecordKey            = "5Jr57Ya8c7"
	PlayGameKey                  = "05Rz1lv"
	PlayGameByH5Key              = "05Rz1lv"
	BetRecordKey                 = "6kqBB1"
	BetRecordByModifiedDate3Key  = "6kqBB1"
	GetJPHistoryKey              = "6kqBB1"
	ForwardGameH5By5Key          = "05Rz1lv"
	WagersRecordBy31Key          = "6kqBB1"
	WagersRecordBy1Key           = "6kqBB1"
	WagersRecordBy38Key          = "6kqBB1"
	ForwardGameH5By30Key         = "05Rz1lv"
	ForwardGameH5By38Key         = "05Rz1lv"
	GetFishEventHistoryKey       = "6kqBB1"
	FishEventUrlKey              = "05Rz1lv"
	GetSKEventHistoryKey         = "6kqBB1"
	SKEventUrlKey                = "05Rz1lv"
	GetWagersSubDetailUrlBy3Key  = "G0v11Mp59d"
	GetWagersSubDetailUrlBy5Key  = "G0v11Mp59d"
	GetWagersSubDetailUrlBy30Key = "G0v11Mp59d"
	GetWagersSubDetailUrlBy38Key = "G0v11Mp59d"
	GetSportEventHistoryKey      = "6kqBB1"
	SportEventUrlKey             = "05Rz1lv"
	GetCasinoEventHistoryKey     = "6kqBB1"
	CasinoEventUrlKey            = "05Rz1lv"
	GetLiveEventHistoryKey       = "6kqBB1"
	LiveEventUrlKey              = "05Rz1lv"
)

type GameBBin struct {
}

//GameType 3
//3001	百家乐
//3003	龙虎斗
//3005	三公
//3006	温州牌九
//3007	轮盘
//3008	骰宝
//3010	德州扑克
//3011	色碟
//3012	牛牛
//3014	无限21点
//3015	番摊
//3016	鱼虾蟹
//3017	保险百家乐
//3018	炸金花

//GameType 5
/* o代表支持 HTML5 FLASH
5005	惑星战记	o	o	老虎机
5007	激爆水果盘		o	老虎机
5008	猴子爬树		o	老虎机
5009	金刚爬楼		o	老虎机
5010	外星战记	o	o	老虎机
5012	外星争霸	o	o	老虎机
5013	传统	o	o	老虎机
5014	丛林	o	o	老虎机
5015	FIFA2010	o	o	老虎机
5016	史前丛林冒险		o	老虎机
5017	星际大战		o	老虎机
5018	齐天大圣		o	老虎机
5019	水果乐园	o	o	老虎机
5025	法海斗白蛇		o	老虎机
5026	2012 伦敦奥运		o	老虎机
5027	功夫龙	o		老虎机
5028	中秋月光派对		o	老虎机
5029	圣诞派对		o	老虎机
5030	幸运财神	o	o	老虎机
5034	王牌5PK		o	大型机台
5035	加勒比扑克		o	桌上游戏
5039	鱼虾蟹	o	o	桌上游戏
5040	百搭二王		o	大型机台
5041	7PK		o	大型机台
5043	钻石水果盘	o	o	老虎机
5044	明星97II	o	o	老虎机
5045	森林舞会	o		大型机台
5046	斗魂	o		大型机台
5054	爆骰	o		桌上游戏
5057	明星97	o	o	老虎机
5058	疯狂水果盘	o	o	老虎机
5060	动物奇观五		o	老虎机
5061	超级7	o	o	老虎机
5062	龙在囧途	o	o	老虎机
5063	水果拉霸	o	o	老虎机
5064	扑克拉霸	o	o	老虎机
5065	筒子拉霸	o	o	老虎机
5066	足球拉霸	o	o	老虎机
5067	大话西游	o		老虎机
5068	酷搜马戏团	o		老虎机
5069	水果擂台	o		老虎机
5070	黄金大转轮		o	大型机台
5073	百家乐大转轮		o	大型机台
5076	数字大转轮	o	o	大型机台
5077	水果大转轮	o	o	大型机台
5078	象棋大转轮		o	大型机台
5079	3D数字大转轮	o	o	大型机台
5080	乐透转轮	o	o	大型机台
5083	钻石列车	o	o	大型机台
5084	圣兽传说	o	o	大型机台
5088	斗大	o	o	桌上游戏
5089	红狗	o	o	桌上游戏
5090	金鸡报喜	o		老虎机
5091	三国拉霸		o	老虎机
5092	封神榜		o	老虎机
5093	金瓶梅	o		老虎机
5094	金瓶梅2	o	o	老虎机
5095	斗鸡	o	o	老虎机
5096	五行	o		老虎机
5097	海底世界	o		老虎机
5098	五福临门	o		老虎机
5099	金狗旺岁	o		老虎机
5100	七夕	o		老虎机
5105	欧式轮盘	o	o	桌上游戏
5106	三国	o	o	老虎机
5107	美式轮盘	o	o	桌上游戏
5108	彩金轮盘	o	o	桌上游戏
5109	法式轮盘	o	o	桌上游戏
5110	夜上海	o		特色游戏
5116	西班牙21点		o	桌上游戏
5117	维加斯21点		o	桌上游戏
5118	奖金21点		o	桌上游戏
5119	神秘岛	o		老虎机
5120	女娲补天	o		老虎机
5123	经典21点	o		桌上游戏
5129	跳高高	o		老虎机
5130	跳起来	o		老虎机
5131	皇家德州扑克		o	桌上游戏
5133	五福临门2	o		老虎机
5201	火焰山		o	老虎机
5202	月光宝盒		o	老虎机
5203	爱你一万年		o	老虎机
5204	2014 FIFA	o	o	老虎机
5402	夜市人生	o	o	老虎机
5404	沙滩排球	o	o	老虎机
5406	神舟27		o	老虎机
5407	大红帽与小野狼	o	o	老虎机
5601	秘境冒险	o	o	老虎机
5701	连连看		o	刮刮乐
5703	发达啰		o	刮刮乐
5704	斗牛		o	刮刮乐
5705	聚宝盆		o	刮刮乐
5706	浓情巧克力		o	刮刮乐
5707	金钱豹		o	刮刮乐
5801	海豚世界	o	o	老虎机
5802	阿基里斯		o	老虎机
5803	阿兹特克宝藏	o	o	老虎机
5804	大明星		o	老虎机
5805	凯萨帝国	o	o	老虎机
5806	奇幻花园		o	老虎机
5808	浪人武士		o	老虎机
5809	空战英豪		o	老虎机
5810	航海时代	o	o	老虎机
5811	狂欢夜		o	老虎机
5821	国际足球		o	老虎机
5823	发大财	o	o	老虎机
5824	恶龙传说	o	o	老虎机
5825	金莲		o	老虎机
5826	金矿工		o	老虎机
5827	老船长		o	老虎机
5828	霸王龙	o	o	老虎机
5832	高速卡车		o	老虎机
5833	沉默武士		o	老虎机
5835	喜福牛年	o	o	老虎机
5836	龙卷风		o	老虎机
5837	喜福猴年	o	o	老虎机
5839	经典高球	o		老虎机
5901	连环夺宝	o	o	特色游戏
5902	糖果派对	o	o	特色游戏
5903	秦皇秘宝	o		特色游戏
5904	蒸气炸弹	o		特色游戏
5907	趣味台球	o		特色游戏
5908	糖果派对2	o		特色游戏
5909	开心消消乐	o		特色游戏
5910	魔法元素	o		特色游戏
5912	连环夺宝2	o		特色游戏
*/

//map检查是否输入合法
//BB视讯
var VideoGameType = map[string]string{
	"3001": "百家乐",
	"3003": "	龙虎斗",
	"3005": "	三公",
	"3006": "温州牌九",
	"3007": "轮盘",
	"3008": "骰宝",
	"3010": "德州扑克",
	"3011": "色碟",
	"3012": "牛牛",
	"3014": "无限21点",
	"3015": "番摊",
	"3016": "鱼虾蟹",
	"3017": "保险百家乐",
	"3018": "炸金花",
}

//BB电子
var ElecGameType = map[string]string{
	"5005": "惑星战记",
	"5007": "激爆水果盘",
	"5008": "猴子爬树",
	"5009": "金刚爬楼",
	"5010": "外星战记",
	"5012": "外星争霸",
	"5013": "传统",
	"5014": "丛林",
	"5015": "FIFA2010",
	"5016": "史前丛林冒险",
	"5017": "星际大战",
	"5018": "齐天大圣",
	"5019": "水果乐园",
	"5025": "法海斗白蛇",
	"5026": "2012 伦敦奥运",
	"5027": "功夫龙",
	"5028": "中秋月光派对",
	"5029": "圣诞派对",
	"5030": "幸运财神",
	"5034": "王牌5PK",
	"5035": "加勒比扑克",
	"5039": "鱼虾蟹",
	"5040": "百搭二王",
	"5041": "7PK",
	"5043": "钻石水果盘",
	"5044": "明星97II",
	"5045": "森林舞会",
	"5046": "斗魂",
	"5054": "爆骰",
	"5057": "明星97",
	"5058": "疯狂水果盘",
	"5060": "动物奇观五",
	"5061": "超级7",
	"5062": "龙在囧途",
	"5063": "水果拉霸",
	"5064": "扑克拉霸",
	"5065": "筒子拉霸",
	"5066": "足球拉霸",
	"5067": "大话西游",
	"5068": "酷搜马戏团",
	"5069": "水果擂台",
	"5070": "黄金大转轮",
	"5073": "百家乐大转轮",
	"5076": "数字大转轮",
	"5077": "水果大转轮",
	"5078": "象棋大转轮",
	"5079": "3D数字大转轮",
	"5080": "乐透转轮",
	"5083": "钻石列车",
	"5084": "圣兽传说",
	"5088": "斗大",
	"5089": "红狗",
	"5090": "金鸡报喜",
	"5091": "三国拉霸",
	"5092": "封神榜",
	"5093": "金瓶梅",
	"5094": "金瓶梅2",
	"5095": "斗鸡",
	"5096": "五行",
	"5097": "海底世界",
	"5098": "五福临门",
	"5099": "金狗旺岁",
	"5100": "七夕",
	"5105": "欧式轮盘",
	"5106": "三国",
	"5107": "美式轮盘",
	"5108": "彩金轮盘",
	"5109": "法式轮盘",
	"5110": "夜上海",
	"5116": "西班牙21点",
	"5117": "维加斯21点",
	"5118": "奖金21点",
	"5119": "神秘岛",
	"5120": "女娲补天",
	"5123": "经典21点",
	"5129": "跳高高",
	"5130": "跳起来",
	"5131": "皇家德州扑克",
	"5133": "五福临门2",
	"5201": "火焰山",
	"5202": "月光宝盒",
	"5203": "爱你一万年",
	"5204": "2014 FIFA",
	"5402": "夜市人生",
	"5404": "沙滩排球",
	"5406": "神舟27",
	"5407": "大红帽与小野狼",
	"5601": "秘境冒险",
	"5701": "连连看",
	"5703": "发达啰",
	"5704": "斗牛",
	"5705": "聚宝盆",
	"5706": "浓情巧克力",
	"5707": "金钱豹",
	"5801": "海豚世界",
	"5802": "阿基里斯",
	"5803": "阿兹特克宝藏",
	"5804": "大明星",
	"5805": "凯萨帝国",
	"5806": "奇幻花园",
	"5808": "浪人武士",
	"5809": "空战英豪",
	"5810": "航海时代",
	"5811": "狂欢夜",
	"5821": "国际足球",
	"5823": "发大财",
	"5824": "恶龙传说",
	"5825": "金莲",
	"5826": "金矿工",
	"5827": "老船长",
	"5828": "霸王龙",
	"5832": "高速卡车",
	"5833": "沉默武士",
	"5835": "喜福牛年",
	"5836": "龙卷风",
	"5837": "喜福猴年",
	"5839": "经典高球",
	"5901": "连环夺宝",
	"5902": "糖果派对",
	"5903": "秦皇秘宝",
	"5904": "蒸气炸弹",
	"5907": "趣味台球",
	"5908": "糖果派对2",
	"5909": "开心消消乐",
	"5910": "魔法元素",
	"5912": "连环夺宝2",
}

var langMap = map[string]string{"cn": "zh-cn"}

func doRequestPost(apiUrl string, data map[string]interface{}) (map[string]interface{}, error) {
	body, err := utils.HttpPostProxy(apiUrl, data)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	return m, nil
}

type SysGame struct {
	Channel    string `json:"channel"`
	Type       int    `json:"type"`
	GameName   string `json:"game_name"`
	GameCode   string `json:"game_code"`
	Status     int    `json:"status"`
	Device     int    `json:"device"`
	Memo       string `json:"memo"`
	ParentId   int    `json:"parent_id"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func AddGame(c *gin.Context) {
	db := model.Db
	sortedKeys := make([]string, 0)
	for k, _ := range ElecGameType {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	timestamp := time.Now().Unix()
	for _, v := range sortedKeys {
		fmt.Println(ElecGameType[v])
		var sg SysGame
		sg.Channel = "BBIN"
		sg.Type = 4
		sg.GameName = ElecGameType[v]
		sg.GameCode = v
		sg.Status = 1
		sg.Device = 2
		sg.Memo = v
		sg.ParentId = 2
		sg.CreateTime = timestamp
		sg.UpdateTime = timestamp
		if err := db.Create(&sg); err != nil {
			fmt.Println(ElecGameType[v])
			fmt.Println(err)
		}
	}
}

func (g *GameBBin) Register(info map[string]interface{}) (interface{}, error) {
	return nil, nil
}

//登录
func (g *GameBBin) Login(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]interface{})
	dataMap := make(map[string]interface{})
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["uppername"] = UpperName
	strA := Str8 //無意義字串長度8碼
	strC := Str1 //無意義字串長度1碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + Login2Key + timeStr)
	data["key"] = strA + strB + strC
	mapRes, err := doRequestPost(API888+"/Login2", data)
	if err != nil {
		return false, err
	}
	dataMap = mapRes["data"].(map[string]interface{})
	if mapRes["result"].(bool) == true && dataMap["Code"].(float64) == 99999 {
		//成功登陆
		switch info["game_code"].(string) {
		case "BBIN_DZ":
			return VideoEventUrl(info)
		case "BBIN_CP":
			return SKEventUrl(info)
		case "BBIN_SX":
			return LiveEventUrl(info)
		case "BBIN_TY":
			return SportEventUrl(info)
		default:
			return SKEventUrl(info)
		}
	} else {
		logrus.Error(mapRes)
		return false, errors.New("net error")
	}
}

//登录方法2
/*func (g *GameBBin) Login(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	//dataMap :=make(map[string]interface{})
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["uppername"] = UpperName
	data["page_site"] = "Ltlottery"
	strA := Str8 //無意義字串長度8碼
	strC := Str1 //無意義字串長度1碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + LoginKey + timeStr)
	data["key"] = strA + strB + strC
	fmt.Println(API888+"/Login"+"?"+utils.Http_build_query(data))
	return nil,nil
	mapRes, err := doRequestPost(API888+"/Login", data)
	if err!=nil{
		return false,err
	}
	fmt.Println(mapRes)
	return nil,nil
}*/

//进入体育
func SportEventUrl(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["lang"] = langMap[info["lang"].(string)]
	strA := Str8 //無意義字串長度8碼
	strC := Str8 //無意義字串長度8碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + SportEventUrlKey + timeStr)
	data["key"] = strA + strB + strC
	urlQuery := utils.Http_build_query(data)
	return API888 + "/SportEventUrl?" + urlQuery, nil
	return nil, nil
}

//进入彩票页
func SKEventUrl(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["lang"] = langMap[info["lang"].(string)]
	strA := Str8 //無意義字串長度8碼
	strC := Str8 //無意義字串長度8碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + SKEventUrlKey + timeStr)
	data["key"] = strA + strB + strC
	urlQuery := utils.Http_build_query(data)
	return API888 + "/SKEventUrl?" + urlQuery, nil
	return nil, nil
}

//进入视讯页
func LiveEventUrl(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["lang"] = langMap[info["lang"].(string)]
	strA := Str8 //無意義字串長度8碼
	strC := Str8 //無意義字串長度8碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + LiveEventUrlKey + timeStr)
	data["key"] = strA + strB + strC
	urlQuery := utils.Http_build_query(data)
	return API888 + "/LiveEventUrl?" + urlQuery, nil
}

//进入电子游戏
func VideoEventUrl(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["gamekind"] = "5"
	data["gametype"] = info["game_type"].(string)
	strA := Str8 //無意義字串長度8碼
	strC := Str8 //無意義字串長度8碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + PlayGameKey + timeStr)
	data["key"] = strA + strB + strC
	urlQuery := utils.Http_build_query(data)
	return API888 + "/PlayGame?" + urlQuery, nil
}

//进入游戏flash
func PlayGame(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["gamekind"] = info["game_kind"].(string)
	data["gametype"] = info["game_type"].(string)
	data["gamecode"] = info["game_code"].(string)
	strA := Str8 //無意義字串長度8碼
	strC := Str8 //無意義字串長度8碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + PlayGameKey + timeStr)
	data["key"] = strA + strB + strC
	urlQuery := utils.Http_build_query(data)
	return API888 + "/PlayGame?" + urlQuery, nil
}

//进入游戏H5
//TODO
func PlayGameByH5(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]string)
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["gamekind"] = info["game_kind"].(string)
	data["gametype"] = info["game_type"].(string)
	data["gamecode"] = info["game_code"].(string)
	strA := Str8 //無意義字串長度8碼
	strC := Str8 //無意義字串長度8碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + PlayGameKey + timeStr)
	data["key"] = strA + strB + strC
	urlQuery := utils.Http_build_query(data)
	return API888 + "/PlayGame?" + urlQuery, nil
	return nil, nil
}

//游戏转至中心账户
func (g *GameBBin) Account2GameTransfer(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]interface{})
	dataMap := make(map[string]interface{})
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["uppername"] = UpperName
	data["remitno"] = info["order_sn"]
	data["action"] = "IN"
	data["remit"] = info["amount"].(float64) //正整数的字符串
	fmt.Println(data)
	strA := Str9 //無意義字串長度9碼
	strC := Str4 //無意義字串長度4碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + data["remitno"].(string) + TransferKey + timeStr)
	data["key"] = strA + strB + strC
	mapRes, err := doRequestPost(APIlink+"/Transfer", data)
	if err != nil {
		return nil, err
	}
	dataMap = mapRes["data"].(map[string]interface{})
	if mapRes["result"].(bool) == true && dataMap["Code"].(float64) == 11100 {
		return true, nil
	}
	return false, errors.New(dataMap["Message"].(string))
}

//游戏转至中心账户
func (g *GameBBin) Game2AccountTransfer(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]interface{})
	dataMap := make(map[string]interface{})
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["uppername"] = UpperName
	data["remitno"] = info["order_sn"]
	data["action"] = "OUT"
	data["remit"] = info["amount"].(float64) //正整数
	strA := Str9                             //無意義字串長度9碼
	strC := Str4                             //無意義字串長度4碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + data["remitno"].(string) + TransferKey + timeStr)
	data["key"] = strA + strB + strC
	mapRes, err := doRequestPost(APIlink+"/Transfer", data)
	if err != nil {
		return nil, err
	}
	dataMap = mapRes["data"].(map[string]interface{})
	if mapRes["result"].(bool) == true && dataMap["Code"].(float64) == 11100 {
		return true, nil
	}
	return false, errors.New(dataMap["Message"].(string))
}

//获取余额
func (g *GameBBin) GetBalance(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]interface{})
	data["website"] = WebSite
	userName := info["game_user_name"].(string)
	data["username"] = userName
	data["uppername"] = UpperName
	strA := Str4 //無意義字串長度4碼
	strC := Str7 //無意義字串長度7碼
	loc, err := time.LoadLocation("America/Caracas")
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().In(loc).Format("20060102")
	strB := utils.Md5V(WebSite + userName + CheckUsrBalanceKey + timeStr)
	data["key"] = strA + strB + strC
	mapRes, err := doRequestPost(APIlink+"/CheckUsrBalance", data)
	if err != nil {
		return nil, err
	}
	if mapRes["result"].(bool) == false {
		return 0,  errors.New("net error")
	}
	fmt.Println(1)
	dataMap, ok := mapRes["data"].([]map[string]interface{})
	if !ok{
		return 0, errors.New("net error")
	}
	//{"result":true,"data":[{"LoginName":"bbinliul","Currency":"RMB","Balance":1,"TotalBalance":1}],
	// "pagination":{"Page":1,"PageLimit":500,"TotalNumber":1,"TotalPage":1}}
	balance, ok := dataMap[0]["Balance"]
	if !ok{
		fmt.Println(2)
		return 0, errors.New("net error")
	}
	return balance,nil
}

//查询记录
func (g *GameBBin) QueryRecord(info map[string]interface{}) (interface{}, error) {
	data := make(map[string]interface{})
	data["startTime"] = info["start_time"].(string) // string
	data["endTime"] = info["end_time"].(string)     //string
	data["channelId"] = "-1"
	data["state"] = "-1"
	data["recordPage"] = strconv.Itoa(info["page"].(int))       //int
	data["recordPage"] = strconv.Itoa(info["page_count"].(int)) //int
	fmt.Println(data)
	mapRes, err := doRequestPost("/MerchantQuery/Bet", data)
	if err != nil {
		return nil, err
	}
	fmt.Println(mapRes)
	return true, nil
}

func (g *GameBBin) GetPrefix() string {
	return UserPrefix
}

func NewBBinGame() game.Game {
	return &GameBBin{}
}

func init() {
	game.Register("bbin", NewBBinGame)
}
