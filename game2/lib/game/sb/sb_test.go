package sb

import (
	"fmt"
	"game2/lib/game"
	"game2/lib/utils"
	"testing"
)

//注册
func TestGameSBRegister(t *testing.T) {
	s, _ := game.NewGame("sb")
	m := make(map[string]interface{})
	m["user_name"] = "liul"
	m["moneysort"] = "RMB"
	resp, err := s.Register(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//登录
func TestGameSBLogin(t *testing.T) {
	s, _ := game.NewGame("sb")
	m := make(map[string]interface{})
	m["game_user_name"] = "jkgsbliul"
	m["gametype"] = "2"
	m["oddtype"] = "A"
	m["gamekind"] = "1"
	m["iframe"] = "0"
	m["platformname"] = "IBC"
	m["lang"] = "cn"
	resp, err := s.Login(m)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", resp)
}

//查询余额
func TestGameSBGetBalance(t *testing.T) {
	s, _ := game.NewGame("sb")
	m := make(map[string]interface{})
	m["game_type"] = "SB_TY"
	m["game_user_name"] = "jkgsbliul"
	m["platformname"] = "IBC"
	resp, err := s.GetBalance(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

//中心专户转到游戏账户
func TestGameSB_Account2GameTransfer(t *testing.T) {
	s, _ := game.NewGame("sb")
	m := make(map[string]interface{})
	random := string(utils.Krand(18, utils.KC_RAND_KIND_ALL))
	fmt.Println(random)
	m["game_user_name"] = "jkgsbliul"
	m["billno"] = random //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	m["usertype"] = "0"  // 1 正常;0 测试
	m["amount"] = "1"    // 操作的额度，请使用整数，如:100
	m["platformname"] = "IBC"
	resp, err := s.Account2GameTransfer(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

//游戏账户转到中心专户
func TestGameSB_Game2AccountTransfer(t *testing.T) {
	s, _ := game.NewGame("sb")
	m := make(map[string]interface{})
	random := string(utils.Krand(18, utils.KC_RAND_KIND_ALL))
	fmt.Println(random)
	m["game_user_name"] = "jkgsbliul"
	m["billno"] = random //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	m["usertype"] = "0"  // 1 正常;0 测试
	m["credit"] = "1"    // 操作的额度，请使用整数，如:100
	m["platformname"] = "IBC"
	resp, err := s.Game2AccountTransfer(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

//查询转帐
func TestGameSB_ConfirmTransferCredit(t *testing.T) {
	//s, _ := game.NewGame("sb")
	m := make(map[string]interface{})
	//random := string(utils.Krand(18, utils.KC_RAND_KIND_ALL))
	//fmt.Println(random)
	m["game_user_name"] = "jkgsbliul"
	m["billno"] = "85uUz97WmAncL3S0Yd" //billno=( sequence), sequence 必须是唯一的数字 (最多18 个数字)
	//m["usertype"] = "0" // 1 正常;0 测试
	resp, err := ConfirmTransferCredit(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
