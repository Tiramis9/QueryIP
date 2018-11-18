package model

import (
	//"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id             int
	MerchantId     int
	UserName       string
	TrueName       string
	Password       string
	Phone          string
	Email          string
	QQ             string
	Salt           string
	Balance        float64
	Status         int
	RegTime        int
	RegIp          string
	LoginTime      int
	LoginIp        string
	Type           int
	ParentId       int
	LastLoginIp    string
	LastLoginTime  int
	Skype          string
	Lang           string
	TimeZone       int
	PayPass        string
	Sex            int
	Birthday       int
	AreaCode       int
	AgentCode      string
	Device         int
	Source         string
	ClassId        int
	GroupId        int
	AttentionLevel int
	FundStatus     int
	Tips           string
	FirstLoginIp   string
	FirstLoginTime int
}

type UserListInfo struct {
	User
	ClassName      string //会员层级(渠道)
	GroupName      string //会员等级
	ParentUserName string //上级代理账号
}

type UserBaseInfo struct {
	User
	ClassName      string //会员层级(渠道)
	GroupName      string //会员等级
	ParentUserName string //上级代理账号
}

type UserBank struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id,omitempty"`
	TrueName string `json:"true_name,omitempty"`
	Phone    int64  `json:"phone,omitempty"`
	CardNo   string `json:"card_no,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	IdCard   string `json:"id_card,omitempty"`
	Status   int    `json:"status,omitempty"`
}

type SysSecurityQuestion struct {
	Id       int
	Question string
	Type     int
	Lang     int
}

type AgentClassGroupInfo struct {
	Agent
	ClassName string `json:"class_name,omitempty"`
	GroupName string `json:"group_name,omitempty"`
}

//用户登录日志
type UserLoginLog struct {
	Id         int
	UserId     int
	System     string //操作系统
	Area       string
	Url        string //登录地址
	Ip         string
	CreateTime int
	UpdateTime int
	MerchantId int
	Isp        string //网络服务商
	Screen     string //屏幕分辨率
	Browser    string //浏览器
	Device     int    //设备1.pc;2.手机
	Source     string //来源
}
type LoginLogInfo struct {
	UserLoginLog
	UserName string
}

//id,user_name,last_login_ip,last_login_time,status,device,source,reg_time,reg_ip
type SubUser struct {
	Id            int    `json:"id"`
	UserName      string `json:"user_name"`
	Status        int    `json:"status"`
	RegTime       int    `json:"reg_time"`
	RegIp         string `json:"reg_ip"`
	LoginTime     int    `json:"login_time"`
	LoginIp       string `json:"login_ip"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int    `json:"last_login_time"`
	Device        int    `json:"device"`
	Source        string `json:"source"`
}

func GetUserById(db *gorm.DB, id int) (*User, error) {
	var user User
	if err := db.Table(`user`).Where(`id=?`, id).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func GetUserBaseInfo(db *gorm.DB, id int) (*User, error) {
	return GetUserById(db, id)
}

func GetUserByMerchantIdAndUserId(db *gorm.DB, id, merchantId int) (*User, error) {
	var user User
	if err := db.Table(`user`).Where(`id=? AND merchant_id=?`, id, merchantId).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u User) GetUserByToken(id int) (user User) {
	return User{}
	/*stmt, errs := Db.Prepare("SELECT id,merch_id,user_name,password,salt FROM user " + "Where id=? ")
	if errs != nil {
		fmt.Println(errs)
		return
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)

	conv := User{}

	row.Scan(&conv.Id, &conv.Merch_id, &conv.User_name, &conv.Password, &conv.Salt)

	return conv*/
}

func (u User) InsterUser(data map[string]interface{}) (bool, interface{}) {
	return true, nil
	/*stmt, err := Db.Prepare("insert into user(merch_id,user_name,true_name,password,phone,email,salt,reg_time,reg_ip,type,parent_id,pay_pass)values(?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	res, err := stmt.Exec(data["merch_id"], data["user_name"], data["true_name"], data["password"], data["phone"], data["email"], data["salt"], data["reg_time"], data["reg_ip"], data["type"], data["parent_id"], data["pay_pass"])
	if err != nil {

		return false, nil
	}
	id, err := res.LastInsertId()
	if err != nil {

		return false, nil
	}
	return true, id*/
}

//修改用户单个字段值
func (u User) UpdateUser(uid int, field string, data string) (bool, int64) {
	return true, 0
	/*stmt, err := Db.Prepare("UPDATE user SET " + field + "=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	res, err := stmt.Exec(data, uid)
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	num, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return false, 0
	} else {
		return true, num
	}*/

}

//修改用户登录ip,登录时间
func (u User) UpdateUser_login_info(login_time int, login_ip string, last_login_time int, last_login_ip string, user_id int) (bool, int64) {
	return true, 0
	/*stmt, err := Db.Prepare("UPDATE user SET login_time=?,login_ip=?,last_login_time=?,last_login_ip=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	res, err := stmt.Exec(login_time, login_ip, last_login_time, last_login_ip, user_id)
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	num, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return false, 0
	} else {
		return true, num
	}*/

}

//修改用户登录ip,登录时间
func (u User) UpdateUser_phone(phone string, area_code int, user_id int) (bool, int64) {
	return true, 0
	/*stmt, err := Db.Prepare("UPDATE user SET phone=?,area_code=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	res, err := stmt.Exec(phone, area_code, user_id)
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	num, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return false, 0
	} else {
		return true, num
	}*/

}

func (u UserBank) GetUserBankList(user_id int) interface{} {
	return nil
	/*var banklist []UserBank
	stmt, errs := Db.Prepare("SELECT id,bank_name,card_no,user_id FROM user_bank WHERE user_id=?")
	if errs != nil {
		fmt.Println(errs)
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(user_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		conv := UserBank{}
		rows.Scan(&conv.Id, &conv.Bank_name, &conv.Card_no, &conv.User_id)
		banklist = append(banklist, conv)
	}
	fmt.Println(banklist)
	return banklist*/
}

func (sq SysSecurityQuestion) GetSecurity(db *gorm.DB) ([]SysSecurityQuestion, error) {

	var securitys []SysSecurityQuestion

	if err := db.Select([]string{"id,question,type,lang"}).Find(&securitys).Error; err != nil {
		return nil, err
	}
	return securitys, nil

}

func GetUserList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]UserListInfo, int, error) {
	whereStr := "u.merchant_id=?"
	condition := []interface{}{merchantId}

	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name=?"
		condition = append(condition, v)
	}

	if v, ok := m["true_name"]; ok {
		whereStr += " AND u.true_name=?"
		condition = append(condition, v)
	}

	if v, ok := m["parent_user_name"]; ok {
		whereStr += " AND a.user_name=?"
		condition = append(condition, v)
	}

	if v, ok := m["status"]; ok {
		whereStr += " AND u.status=?"
		condition = append(condition, v)
	}

	if v, ok := m["phone"]; ok {
		whereStr += " AND u.phone=?"
		condition = append(condition, v)
	}

	if v, ok := m["class_id"]; ok {
		whereStr += " AND uc.id=?"
		condition = append(condition, v)
	}

	if v, ok := m["group_id"]; ok {
		whereStr += " AND ug.id=?"
		condition = append(condition, v)
	}

	var userList []UserListInfo

	if err := db.LogMode(true).Table(`user AS u`).Joins(`
		LEFT JOIN agent AS a ON u.parent_id=a.id
	`).Joins(`
		LEFT JOIN merchant_user_group AS ug ON u.group_id=ug.id
	`).Joins(`
		LEFT JOIN merchant_user_class AS uc ON u.class_id=uc.id
	`).Select(`
		u.*,
		a.user_name as parent_user_name,
		ug.group_name,
		uc.class_name
	`).Where(whereStr, condition...).Offset((page - 1) * pageCount).Limit(pageCount).Find(&userList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`user AS u`).Joins(`
		LEFT JOIN agent AS a ON u.parent_id=a.id
	`).Joins(`
		LEFT JOIN merchant_user_group AS ug ON u.group_id=ug.id
	`).Joins(`
		LEFT JOIN merchant_user_class AS uc ON u.class_id=uc.id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return userList, count, nil
}

//编辑用户信息
func (user User) UserEdit(db *gorm.DB, params map[string]interface{}) error {
	if err := db.Debug().Model(&user).Updates(params).Error; err != nil {
		return err
	}
	return nil
}

//获取用户信息
func GetUserInfo(db *gorm.DB, id int, merchantId int) (*UserBaseInfo, error) {
	var userInfo UserBaseInfo
	if err := db.Table(`user AS u`).Joins(`
		LEFT JOIN agent AS a ON u.parent_id=a.id`).Joins(`
		LEFT JOIN merchant_user_group AS ug ON u.group_id=ug.id`).Joins(`
		LEFT JOIN merchant_user_class AS uc ON u.group_id=uc.id`).Select(`
		u.id,u.user_name,u.true_name,u.email,u.phone,u.birthday,u.sex,u.class_id,u.group_id,u.attention_level,u.fund_status,u.status,
		u.reg_time,u.reg_ip,u.first_login_ip,u.first_login_time,a.user_name as parent_user_name,ug.group_name,uc.class_name
	`).Where(`u.id=? AND u.merchant_id=?`, id, merchantId).Find(&userInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}
	return &userInfo, nil
}

func UserLoginLogList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]LoginLogInfo, int, error) {
	whereStr := "u.merchant_id=?"
	condition := []interface{}{merchantId}

	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name=?"
		condition = append(condition, v)
	}

	if v, ok := m["ip"]; ok {
		whereStr += " AND ul.ip=?"
		condition = append(condition, v)
	}

	var logList []LoginLogInfo

	if err := db.Table(`user_login_log AS ul`).Joins(`
		LEFT JOIN user AS u ON ul.user_id=u.id
	`).Select(`
		ul.*,
		u.user_name 
	`).Where(whereStr, condition...).Offset((page - 1) * pageCount).Limit(pageCount).Find(&logList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`user_login_log AS ul`).Joins(`
		LEFT JOIN user AS u ON ul.user_id=u.id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return logList, count, nil
}

//查询代理下的子用户
func GetSubUserList(db *gorm.DB, parentId int, merchantId int, page int, pageCount int) ([]SubUser, error) {
	var subUserList []SubUser
	if err := db.Table(`user`).Select("id,user_name,last_login_ip,last_login_time,status,device,source,reg_time,reg_ip").
		Where("parent_id=? AND merchant_id=?", parentId, merchantId).Offset((page - 1) * pageCount).
		Limit(pageCount).Find(&subUserList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return subUserList, nil
}

//查询代理下的子用户记录条数
func GetSubUserCount(db *gorm.DB, parentId int, merchantId int) (int, error) {
	var total int
	if err := db.Table(`user`).Where("parent_id=? AND merchant_id=?", parentId, merchantId).Count(&total).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return total, nil
		}
		return 0, err
	}
	return total, nil
}

func IsExistWithMerchantIdAndUserId(db *gorm.DB, merchantId, userId int) (bool, error) {
	var user User
	if err := db.Table(`user`).Where(`merchant_id=? AND id=?`, merchantId, userId).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// 中心账户加扣款，确保
func (u *User) MerchantUpdateUserBalance(db *gorm.DB, money float64) error {
	if money != 0 {
		temp := db.Table(`user`).Update("balance", gorm.Expr("balance+?", money)).Where(`
			merchant_id=? AND id=? AND balance+?>=0
		`, u.MerchantId, u.Id, money)
		rowsAffected, err := temp.RowsAffected, temp.Error
		if err != nil {
			return err
		}
		if rowsAffected != 1 {
			//须确保能根据u.merchantId和u.Id能查出唯一用户
			return ErrNoEnoughMoney
		}
	}
	return nil
}
