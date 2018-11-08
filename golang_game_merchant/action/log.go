package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OpeationLog(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": "1", "user_name": "葫芦娃", "main_systematics": "admin", "sub_systematics": "login", "content": "admin/login", "ip": "127.0.0.1", "create_time": "1539601066"},
		{"id": "2", "user_name": "葫芦娃2", "main_systematics": "admin", "sub_systematics": "login", "content": "admin/login", "ip": "127.0.0.1", "create_time": "1539601066"},
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}
