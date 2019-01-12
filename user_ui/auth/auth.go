package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"game2/action"
	"game2/lib/utils"
	"game2/logic"
	"game2/service"
	"io/ioutil"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignCheck(c *gin.Context) {
	var signStr string
	strs := []string{}
	//所有表单参数map
	form := make(map[string]interface{})
	body, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		action.RespParamErr(c)
		return
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &form)
		if err != nil {
			logrus.Error("auth32:", err)
			action.RespServerErr(c)
			return
		}
	}
	//fmt.Println(c.Request.Header)
	signS, ok := c.Request.Header["Sign"]
	if !ok || signS[0] == "" {
		action.RespParamErr(c)
		return
	}
	sign := signS[0]
	fmt.Println(sign)
	appidS, ok := c.Request.Header["Appid"]
	if !ok || appidS[0] == "" {
		action.RespParamErr(c)
		return
	}
	appid := appidS[0]
	secretS, ok := c.Request.Header["Secret"]
	if !ok || secretS[0] == "" {
		action.RespParamErr(c)
		return
	}
	secret := secretS[0]
	//从redis中查询appid与secret是否匹配
	redisKey := "appid_" + appid
	//fmt.Println(redisKey)
	info, err := service.RedisGetMap(redisKey)
	if err != nil || secret != info["secret"] {
		logrus.Error("secret:", secret)
		logrus.Error("info[secret]:", info["secret"])
		action.RespUnauthorized(c)
		return
	}
	merchId := info["merchant_id"]
	form["appid"] = appid
	form["secret"] = secret
	for k, _ := range form {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	for _, val := range strs {
		var str string
		switch form[val].(type) {
		case int:
			in := form[val].(int)
			str = strconv.Itoa(in)
		case string:
			str = form[val].(string)
		case float64:
			fl := form[val].(float64)
			str = strconv.FormatFloat(fl, 'E', -1, 64)
		}
		signStr += val + "=" + str + "&"
	}
	//fmt.Println(signStr)
	//对字符串签名,与sign对比
	/*signSlice := signStr[0 : len(signStr)-1]
	md5Str := utils.Md5V(string(signSlice))
	//fmt.Println(md5Str)
	if md5Str != sign {
		utils.Log("sgin error", "debug","")
		action.RespUnauthorized(c)
		return
	}*/
	c.Set("merchant_id", merchId)
	fmt.Println("signCheck")
}

//Token中间件
func TokenRequired(c *gin.Context) {
	tokenS, ok := c.Request.Header["Token"]
	if !ok {
		logrus.Error("no token")
		action.RespParamErr(c)
		return
	}
	token := tokenS[0]
	if token == "" {
		logrus.Error("token empty")
		action.RespParamErr(c)
		return
	}
	userInfo := logic.UserInfoByRedis(token)
	if userInfo == nil {
		logrus.Error("token expired")
		action.RespTokenExpiredErr(c)
		return
	}
	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		action.RespServerErr(c)
		return
	}
	//用户信息存在,延长在线时间
	service.RedisSet(token, string(jsonData), utils.LOGIN_EXPIRED_TIME)
	userId := userInfo["id"]
	c.Set("user_id", userId)
	fmt.Println("TokenRequired")
}
