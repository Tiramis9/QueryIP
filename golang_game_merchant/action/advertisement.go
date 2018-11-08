package action

import (
	"github.com/gin-gonic/gin"
	"golang_game_merchant/global/status"
)

func AdvertisementList(c *gin.Context) {
	type_code := c.PostForm("type")
	var data interface{}
	switch type_code {
	case "1":
		data = [...]map[string]interface{}{
			{"id": 1, "sort": 0, "name": "rng强无敌", "location": 1, "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "url": "www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "exist_time": 0, "status": 1, "type": 1},
			{"id": 2, "sort": 0, "name": "rng强无敌2", "location": 1, "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "url": "www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "exist_time": 0, "status": 1, "type": 1},
		}
	case "2":
		data = [...]map[string]interface{}{
			{"id": 3, "sort": 0, "name": "rng强无敌", "location": 1, "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "url": "www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "exist_time": 0, "status": 1, "type": 2},
			{"id": 4, "sort": 0, "name": "rng强无敌2", "location": 1, "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "url": "www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "exist_time": 0, "status": 1, "type": 2},
		}
	case "3":
		data = [...]map[string]interface{}{
			{"id": 5, "sort": 0, "name": "rng强无敌", "location": 1, "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "url": "www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "exist_time": 0, "status": 1, "type": 3},
			{"id": 6, "sort": 0, "name": "rng强无敌2", "location": 1, "image": "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1538135023137&di=5385fa88295bbe597762d47f80117f3e&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20170621%2Fa5b8c979d8674a098c2b10abd8d5ec31_th.jpg", "url": "www.baidu.com", "start_time": "1539601066", "end_time": "1539601066", "exist_time": 0, "status": 1, "type": 3},
		}
	default:
		RespParamErr(c)
		return
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	RespJson(c, status.OK, res)
}

func AdvertisementAdd(c *gin.Context) {
	RespSuccess(c)
}

func AdvertisementEdit(c *gin.Context) {
	RespSuccess(c)
}

func AdvertisementDel(c *gin.Context) {
	RespSuccess(c)
}
