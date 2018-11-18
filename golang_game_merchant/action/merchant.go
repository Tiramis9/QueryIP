package action

import (
	"golang_game_merchant/global/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MerchantLogin(c *gin.Context) {
	data := map[string]interface{}{"token": "xxxxxxxxxxxxxxxxxxxxxxxxx"}
	RespJson(c, status.OK, data)
}

func MerchantLogout(c *gin.Context) {
	data := [...]map[string]interface{}{}
	RespJson(c, status.OK, data)
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
