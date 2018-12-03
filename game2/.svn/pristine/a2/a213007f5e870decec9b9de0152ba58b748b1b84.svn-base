package action

import (
	"game2/global/status"
	"game2/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	AnnouncementInfo struct {
		Content    interface{} `json:"content"`
		Device     interface{} `json:"device"`
		StartTime  interface{} `json:"start_time"`
		EndTime    interface{} `json:"end_time"`
		Sort       interface{} `json:"sort"`
		Type       interface{} `json:"type"`
		Title      interface{} `json:"title"`
		Url        interface{} `json:"url"`
		Status     interface{} `json:"status"`
		MerchantID interface{} `json:"merchant_id"`
	}
	AnnouncementResponse struct {
		List  []AnnouncementInfo `json:"list"`
	}
)

func AnnouncementList(c *gin.Context) {
	id, ok := c.Get("merchant_id")
	if !ok {
		logrus.Error("merchant_id过期")
		RespServerErr(c)
		return
	}
	merchantId := int(id.(float64))
	msg := map[string]interface{}{"merchant_id":merchantId}
	list, err := model.GetMerchantAnnouncementList(model.Db,msg)
	if err != nil {
		RespServerErr(c)
		logrus.Error(err)
		return
	}
	resp := AnnouncementResponse{
		List:  make([]AnnouncementInfo, 0),
	}
	for i := range list {
		temp := AnnouncementInfo{
			Content:   list[i].Content,
			Device:    list[i].Device,
			StartTime: list[i].StartTime,
			EndTime:   list[i].EndTime,
			Url:       list[i].Url,
			Type:      list[i].Type,
			Title:     list[i].Title,
			Sort:      list[i].Sort,
			Status:    list[i].Status,
			MerchantID: list[i].MerchantID,
		}
		resp.List = append(resp.List, temp)
	}
	RespJson(c, status.OK, resp)
}