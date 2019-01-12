//本包仅用于main函数初始化logrus和日志轮转
//main函数调用InitLog初始化完成后，其他包无需import本包，其他包只需要import logrus即可
package log

import (
	"game2/global"
	"github.com/sirupsen/logrus"
)

func InitLog() {
	fmtr := new(logrus.TextFormatter)
	fmtr.FullTimestamp = true                    // 显示完整时间
	fmtr.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	fmtr.DisableTimestamp = false                // 禁止显示时间
	fmtr.DisableColors = false                   // 禁止颜色显示

	hook := NewHook()
	hook.Field = "line"
	logrus.AddHook(hook)

	jack := &global.AppConfig.LogConf //日志路径，轮转，压缩等
	logrus.SetFormatter(fmtr)
	logrus.SetOutput(jack)
	logrus.SetLevel(logrus.DebugLevel)
}
