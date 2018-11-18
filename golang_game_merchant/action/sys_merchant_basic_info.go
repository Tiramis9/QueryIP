package action

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)


func QueryMerchantBasicInfo(c *gin.Context) {

	merchantID := 1

	var err error
	var resp model.BasicInfo
	resp, err = model.QueryMerchantBasicInfo(model.Db, merchantID)
	if err != nil {
		logrus.Errorf("model.QueryMerchantBasicInfo,%v\n", err)
		RespServerErr(c)
		return
	}

	RespJson(c, status.OK, resp)
}

//{"withdraw_single_max":"withdraw_max","withdraw_single_min":"withdraw_min","agent_plat_url":"a.com/plat",
//"agent_spread_url":"a.com/spread","app_download_url":"a.com/app_dld","service_online_url":"a.com/osrv",
//"app_logo":"logo_A","allow_ip_minute":10,"active_status":1,"reg_status":1}
func ModifyMerchantBasicInfo(c *gin.Context) {
	var req model.BasicInfo
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	//todo  校验,记得token
	if req.AgentPlatUrl == "" || req.AgentSpreadUrl == "" || req.AppDownloadUrl == "" ||
		req.AppLogo == "" || req.ServiceOnlineUrl == "" {
		logrus.Errorf("one of parameter is empty,%v", req)
		RespParamErr(c)
		return
	}

	merchantID := 1
	err := model.ModifyMerchantBasicInfo(model.Db, &req, merchantID)
	if err != nil {
		logrus.Errorf("ModifyMerchantBasicInfo(),%v", err)
		RespServerErr(c)
		return
	}

	RespSuccess(c)
}
