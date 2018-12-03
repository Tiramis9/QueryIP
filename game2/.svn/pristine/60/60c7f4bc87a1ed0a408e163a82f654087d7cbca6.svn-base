package crontab

import "github.com/sirupsen/logrus"

type Director struct{}

func newDirector() *Director {
	return &Director{}
}

func (d *Director) Execute(task Task) {
	// 获取注单记录
	recordList, err := task.QueryRecord()
	if err != nil {
		logrus.Error(err)
		return
	}

	// 注单记录写数据库
	if err := task.RecordList2Db(recordList); err != nil {
		logrus.Error(err)
		return
	}
}

var DInstance = newDirector()
