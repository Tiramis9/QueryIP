package action

import (
	"fmt"
	"golang_game_merchant/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//账户列表
func AccountList(c *gin.Context) {
	var data interface{}
	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	user_id := userid.(int)
	account := model.UserAccount{}
	//转为整形
	//获取列表
	accountlist := account.GetAccountListByUserId(user_id)
	data = accountlist
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

//账户详情
func AccountInfo(c *gin.Context) {
	var data interface{}
	game_name := c.PostForm("game_name")
	//调用第三方接口
	data = map[string]interface{}{"game_balance": "1000.00", "game_name": game_name}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

//
func AccountTrans(c *gin.Context) {
	var data interface{}

	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	user_id := userid.(int)
	from := c.PostForm("from")
	to := c.PostForm("to")
	amount := c.PostForm("amount")

	//转为整形
	//金额转为float64
	amount_float, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println(err)
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	if (from == "0" && to == "0") || (from != "0" && to != "0") {
		res := gin.H{"code": 0, "data": data, "msg": "field error"}
		c.JSON(http.StatusOK, res)
		return
	}
	account_str := from
	if from == "0" {
		account_str = to
	}
	account_name := "AG"
	switch account_str {
	case "1":
	case "2":
		account_name = "BBIN"
	case "3":
		account_name = "SB"
	default:

	}
	account := model.UserAccount{Account_name: account_name, User_id: user_id, Money: amount_float}
	//获取列表
	ok, msg := account.TransAccount(user_id, from, to, amount_float)
	fmt.Println(ok, msg)
	res := gin.H{"code": 1, "data": nil, "msg": "ok"}
	if !ok {
		res = gin.H{"code": 0, "data": nil, "msg": msg}
	}
	c.JSON(http.StatusOK, res)
}
