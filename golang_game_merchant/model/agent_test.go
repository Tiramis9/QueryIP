package model

import (
	"testing"
)

func TestGetAgentList(t *testing.T) {
	db := connectDb()
	defer db.Close()

	list, err := GetAgentList(db, 1, "=", 1, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("%#v", m)
	}

	list, err = GetAgentList(db, 1, "=", -1, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("%#v", m)
	}

	list, err = GetAgentList(db, 1, ">=", 0, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf(">=%#v", m)
	}
}
