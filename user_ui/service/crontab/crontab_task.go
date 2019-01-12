package crontab

import "fmt"

type Task interface {
	// 获取注单列表数据
	QueryRecord() (list interface{}, err error)

	// 注单列表数据写数据库
	RecordList2Db(list interface{}) error
}

type Instance func() Task

var adapter = make(map[string]Instance)

func Register(name string, task Instance) {
	if _, ok := adapter[name]; ok {
		panic("Task: Register called twice for adapter " + name)
	}
	adapter[name] = task
}

func NewTask(name string) (task Task, err error) {
	instanceFunc, ok := adapter[name]
	if !ok {
		err = fmt.Errorf("task: unknown adapter name %v (forgot to import?)", name)
		return
	}

	return instanceFunc(), nil
}
