package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang_game_merchant/global/status"
)

func GameTypeList(c *gin.Context) {
	//user_id := c.PostForm("user_id")
	data := map[string]interface{}{}
	data_ag := map[string]interface{}{}
	data_pt := map[string]interface{}{}
	data_abc := map[string]interface{}{}
	data1 := [...]map[string]interface{}{{"id": 1, "url": "http://www.baidu.com", "cover": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "game_name": "AG捕鱼王", "type": 6, "desc": "666", "group": "1", "group_name": "AG捕鱼王"}}
	data2 := [...]map[string]interface{}{{"id": 2, "url": "http://www.baidu.com", "cover": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538115150044&di=5436bd6cf7658d869400d439c208112b&imgtype=0&src=http%3A%2F%2Fimg5.duitang.com%2Fuploads%2Fitem%2F201409%2F19%2F20140919234215_RmBdj.png", "game_name": "弓兵", "type": 6, "desc": "666", "group": "2", "group_name": "PT电子游戏"},
		{"id": 3, "url": "http://www.baidu.com", "cover": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538115150044&di=5436bd6cf7658d869400d439c208112b&imgtype=0&src=http%3A%2F%2Fimg5.duitang.com%2Fuploads%2Fitem%2F201409%2F19%2F20140919234215_RmBdj.png", "game_name": "五虎将", "type": 6, "desc": "666", "group": "2", "group_name": "PT电子游戏"},
	}

	data3 := [...]map[string]interface{}{{"id": 4, "url": "http://www.baidu.com", "cover": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538115150044&di=5436bd6cf7658d869400d439c208112b&imgtype=0&src=http%3A%2F%2Fimg5.duitang.com%2Fuploads%2Fitem%2F201409%2F19%2F20140919234215_RmBdj.png", "game_name": "弓兵2", "type": 6, "desc": "666", "group": "2", "group_name": "ABC电子游戏"},
		{"id": 5, "url": "http://www.baidu.com", "cover": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538115150044&di=5436bd6cf7658d869400d439c208112b&imgtype=0&src=http%3A%2F%2Fimg5.duitang.com%2Fuploads%2Fitem%2F201409%2F19%2F20140919234215_RmBdj.png", "game_name": "五虎将2", "type": 6, "desc": "666", "group": "2", "group_name": "ABC电子游戏"},
	}

	data_ag["game_data"] = data1
	data_ag["icon"] = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg"
	data_ag["cid"] = "fish"

	data_pt["game_data"] = data2
	data_pt["icon"] = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538115150044&di=5436bd6cf7658d869400d439c208112b&imgtype=0&src=http%3A%2F%2Fimg5.duitang.com%2Fuploads%2Fitem%2F201409%2F19%2F20140919234215_RmBdj.png"
	data_pt["cid"] = "pt_game"

	data_abc["game_data"] = data3
	data_abc["icon"] = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538115150044&di=5436bd6cf7658d869400d439c208112b&imgtype=0&src=http%3A%2F%2Fimg5.duitang.com%2Fuploads%2Fitem%2F201409%2F19%2F20140919234215_RmBdj.png"
	data_abc["cid"] = "abc_game"

	data["AG捕鱼王"] = data_ag
	data["PT电子游戏"] = data_pt
	data["ABC电子游戏"] = data_abc
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func GameAllList(c *gin.Context) {
	data_sport_list := [...]map[string]interface{}{
		{"game_name": "沙巴体育", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
		{"game_name": "NEW BB体育", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
	}
	data_sport_total := map[string]interface{}{"name": "体育总计", "effective_member": 40, "bet_num": 40, "total_bet": 40, "effective_bet": 40, "win_lost_amount": 40, "win_rate": "50%"}

	data_real_list := [...]map[string]interface{}{
		{"game_name": "AG国际厅", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
		{"game_name": "欧博视讯", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
	}
	data_real_total := map[string]interface{}{"name": "真人总计", "effective_member": 40, "bet_num": 40, "total_bet": 40, "effective_bet": 40, "win_lost_amount": 40, "win_rate": "50%"}

	data_lottery_list := [...]map[string]interface{}{
		{"game_name": "VR彩票", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
		{"game_name": "BBIN彩票", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
	}
	data_lottery_total := map[string]interface{}{"name": "彩票总计", "effective_member": 40, "bet_num": 40, "total_bet": 40, "effective_bet": 40, "win_lost_amount": 40, "win_rate": "50%"}

	data_chess_list := [...]map[string]interface{}{
		{"game_name": "开元棋牌", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
		{"game_name": "BBIN彩票", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
	}

	data_chess_total := map[string]interface{}{"name": "棋牌游戏总计", "effective_member": 40, "bet_num": 40, "total_bet": 40, "effective_bet": 40, "win_lost_amount": 40, "win_rate": "50%"}

	data_game_list := [...]map[string]interface{}{
		{"game_name": "三国", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
		{"game_name": "仁王", "effective_member": 20, "bet_num": 20, "total_bet": 20, "effective_bet": 20, "win_lost_amount": 20, "win_rate": "50%"},
	}

	game_active1 := [...]map[string]interface{}{
		{"award": 0, "pool": 0, "game_name": "jackpot"},
		{"award": 0, "pool": 0, "game_name": "jackpot2"},
	}

	tips_list := [...]map[string]interface{}{
		{"tips_num": 0, "game_name": "彩播"},
		{"tips_num": 0, "game_name": "斗鱼"},
	}

	data_game_total := map[string]interface{}{"name": "电子游戏总计", "effective_member": 40, "bet_num": 40, "total_bet": 40, "effective_bet": 40, "win_lost_amount": 40, "win_rate": "50%"}

	data_total := map[string]interface{}{"name": "总计", "effective_member": 200, "bet_num": 200, "total_bet": 200, "effective_bet": 200, "win_lost_amount": 200, "win_rate": "50%"}

	data_sport := map[string]interface{}{"list": data_sport_list, "total": data_sport_total}
	data_lottery := map[string]interface{}{"list": data_lottery_list, "total": data_lottery_total}
	data_real := map[string]interface{}{"list": data_real_list, "total": data_real_total}
	data_chess := map[string]interface{}{"list": data_chess_list, "total": data_chess_total}
	data_game := map[string]interface{}{"list": data_game_list, "total": data_game_total}
	datas := map[string]interface{}{}
	game_active_list := map[string]interface{}{}
	game_active_list["pt_game"] = game_active1

	datas["game_active_list"] = game_active_list
	datas["tips_list"] = tips_list
	datas["sport"] = data_sport
	datas["lottery"] = data_lottery
	datas["real"] = data_real
	datas["chess"] = data_chess
	datas["game"] = data_game
	datas["total"] = data_total
	/*datas["ag_game_jackpot_award"] = 0
	datas["ag_game_jackpot_pool"] = 0
	datas["bg_tips"] = 0
	datas["caibo_tips_cash"] = 0
	datas["dg_tips"] = 0*/
	datas["dushen_bonus"] = 0
	datas["dushen_entry_fee"] = 0
	datas["dushen_refund"] = 0
	/*datas["pt_game_jackpot_award"] = 0
	datas["pt_game_jackpot_pool"] = 0*/

	RespJson(c, status.OK, datas)
}

func GameList(c *gin.Context) {
	game_type := c.PostForm("type")
	var data interface{}
	switch game_type {
	//类型 1.真人视讯;2.彩票游戏;3.棋牌游戏;4.电子游戏;5.体育赛事;
	case "1":
		data = [...]map[string]interface{}{
			{"game_code": "AG_SX", "game_name": "AG视讯", "status": 1, "type": 1, "channel": "AG"},
			{"game_code": "BBIN_SX", "game_name": "BBIN视讯", "status": 1, "type": 1, "channel": "AG"},
		}
	case "2":
		data = [...]map[string]interface{}{
			{"game_code": "VR_CP", "game_name": "VR彩票", "status": 1, "type": 2, "channel": "VR"},
			{"game_code": "SSC_CP", "game_name": "时时彩", "status": 1, "type": 2, "channel": "SSC"},
		}
	case "3":
		data = [...]map[string]interface{}{
			{"game_code": "KY_QP", "game_name": "开元棋牌", "status": 1, "type": 3, "channel": "KY"},
			{"game_code": "CS_QP", "game_name": "财神棋牌", "status": 1, "type": 3, "channel": "CS"},
		}
	case "4":
		data = [...]map[string]interface{}{
			{"game_code": "BYW_BYW", "game_name": "捕鱼王", "status": 1, "type": 4, "channel": "BYW"},
			{"game_code": "PT_DZ", "game_name": "PT电子游戏", "status": 1, "type": 4, "channel": "PT"},
		}
	case "5":
		data = [...]map[string]interface{}{
			{"game_code": "SB_TY", "game_name": "沙巴体育", "status": 1, "type": 5, "channel": "SB"},
			{"game_code": "BBIN_TY", "game_name": "BBIN体育", "status": 1, "type": 5, "channel": "BBIN"},
			{"game_code": "BBIN_XTY", "game_name": "BBIN新体育", "status": 1, "type": 5, "channel": "BBIN"},
		}
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

func GameSingleReport(c *gin.Context) {
	game_code := c.PostForm("game_code")
	var data interface{}
	switch game_code {
	//类型 1.真人视讯;2.彩票游戏;3.棋牌游戏;4.电子游戏;5.体育赛事;
	case "AG_SX":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "BBIN_SX":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "VR_CP":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "SSC_CP":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "KY_QP":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_amount": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_amount": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "CS_QP":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_amount": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_amount": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "BYW_BYW":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "win_lost_amount": 1000, "win_rate": "50%", "extract_amount": 5, "bullet_value": 100, "bonus": 20, "mission rewards": 20},
			{"user_name": "dfdf", "bet_num": 2, "win_lost_amount": 1000, "win_rate": "50%", "extract_amount": 5, "bullet_value": 100, "bonus": 20, "mission rewards": 20},
		}
	case "PT_DZ":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%"},
		}
	case "SB_TY":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%", "create_time": 1539601076},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%", "create_time": 1539601076},
		}
	case "BBIN_TY":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%", "create_time": 1539601076},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%", "create_time": 1539601076},
		}
	case "BBIN_XTY":
		data = [...]map[string]interface{}{
			{"user_name": "aaafd", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%", "create_time": 1539601076},
			{"user_name": "dfdf", "bet_num": 2, "total_bet": 2000, "effective_bet": 2000, "win_lost_amount": 1000, "win_rate": "50%", "create_time": 1539601076},
		}
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}