package action

import (
	"fmt"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GlobalReport(c *gin.Context) {
	fmt.Println("welcome to GlobalReport")

	begin_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")
	/*
		token := c.PostForm("token")
		conn := redisclient.Get() //连接redis
		defer conn.Close()        //销毁本次链连接
		_, err := conn.Do("GET", token)
		if err != nil {
			res := gin.H{"code": 0, "data": nil, "msg": "test token inviald"}
			RespJson(c, status.OK, res)
			return
		}
	*/
	date := make(map[string]int)
	//responseData := map[string]interface{}{}
	begin_date := model.StringtoInt(begin_time)
	end_date := model.StringtoInt(end_time)
	date["start_time"] = begin_date
	date["end_time"] = end_date
	responseData, err := model.GetMerchantAnnouncement(model.Db, date)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, responseData)
	//	fmt.Println(data)
	//	RespJson(c, status.OK, data)
	/*

		data1 := [...]map[string]interface{}{
			{"game_name": "沙巴体育", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
			{"game_name": "BB体育", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
			{"game_name": "NEW BB体育", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
		}
		data2 := [...]map[string]interface{}{
			{"game_name": "BBIN彩票", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
			{"game_name": "VR彩票", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
		}
		data3 := [...]map[string]interface{}{
			{"game_name": "AG国际厅", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
			{"game_name": "DG视讯", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
			{"game_name": "欧博视讯", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
		}
		data4 := [...]map[string]interface{}{
			{"game_name": "开元棋牌", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
		}
		data5 := [...]map[string]interface{}{
			{"game_name": "捕鱼达人", "effective_member": "1", "effective_bet": "500.00", "win": "-500.00", "win_rate": "-100.00%"},
		}

		data_all := map[string]interface{}{}
		data := map[string]interface{}{}
		data["sport"] = data1
		data["lottery"] = data2
		data["real"] = data3
		data["chess"] = data4
		data["game"] = data5
		data_all["new_register"] = 200
		data_all["recharge_member"] = 200
		data_all["effective_bet"] = 2634.69
		data_all["win_lost_amount"] = -200
		data_all["recharge_amount"] = 200
		data_all["withdraw_amount"] = 200
		data_all["bonus_amount"] = 200 //红利
		data_all["rebate"] = 200       //反水
		data_all["data"] = data
		suc["data"] = data_all

		c.JSON(http.StatusOK, suc)
	*/
}

func FinanceList(c *gin.Context) {
	data_all := map[string]interface{}{}
	datas := [...]map[string]interface{}{{"id": 1, "user_name": "aaabc", "true_name": "大娃", "effective_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6}, {"id": 2, "user_name": "cccbc", "true_name": "二娃", "effective_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6}}
	data_all["data"] = datas
	data_all["members"] = 10
	data_all["effective_bet"] = 2000
	data_all["win_lost_amount"] = 500
	data_all["recharge_amount"] = 500
	data_all["bonus_amount"] = 500
	data_all["rebate"] = 500
	data_all["fee"] = 0
	res := gin.H{"code": 1, "data": data_all, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}
