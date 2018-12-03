package og

import (
	"game2/lib/game"
	"testing"
)

// test 注册
func TestRegister(t *testing.T) {
	g, _ := game.NewGame("og")
	msg := make(map[string]interface{})
	token, err := GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	msg["username"] = "myuser123"
	msg["country"] = "China"
	msg["fullname"] = "MyUser"
	msg["language"] = "en"
	msg["email"] = "myuser123@test.com"
	msg["birthdate"] = "1992-02-18"

	resp, err := g.Register(msg) // ok
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}

// test Transaction
/*************测试注单*******************************************/

func TestRecordPAYAndPUT(t *testing.T) {
	const (
		OPERATOR_NAME = "mog074jk"             // X-Operator 经营方代码
		OPERATORKEY   = "D4NDSJujqQwkZvQaadBj" //  X-Key
	)
	data := make(map[string]interface{})
	//data["start_time"] = "2018-11-29 10:52:46"
	//data["end_time"] = "2018-11-29 10:59:46"
	data["start_time"] = "2018-11-29 06:50:46"
	data["end_time"] = "2018-11-29 06:59:46"
	g, _ := game.NewGame("og")

	result, err := g.QueryRecord(data) // 同如上 api 间隔10秒钟查询一次
	if err != nil {
		t.Log("err:", err)
	}
	t.Log(result)
}

// test  balance  updata or get
func TestBalance(t *testing.T) {
	token, err := GetToken()
	if err != nil {
		t.Fatal(err)
	}
	g, _ := game.NewGame("og")
	msg := make(map[string]interface{})
	msg["game_user_name"] = "myuser123"
	//msg["providerId"] = 1
	msg["X-Token"] = token
	msg["amount"] = 1
	msg["order_sn"] = "sampleTransferc00010"
	resp, err := g.GetBalance(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

	resp, err = g.Account2GameTransfer(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

/*****************玩游戏
   providerId 游戏供应商 id : 1
   key 游戏金钥
	  type  产生链接类型预设desktop允许值: "desktop", "mobile"
  ******************************/

func TestPlay(t *testing.T) {
	g, _ := game.NewGame("og")
	msg := make(map[string]interface{})
	//msg["providerId"] = 1
	msg["game_user_name"] = "myuser123"
	msg["type"] = "desktop"
	g.Login(msg)
	url, err := g.Login(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)

}
