package bbin

import (
	"fmt"
	"game2/lib/game"
	"game2/lib/utils"
	"testing"
)

/*func TestGameVRRegister(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"li1"}
	res, err:= Register(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}*/

func TestLogin(t *testing.T) {
	g,err := game.NewGame("bbin")
	userInfo := map[string]interface{}{"user_name":"liul"}
	res, err:= g.Login(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestPlayGame(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"liul","game_kind":"3","game_type":"3016","game_code":"1"}
	res, err:= PlayGame(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestAccount2GameTransfer(t *testing.T) {
	g,err := game.NewGame("bbin")
	userInfo := map[string]interface{}{"user_name":"liul","bill_no":"12562362372673","amount":"1"}
	res, err:= g.Account2GameTransfer(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestGame2AccountTransfer(t *testing.T) {
	g,err := game.NewGame("bbin")
	userInfo := map[string]interface{}{"user_name":"liul","bill_no":utils.CreateOrderNo(1),"amount":"1"}
	res, err:= g.Game2AccountTransfer(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestGetBalance(t *testing.T) {
	g,err := game.NewGame("bbin")
	userInfo := map[string]interface{}{"user_name":"liul"}
	res, err:= g.GetBalance(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestAddGame(t *testing.T) {
	//AddGame()
}

func TestSportEventUrl(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"liul"}
	res, err:= SportEventUrl(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}
