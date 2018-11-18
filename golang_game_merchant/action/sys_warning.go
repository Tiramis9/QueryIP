package action

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)

type warningLogReq struct {
	//Token     string `json:"token"`
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

type warningLogRESP struct {
	List  []model.WarnLog `json:"list"`
	Total int             `json:"total"`
}

type warningPolicyReq struct {
	model.WarningPolicy
	//Token string `json:"token"`
}

type warningPolicyRESP struct {
	List  []model.WarningPolicy `json:"list"`
}

type sysWarningPolicyRESP struct {
	List []model.SysWarningPolicy `json:"list"`
}


// Post: {"page":1,"page_count":3}
func QueryWarningLog(c *gin.Context) {
	merchantID := 1

	var req warningLogReq
	err := c.Bind(&req)
	if err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)

	var resp warningLogRESP
	resp.List, resp.Total, err = model.QueryWarningLog(model.Db, merchantID, req.Page, req.PageCount)
	if err != nil {
		logrus.Errorf("model.QueryWarningLog(),%v", err)
		RespServerErr(c)
		return
	}

	RespJson(c, status.OK, resp)
}

func QueryWarningPolicy(c *gin.Context) {
	var req warningPolicyReq //取token
	err := c.Bind(&req)
	if err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	merchantID := 1

	var resp  warningPolicyRESP
	resp.List, err = model.QueryWarningPolicy(model.Db, merchantID)
	if err != nil {
		RespServerErr(c)
		return
	}

	RespJson(c, status.OK, resp)
}

func QuerySysWarningPolicy(c *gin.Context) {
	var resp  sysWarningPolicyRESP
	var err error
	resp.List, err = model.QuerySysWarningPolicy(model.Db)
	if err != nil {
		RespServerErr(c)
		return
	}

	RespJson(c, status.OK, resp)
}



func ModifyWarningPolicy(c *gin.Context) {
	var req warningPolicyReq //取token
	err := c.Bind(&req)
	if err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	//校验参数
	if req.Status != 1 && req.Status != -1 {
		logrus.Error("action.ModifyWarningPolicy(),invalid argument 'status'")
		RespParamErr(c)
		return
	}
	
	merchantID := 1

	err = model.ModifyWarningPolicy(model.Db, merchantID, &(req.WarningPolicy))
	if err != nil {
		logrus.Errorf("model.ModifyWarningPolicy(),%v", err)
		RespServerErr(c)
		return
	}

	RespSuccess(c)
}


func AddWarningPolicy(c *gin.Context){
	var req warningPolicyReq //取token
	err := c.Bind(&req)
	if err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	//参数校验
	if req.SysWarningId <= 0 {
		logrus.Error("action.AddWarningPolicy(), invalid sys_warning_id")
		RespParamErr(c)
		return
	}

	merchantID := 1
	if err:= model.AddWarningPolicy(model.Db,merchantID,req.SysWarningId,req.Value); err != nil {
		logrus.Errorf("model.AddWarningPolicy(),%v", err)
		RespServerErr(c)
		return
	}

	RespSuccess(c)
}

func DelWarningPolicy(c *gin.Context){
	var req warningPolicyReq //取token
	err := c.Bind(&req)
	if err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	//参数校验
	if req.SysWarningId <= 0 {
		logrus.Error("action.DelWarningPolicy(), invalid sys_warning_id")
		RespParamErr(c)
		return
	}

	merchantID := 1
	if err:= model.DelWarningPolicy(model.Db,merchantID,req.SysWarningId); err != nil {
		logrus.Errorf("model.DelWarningPolicy(),%v", err)
		RespServerErr(c)
		return
	}

	RespSuccess(c)
}
