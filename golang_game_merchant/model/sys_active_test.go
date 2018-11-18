package model

import (
	"testing"
)

func TestGetPayActiveInfo(t *testing.T) {
	db := connectDb()
	defer db.Close()

	info, err := GetPayActiveInfo(db, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", info)
}
