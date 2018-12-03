package ky

import (
	"testing"
	"time"
	"fmt"
)

func TestLogin(t *testing.T) {
	game := NewGameKY()
	m := make(map[string]interface{})
	m["game_user_name"] = "test"
	m["login_ip"] = "127.0.0.1"
	m["game_code"] = "0"
	resp, err := game.Login(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestGetBalance(t *testing.T) {
	game := NewGameKY()
	m := make(map[string]interface{})
	m["game_user_name"] = "test"
	resp, err := game.GetBalance(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestAccount2GameTransfer(t *testing.T) {
	game := NewGameKY()
	m := make(map[string]interface{})
	m["game_user_name"] = "test"
	m["amount"] = "1"
	resp, err := game.Account2GameTransfer(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestGame2AccountTransfer(t *testing.T) {
	game := NewGameKY()
	m := make(map[string]interface{})
	m["game_user_name"] = "test"
	m["amount"] = "1"
	resp, err := game.Game2AccountTransfer(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestGetGameRecord(t *testing.T) {
	game := NewGameKY()
	m := make(map[string]interface{})
	now := time.Now()
	m["start_time"] = fmt.Sprint(now.Add(-5*time.Second).UnixNano()/1e6)
	m["end_time"] = fmt.Sprint(now.UnixNano()/1e6)
	resp, err := game.QueryRecord(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
