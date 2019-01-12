package mg

import (
	"game2/lib/game"
	"testing"
)

//获取token
func TestGameMG_DoLogin(t *testing.T) {
	//mg, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["time_zone"] = "UTC"
	m["currency"] = "CNY"
	m["lang"] = "cn"
	resp, err := DoLogin(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//刷新token
func TestGameMG_DoRefreshToken(t *testing.T) {
	//mg, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["time_zone"] = "UTC+8"
	m["currency"] = "CNY"
	m["lang"] = "cn"
	m["refresh_token"] = "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX25hbWUiOiJKS0dDTllBX2FwaSIsImN0eCI6MzcsInBpZCI6MjI0MCwiYW4iOiJKS0dDTllBIiwiY2xpZW50X2lkIjoiQVZJQUNOWV9hdXRoIiwiYXAiOiIxLDEwMDYsODg2ODE2MCwzNjIwNzgyNCIsInVpZCI6MzUxNzY5NDMsImF0Ijo0LCJzY29wZSI6WyJhdWRpdDpyIiwibGF1bmNoZXJfaXRlbTpyIiwidHg6ciIsImFwcF9uYW1lOnIiLCJleGNoYW5nZV9yYXRlczpyIiwiY2FtcGFpZ246dyIsImFwcF9pbnN0YWxsZWQ6ciIsInVzZXI6dyIsIndhbGxldDpyIiwiY2FtcGFpZ246ciIsInRva2VuOnciLCJyZXBvcnQ6ciIsInVzZXI6ciIsImFjY2FwcDpyIiwiY2F0ZWdvcnk6ciIsImFjY291bnQ6dyIsIml0ZW06ciIsInR4OnciLCJhY2NvdW50OnIiXSwiYXRpIjoiN2VkZWM3YzItOGZiZS00M2Q3LWI4OWQtM2JkNTE5ZTQxYTIzIiwiZXhwIjoxNTQzNDAzODkxLCJhaWQiOjM2MjA3ODI0LCJ1ciI6MywianRpIjoiNGU1ZmMyNTYtNTU3YS00ZjU2LTg5MTctZGE0YWZjYzQ1YmZiIn0.oeb5Y9eL_gQyL7LQx9NQecit3r_Bz8CwKDfzBzlXhUM"
	resp, err := DoRefreshToken(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//登录游戏
func TestGameMG_Login(t *testing.T) {
	mg, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["time_zone"] = "UTC"
	m["currency"] = "CNY"
	m["lang"] = "cn"
	m["user_name"] = "test8"
	m["user_id"] = 50015
	m["account_id"] = "38814022" //会员游戏id
	m["item_id"] = "1028"        //游戏id
	m["app_id"] = "1001"         //游戏id
	m["access_token"] = "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX25hbWUiOiJKS0dDTllBX2FwaSIsImN0eCI6MzgsInBpZCI6MjI0MCwiYW4iOiJKS0dDTllBIiwiY2xpZW50X2lkIjoiQVZJQUNOWV9hdXRoIiwiYXAiOiIxLDEwMDYsODg2ODE2MCwzNjIwNzgyNCIsInVpZCI6MzUxNzY5NDMsImF0Ijo0LCJzY29wZSI6WyJhdWRpdDpyIiwibGF1bmNoZXJfaXRlbTpyIiwidHg6ciIsImFwcF9uYW1lOnIiLCJleGNoYW5nZV9yYXRlczpyIiwiY2FtcGFpZ246dyIsImFwcF9pbnN0YWxsZWQ6ciIsInVzZXI6dyIsIndhbGxldDpyIiwiY2FtcGFpZ246ciIsInRva2VuOnciLCJyZXBvcnQ6ciIsInVzZXI6ciIsImFjY2FwcDpyIiwiY2F0ZWdvcnk6ciIsImFjY291bnQ6dyIsIml0ZW06ciIsInR4OnciLCJhY2NvdW50OnIiXSwiZXhwIjoxNTQzNDAwNTA0LCJhaWQiOjM2MjA3ODI0LCJ1ciI6MywianRpIjoiNDY2MjgyNGYtMGExNy00MmRlLWI5NTgtMWExYjMxNmQzNmRlIn0.WqOkNe-Al1bql-Z495nlshI03vnxJGyhHlh4zOb0njU"
	resp, err := mg.Login(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//注册
func TestGameMGRegister(t *testing.T) {
	mg, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["access_token"] = "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX25hbWUiOiJKS0dDTllBX2FwaSIsImN0eCI6MzksInBpZCI6MjI0MCwiYW4iOiJKS0dDTllBIiwiY2xpZW50X2lkIjoiQVZJQUNOWV9hdXRoIiwiYXAiOiIxLDEwMDYsODg2ODE2MCwzNjIwNzgyNCIsInVpZCI6MzUxNzY5NDMsImF0Ijo0LCJzY29wZSI6WyJhdWRpdDpyIiwibGF1bmNoZXJfaXRlbTpyIiwidHg6ciIsImFwcF9uYW1lOnIiLCJleGNoYW5nZV9yYXRlczpyIiwiY2FtcGFpZ246dyIsImFwcF9pbnN0YWxsZWQ6ciIsInVzZXI6dyIsIndhbGxldDpyIiwiY2FtcGFpZ246ciIsInRva2VuOnciLCJyZXBvcnQ6ciIsInVzZXI6ciIsImFjY2FwcDpyIiwiY2F0ZWdvcnk6ciIsImFjY291bnQ6dyIsIml0ZW06ciIsInR4OnciLCJhY2NvdW50OnIiXSwiZXhwIjoxNTQzNDA4MTU2LCJhaWQiOjM2MjA3ODI0LCJ1ciI6MywianRpIjoiZmNlYmQ2MDItMDRkOS00NWZmLWJhYTgtMjQxZDc1NWY2ZTljIn0.g4AMBOO5kV2FneA_c0pocPqRCmviDW8A_yVh9g7DlIA"
	m["time_zone"] = "UTC+8"
	m["currency"] = "CNY"
	m["lang"] = "cn"
	m["user_name"] = "test8"
	//m["password"] = "123456"
	//m["ext_ref"] = "0"
	resp, err := mg.Register(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//获取余额
func TestGameMG_GetBalance(t *testing.T) {
	s, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["account_id"] = "38676990" //游戏账户id
	m["access_token"] = "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX25hbWUiOiJKS0dDTllBX2FwaSIsImN0eCI6MzcsInBpZCI6MjI0MCwiYW4iOiJKS0dDTllBIiwiY2xpZW50X2lkIjoiQVZJQUNOWV9hdXRoIiwiYXAiOiIxLDEwMDYsODg2ODE2MCwzNjIwNzgyNCIsInVpZCI6MzUxNzY5NDMsImF0Ijo0LCJzY29wZSI6WyJhdWRpdDpyIiwibGF1bmNoZXJfaXRlbTpyIiwidHg6ciIsImFwcF9uYW1lOnIiLCJleGNoYW5nZV9yYXRlczpyIiwiY2FtcGFpZ246dyIsImFwcF9pbnN0YWxsZWQ6ciIsInVzZXI6dyIsIndhbGxldDpyIiwiY2FtcGFpZ246ciIsInRva2VuOnciLCJyZXBvcnQ6ciIsInVzZXI6ciIsImFjY2FwcDpyIiwiY2F0ZWdvcnk6ciIsImFjY291bnQ6dyIsIml0ZW06ciIsInR4OnciLCJhY2NvdW50OnIiXSwiZXh0dyI6ZmFsc2UsImV4cCI6MTU0MzM5NzAxNCwiYWlkIjozNjIwNzgyNCwidXIiOjMsImp0aSI6ImFiYTZiZTZiLWI4YzItNDQzYi04NjljLTIwN2VlODZmM2RiYiJ9.bGXFay01QIRsRrbxG4XygPxnOqh8CecvabIl1wTHUf0"
	m["time_zone"] = "UTC"
	m["currency"] = "CNY"
	m["lang"] = "cn"

	resp, err := s.GetBalance(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//中心账户转游戏账户
func TestGameMG_Account2GameTransfer(t *testing.T) {
	s, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["account_id"] = "38676990" //游戏账户id
	m["access_token"] = "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX25hbWUiOiJKS0dDTllBX2FwaSIsImN0eCI6MzcsInBpZCI6MjI0MCwiYW4iOiJKS0dDTllBIiwiY2xpZW50X2lkIjoiQVZJQUNOWV9hdXRoIiwiYXAiOiIxLDEwMDYsODg2ODE2MCwzNjIwNzgyNCIsInVpZCI6MzUxNzY5NDMsImF0Ijo0LCJzY29wZSI6WyJhdWRpdDpyIiwibGF1bmNoZXJfaXRlbTpyIiwidHg6ciIsImFwcF9uYW1lOnIiLCJleGNoYW5nZV9yYXRlczpyIiwiY2FtcGFpZ246dyIsImFwcF9pbnN0YWxsZWQ6ciIsInVzZXI6dyIsIndhbGxldDpyIiwiY2FtcGFpZ246ciIsInRva2VuOnciLCJyZXBvcnQ6ciIsInVzZXI6ciIsImFjY2FwcDpyIiwiY2F0ZWdvcnk6ciIsImFjY291bnQ6dyIsIml0ZW06ciIsInR4OnciLCJhY2NvdW50OnIiXSwiZXh0dyI6ZmFsc2UsImV4cCI6MTU0MzM5NzAxNCwiYWlkIjozNjIwNzgyNCwidXIiOjMsImp0aSI6ImFiYTZiZTZiLWI4YzItNDQzYi04NjljLTIwN2VlODZmM2RiYiJ9.bGXFay01QIRsRrbxG4XygPxnOqh8CecvabIl1wTHUf0"
	m["time_zone"] = "UTC+8"
	m["currency"] = "CNY"
	m["lang"] = "cn"
	m["amount"] = "2"

	resp, err := s.Account2GameTransfer(m)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//游戏账户转中心账户
func TestGameMG_Game2AccountTransfer(t *testing.T) {
	s, _ := game.NewGame("mg")
	m := make(map[string]interface{})
	m["account_id"] = "38676990" //游戏账户id
	m["access_token"] = "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX25hbWUiOiJKS0dDTllBX2FwaSIsImN0eCI6MzcsInBpZCI6MjI0MCwiYW4iOiJKS0dDTllBIiwiY2xpZW50X2lkIjoiQVZJQUNOWV9hdXRoIiwiYXAiOiIxLDEwMDYsODg2ODE2MCwzNjIwNzgyNCIsInVpZCI6MzUxNzY5NDMsImF0Ijo0LCJzY29wZSI6WyJhdWRpdDpyIiwibGF1bmNoZXJfaXRlbTpyIiwidHg6ciIsImFwcF9uYW1lOnIiLCJleGNoYW5nZV9yYXRlczpyIiwiY2FtcGFpZ246dyIsImFwcF9pbnN0YWxsZWQ6ciIsInVzZXI6dyIsIndhbGxldDpyIiwiY2FtcGFpZ246ciIsInRva2VuOnciLCJyZXBvcnQ6ciIsInVzZXI6ciIsImFjY2FwcDpyIiwiY2F0ZWdvcnk6ciIsImFjY291bnQ6dyIsIml0ZW06ciIsInR4OnciLCJhY2NvdW50OnIiXSwiZXh0dyI6ZmFsc2UsImV4cCI6MTU0MzM5NzAxNCwiYWlkIjozNjIwNzgyNCwidXIiOjMsImp0aSI6ImFiYTZiZTZiLWI4YzItNDQzYi04NjljLTIwN2VlODZmM2RiYiJ9.bGXFay01QIRsRrbxG4XygPxnOqh8CecvabIl1wTHUf0"
	m["time_zone"] = "UTC+8"
	m["currency"] = "CNY"
	m["lang"] = "cn"
	m["amount"] = "2"

	resp, err := s.Game2AccountTransfer(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}
