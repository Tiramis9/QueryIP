package vr

import (
	"fmt"
	"testing"
)

func TestGameVRRegister(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"li1"}
	res, err:= Register(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestForwardGame(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"li1"}
	res, err:= ForwardGame(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestGetBalance(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"li1"}
	res, err:= GetBalance(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestGame2AccountTransfer(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"li1","id":1,"amount":0.008}
	res, err:= Game2AccountTransfer(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestAccount2GameTransfer(t *testing.T) {
	userInfo := map[string]interface{}{"user_name":"li1","id":1,"amount":0.001}
	res, err:= Account2GameTransfer(userInfo)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}
