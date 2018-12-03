package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type User struct {
	Id            int     `json:"id"`
	MerchantId    int     `json:"merchant_id"`
	UserName      string  `json:"user_name"`
	Password      string  `json:"password,omitempty"`
	Salt          string  `json:"salt,omitempty"`
	TrueName      string  `json:"true_name"`
	Phone         string  `json:"phone"`
	Email         string  `json:"email,omitempty"`
	QQ            string  `json:"qq,omitempty"`
	Birthday      string  `json:"birthday,omitempty"`
	LoginTime     int64   `json:"login_time,omitempty"`
	LoginIp       string  `json:"login_ip,omitempty"`
	LastLoginIp   string  `json:"last_login_ip,omitempty"`
	LastLoginTime int64   `json:"last_login_time"`
	Balance       float64 `json:"balance"`
	Skype         string  `json:"skype,omitempty"`
	Device        int     `json:"device,omitempty"`
	Source        string  `json:"source,omitempty"`
	Status        int     `json:"status"`
	ClassId       int     `json:"class_id,omitempty"`
	GroupId       int     `json:"group_id,omitempty"`
	TimeZone      int     `json:"time_zone,omitempty"`
	Lang          string  `json:"lang,omitempty"`
	AreaCode      int     `json:"area_code,omitempty"`
	RegTime       int64   `json:"reg_time"`
	RegIp         string  `json:"reg_ip"`
	AgentCode     string  `json:"agent_code"`
	ParentId      int     `json:"parent_id"`
	PayPass       string  `json:"pay_pass"`
	ResetPassword string `json:"reset_password"`
	ResetPaypass string `json:"reset_paypass"`
}

type UserBank struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	TrueName   string `json:"true_name"`
	CardNo     string `json:"card_no"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type SysSecurityQuestion struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Lang     int    `json:"lang"`
}

type UserAnswers struct {
	Id         int    `json:"id"`
	Answer     string `json:"answer"`
	UserId     int    `json:"user_id"`
	QuestionId int    `json:"question_id"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	MerchantId int    `json:"merchant_id"`
}

type UserAnswersInfo struct {
	Id         int    `json:"id"`
	Question string `json:"question"`
	Type       int    `json:"type"`
	QuestionId int    `json:"question_id"`
}

type UserInterface interface {
	GetUserById()
	GetUserByName()
}

func GetUserById(db *gorm.DB, id int) (*User, error) {
	var u User
	if err := db.Table("user").Select("id,merchant_id,user_name,true_name,phone,email,qq,birthday,balance,last_login_time,last_login_ip,password,salt,pay_pass,reset_paypass").
		Where("id=?", id).Find(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &u, nil
}

func GetUserByPhone(db *gorm.DB, phone string, areaCode int) (*User, error) {
	var u User
	if err := db.Table("user").Select("id,merchant_id,user_name,true_name,phone,email,qq,birthday,balance,last_login_time,last_login_ip,password,salt").
		Where("phone=? AND area_code=?", phone, areaCode).Find(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &u, nil
}

func GetUserByName(db *gorm.DB, name string) (*User, error) {
	var m User
	if err := db.Table("user").Select("id,merchant_id,user_name,password,salt,balance,login_time,login_ip,"+
		"last_login_ip,last_login_time,true_name,phone,email,qq,skype,lang,time_zone,birthday,area_code,"+
		"device,source,class_id,group_id,status,reset_password,reset_paypass").Where("user_name=?", name).Or("phone=?", name).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &m, nil
}

func (u *User) InsertUser(db *gorm.DB) (int, error) {
	if err := db.Create(&u).Error; err != nil {
		logrus.Error(err)
		return 0, err
	}
	return u.Id, nil
}

//修改用户单个字段值
func UpdateUser(db *gorm.DB, userId int, fields map[string]interface{}) (bool, error) {
	if err := db.Table("user").Where("id=?", userId).Update(fields).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

//修改用户登录ip,登录时间
func UpdateUserLoginInfo(db *gorm.DB, userId int, loginTime int64, loginIp string, lastLoginTime int64, lastLoginIp string) (bool, error) {
	m := make(map[string]interface{})
	m["login_time"] = loginTime
	m["login_ip"] = loginIp
	m["last_login_time"] = lastLoginTime
	m["last_login_ip"] = lastLoginIp
	if err := db.Table("user").Where("id=?", userId).Update(m).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

//获取用户银行列表
func GetUserBankList(db *gorm.DB, userId int) ([]UserBank, error) {
	var bankList []UserBank
	if err := db.Table("user_bank").Where("user_id=?", userId).Find(&bankList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return bankList, nil
}

//获取单个用户银行信息
func GetUserBankInfo(db *gorm.DB, id int, userId int) (*UserBank, error) {
	var u UserBank
	if err := db.Table("user_bank").Select("id,bank_name,card_no,user_id").Where("id=? AND user_id=?", id, userId).Find(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &u, nil
}

//添加银行信息
func (a UserBank) AddUserBank(db *gorm.DB) (int, error) {
	if err := db.Table("user_bank").Create(&a).Error; err != nil {
		return 0, err
	}
	return a.Id, nil
}

//获取系统安全密保问题列表
func GetSecurityList(db *gorm.DB, lang int) ([]SysSecurityQuestion, error) {
	var securityList []SysSecurityQuestion
	if err := db.Table("sys_security_question").Select("id,question,lang").Where("lang=?", lang).Find(&securityList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return securityList, nil
}

//获取系统安全密保问题详情
func GetSecurityInfo(db *gorm.DB, id int) (*SysSecurityQuestion, error) {
	var security SysSecurityQuestion
	if err := db.Table("sys_security_question").Select("id,question,lang").Where("id=?", id).Find(&security).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &security, nil
}

//设置安全密保答案
func (ua UserAnswers) SetSecurity(db *gorm.DB) (int, error) {
	if err := db.Create(&ua).Error; err != nil {
		return 0, err
	}
	return ua.Id, nil
}

//获取安全密保答案
func GetUserSecurity(db *gorm.DB, questionId, userId int) (*UserAnswers, error) {
	var m UserAnswers
	if err := db.Table("user_answers").Where("question_id=? And user_id=? ", questionId, userId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//获取用户的密保
func GetUserSecurityList(db *gorm.DB, userId int) ([]UserAnswersInfo, error) {
	var m []UserAnswersInfo
	if err := db.Table("user_answers as ua").Select("ua.id,ua.question_id,sq.question").
		Joins("LEFT JOIN sys_security_question sq ON ua.question_id=sq.id").
		Where("ua.user_id=?", userId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
