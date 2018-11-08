package action

import (
	"crypto/md5"
	"fmt"
	"golang_game_merchant/lib/utils"
	"golang_game_merchant/model"
	"io"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	//"github.com/gomodule/redigo/redisclient"
	"golang_game_merchant/global/redisclient"
)

//登录
func UserLogin(c *gin.Context) {
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

		last_login_time, err := strconv.Atoi(user.LoginTime)
		if err != nil {
			res := gin.H{"code": 0, "data": nil, "msg": "system errorlo"}
			c.JSON(http.StatusOK, res)
			return
		}
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
}

//获取注册渲染页面配置
func UserRegisterPageInfo(c *gin.Context) {
	//获取商户id
	merchantIdStr := c.PostForm("merchant_id")
	merchantId, err := strconv.Atoi(merchantIdStr)
	if err != nil {
		panic(err)
	}
	merchant, err := model.GetMerchantWebsiteReg(model.Db, merchantId)
	if err == model.ErrRecordNotFound {
		res := gin.H{"code": 0, "data": nil, "msg": "Merchant does not exist"}
		c.JSON(http.StatusOK, res)
		return
	}

	res := gin.H{"code": 1, "data": merchant, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

//注册
func UserRegister(c *gin.Context) {
	//var data interface{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	reg_ip := c.PostForm("reg_ip")
	question1 := c.PostForm("question1")
	question2 := c.PostForm("question2")
	answer1 := c.PostForm("answer1")
	answer2 := c.PostForm("answer2")
	true_name := c.PostForm("true_name")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	pay_pass := c.PostForm("pay_pass")
	parent_id := c.PostForm("parent_id")
	member_type := c.PostForm("member_type")

	//获取商户的网站配置信息
	merchant_id := 1 //先写死
	merchant, _ := model.GetMerchantWebsiteReg(model.Db, merchant_id)

	question1_len := utf8.RuneCountInString(question1)
	question2_len := utf8.RuneCountInString(question2)
	answer1_len := utf8.RuneCountInString(answer1)
	answer2_len := utf8.RuneCountInString(answer2)
	true_name_len := utf8.RuneCountInString(true_name)
	phone_len := utf8.RuneCountInString(phone)
	email_len := utf8.RuneCountInString(email)
	pay_pass_len := utf8.RuneCountInString(pay_pass)
	parent_id_len := utf8.RuneCountInString(parent_id)
	member_type_len := utf8.RuneCountInString(member_type)

	reg_ip_check := utils.RegexpMatch("ip", reg_ip)

	if reg_ip_check == false {
		res := gin.H{"code": 0, "data": nil, "msg": "Please enter the correct IP address"}
		c.JSON(http.StatusOK, res)
		return
	}

	if parent_id_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": "Parameter incompleteness"}
		c.JSON(http.StatusOK, res)
		return
	}
	if member_type_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": "Parameter incompleteness"}
		c.JSON(http.StatusOK, res)
		return
	}
	if merchant.RegPayPass == 3 && pay_pass_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": "Payment password cannot be empty"}
		c.JSON(http.StatusOK, res)
		return
	}

	if merchant.RegSecurityQuestion == 3 && (question1_len < 1 || question2_len < 1 || answer1_len < 1 || answer2_len < 1) {
		res := gin.H{"code": 0, "data": nil, "msg": "The confidentiality issue must be complete"}
		c.JSON(http.StatusOK, res)
		return
	}
	if merchant.RegTrueName == 3 && true_name_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": "Real names cannot be empty"}
		c.JSON(http.StatusOK, res)
		return
	}
	if merchant.RegPhone == 3 && phone_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": "The phone cannot be empty"}
		c.JSON(http.StatusOK, res)
		return
	}
	if merchant.RegEmail == 3 && email_len < 1 {
		res := gin.H{"code": 0, "data": nil, "msg": "email cannot be empty"}
		c.JSON(http.StatusOK, res)
		return
	}

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

	u := model.User{}
	user := u.GetUserByName(username)

	if user.Id != 0 {
		res := gin.H{"code": 0, "data": nil, "msg": "User name has been registered"}
		c.JSON(http.StatusOK, res)
		return
	} else {
		var err interface{}
		w := md5.New()

		salt := string(utils.Krand(4, utils.KC_RAND_KIND_ALL)) //3表示随机数包含数字、大小写字母

		str := password + salt

		io.WriteString(w, str)
		password_md5 := fmt.Sprintf("%x", w.Sum(nil))
		data := make(map[string]interface{})

		data["merch_id"] = 1
		data["user_name"] = username
		data["true_name"] = true_name
		data["password"] = password_md5
		phone = utils.CheckEmptyStr(phone, "0")
		data["phone"], err = strconv.Atoi(phone)
		if err != nil {
			res := gin.H{"code": 1, "data": nil, "msg": "system error"}
			c.JSON(http.StatusOK, res)
			return
		}
		data["email"] = email
		data["salt"] = salt
		data["reg_time"] = int(time.Now().Unix())
		data["reg_ip"] = reg_ip
		data["type"] = member_type
		data["parent_id"] = parent_id
		data["pay_pass"] = pay_pass
		res, _ := u.InsterUser(data)

		if res == true {
			res := gin.H{"code": 1, "data": nil, "msg": "ok"}
			c.JSON(http.StatusOK, res)
		} else {
			res := gin.H{"code": 0, "data": nil, "msg": "register fail"}
			c.JSON(http.StatusOK, res)
		}
	}
	/*res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)*/
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
		res := gin.H{"code": 0, "data": nil, "msg": "system error"}
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
		res := gin.H{"code": 0, "data": nil, "msg": "system error"}
		c.JSON(http.StatusOK, res)
	} else {
		res := gin.H{"code": 1, "data": nil, "msg": "ok"}
		c.JSON(http.StatusOK, res)
	}

}

//绑定QQ
func UserBindQQ(c *gin.Context) {
	userid, ok := c.Get("user_id")
	if !ok {
		res := gin.H{"code": 0, "data": nil, "msg": "fail"}
		c.JSON(http.StatusOK, res)
		return
	}
	userid_int := userid.(int)

	qq := c.PostForm("qq")

	reg_check := utils.RegexpMatch("qq", qq)
	if reg_check == false {
		res := gin.H{"code": 0, "data": nil, "msg": "Wrong qq format"}
		c.JSON(http.StatusOK, res)
		return
	}
	u := model.User{}

	update_res, _ := u.UpdateUser(userid_int, "qq", qq)

	if update_res == false {
		res := gin.H{"code": 0, "data": nil, "msg": "system error"}
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

	q_type := c.PostForm("type")
	lang := c.PostForm("lang")

	q_type = utils.CheckEmptyStr(q_type, "1")
	lang = utils.CheckEmptyStr(lang, "1")

	q_type_int, err := strconv.Atoi(q_type)
	if err != nil {
		res := gin.H{"code": 0, "data": nil, "msg": "system errorlo"}
		c.JSON(http.StatusOK, res)
		return
	}

	lang_int, err := strconv.Atoi(lang)
	if err != nil {
		res := gin.H{"code": 0, "data": nil, "msg": "system errorlo"}
		c.JSON(http.StatusOK, res)
		return
	}

	u := model.SysSecurityQuestion{}
	securitys := u.GetSecurity(q_type_int, lang_int)
	res := gin.H{"code": 1, "data": securitys, "msg": "ok"}
	if securitys == nil {
		res = gin.H{"code": 0, "data": nil, "msg": "system error"}
	}
	c.JSON(http.StatusOK, res)
}
