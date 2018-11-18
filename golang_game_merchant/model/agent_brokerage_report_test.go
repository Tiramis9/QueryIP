package model

import (
	"testing"
	"time"
)

func TestGetAgentBrokerageReportList(t *testing.T) {
	db := connectDb()
	defer db.Close()
	where := make(map[string]interface{})
	where["start_time"] = 0
	where["end_time"] = time.Now().Unix()
	list, err := GetAgentBrokerageReportList(db, 1, where, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("1%#v", m)
	}

	list, err = GetAgentBrokerageReportList(db, 1, where, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("2%#v", m)
	}
}

func TestUpdateAgentBrokerageReport(t *testing.T) {
	db := connectDb()
	defer db.Close()
	fields := make(map[string]interface{})
	timestamp := time.Now().Unix()
	var abReport AgentBrokerageReport
	abReport.Id = 1
	fields["deal_time"] = timestamp
	fields["org_expense"] = 1000
	fields["brokerage_total"] = 10000
	info, err := abReport.UpdateAgentBrokerageReport(db, 1, fields)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("1%#v", info)
}

func TestGetAgentReportList(t *testing.T) {
	db := connectDb()
	defer db.Close()
	list, err := GetAgentReportList(db, 1, 0, time.Now().Unix(), 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("1%#v", m)
	}

	//总条数
	count, err := GetAgentReportCount(db, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("2[%#v]", count)
}

func TestGetSubUserReportList(t *testing.T) {
	db := connectDb()
	defer db.Close()

	//列表
	list, err := GetSubUserReportList(db, 50010, 1, 0, time.Now().Unix(), 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range list {
		t.Logf("1[%#v]", m)
	}

	//总条数
	count, err := GetSubUserReportCount(db, 50010, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("2[%#v]", count)

	//总计
	info, err := GetSubUserReportTotal(db, 50010, 1, 0, time.Now().Unix())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("3[%#v]", info)
}
