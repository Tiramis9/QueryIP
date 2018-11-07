package server

import (
	"testing"
)

//curl -X POST "http://172.16.101.8:8080/ip"  -d '{99.9.0.1}'
func TestBenchmark_Must(t *testing.T) {
	city := new(CityANDCountry)
	city.Addr = "127.0.0.1:80800"
	city.Continent = "亚洲"
	city.Country = "中国"
	city.Subdivisions = "广东"
	city.Locationcity = "深圳市"
	city.TimeZone = "Asia/Shanghai"
	city.Latitude = "22.5333"
	city.Longitude = "114.1333"
	result := StructTOJson(city)
	if result == nil {
		t.Error("StructTOJson")
	}
	iP := MatchIP("121.1.2.2")
	if iP == "" {
		t.Error("IP Matching failed..")
	}
	/*
			//语言编码
		const (
			China     = "CN"  = "zh-CN"
			America    ="US" = "en"
			Japanese = "JP"   = "ja"
			Korea     = "KR"  = "ru"
			Filipino  = "PH"  = "fr"
			Brazil    = "BR"   = "pt-BR"
			Spain     = "ES"  = "es"
			Deutschland = "DE" = "de"
		)
	*/
	//if err := server01.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//	t.Error("测试没通过..")
	//}

}
func TestBenchmark_DealHander(t *testing.T) {

}
