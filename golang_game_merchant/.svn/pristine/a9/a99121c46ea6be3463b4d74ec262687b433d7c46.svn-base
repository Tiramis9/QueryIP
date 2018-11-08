package model

import (
	"testing"
)

func TestGetMerchantWebsiteReg(t *testing.T) {
	db := connectDb()
	defer db.Close()

	mw, err := GetMerchantWebsiteReg(db, 1)
	if err != nil {
		if err == ErrRecordNotFound {
			t.Log("not found")
		}

		t.Fatal(err)
	}

	t.Logf("%#v", mw)
}
