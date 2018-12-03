package action

import (
	"fmt"
	"game2/global/status"
	"game2/lib/utils"
	"game2/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type UserSecurityReq struct {
	Lang int `json:"lang" binding:"required"`
}

//设置密保
type UserSetSecurityReq struct {
	QuestionId []int    `json:"question_id" binding:"required"`
	Answer     []string `json:"answer" binding:"required"`
}

type ValidateUserReq struct {
	UserName string `json:"user_name" binding:"required"`
}

//登录密保验证
type PasswordValidateSecurityReq struct {
	UserName string   `json:"user_name" binding:"required"`
	QuestionId       []int    `json:"question_id" binding:"required"`
	Answer   []string `json:"answer" binding:"required"`
}

//资金密保验证
type PayPassValidateSecurityReq struct {
	QuestionId     []int    `json:"question_id" binding:"required"`
	Answer []string `json:"answer" binding:"required"`
}

//重置登录密码
type ResetPasswordReq struct {
	UserName      string `json:"user_name" binding:"required"`
	ResetPassword string `json:"reset_password" binding:"required"`
	Password      string `json:"password" binding:"required"`
}

//重置资金密码
type ResetPayPassReq struct {
	PayPass      string `json:"pay_pass" binding:"required"`
	ResetPaypass string `json:"reset_paypass" binding:"required"`
}

func ValidateUser(c *gin.Context) {
	var m ValidateUserReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	info, err := model.GetUserByName(model.Db, m.UserName)
	if err != nil {
		RespServerErr(c)
		return
	}
	if info == nil {
		RespJson(c, status.ErrUserNotExist, nil)
		return
	}
	RespSuccess(c)
}

//获取密保问题
func UserGetSecurity(c *gin.Context) {
	var m UserSecurityReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}

	list, err := model.GetSecurityList(model.Db, m.Lang)
	if err != nil {
		RespServerErr(c)
		return
	}
	res := map[string]interface{}{"list": list}
	RespJson(c, status.OK, res)
}

//设置密保答案
func UserSetSecurity(c *gin.Context) {
	var m UserSetSecurityReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		logrus.Error(c)
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	mid, ok := c.Get("merchant_id")
	if !ok {
		logrus.Error(c)
		RespServerErr(c)
		return
	}
	MerchantId := int(mid.(float64))
	nowTime := time.Now().Unix()
	list, err := model.GetUserSecurityList(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	//判断用户的密保问题是否设置过,
	//设置过的不可再设置
	if len(list) > 0 {
		RespJson(c, status.ErrSecurityExist, nil)
		return
	}
	tx := model.Db.Begin()
	for k, v := range m.QuestionId {
		var ua = model.UserAnswers{
			UserId:     userId,
			MerchantId: MerchantId,
			QuestionId: v,
			Answer:     m.Answer[k],
			CreateTime: nowTime,
			UpdateTime: nowTime,
		}
		_, err := ua.SetSecurity(tx)
		if err != nil {
			tx.Rollback()
			RespServerErr(c)
			return
		}
	}
	tx.Commit()
	RespJson(c, status.OK, nil)
}

//查看用户设置过的密保
func UserAnswerList(c *gin.Context) {
	var m ValidateUserReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	info, err := model.GetUserByName(model.Db, m.UserName)
	if err != nil {
		RespServerErr(c)
		return
	}
	if info == nil {
		RespJson(c, status.ErrUserNotExist, nil)
		return
	}
	data := make(map[string]interface{})
	userId := info.Id
	list, err := model.GetUserSecurityList(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	data["list"] = list
	RespJson(c, status.OK, data)
}

//检验登录密保答案（不需要登录）
func PasswordValidateSecurity(c *gin.Context) {
	var m PasswordValidateSecurityReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	info, err := model.GetUserByName(model.Db, m.UserName)
	if err != nil {
		RespServerErr(c)
		return
	}
	if info == nil {
		RespParamErr(c)
		return
	}
	userId := info.Id
	for k, v := range m.QuestionId {
		info, err := model.GetUserSecurity(model.Db, v, userId)
		if err != nil {
			RespServerErr(c)
			return
		}
		if info == nil{
			//未设置密保
			//TODO
			RespUnauthorized(c)
			return
		}
		if info.Answer != m.Answer[k] {
			RespJson(c, status.ErrAnswerError, nil)
			return
		}
	}
	fields := make(map[string]interface{})
	//成功的话可以重置密码、资金密码
	nowTime := time.Now().Unix()
	fields["update_time"] = nowTime
	fields["reset_password"] = string(utils.Krand(10, utils.KC_RAND_KIND_ALL))
	_, err = model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	fmt.Println(fields["reset_password"])
	RespJson(c, status.OK, fields["reset_password"])
}

//检验资金密保答案（需要登录）
func PayPassValidateSecurity(c *gin.Context) {
	var m PayPassValidateSecurityReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	id, ok := c.Get("user_id")
	if !ok {
		logrus.Error(c)
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	for k, v := range m.QuestionId {
		info, err := model.GetUserSecurity(model.Db, v, userId)
		if err != nil {
			RespServerErr(c)
			return
		}
		if info.Answer != m.Answer[k] {
			RespJson(c, status.ErrAnswerError, nil)
			return
		}
	}
	fields := make(map[string]interface{})
	//成功的话可以重置密码、资金密码
	nowTime := time.Now().Unix()
	fields["update_time"] = nowTime
	fields["reset_paypass"] = string(utils.Krand(10, utils.KC_RAND_KIND_ALL))
	_, err := model.UpdateUser(model.Db, userId, fields)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespJson(c, status.OK, fields["reset_paypass"])
}

//重置资金密码
func ResetPayPass(c *gin.Context) {
	var m ResetPayPassReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	id, ok := c.Get("user_id")
	if !ok {
		logrus.Error(c)
		RespServerErr(c)
		return
	}
	userId := int(id.(float64))
	user, err := model.GetUserById(model.Db, userId)
	if err != nil {
		RespServerErr(c)
		return
	}
	if user.ResetPaypass == "" || user.ResetPaypass != m.ResetPaypass {
		RespUnauthorized(c)
		return
	}
	fields := make(map[string]interface{})
	//重置资金密码
	nowTime := time.Now().Unix()
	fields["update_time"] = nowTime
	str := m.PayPass + user.Salt
	passwordMd5 := utils.Md5V(utils.Md5V(str))
	//重置资金密码
	fields["pay_pass"] = passwordMd5
	fields["reset_paypass"] = ""
	_, err = model.UpdateUser(model.Db, user.Id, fields)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

//重置登录密码
func ResetPassword(c *gin.Context) {
	var m ResetPasswordReq
	if err := c.BindJSON(&m); err != nil {
		RespParamErr(c)
		return
	}
	info, err := model.GetUserByName(model.Db, m.UserName)
	if err != nil {
		RespServerErr(c)
		return
	}
	if info == nil {
		RespParamErr(c)
		return
	}
	if info.ResetPassword == "" || info.ResetPassword != m.ResetPassword {
		RespUnauthorized(c)
		return
	}

	fields := make(map[string]interface{})
	//成功的话可以重置密码、资金密码
	nowTime := time.Now().Unix()
	fields["update_time"] = nowTime
	//密码
	str := m.Password + info.Salt
	passwordMd5 := utils.Md5V(utils.Md5V(str))
	fields["password"] = passwordMd5
	fields["reset_password"] = ""
	_, err = model.UpdateUser(model.Db, info.Id, fields)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}