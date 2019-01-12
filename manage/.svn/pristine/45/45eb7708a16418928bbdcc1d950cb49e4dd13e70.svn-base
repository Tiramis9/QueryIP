package action

//{
//"user_name":"liul",
//"main_menu":"admin",
//"sub_menu":"login",
//"start_time":1510000222,
//"end_time":1550000222,
//"page":1,
//"page_count":2
//}

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)

type OptLogResp struct {
	List  []model.OperationLog `json:"list"`
	Total int                  `json:"total"`
}

func QueryOperationLog(c *gin.Context) {
	merchantID := 1 //todo token->id

	var req model.OperationLogReq
	err := c.Bind(&req)
	if err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)

	var resp OptLogResp
	resp.List, resp.Total, err = model.QueryOperationLog(model.Db, &req, merchantID)
	if err != nil {
		logrus.Errorf("model.QueryOperationLog(),%v", err)
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, resp)
}
