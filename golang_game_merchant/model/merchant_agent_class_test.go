package model

import (
	"testing"
)

func TestGetMerchantAgentClassList(t *testing.T) {
	db := connectDb()
	defer db.Close()

	list, err := GetMerchantAgentClassList(db, 1)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("%#v", m)
	}
}

func TestGetMerchantAgentClassInfo(t *testing.T) {
	db := connectDb()
	defer db.Close()

	info, err := GetMerchantAgentClassInfo(db, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", info)
}

func TestUpdateMerchantAgentClass(t *testing.T) {
	db := connectDb()
	defer db.Close()
	m := MerchantAgentClass{Id: 1}
	info, err := m.UpdateMerchantAgentClass(db, 1, map[string]interface{}{"mode": 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", info)
}

func TestDelMerchantAgentClass(t *testing.T) {
	db := connectDb()
	defer db.Close()

	m := MerchantAgentClass{Id: 2}
	info, err := m.DelMerchantAgentClass(db, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", info)
}
