package action

import (
	"github.com/gin-gonic/gin"
	"golang_game_merchant/model"
	"net/http"
	"strconv"
)

//获取下载url
func MerchantWebsiteAppDownload(c *gin.Context) {
	//获取商户id
	merchantIdStr := c.PostForm("merchant_id")
	merchantId, err := strconv.Atoi(merchantIdStr)
	if err != nil {
		panic(err)
	}
	merchant, err := model.GetMerchantWebsiteAppDownload(model.Db, merchantId)
	if err != nil {
		if err == model.ErrRecordNotFound {
			res := gin.H{"code": 0, "data": nil, "msg": "Merchant does not exist"}
			c.JSON(http.StatusOK, res)
			return
		}

	}

	download_url := map[string]string{
		"Android": merchant.AppDownloadUrl,
		"IOS":     merchant.AppDownloadUrl,
	}

	res := gin.H{"code": 1, "data": download_url, "msg": "ok"}
	c.JSON(http.StatusOK, res)

}
