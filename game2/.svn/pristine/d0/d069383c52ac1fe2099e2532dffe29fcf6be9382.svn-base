package crontab

import (
	"testing"
)

// test 匹配时间戳
func TestMatch(t *testing.T) {
	msg := `"\/Date(1543488845000)\/"`
	resp, err := MatchDate(msg) // ok
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}
func testCrontab(t *testing.T) {
	data := make(map[string]interface{})
	data["start_time"] = "2018-11-29 06:50:46"
	data["end_time"] = "2018-11-29 06:59:46"
	_, err := GameRecord("OG", data)
	//err := CronTabPrepare()
	if err != nil {
		t.Fatal(err)
	}
}

// test 匹配date
func TestMatchStringTime(t *testing.T) {
	s, e := newOGTimeDate()
	t.Log(s, e)

}

// test Crontab
func TestCrontabtext(t *testing.T) {

}
