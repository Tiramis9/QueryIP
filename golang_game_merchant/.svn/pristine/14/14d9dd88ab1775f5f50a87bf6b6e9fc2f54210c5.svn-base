package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang_game_merchant/global/status"
)

func MerchantLogin(c *gin.Context) {
	data := map[string]interface{}{"token": "xxxxxxxxxxxxxxxxxxxxxxxxx"}
	RespJson(c, status.OK, data)
}

func MerchantLogout(c *gin.Context) {
	data := [...]map[string]interface{}{}
	RespJson(c, status.OK, data)
}

func MerchantBaseInfo(c *gin.Context) {
	data := map[string]interface{}{"id": 1, "name": "巨星国际", "code": "jx", "bail": 499102.00, "withdraw_simple_min": 100, "withdraw_simple_max": 10000, "service_online_url": "www.baidu.com", "agent_plat_url": "www.baidu.com", "agent_spread_url": "www.baidu.com", "app_download_url": "www.baidu.com", "app_logo": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539597282176&di=48932ce83f9a0dd16c32346831825613&imgtype=0&src=http%3A%2F%2Fwww.jituwang.com%2Fuploads%2Fallimg%2F151208%2F258057-15120R3593535.jpg", "allow_ip_minute": 1000, "reg_status": 1, "active_status": 0}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

func MerchantBaseInfoEdit(c *gin.Context) {
	//data := [...]map[string]interface{}{}
	RespSuccess(c)
}

func RegisterPageConf(c *gin.Context) {
	data := map[string]interface{}{"reg_pay_pass": 1, "reg_security_question": 1, "reg_true_name": 1, "reg_phone": 1, "reg_email": 1}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

func RegisterPageConfEdit(c *gin.Context) {
	RespSuccess(c)
}

func MerchantGameStatus(c *gin.Context) {
	data1 := [...]map[string]interface{}{
		{"game_id": 1, "game_name": "啦啦真人视讯", "status": 1},
		{"game_id": 2, "game_name": "哈哈真人视讯", "status": 1},
	}
	data2 := [...]map[string]interface{}{
		{"game_id": 3, "game_name": "重庆时时彩", "status": 1},
		{"game_id": 4, "game_name": "广东快乐十分", "status": 1},
	}
	data3 := [...]map[string]interface{}{
		{"game_id": 5, "game_name": "斗地主", "status": 1},
		{"game_id": 6, "game_name": "跑得快", "status": 1},
	}
	data4 := [...]map[string]interface{}{
		{"game_id": 7, "game_name": "仁王", "status": 1},
		{"game_id": 8, "game_name": "三国演义", "status": 1},
	}
	data5 := [...]map[string]interface{}{
		{"game_id": 9, "game_name": "沙巴体育", "status": 1},
		{"game_id": 10, "game_name": "BB体育", "status": 1},
	}
	data := [...]map[string]interface{}{{"list": data1, "type": 1}, {"list": data2, "type": 2}, {"list": data3, "type": 3}, {"list": data4, "type": 4}, {"list": data5, "type": 5}}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

func MerchantGameStatusEdit(c *gin.Context) {
	RespSuccess(c)
}

func ShieldIpAddress(c *gin.Context) {
	RespSuccess(c)
}

func ShieldAreaAddress(c *gin.Context) {
	RespSuccess(c)
}
