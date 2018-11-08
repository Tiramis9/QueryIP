package action

import (
	"github.com/gin-gonic/gin"
	"golang_game_merchant/global/status"
)

func AnnouncementList(c *gin.Context) {
	type_code := c.PostForm("type")
	var data interface{}
	switch type_code {
	case "1":
		data = [...]map[string]interface{}{
			{"id": "1", "sort": 0, "title": "rng强无敌", "content": "小组第一出线", "url": "http://www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "device": 1, "status": 1, "type": 1},
			{"id": "2", "sort": 0, "title": "rng强无敌2", "content": "小组第一出线2", "url": "http://www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "device": 1, "status": 1, "type": 1},
		}
	case "2":
		data = [...]map[string]interface{}{
			{"id": "3", "sort": 0, "title": "rng强无敌", "content": "小组第一出线", "url": "http://www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "device": 1, "status": 1, "type": 2},
			{"id": "4", "sort": 0, "title": "rng强无敌2", "content": "小组第一出线2", "url": "http://www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "device": 1, "status": 1, "type": 2},
		}
	default:
		RespParamErr(c)
		return
	}

	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	RespJson(c, status.OK, res)
}

func AnnouncementAdd(c *gin.Context) {
	RespSuccess(c)
}

func AnnouncementEdit(c *gin.Context) {
	RespSuccess(c)
}

func AnnouncementDel(c *gin.Context) {
	RespSuccess(c)
}
