package action

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)

type GameDescriptionReq struct {
	model.GameDescription
	//Token int `json:"token"`
}

func QueryMerchantGameStatus(c *gin.Context) {
	merchantID := 1
	gds, err := model.QueryMerchantGameStatus(model.Db, merchantID)
	if err != nil {
		RespServerErr(c)
		logrus.Errorf("model.QueryMerchantGameStatus(), %v", err)
		return
	}

	RespJson(c, status.OK, gds)
}

// Post: {"game_id":1,"status":0}
func ModifyMerchantGameStatus(c *gin.Context) {
	var req GameDescriptionReq
	if err := c.Bind(&req); err != nil {
		logrus.Errorf("c.Bind(),%v", err)
		RespParamErr(c)
		return
	}

	if req.Status != 1 && req.Status != 2 {
		logrus.Error("post status should be 1 or 2")
		RespParamErr(c)
		return
	}

	merchantID := 1
	err := model.ModifyMerchantGameStatus(model.Db, &(req.GameDescription), merchantID)
	if err != nil {
		RespServerErr(c)
		logrus.Errorf("model.ModifyMerchantGameStatus(), %v", err)
		return
	}

	RespSuccess(c)
}
