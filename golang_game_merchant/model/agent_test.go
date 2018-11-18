package model

import (
	"testing"
)

func TestGetAgentList(t *testing.T) {
	db := connectDb()
	defer db.Close()

	where := map[string]interface{}{"user_name": "", "status": -1}
	list, err := GetAgentList(db, 1, 1, where, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("%#v", m)
	}

	where = map[string]interface{}{"user_name": "", "status": -1}
	list, err = GetAgentList(db, 1, 0, where, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("%#v", m)
	}
}

func TestGetAgentCount(t *testing.T) {
	db := connectDb()
	defer db.Close()

	where := map[string]interface{}{"user_name": "", "status": -1}
	list, err := GetAgentCount(db, 1, 1, where)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", list)
}

func TestUpdateAgent(t *testing.T) {
	db := connectDb()
	defer db.Close()

	//agent := Agent{Id:50016, MerchantId:2}
	where := map[string]interface{}{"id": 50016, "merchant_id": 1}
	err := UpdateAgent(db, where, map[string]interface{}{"status": -1})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", "success")
}

func TestGetAgentInfo(t *testing.T) {
	db := connectDb()
	defer db.Close()

	//agent := Agent{Id:50016, MerchantId:2}
	res, err := GetAgentInfo(db, 50010, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", res)
}
