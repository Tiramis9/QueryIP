package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang_game_merchant/global"
	"time"
)

type BaseModel struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Db *gorm.DB

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

	if err = Db.DB().Ping(); err != nil {
		panic(err)
	}

	/*if err = autoCreateTables(Db); err != nil {
		panic(err)
	}*/
}

func TxBegin() *gorm.DB {
	return Db.Begin()
}

func TxCommit(tx *gorm.DB) {
	tx.Commit()
}

//自动建表
func autoCreateTables(db *gorm.DB) error {
	return db.Set("gorm: table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci").AutoMigrate(
		&MerchantWebsite{},
		&Message{},
		&User{},
		&UserBank{},
		&SysSecurityQuestion{},
		&UserAccount{},
		&UserBill{},
		&UserWithdraw{},
	).Error
}
