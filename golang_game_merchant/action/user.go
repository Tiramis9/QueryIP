package action

import (
	"crypto/md5"
	//"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/lib/utils"
	"golang_game_merchant/model"
	"io"
	"net/http"
	//utils2 "src/golang_game_merchant/lib/utils"
	"strconv"
	//"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	//"github.com/gomodule/redigo/redisclient"
	"golang_game_merchant/global/redisclient"
)

type (
	//会员列表
	MemberListReq struct {
		ClassId        int    `json:"class_id"`                  //会员层级(渠道)
		GroupId        int    `json:"group_id"`                  //会员等级
		Status         int    `json:"status" binding:"required"` //账号状态 1启用. 0.禁用 必填
		TrueName       string `json:"true_name"`                 //会员姓名
		UserName       string `json:"user_name"`                 //会员账号
		Phone          string `json:"phone,omitempty"`           //会员手机号
		ParentUserName string `json:"parent_user_name"`          //上级代理账号
		Page           int    `json:"page"`                      //页码
		PageCount      int    `json:"page_count"`                //每页显示的数量
	}

	MemberList struct {
		ClassName      interface{} `json:"class_name"`       //会员层级(渠道)
		GroupName      interface{} `json:"group_name"`       //会员等级
		Id             interface{} `json:"id"`               //记录id
		TrueName       interface{} `json:"true_name"`        //会员姓名
		UserName       interface{} `json:"user_name"`        //会员账号
		ParentUserName interface{} `json:"parent_user_name"` //上级代理账号
		Balance        interface{} `json:"balance"`          //账户余额
		LastLoginIp    interface{} `json:"last_login_ip"`    //最后登录ip
		LastLoginTime  interface{} `json:"last_login_time"`  //最后登录时间
		RegTime        interface{} `json:"reg_time"`         //注册时间
		RegIp          interface{} `json:"reg_ip"`           //注册ip
		Phone          interface{} `json:"phone"`            //会员手机号
		Email          interface{} `json:"email"`            //邮箱
		FirstLoginIp   interface{} `json:"first_login_ip"`   //首次登录ip
		FirstLoginTime interface{} `json:"first_login_time"` //首次登录时间
	}

	MemberListResp struct {
		List  []MemberList `json:"list"`  //会员列表
		Total interface{}  `json:"total"` //总数
	}
)

type (
	//编辑用户信息
	UserReq struct {
		Id             int    `json:"id" binding:"required"`
		TrueName       string `json:"true_name"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
		Birthday       int    `json:"birthday"`
		Password       string `json:"password"`        //登录密码
		PayPass        string `json:"pay_pass"`        //资金密码
		Status         int    `json:"status"`          //账号状态 状态 1启用. 0.禁用
		Sex            int    `json:"sex"`             //性别 0.未知；1.男；2，女
		ClassId        int    `json:"class_id"`        //会员层级(渠道)
		GroupId        int    `json:"group_id"`        //会员等级
		AttentionLevel int    `json:"attention_level"` //会员关注级别 1.正常; 2.可疑; 3.危险
		FundStatus     int    `json:"fund_status"`     //资金状态 1.正常; 2.锁定
		Tips           string `json:"tips"`            //备注
	}

	UserResp struct {
		Id             interface{} `json:"id"`
		Phone          interface{} `json:"phone"`
		Email          interface{} `json:"email"`
		TrueName       interface{} `json:"true_name"`        //真实姓名
		UserName       interface{} `json:"user_name"`        //用户名
		Birthday       interface{} `json:"birthday"`         //生日
		Status         interface{} `json:"status"`           //账号状态 状态 1启用. 0.禁用
		Sex            interface{} `json:"sex"`              //性别 0.未知；1.男；2，女
		ClassId        interface{} `json:"class_id"`         //会员层级id(渠道)
		GroupId        interface{} `json:"group_id"`         //会员等级id
		ClassName      interface{} `json:"class_name"`       //会员层级(渠道)
		GroupName      interface{} `json:"group_name"`       //会员等级
		AttentionLevel interface{} `json:"attention_level"`  //会员关注级别 1.正常; 2.可疑; 3.危险
		FundStatus     interface{} `json:"fund_status"`      //资金状态 1.正常; 2.锁定
		Tips           interface{} `json:"tips"`             //备注
		FirstLoginIp   interface{} `json:"first_login_ip"`   //首次登录ip
		FirstLoginTime interface{} `json:"first_login_time"` //首次登录时间
		ParentUserName interface{} `json:"parent_user_name"` //上级代理账号
		Balance        interface{} `json:"balance"`          //账户余额
		RegTime        interface{} `json:"reg_time"`         //注册时间
		RegIp          interface{} `json:"reg_ip"`           //注册ip
	}
)
type (
	//密保问题
	SecurityQuestionReq struct {
		Type int `json:"type"`
		Lang int `json:"lang"`
	}

	SecurityQuestions struct {
		Id       interface{} `json:"id"`
		Question interface{} `json:"question"`
		Type     interface{} `json:"type"`
		Lang     interface{} `json:"lang"`
	}

	SecurityQuestionResp struct {
		List []SecurityQuestions `json:"list"` //密保问题列表
	}
)
type (
	//用户登录日志
	UserLoginLogReq struct {
		UserName  string `json:"user_name"`  //会员账号
		Ip        string `json:"ip"`         //登录ip
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示的数量
	}
	UserLoginLogList struct {
		Id         interface{} `json:"id"`          //id
		UserName   interface{} `json:"user_name"`   //会员账号
		Ip         interface{} `json:"ip"`          //登录ip
		Device     interface{} `json:"device"`      //设备1.pc;2.手机
		CreateTime interface{} `json:"create_time"` //登录时间
		Source     interface{} `json:"source"`      //来源
		Url        interface{} `json:"url"`         //登录网址
		Area       interface{} `json:"area"`        //登录区域
		System     interface{} `json:"system"`      //操作系统
		Isp        interface{} `json:"isp"`         //网络服务商
		Screen     interface{} `json:"screen"`      //分辨率
		Browser    interface{} `json:"browser"`     //浏览器
	}
	UserLoginLogResp struct {
		List  []UserLoginLogList `json:"list"`  //登录日志列表
		Total interface{}        `json:"total"` //总数
	}
)

//登录
/*func UserLogin(c *gin.Context) {
	var data interface{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	ip := c.PostForm("ip")
	user_len := utf8.RuneCountInString(username)
	pass_len := utf8.RuneCountInString(password)
	if int(user_len) < 3 || int(user_len) > 15 {
		res := gin.H{"code": 0, "data": nil, "msg": "The user name needs to be between 3 and 15 characters long"}
		c.JSON(http.StatusOK, res)
		return
	}
	if int(pass_len) < 5 || int(pass_len) > 20 {
		res := gin.H{"code": 0, "data": nil, "msg": "The password should be between 5 and 20 characters long"}
		c.JSON(http.StatusOK, res)
		return
	}

	ip_check := utils.RegexpMatch("ip", ip)

	if ip_check == false {
		res := gin.H{"code": 0, "data": nil, "msg": "Please enter the correct IP address"}
		c.JSON(http.StatusOK, res)
		return
	}
	u := model.User{}
	user := u.GetUserByName(username)

	if user.Id == 0 {
		res := gin.H{"code": 0, "data": nil, "msg": "The username does not exist"}
		c.JSON(http.StatusOK, res)
		return
	} else {
		ex_time := 3600
		//盐值和传递的密码md5的加密
		check_pass := md5.New()
		check_str := password + string(user.Salt)
		io.WriteString(check_pass, check_str)
		check_str = fmt.Sprintf("%x", check_pass.Sum(nil))
		if check_str != user.Password {
			res := gin.H{"code": 0, "data": nil, "msg": "User name or password error"}
			c.JSON(http.StatusOK, res)
			return
		}
		w := md5.New()

		rand_str := string(utils.Krand(16, utils.KC_RAND_KIND_ALL)) //3表示随机数包含数字、大小写字母
		token_v := user.Id

		token_v2 := strconv.Itoa(token_v)

		str := rand_str + token_v2

		io.WriteString(w, str)
		token := fmt.Sprintf("%x", w.Sum(nil))

		data = map[string]string{"token": token}
		conn := redisclient.Get()

		//记得销毁本次链连接
		defer conn.Close()

		_, err := conn.Do("SET", token, token_v, "EX", ex_time)
		if err != nil {
			utils.Log(err, "debug", "")
			res := gin.H{"code": 0, "data": nil, "msg": "redisclient set error"}
			c.JSON(http.StatusOK, res)
			return
		}
		//更新登录时间,ip
		login_time := int(time.Now().Unix())
		login_ip := ip

		last_login_time := user.LoginTime

		last_login_ip := string(user.LoginIp)
		update_res, _ := u.UpdateUser_login_info(login_time, login_ip, last_login_time, last_login_ip, user.Id)
		if update_res == false {
			res := gin.H{"code": 0, "data": nil, "msg": "system error"}
			c.JSON(http.StatusOK, res)
			return
		}

	}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}*/

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
	var data interface{}
	user_id, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	//user_id := 1
	userid := user_id.(int)
	user, _ := model.GetUserById(model.Db, userid)
	//"security": "101", "email": "", "qq": "", "receiving_address": "深圳市宝安区", "birth": "1997-07-29"
	if user.Id != 0 {
		data = user
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

//用户基本信息
func UserBaseInfo(c *gin.Context) {
	var data interface{}
	user_id, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": data, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	//user_id := 1
	userid := user_id.(int)
	user, _ := model.GetUserBaseInfo(model.Db, userid)
	//"security": "101", "email": "", "qq": "", "receiving_address": "深圳市宝安区", "birth": "1997-07-29"
	if user.Id != 0 {
		data = user
	}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

//绑定邮箱
func UserBindEmail(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": nil, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}

	userid_int := userid.(int)

	email := c.PostForm("email")

	reg_check := utils.RegexpMatch("email", email)
	if reg_check == false {
		res := gin.H{"code": 0, "data": nil, "msg": "Wrong email format"}
		c.JSON(http.StatusOK, res)
		return
	}
	u := model.User{}

	update_res, _ := u.UpdateUser(userid_int, "email", email)

	if update_res == false {
		res := gin.H{"code": 0, "data": nil, "msg": utils.ERR_SYSTEM_ERROR}
		c.JSON(http.StatusOK, res)
	} else {
		res := gin.H{"code": 1, "data": nil, "msg": "ok"}
		c.JSON(http.StatusOK, res)
	}

}

//绑定手机
func UserBindPhone(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": nil, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	userid_int := userid.(int)

	phone := c.PostForm("phone")
	area_code := c.PostForm("area_code")

	phone_len := utf8.RuneCountInString(phone)
	if phone_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": utils.ERR_LACK_PARAMETERS}
		c.JSON(http.StatusOK, res)
		return
	}

	area_code_len := utf8.RuneCountInString(area_code)

	if area_code_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": utils.ERR_LACK_PARAMETERS}
		c.JSON(http.StatusOK, res)
		return
	}

	area_code_int, error := strconv.Atoi(area_code)
	if error != nil {
		res := gin.H{"code": 0, "data": nil, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		fmt.Println(error)
		return
	}

	u := model.User{}

	update_res, _ := u.UpdateUser_phone(phone, area_code_int, userid_int)

	if update_res == false {
		res := gin.H{"code": 0, "data": nil, "msg": "system error"}
		c.JSON(http.StatusOK, res)
	} else {
		res := gin.H{"code": 1, "data": nil, "msg": "ok"}
		c.JSON(http.StatusOK, res)
	}

}

//重置密码
func ResetPass(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": nil, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	userid_int := userid.(int)

	old_pass := c.PostForm("old_pass")
	new_pass := c.PostForm("new_pass")
	check_pass := c.PostForm("check_pass")

	if new_pass != check_pass {
		res := gin.H{"code": 0, "data": nil, "msg": "The two passwords do not match"}
		c.JSON(http.StatusOK, res)
		return
	}

	u := model.User{}
	user := u.GetUserByToken(userid_int)
	//盐值和传递的密码md5的加密
	w := md5.New()
	check_str := old_pass + string(user.Salt)
	io.WriteString(w, check_str)
	check_str = fmt.Sprintf("%x", w.Sum(nil))
	if check_str != user.Password {
		res := gin.H{"code": 0, "data": nil, "msg": "password error"}
		c.JSON(http.StatusOK, res)
		return
	}
	//生成新密码
	w2 := md5.New()
	check_str_new := new_pass + string(user.Salt)
	io.WriteString(w2, check_str_new)
	check_str_new = fmt.Sprintf("%x", w2.Sum(nil))

	update_res, _ := u.UpdateUser(userid_int, "password", check_str_new)

	if update_res == false {
		res := gin.H{"code": 0, "data": nil, "msg": "system error"}
		c.JSON(http.StatusOK, res)
	} else {
		res := gin.H{"code": 1, "data": nil, "msg": "ok"}
		c.JSON(http.StatusOK, res)
	}

}

func UserBankList(c *gin.Context) {

	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": nil, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	userid_int := userid.(int)
	u := model.UserBank{}

	banklist := u.GetUserBankList(userid_int)
	res := gin.H{"code": 1, "data": banklist, "msg": "ok"}
	if banklist == nil {
		res = gin.H{"code": 0, "data": nil, "msg": "system error"}
	}
	c.JSON(http.StatusOK, res)
}

//获取密保问题
func UserGetSecurity(c *gin.Context) {
	var req SecurityQuestionReq
	var data SecurityQuestionResp
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	if req.Lang == 0 {
		req.Lang = 1
	}
	if req.Type == 0 {
		req.Type = 1
	}

	sq := model.SysSecurityQuestion{Type: req.Type, Lang: req.Type}
	securitys, err := sq.GetSecurity(model.Db)
	if err != nil {
		RespServerErr(c)
		return
	}
	for i := range securitys {
		temp := SecurityQuestions{
			Id:       securitys[i].Id,
			Type:     securitys[i].Type,
			Lang:     securitys[i].Lang,
			Question: securitys[i].Question,
		}
		data.List = append(data.List, temp)
	}
	RespJson(c, status.OK, data)

}

//验证会员列表请求数据
func memberListReqCheck(req *MemberListReq) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if req.TrueName != "" {
		m["true_name"] = req.TrueName
	}
	if req.ParentUserName != "" {
		m["parent_user_name"] = req.ParentUserName
	}

	if req.ClassId != 0 {
		m["class_id"] = req.ClassId
	}
	if req.GroupId != 0 {
		m["group_id"] = req.GroupId
	}
	if req.Status != 999 {
		m["status"] = req.Status
	}
	if req.Phone != "" {
		m["phone"] = req.Phone
	}
	if req.Page < 1 {
		req.Page = 1
	}

	if req.PageCount <= 0 {
		req.PageCount = 10
	}
	return m, nil
}

//获取用户列表
func GetUserList(c *gin.Context) {
	var req MemberListReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := memberListReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	// 数据库查询数据
	//todo: get merchantId from token

	merchantId := 1 //暂定为1,之后从登陆那里获取
	list, count, err := model.GetUserList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	// 组装数据返回给前端显示
	resp := MemberListResp{
		List:  make([]MemberList, 0),
		Total: count,
	}

	for i := range list {
		temp := MemberList{
			ClassName:      list[i].ClassName,
			GroupName:      list[i].GroupName,
			Id:             list[i].Id,
			TrueName:       list[i].TrueName,
			UserName:       list[i].UserName,
			ParentUserName: list[i].ParentUserName,
			Balance:        list[i].Balance,
			LastLoginIp:    list[i].LastLoginIp,
			LastLoginTime:  list[i].LastLoginTime,
			RegTime:        list[i].RegTime,
			RegIp:          list[i].RegIp,
			Phone:          list[i].Phone,
			Email:          list[i].Email,
			FirstLoginIp:   list[i].FirstLoginIp,
			FirstLoginTime: list[i].FirstLoginTime,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)

}

//获取用户信息
func GetUserInfo(c *gin.Context) {
	var req UserReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	//todo 从token中获取merchant_id
	merchantId := 1 //暂定为1,之后从登陆那里获取
	data, err := model.GetUserInfo(model.Db, req.Id, merchantId)

	if err != nil {
		RespServerErr(c)
		return
	}

	if data == nil {
		RespSuccess(c)
		return
	}
	res := UserResp{
		ClassId:        data.ClassId,
		GroupId:        data.GroupId,
		ClassName:      data.ClassName,
		GroupName:      data.GroupName,
		Id:             data.Id,
		TrueName:       data.TrueName,
		UserName:       data.UserName,
		ParentUserName: data.ParentUserName,
		Balance:        data.Balance,
		RegTime:        data.RegTime,
		RegIp:          data.RegIp,
		Phone:          data.Phone,
		Email:          data.Email,
		FirstLoginIp:   data.FirstLoginIp,
		FirstLoginTime: data.FirstLoginTime,
	}
	RespJson(c, status.OK, res)
}

//验证编辑会员请求数据
func userEditReqCheck(req *UserReq) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	if req.Password != "" {
		m["password"] = req.Password
	}
	if req.PayPass != "" {
		m["pay_pass"] = req.PayPass
	}
	if req.ClassId != 0 {
		m["class_id"] = req.ClassId
	}
	if req.GroupId != 0 {
		m["group_id"] = req.GroupId
	}
	if req.TrueName != "" {
		m["true_name"] = req.TrueName
	}
	if req.Phone != "" {
		m["phone"] = req.Phone
	}
	if req.Email != "" {
		m["email"] = req.Email
	}
	if req.Birthday != 0 {
		m["birthday"] = req.Birthday
	}

	if req.AttentionLevel != 0 {
		m["attention_level"] = req.AttentionLevel
	}
	if req.FundStatus != 0 {
		m["fund_status"] = req.FundStatus
	}
	if (req.Sex == 0) || (req.Sex == 1) || (req.Sex == 2) {
		m["sex"] = req.Sex
	}
	if (req.Status == 1) || (req.Status == 0) {
		m["status"] = req.Status
	}

	return m, nil
}

//编辑用户
func UserEdit(c *gin.Context) {
	var req UserReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	//验证参数合法性
	m, err := userEditReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo 从token中获取merchant_id
	merchantId := 1 //暂定为1,之后从登陆那里获取

	if v, ok := m["password"]; ok {
		userInfo, err := model.GetUserInfo(model.Db, req.Id, merchantId)
		if err != nil {
			logrus.Error(err)
			RespServerErr(c)
			return
		}
		salt := userInfo.Salt
		//将密码加密
		password, _ := v.(string)
		password_str := password + salt
		password_str = utils.Md5S(password_str)
		m["password"] = password_str
	}
	if v, ok := m["pay_pass"]; ok {
		userInfo, err := model.GetUserInfo(model.Db, req.Id, merchantId)
		if err != nil {
			logrus.Error(err)
			RespServerErr(c)
			return
		}
		salt := userInfo.Salt
		//将密码加密
		paypass, _ := v.(string)
		paypass_str := paypass + salt
		paypass_str = utils.Md5S(paypass_str)
		m["pay_pass"] = paypass_str
	}

	user := model.User{Id: req.Id, MerchantId: merchantId}
	err = user.UserEdit(model.Db, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

func userLoginLogReqCheck(req *UserLoginLogReq) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if req.Ip != "" {
		m["ip"] = req.Ip
	}
	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if req.Page < 1 {
		req.Page = 1
	}

	if req.PageCount <= 0 {
		req.PageCount = 10
	}

	return m, nil
}

//用户登录日志
func GetUserLoginLogList(c *gin.Context) {
	var req UserLoginLogReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}
	//验证参数合法性
	m, err := userLoginLogReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 数据库查询数据
	//todo: get merchantId from token

	merchantId := 1 //暂定为1,之后从登陆那里获取
	list, count, err := model.UserLoginLogList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	// 组装数据返回给前端显示
	resp := UserLoginLogResp{
		List:  make([]UserLoginLogList, 0),
		Total: count,
	}

	for i := range list {
		temp := UserLoginLogList{
			Id:         list[i].Id,
			UserName:   list[i].UserName,
			Ip:         list[i].Ip,
			Device:     list[i].Device,
			CreateTime: list[i].CreateTime,
			Source:     list[i].Source,
			Url:        list[i].Url,
			Area:       list[i].Area,
			System:     list[i].System,
			Isp:        list[i].Isp,
			Screen:     list[i].Screen,
			Browser:    list[i].Browser,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)

}
