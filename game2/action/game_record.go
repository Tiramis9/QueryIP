package action

import (
	"game2/service/crontab"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GameRecordReq struct {
	Name string `json:"name" binding:"required"`
}

func GameRecord(c *gin.Context) {
	var req GameRecordReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}
	task, err := crontab.NewTask(req.Name)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	logrus.Infof("crontab task [%v]: get game record start.", req.Name)
	defer logrus.Infof("crontab task [%v]: get game record end.", req.Name)
	crontab.DInstance.Execute(task)
	RespSuccess(c)
}
