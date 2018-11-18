package action

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)

type MemberRegistrationReq struct {
	model.MemberRegistration
	//Token int `json:"token"`
}

func QueryMemberRegistration(c *gin.Context) {
	merchantID := 1
	resp, err := model.QueryMemberRegistration(model.Db, merchantID)
	if err != nil {
		logrus.Errorf("model.QueryMemberRegistration(), %v", err)
		RespServerErr(c)
		return
	}

	RespJson(c, status.OK, resp)
}

func ModifyMemberRegistration(c *gin.Context) {
	var req MemberRegistrationReq
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("ModifyMemberRegistration,c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	for field, value := range map[string]int{"RegEmail": req.RegEmail,
		"RegPayPass":          req.RegPayPass,
		"RegPhone":            req.RegPhone,
		"RegSecurityQuestion": req.RegSecurityQuestion,
		"RegTrueName":         req.RegTrueName} {
		if value < 1 || value > 3 {
			logrus.Errorf("invalid arguments, action.ModifyMemberRegistration(),invalid value of '%s'", field)
			RespParamErr(c)
			return
		}
	}

	merchantID := 1
	err := model.ModifyMemberRegistration(model.Db, &(req.MemberRegistration), merchantID)
	if err != nil {
		logrus.Errorf(" model.ModifyMemberRegistration(), %v", err)
		RespParamErr(c)
		return
	}

	RespSuccess(c)
}
