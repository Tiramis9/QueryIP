package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

//处理client POST 消息，regexp MustIP,if OK returns IP
func MatchIP(Msg string) (string, error) {
	var str string
	myexp := regexp.MustCompile(`\d+\.\d+.\d+\.\d+`)
	result := myexp.FindAllStringSubmatch(Msg, 1)
	if result == nil {
		return "", errors.New("Please enter valid IP information!")
	}
	for _, value := range result {
		if value != nil {
			for _, str = range value {
			}
		}
	}
	return str, nil
}

/*
//处理client GET 消息，regexp MustIP,if OK returns IP
func MatchGETIP(Msg map[string][]string) (string, error) {
	//	var temp string
	var str string
	value, err := MapTOString(Msg)
	if err != nil {
		return "", err
	}
	myexp := regexp.MustCompile(`\d+\.\d+.\d+\.\d+`)
	result := myexp.FindAllStringSubmatch(value, 1)
	if result == nil {
		return "", errors.New("Please enter valid IP information! " + "[" + value + "]")
	}
	for _, value := range result {
		for _, str = range value {
		}
	}
	return str, nil
}
*/
// map[string][]string to string
func MapTOString(Msg map[string][]string) (string, error) {
	var template string
	if Msg == nil {
		return "", errors.New("Message Body is empty")
	}
	for _, value := range Msg {
		for _, template = range value {
		}
	}
	return template, nil
}

//结构体转换JSON
func StructTOJson(city *CityANDCountry) []byte {
	result, err := json.MarshalIndent(city, "", "")
	//result, err := json.Marshal(city)
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println(string(result))
	return result
}

//float64转换JSON
func FloatTOJson(location float64) []byte {
	//result, err := json.Marshal(Msg)
	result, err := json.MarshalIndent(location, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println("FloatTOJson", string(result))
	return result
}

//string转换JSON
func StringTOJson(TimeZone string) []byte {
	//result, err := json.Marshal(Msg)
	result, err := json.MarshalIndent(TimeZone, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	//	fmt.Println("FloatTOJson", string(result))
	return result
}
