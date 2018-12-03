package agin

import (
	"testing"
	"game2/lib/utils"
)

func TestCheckOrCreateGameAccount(t *testing.T) {
	game := newGameAGin()

	m := make(map[string]interface{})
	m["ac_type"] = AcTypeRealMoney
	//m["ac_type"] = AcTypeTryPlay
	//m["game_user_name"] = "DJ9test1"
	m["game_user_name"] = "DJ9test2"
	m["lang"] = "cn"
	resp, err := game.Login(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestGenerateForwardGameUrl(t *testing.T) {
	game := newGameAGin()

	m := make(map[string]interface{})
	m["ac_type"] = AcTypeRealMoney
	//m["ac_type"] = AcTypeTryPlay
	//m["game_user_name"] = "DJ9test1"
	m["game_user_name"] = "DJ9test2"
	m["lang"] = "cn"
	resp, err := game.generateForwardGameUrl(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestGetBalance(t *testing.T) {
	game := newGameAGin()

	m := make(map[string]interface{})
	m["ac_type"] = AcTypeRealMoney
	//m["ac_type"] = AcTypeTryPlay
	//m["game_user_name"] = "DJ9test1"
	m["game_user_name"] = "DJ9test2"
	m["lang"] = "cn"
	resp, err := game.GetBalance(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestAccount2GameTransfer(t *testing.T) {
	game := newGameAGin()

	orderSn := utils.CreateOrderNo(10)
	t.Log(orderSn)
	// DJ9_AGIN154348181010698 —— +1元
	// DJ9_AGIN154348285810698 —— +1元
	// DJ9_AGIN154348293910698 —— +1元
	m := make(map[string]interface{})
	m["game_user_name"] = "DJ9test2"
	m["lang"] = "cn"
	m["order_sn"] = game.Agent+orderSn
	m["type"] = "IN"
	m["amount"] = "1"
	resp, err := game.Account2GameTransfer(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestGame2AccountTransfer(t *testing.T) {
	game := newGameAGin()

	orderSn := utils.CreateOrderNo(10)
	t.Log(orderSn)
	// DJ9_AGIN154348300810698 —— -1元
	// DJ9_AGIN154348407810698 —— -1元
	m := make(map[string]interface{})
	m["game_user_name"] = "DJ9test2"
	m["lang"] = "cn"
	m["order_sn"] = game.Agent+orderSn
	m["type"] = "IN"
	m["amount"] = "1"
	resp, err := game.Game2AccountTransfer(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestQueryOrderStatus(t *testing.T) {
	game := newGameAGin()
	m := make(map[string]interface{})
	m["order_sn"] = "DJ9_AGIN154348300810698"
	m["lang"] = "cn"
	err := game.queryOrderStatus(m)
	t.Log(err)
}
