package action

import (
	"game2/global/status"
	"game2/model"

	"time"

	"github.com/gin-gonic/gin"
)

type MessageListReq struct {
	Page int `json:"page"`
	PageCount int `json:"page_count"`
}

type MessageListResp struct {
	Id         int    `json:"id"`
	MsgId      int    `json:"msg_id"`
	HaveRead   int    `json:"have_read"`
	CreateTime string `json:"create_time"`
	ReadTime   string `json:"read_time"`
	UserId     int    `json:"user_id"`
	Title      string `json:"title"`
}

type MessageInfoReq struct {
	Id int `json:"id" binding:"required"`
}
 
//消息列表
func MessageList(c *gin.Context) {
	var m MessageListReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m);err!=nil{
		RespParamErr(c)
		return
	}
	page,pageCount := InitPage(m.Page, m.PageCount)
	data := make(map[string]interface{})
	//获取列表
	messList,err := model.GetMessageList(model.Db,userId, page, pageCount)
	if err!=nil{
		RespServerErr(c)
		return
	}
	var dataList []MessageListResp
	for i:=range messList{
		var m MessageListResp
		m.Id = messList[i].Id
		m.MsgId = messList[i].MsgId
		m.HaveRead = messList[i].HaveRead
		m.CreateTime = messList[i].CreateTime
		m.ReadTime = messList[i].ReadTime
		m.UserId = messList[i].UserId
		m.Title = messList[i].Title
		dataList = append(dataList, m)
	}
	data["list"] = dataList
	//获取总数
	total,err := model.GetMessageCount(model.Db, userId)
	if err!=nil{
		RespServerErr(c)
		return
	}
	data["total"] = total
	RespJson(c,status.OK, data)
}

//消息详情
func MessageInfo(c *gin.Context) {
	var m MessageInfoReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m);err!=nil{
		RespParamErr(c)
		return
	}

	info,err := model.GetMessageInfo(model.Db, m.Id, userId)
	if err!=nil{
		RespServerErr(c)
		return
	}
	if info == nil{
		RespParamErr(c)
		return
	}
	if info.HaveRead == 0{
		model.ReadMessage(model.Db, m.Id, time.Now().Unix())
	}
	RespJson(c, status.OK, info)
}