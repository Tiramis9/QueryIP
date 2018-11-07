package config

import (
	"time"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var (
	Path     = string("./log/")
	File     = string("std.log")
	MaxAge   = time.Hour * 24
	Interval = time.Second * 2000
)
var (
	//ADDRPORT = "172.16.200.21:8080"
	ADDRPORT = ":8080"
)

var (
	MmdbFile = "./GeoIP2_data/GeoLite2-City_20181023/GeoLite2-City.mmdb"
	DbHander *geoip2.Reader
)

//语言编码
const (
	China       = "zh-CN"
	America     = "en"
	Japanese    = "ja"
	Korea       = "ru"
	Filipino    = "fr"
	Brazil      = "pt-BR"
	Spain       = "es"
	Deutschland = "de"
)

//ip编码
const (
	CN = "CN"
	US = "US"
	PH = "PH"
	KR = "KR"
	JP = "JP"
	BR = "BR"
	ES = "ES"
	DE = "DE"
)
