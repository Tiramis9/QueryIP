package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	geoip2 "github.com/oschwald/geoip2-golang"
	"public.info/config"
)

// 查询IP库返回详细信息
func GeoliteLookup(IP string) (*geoip2.City, error) {
	log := config.Log.WithField("package", "server")
	log.Info("start running GeoliteLookup ")
	log.Info(IP)
	ip := net.ParseIP(IP)
	record, err := config.DbHander.City(ip)

	if err != nil {
		log.Println(err)
		return record, errors.New("Query IP information does not exist!" + IP)
	}
	if record.City.GeoNameID == 0 {
		log.Println("Query IP information does not exist!:" + IP)
		return record, errors.New("Query IP information does not exist!:" + IP)
	}
	log.Info("GeoliteLookup running over")
	return record, nil
}

//根据所查询的ip地址，返回本地语言
func GeoipTOCity(City *geoip2.City, nativeLange string) (*CityANDCountry, error) {
	log := config.Log.WithField("package", "server")
	log.Info("start matching City proper ")
	if nativeLange == "" {
		log.Println("nativeLange is nil or Invalid")
		return nil, errors.New("nativeLange is nil or Invalid")
	}
	ContinEnt := City.Continent.Names[nativeLange]
	country := City.Country.Names[nativeLange]
	subdivisions := City.Subdivisions[0].Names[nativeLange]
	locationcity := City.City.Names[nativeLange]
	latitude := FloatTOJson(City.Location.Latitude)
	longitude := FloatTOJson(City.Location.Longitude)
	countryMessage := new(CityANDCountry)
	countryMessage.Continent = ContinEnt
	countryMessage.Country = country
	countryMessage.Subdivisions = subdivisions
	countryMessage.Locationcity = locationcity
	countryMessage.Latitude = string(latitude)
	countryMessage.Longitude = string(longitude)
	countryMessage.TimeZone = City.Location.TimeZone
	return countryMessage, nil

}

//获取客户端的ip，给当前客户端返回IP详细信息
func DealQuery(conn http.ResponseWriter, connMsg string) {
	log := config.Log.WithField("package", "server")
	log.Info("start running DealQuerycmd ")
	log.Infof("start matching information: [%s]", connMsg)
	IP, err := MatchIP(connMsg)
	if err != nil {
		fmt.Println(err, string(connMsg))
		fmt.Fprintf(conn, "%v ![%v]", err, string(connMsg))
		return
	}
	City, err := GeoliteLookup(IP)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(conn, err)
		return
	}
	switch {
	case City.Country.IsoCode == config.CN:
		log.Info("dealQuerycmd  Country CN(China) ")
		country, err := GeoipTOCity(City, config.China)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.US:
		log.Info("dealQuerycmd  Country US(America)")
		country, err := GeoipTOCity(City, config.America)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.PH:
		log.Info("dealQuerycmd  Country PH (Filipino) ")
		country, err := GeoipTOCity(City, config.Filipino)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.KR:
		log.Info("dealQuerycmd  Country KR(Korea) ")
		country, err := GeoipTOCity(City, config.Korea)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.JP:
		log.Info("dealQuerycmd  Country JP(Japanese) ")
		country, err := GeoipTOCity(City, config.Japanese)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.BR:
		log.Info("dealQuerycmd  Country BR(Brazil) ")
		country, err := GeoipTOCity(City, config.Brazil)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.ES:
		log.Info("dealQuerycmd  Country ES(Spain) ")
		country, err := GeoipTOCity(City, config.Spain)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	case City.Country.IsoCode == config.DE:
		log.Info("dealQuerycmd  Country DE(Deutschland) ")
		country, err := GeoipTOCity(City, config.Deutschland)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(conn, IP, err)
		}
		country.Addr = IP
		result := StructTOJson(country)
		fmt.Println(string(result))
		fmt.Fprintf(conn, string(result))
		break
	default:
		log.Info("dealQuerycmd  Country other ")
		fmt.Fprintf(conn, "Query IP information does not exist!:%v\n", string(IP))
		log.Info(City)
		log.Infof("Query IP information does not exist!:%v\n", string(IP))
		break
	}
}
