package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id            int    `json:"id"`
	MerchId       int    `json:"merch_id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password,omitempty"`
	Salt          string `json:"salt,omitempty"`
	TrueName      string `json:"true_name"`
	Phone         string `json:"phone"`
	Email         string `json:"email,omitempty"`
	QQ            string `json:"qq,omitempty"`
	Birthday      string `json:"birthday,omitempty"`
	LoginTime     string `json:"login_time,omitempty"`
	LoginIp       string `json:"login_ip,omitempty"`
	LastLoginIp   string `json:"last_login_ip,omitempty"`
	LastLoginTime string `json:"last_login_time"`
	Balance       string `json:"balance"`
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
	Id       int    `json:"id"`
	Question string `json:"question"`
	Type     int    `json:"type"`
	Lang     int    `json:"lang"`
}

// select * from user where id=?
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

func (u User) GetUserByName(name string) (user User) {
	return User{}
	/*stmt, errs := Db.Prepare("SELECT id,merch_id,user_name,password,salt,login_time,login_ip,last_login_ip,last_login_time FROM user " + "Where user_name=? or phone=? ")
	if errs != nil {
		fmt.Println(errs)
		return
	}

	defer stmt.Close()
	row := stmt.QueryRow(name, name)

	conv := User{}

	row.Scan(&conv.Id, &conv.Merch_id, &conv.User_name, &conv.Password, &conv.Salt, &conv.Login_time, &conv.Login_ip, &conv.Last_login_ip, &conv.Last_login_time)

	return conv*/
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

func (u SysSecurityQuestion) GetSecurity(q_type int, lang int) interface{} {
	return nil
	/*var securitys []SysSecurityQuestion
	stmt, errs := Db.Prepare("SELECT id,question,type,lang FROM sys_security_question where type=? and lang=?")
	if errs != nil {
		fmt.Println(errs)
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(q_type, lang)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		conv := SysSecurityQuestion{}
		rows.Scan(&conv.Id, &conv.Question, &conv.Q_type, &conv.Lang)
		securitys = append(securitys, conv)
	}

	return securitys*/
}
