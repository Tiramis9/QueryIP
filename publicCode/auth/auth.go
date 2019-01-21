package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang_game_merchant/action"
	"golang_game_merchant/global"
	"golang_game_merchant/global/redisclient"
	"golang_game_merchant/global/status"
	"io/ioutil"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignCheck(c *gin.Context) {
	//所有表单参数map
	form := make(map[string]interface{})
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err)
		action.RespParamErr(c)
		return
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &form)
		if err != nil {
			logrus.Error("auth32:", err)
			action.RespParamErr(c)
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	signs, ok := c.Request.Header["Sign"]
	if !ok || signs[0] == "" {
		logrus.Error("Sign not exist")
		action.RespParamErr(c)
		return
	}
	sign := signs[0]
	logrus.Debug(sign)

	appIds, ok := c.Request.Header["Appid"]
	if !ok || appIds[0] == "" {
		logrus.Error("Appid not exist")
		action.RespParamErr(c)
		return
	}
	appId := appIds[0]

	secrets, ok := c.Request.Header["Secret"]
	if !ok || secrets[0] == "" {
		logrus.Error("Secret not exist")
		action.RespParamErr(c)
		return
	}
	secret := secrets[0]

	//从redis中查询appid与secret是否匹配
	redisKey := "appid_" + appId
	//fmt.Println(redisKey)
	info, err := redisclient.RedisGetMap(redisKey)
	fmt.Println(info)
	if err != nil || len(info) == 0 || secret != info["secret"] {
		logrus.Error("secret:", secret)
		logrus.Error("info[secret]:", info["secret"])
		action.RespNoWebSiteInfo(c)
		return
	}
	merchId := info["merchant_id"]
	if int(info["background_status"].(float64)) == -1 { //三级后台网站关闭，(系统维护中)
		action.RespJson(c, status.ErrWebSiteMaintenance, nil)
		return
	}
	form["appid"] = appId
	form["secret"] = secret
	var strList []string
	for k, _ := range form {
		strList = append(strList, k)
	}
	sort.Strings(strList)
	var signStr string
	for _, val := range strList {
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

	//对字符串签名,与sign对比
	/*signSlice := signStr[0 : len(signStr)-1]
	md5Str := utils.Md5V(string(signSlice))
	//fmt.Println(md5Str)
	if md5Str != sign {
		logrus.Error("sign check failed")
		action.RespUnauthorized(c)
		return
	}*/
	c.Set("merchant_id", merchId)
	c.Set("ip", "127.0.0.1") //todo: 测试
}

func Auth(c *gin.Context) {
	tokenS, ok := c.Request.Header["Token"]
	if !ok {
		logrus.Error("no token")
		action.RespUnauthorized(c)
		return
	}
	token := tokenS[0]
	if token == "" {
		logrus.Error("token empty")
		action.RespUnauthorized(c)
		return
	}

	userInfo := redisclient.UserInfoByRedis(token)
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
	redisclient.RedisSet(token, string(jsonData), 3600)
	userId := userInfo["id"]
	c.Set("user_id", userId)
	//判断管理员是否为超管
	if int(userInfo["parent_id"].(float64)) == 0 {
		//当为超管时不处理
	} else {
		url := c.Request.URL.String()
		fmt.Println(url)
		//查询用户的角色id下的url Map
		authMap, err := redisclient.HgetFromRedis(global.RedisAuthKey, strconv.Itoa(int(userInfo["merchant_role_id"].(float64))))
		fmt.Println("authMap:", authMap)
		if err != nil {
			action.RespServerErr(c)
			return
		}
		/*_, ok = authMap[url]
		if !ok{//权限map无记录
			action.RespUnauthorized(c)
			return
		}*/
	}
}
