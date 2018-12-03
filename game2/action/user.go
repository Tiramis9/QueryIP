package action

import (
	"game2/global/status"
	"game2/lib/utils"
	"game2/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	//"github.com/gomodule/redigo/redis"
	"encoding/json"
	"game2/lib/redisclient"
	"game2/service"
)

type UserLoginReq struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Ip       string `json:"ip" binding:"required"`
}

type UserRegisterReq struct {
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RegIp      string `json:"reg_ip" binding:"required"`
	QuestionId int    `json:"question_id"`
	Answer     string `json:"answer"`
	TrueName   string `json:"true_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	PayPass    string `json:"pay_pass"`
	AgentCode  string `json:"agent_code"`
	//ParentId int `json:"parent_id"`
}

type UserBankAddReq struct {
	CardNo     string `json:"card_no"`
	BankBranch string `json:"bank_branch"`
	BankName   string `json:"bank_name"`
	PayPass    string `json:"pay_pass"` //资金密码
}

type UserEmailReq struct {
	Email string `json:"email" binding:"required"`
}

type UserQqReq struct {
	Qq string `json:"qq" binding:"required"`
}

type UserPhoneReq struct {
	AreaCode int    `json:"area_code" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UserPassReq struct {
	OldPass   string `json:"old_pass" binding:"required"`
	NewPass   string `json:"new_pass" binding:"required"`
	CheckPass string `json:"check_pass" binding:"required"`
}

type SetPayPassReq struct {
	PayPass string `json:"pay_pass" binding:"required"`
}

type UpdatePayPassReq struct {
	OldPaypass string `json:"old_paypass"`
	NewPaypass string `json:"new_paypass"`
}

//登录
func UserLogin(c *gin.Context) {
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
		return
	}
	merchId := int(merchantId.(float64))
	var m UserLoginReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	userName := m.UserName
	password := m.Password
	ip := m.Ip
	userLen := utf8.RuneCountInString(userName)
	passLen := utf8.RuneCountInString(password)
	if int(userLen) < 3 || int(userLen) > 15 {
		RespParamErr(c)
		return
	}
	if int(passLen) < 5 || int(passLen) > 20 {
		RespParamErr(c)
		return
	}
	//验证ip
	ipCheck := utils.RegexpMatch("ip", ip)
	if ipCheck == false {
		RespParamErr(c)
		return
	}
	//ToDo
	//检测ip地址详情
	//utils.HttpGet(global.AppConfig.IpUrl+"?ip="+ip)

	area := "广东省"
	//检查ip是否在黑名单中
	info, err := model.GetIpInfo(model.Db, merchId, ip, area)
	if err != nil {
		RespServerErr(c)
		return
	}
	if info != nil { //ip被禁止
		RespJson(c, status.ErrIpBanned, nil)
		return
	}

	user, err := model.GetUserByName(model.Db, userName)
	if err != nil {
		RespServerErr(c)
		return
	}
	if user == nil {
		RespNotFoundErr(c)
		return
	}
	//检查用户状态
	if user.Status != 1 {
		logrus.Error("用户状态", user.Status)
		RespUnauthorized(c)
		return
	}

	//ex_time := 3600
	//盐值和传递的密码md5的加密
	checkStr := utils.Md5V(utils.Md5V(password + string(user.Salt)))
	//fmt.Println(checkStr)
	if checkStr != user.Password {
		RespUserOrPassErr(c)
		return
	}

	//生成token
	randStr := string(utils.Krand(16, utils.KC_RAND_KIND_ALL)) //3表示随机数包含数字、大小写字母
	tokenV := strconv.Itoa(user.Id)
	str := randStr + tokenV
	token := utils.Md5V(str)

	data := make(map[string]interface{})
	data["token"] = token
	data["user_name"] = user.UserName
	data["true_name"] = user.TrueName
	data["phone"] = user.Phone
	data["email"] = user.Email
	data["qq"] = user.QQ
	data["skype"] = user.Skype
	data["lang"] = user.Lang
	data["time_zone"] = user.TimeZone
	data["birthday"] = user.Birthday
	data["area_code"] = user.AreaCode
	data["device"] = user.Device
	data["source"] = user.Source
	data["class_id"] = user.ClassId
	data["group_id"] = user.GroupId
	data["last_login_time"] = user.LastLoginTime
	data["last_login_ip"] = user.LastLoginIp
	data["id"] = user.Id
	data["merchant_id"] = user.MerchantId
	data["login_ip"] = user.LoginIp
	data["login_time"] = user.LoginTime
	jsonData, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	err = service.RedisSet(token, string(jsonData), utils.LOGIN_EXPIRED_TIME)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	data["balance"] = user.Balance
	//检查商户预警策略
	loginIpWarn, err := model.GetMerchantWarn(model.Db, merchId, "last_login_ip")
	if err != nil {
		RespServerErr(c)
		return
	}
	if loginIpWarn != nil { //无ip警告
		if loginIpWarn.SysStatus == 1 && loginIpWarn.Status == 1 { //ip警告开启
			//检测上次登录ip地区
			//ToDo
		}
	}
	//更新登录时间,ip
	loginTime := time.Now().Unix()
	loginIp := ip
	lastLoginTime := user.LoginTime
	if err != nil {
		RespServerErr(c)
		return
	}
	lastLoginIp := user.LoginIp
	updateRes, err := model.UpdateUserLoginInfo(model.Db, user.Id, loginTime, loginIp, lastLoginTime, lastLoginIp)
	if updateRes == false {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, data)
}

//获取注册渲染页面配置
func UserRegisterPageInfo(c *gin.Context) {
	//获取商户id
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
		return
	}
	merchId := int(merchantId.(float64))
	merchant, err := model.GetMerchantWebsiteReg(model.Db, merchId)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, merchant)
}

//注册
func UserRegister(c *gin.Context) {
	tx := model.Db.Begin()
	//var data interface{}
	var u UserRegisterReq
	if err := c.BindJSON(&u); err != nil {
		RespParamErr(c)
		return
	}
	//获取商户的网站配置信息
	merchantId, ok := c.Get("merchant_id")
	if !ok {
		RespServerErr(c)
	}
	merchId := int(merchantId.(float64))
	merchant, err := model.GetMerchantWebsiteReg(tx, merchId)
	if err != nil {
		RespServerErr(c)
		return
	}
	questionId := u.QuestionId
	answerLen := utf8.RuneCountInString(u.Answer)
	trueNameLen := utf8.RuneCountInString(u.TrueName)
	phoneLen := utf8.RuneCountInString(u.Phone)
	emailLen := utf8.RuneCountInString(u.Email)
	payPassLen := utf8.RuneCountInString(u.PayPass)

	regIpCheck := utils.RegexpMatch("ip", u.RegIp)
	if regIpCheck == false {
		RespParamErr(c)
		return
	}

	if merchant.RegPayPass == 3 && payPassLen < 1 {
		RespParamErr(c)
		return
	}

	if merchant.RegSecurityQuestion == 3 {
		if questionId == 0 || answerLen < 1 {
			RespParamErr(c)
			return
		}
		//查看QuestionId是否正确
		res, err := model.GetSecurityInfo(tx, u.QuestionId)
		if err != nil {
			RespServerErr(c)
			return
		}
		if res == nil {
			RespParamErr(c)
			return
		}
	}
	if merchant.RegTrueName == 3 && trueNameLen < 1 {
		RespParamErr(c)
		return
	}
	if merchant.RegPhone == 3 && phoneLen < 1 {
		RespParamErr(c)
		return
	}
	if merchant.RegEmail == 3 && emailLen < 1 {
		RespParamErr(c)
		return
	}

	userLen := utf8.RuneCountInString(u.UserName)
	passLen := utf8.RuneCountInString(u.Password)

	if int(userLen) < 3 || int(userLen) > 15 {
		RespParamErr(c)
		return
	}
	if int(passLen) < 5 || int(passLen) > 20 {
		RespParamErr(c)
		return
	}

	//查看用户是否注册
	user, err := model.GetUserByName(tx, u.UserName)
	if err != nil {
		RespNotFoundErr(c)
		return
	}
	if user != nil {
		RespUserExistErr(c)
		return
	}
	var userModel model.User
	//密码
	salt := string(utils.Krand(4, utils.KC_RAND_KIND_ALL)) //3表示随机数包含数字、大小写字母
	str := u.Password + salt
	passwordMd5 := utils.Md5V(utils.Md5V(str))

	//支付密码
	if u.PayPass != "" {
		str2 := u.PayPass + salt
		payPassMd5 := utils.Md5V(utils.Md5V(str2))
		userModel.PayPass = payPassMd5
	}

	timestamp := time.Now().Unix()

	//ToDo
	if u.AgentCode != "" {
		// parentId
		agent, err := model.GetAgentByAgentCode(tx, u.AgentCode)
		if err != nil {
			RespServerErr(c)
			return
		}
		//代理码错误
		if agent == nil {
			RespJson(c, status.ErrAgentCodeError, nil)
			return
		}
		parentId := agent.Id
		userModel.ParentId = parentId
		userModel.AgentCode = u.AgentCode
	}
	userModel.MerchantId = merchId
	userModel.UserName = u.UserName
	userModel.TrueName = u.TrueName
	userModel.Password = passwordMd5
	userModel.Phone = u.Phone
	userModel.Email = u.Email
	userModel.Salt = salt
	userModel.RegTime = timestamp
	userModel.RegIp = u.RegIp
	userModel.Status = 1

	id, err := userModel.InsertUser(tx)
	if err != nil {
		tx.Rollback()
		RespServerErr(c)
		return
	}
	if merchant.RegSecurityQuestion == 3 {
		//添加密保问题
		var ua = model.UserAnswers{
			UserId:     id,
			MerchantId: merchId,
			QuestionId: u.QuestionId,
			Answer:     u.Answer,
			CreateTime: timestamp,
			UpdateTime: timestamp,
		}
		_, err := ua.SetSecurity(tx)
		if err != nil {
			tx.Rollback()
			RespServerErr(c)
			return
		}
	}
	tx.Commit()
	RespJson(c, status.OK, id)
}

//登出
func UserLogout(c *gin.Context) {
	token := c.PostForm("token")

	//连接redis
	conn := redisclient.Get()

	//记得销毁本次链连接
	defer conn.Close()
	_, err := conn.Do("DEL", token)
	if err != nil {
		res := gin.H{"code": 0, "data": nil, "msg": utils.ERR_SYSTEM_ERROR}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{"code": 1, "data": nil, "msg": "ok"}
	c.JSON(http.StatusOK, res)
	return

}

//用户信息
func UserInfo(c *gin.Context) {
	resp := make(map[string]interface{})
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	user, err := model.GetUserById(model.Db, userId)
	resp["id"] = user.Id
	resp["merchant_id"] = user.MerchantId
	resp["user_name"] = user.UserName
	resp["true_name"] = user.TrueName
	resp["phone"] = user.Phone
	resp["email"] = user.Email
	resp["qq"] = user.QQ
	resp["birthday"] = user.Birthday
	resp["last_login_ip"] = user.LastLoginIp
	resp["last_login_time"] = user.LastLoginTime
	resp["balance"] = user.Balance
	if user.PayPass!=""{
		resp["set_paypass"] = 1
	}else{
		resp["set_paypass"] = 0
	}
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, resp)
}

//绑定邮箱
func UserBindEmail(c *gin.Context) {
	var m UserEmailReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	//检查邮箱格式
	regCheck := utils.RegexpMatch("email", m.Email)
	if regCheck == false {
		RespParamErr(c)
		return
	}
	fields := make(map[string]interface{})
	fields["email"] = m.Email
	_, err := model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

//绑定QQ
func UserBindQQ(c *gin.Context) {
	var m UserQqReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	regCheck := utils.RegexpMatch("qq", m.Qq)
	if regCheck == false {
		RespParamErr(c)
		return
	}
	fields := make(map[string]interface{})
	fields["qq"] = m.Qq
	_, err := model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

//绑定手机
func UserBindPhone(c *gin.Context) {
	var m UserPhoneReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	phone := m.Phone
	areaCode := m.AreaCode
	tokenS, ok := c.Request.Header["Token"]
	if !ok {
		RespParamErr(c)
		return
	}
	token := tokenS[0]
	phoneLen := utf8.RuneCountInString(phone)
	if phoneLen < 1 {
		RespParamErr(c)
		return
	}

	//ToDo
	//验证areaCode准确性

	//验证手机号与区号是否存在
	phoneUser, err := model.GetUserByPhone(model.Db, phone, areaCode)
	if err != nil {
		RespServerErr(c)
		return
	}
	if phoneUser != nil {
		RespJson(c, status.ErrPhoneExist, nil)
		return
	}
	fields := make(map[string]interface{})
	fields["area_code"] = m.AreaCode
	fields["phone"] = m.Phone
	_, err = model.UpdateUser(model.Db, userId, fields)

	if err != nil {
		RespServerErr(c)
		return
	} else {
		//获取redis的用户数据
		userInfo, error := service.RedisGetMap(token)
		if error != nil {
			RespServerErr(c)
			return
		}
		userInfo["area_code"] = areaCode
		userInfo["phone"] = phone

		jsonData, err := json.Marshal(userInfo)
		if err != nil {
			RespServerErr(c)
			return
		}
		err = service.RedisSet(token, string(jsonData), utils.LOGIN_EXPIRED_TIME)
		if err != nil {
			RespServerErr(c)
			return
		}
		RespSuccess(c)
	}
}

//修改密码
func UpdatePass(c *gin.Context) {
	var m UserPassReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	//检查新密码与确认密码
	if m.NewPass != m.CheckPass {
		RespParamErr(c)
		return
	}

	user, err := model.GetUserById(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	//盐值和传递的密码md5的加密
	str := m.OldPass + string(user.Salt)
	checkStr := utils.Md5V(utils.Md5V(str))
	if checkStr != user.Password {
		RespJson(c, status.ErrPassError, nil)
		return
	}
	//生成新密码
	strNew := m.NewPass + string(user.Salt)
	newPass := utils.Md5V(utils.Md5V(strNew))
	fields := make(map[string]interface{})
	fields["password"] = newPass
	_, err = model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

//修改资金密码
func UpdatePayPass(c *gin.Context){
	var m UpdatePayPassReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	user, err := model.GetUserById(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	//盐值和传递的密码md5的加密
	str := m.OldPaypass + string(user.Salt)
	checkStr := utils.Md5V(utils.Md5V(str))
	if checkStr != user.PayPass {
		RespJson(c, status.ErrPassError, nil)
		return
	}
	//生成新密码
	strNew := m.NewPaypass + string(user.Salt)
	newPass := utils.Md5V(utils.Md5V(strNew))
	fields := make(map[string]interface{})
	fields["pay_pass"] = newPass
	_, err = model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)

}

//设置资金密码
func SetPayPass(c *gin.Context) {
	var m SetPayPassReq
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	if err := c.BindJSON(&m); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	user, err := model.GetUserById(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	if len(user.PayPass) > 0 {
		RespJson(c, status.ErrPayPassHadSet, nil)
		return
	}
	//生成新密码
	strNew := m.PayPass + string(user.Salt)
	newPass := utils.Md5V(utils.Md5V(strNew))
	fields := make(map[string]interface{})
	fields["pay_pass"] = newPass
	_, err = model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

func UserBankList(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	bankList, err := model.GetUserBankList(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	res := map[string]interface{}{"list": bankList}
	RespJson(c, status.OK, res)
}

func UserBankAdd(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		logrus.Error(c)
		RespServerErr(c)
		return
	}
	var ub UserBankAddReq
	if err := c.BindJSON(&ub); err != nil {
		RespParamErr(c)
		return
	}
	userId := int(id.(float64))
	//取代理真实姓名
	user, err := model.GetUserById(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	if user.TrueName == "" { //未实名认证,则重新不允许添加银行卡
		RespJson(c, status.ErrNoTrueName, nil)
		return
	}
	//未设置资金密码
	if user.PayPass == "" { //未设置资金密码,则不允许添加银行卡
		RespJson(c, status.ErrNoPayPass, nil)
		return
	}
	//判断资金密码
	checkStr := utils.Md5V(utils.Md5V(ub.PayPass + string(user.Salt)))
	//fmt.Println(checkStr)
	if checkStr != user.PayPass {
		RespJson(c, status.ErrPayPassError, nil)
		return
	}
	nowTime := time.Now().Unix()
	trueName := user.TrueName
	userBankModel := model.UserBank{
		UserId:     userId,
		TrueName:   trueName,
		CardNo:     ub.CardNo,
		BankName:   ub.BankName,
		BankBranch: ub.BankBranch,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	re, err := userBankModel.AddUserBank(model.Db)
	if err != nil {
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, re)
}
