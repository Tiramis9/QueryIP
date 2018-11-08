package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

func connectDb() *gorm.DB {
	dialect := fmt.Sprintf("juxing:MxLkdfwJdKzBRBPJ@tcp(192.168.255.41:3306)/juxing?charset=utf8&loc=Local&parseTime=true")
	db, err := gorm.Open("mysql", dialect)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err.Error())
	}

	db.SingularTable(true) //表名非复数形式
	return db
}

func TestConnectDb(t *testing.T) {
	db := connectDb()
	defer db.Close()

	if db == nil {
		t.Fatal("connect db error")
	}

	t.Log("connect db success")
}
