package model

import (
	"fmt"
	"game2/global"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type DbLogger struct {
	gorm.Logger
}

var Db *gorm.DB

func (l DbLogger) Print(values ...interface{}) {
	logrus.Info(values...)
}

func DbInit() {
	var err error
	conf := global.AppConfig
	source := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&loc=Local&parseTime=true", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	Db, err = gorm.Open(conf.DbDriver, source)
	if err != nil {
		panic(err)
	}

	Db.DB().SetMaxOpenConns(50)
	Db.DB().SetMaxIdleConns(50)
	Db.DB().SetConnMaxLifetime(10 * time.Second)
	Db.SingularTable(true) //表名非复数形式
	Db.SetLogger(DbLogger{})
	mysql.SetLogger(DbLogger{})

	if err = Db.DB().Ping(); err != nil {
		panic(err)
	}

	/*if err = autoCreateTables(Db); err != nil {
		panic(err)
	}*/
}
