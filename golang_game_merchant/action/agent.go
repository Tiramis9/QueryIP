package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AgentList(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "class_name": "代理层级1", "team_member_num": 10, "last_login_time": 1539601066, "last_login_ip": "127.0.0.1", "reg_time": "127.0.0.1", "user_name": "代理1", "status": 1},
		{"id": 2, "class_name": "代理层级1", "team_member_num": 8, "last_login_time": 1539601066, "last_login_ip": "127.0.0.1", "reg_time": "127.0.0.1", "user_name": "代理1", "status": 1},
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func AgentReportList(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "user_name": "aaabc", "total_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6},
		{"id": 2, "user_name": "cccdf", "total_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6},
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func AgentInfo(c *gin.Context) {
	data := map[string]interface{}{"user_name": "ssbv", "true_name": "大娃", "class_name": "代理层级1", "phone": "13111110101", "team_member_num": 10, "email": "4344@gmail.com", "reg_time": 1539601066, "qq": 46464454, "last_login_time": 1539601066, "skype": "232323"}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

func SubordinateList(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "user_name": "aaabc", "last_login_time": 1539601066, "reg_time": 1539601066, "device": 1, "source": "http://www.jx550.com/preferential.html", "status": 1},
		{"id": 2, "user_name": "dddfs", "last_login_time": 1539601066, "reg_time": 1539601066, "device": 1, "source": "http://www.jx550.com/preferential.html", "status": 1},
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func CommissionReport(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "user_id": 1, "start_time": 1539601066, "end_time": 1539601076, "user_name": "aaabc", "effective_member": 20, "total_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6, "org_expense": 2, "brokerage_total": 2},
		{"id": 2, "user_id": 1, "start_time": 1539601066, "end_time": 1539601076, "user_name": "aaabc", "effective_member": 20, "total_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6, "org_expense": 2, "brokerage_total": 2},
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func CommissionHistory(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "user_id": 1, "start_time": 1539601066, "end_time": 1539601076, "user_name": "aaabc", "effective_member": 20, "total_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6, "org_expense": 2, "brokerage_total": 2},
		{"id": 2, "user_id": 1, "start_time": 1539601066, "end_time": 1539601076, "user_name": "aaabc", "effective_member": 20, "total_bet": 2000, "win_lost_amount": 987, "recharge_amount": 954616, "withdraw_amount": 4897, "bonus_amount": 2000, "rebate": 15, "fee": 6, "org_expense": 2, "brokerage_total": 2},
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func AgentCheck(c *gin.Context) {
	data := [...]map[string]interface{}{{"id": 1, "user_name": "ssbv", "true_name": "大娃", "phone": "13111110101", "email": "4344@gmail.com", "reg_time": 1539601066, "remark": "aaa", "qq": 46464454, "skype": "232323", "status": 1}, {"id": 2, "user_name": "ssbv", "true_name": "大娃", "phone": "13111110101", "email": "4344@gmail.com", "reg_time": 1539601066, "remark": "aaa", "qq": 46464454, "skype": "232323", "status": 1}}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}
