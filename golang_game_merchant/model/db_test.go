package model

import (
	"testing"
)

func TestConnectDb(t *testing.T) {
	db := connectDb()
	defer db.Close()

	if db == nil {
		t.Fatal("connect db error")
	}

	t.Log("connect db success")
}

func TestGormUpdate(t *testing.T) {
	db := connectDb()
	defer db.Close()

	ndb := db.Debug().Table(`user_rate`).Where(`id=?`, 1).Updates(map[string]interface{}{
		"user_id":     10,
		"merchant_id": 9,
	})
	row, err := ndb.RowsAffected, ndb.Error
	t.Log(row)
	t.Log(err)
}
