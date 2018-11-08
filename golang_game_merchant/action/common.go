package action

import (
	"github.com/gin-gonic/gin"
	"golang_game_merchant/global/status"
)

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespJson(c *gin.Context, code int, data interface{}) {
	result := RestResponse{
		Code: code,
		Msg:  status.Description(code),
		Data: data,
	}

	c.JSON(status.OK, result)
	c.Abort()
}

func RespSuccess(c *gin.Context) {
	RespJson(c, status.OK, nil)
}

func RespParamErr(c *gin.Context) {
	RespJson(c, status.ErrParamError, nil)
}

func RespServerErr(c *gin.Context) {
	RespJson(c, status.ErrServerError, nil)
}

func RespNotFoundErr(c *gin.Context) {
	RespJson(c, status.ErrNotFound, nil)
}

func RespUnauthorized(c *gin.Context) {
	RespJson(c, status.ErrUnauthorized, nil)
}
