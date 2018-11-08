package action

import (
	"fmt"
	"golang_game_merchant/model"
	"net/http"
	//"reflect"
	"golang_game_merchant/lib/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//消息列表
func MessageList(c *gin.Context) {
	var data interface{}
	total := 0
	next_page := 1
	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	//user_id := 1
	user_id := userid.(int)
	page := c.PostForm("page")
	pagecount := c.PostForm("page_count")
	//检查page、pagecount是否为""
	page = utils.CheckEmptyStr(page, utils.DEFAULT_PAGE)
	pagecount = utils.CheckEmptyStr(pagecount, utils.DEFAULT_PAGECOUNT)
	//转为整形
	page_i, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	page_count_i, err := strconv.Atoi(pagecount)
	if err != nil {
		fmt.Println(err)
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	ch := make(chan int)
	defer close(ch)
	//获取列表
	go func() {
		messlist, _ := model.GetMessageList(model.Db, user_id, page_i, page_count_i)
		data = messlist
		ch <- 1
	}()
	//获取总数
	go func() {
		total, _ = model.GetMessageCount(model.Db, user_id)
		next_page = page_i + 1
		ch <- 1
	}()
	//等待通道数结束
	for i := 0; i < 2; i++ {
		<-ch
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": total, "next_page": next_page}
	c.JSON(http.StatusOK, res)
}

//读取消息
func MessageRead(c *gin.Context) {
	id := c.PostForm("id")
	id_i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "fail"})
		return
	}
	ok, _ := model.ReadMessage(model.Db, id_i, 1121212)
	res := gin.H{"code": "0", "msg": "fail"}
	if ok {
		res = gin.H{"code": "1", "msg": "ok"}
	}
	c.JSON(http.StatusOK, res)
}
