package model

import (
	"testing"
)

func TestAnnouncementADDcheck(t *testing.T) {
	db := connectDb()
	defer db.Close()
	msg := make(map[string]interface{})
	msg["type"] = "1"
	_, _, err := GetMerchantAnnouncementList(Db, 1, 2, msg)
	if err != nil {
		t.Fatal(err)
	}

}
