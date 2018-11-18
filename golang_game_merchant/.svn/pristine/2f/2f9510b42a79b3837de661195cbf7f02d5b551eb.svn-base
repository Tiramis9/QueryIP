package model

import "github.com/jinzhu/gorm"

//位置：merchant_website
type MemberRegistration struct {
	RegEmail            int `json:"reg_email"`  //in,out
	RegPayPass          int `json:"reg_pay_pass"` //in,out
	RegPhone            int `json:"reg_phone"` //in,out
	RegSecurityQuestion int `json:"reg_security_question"` //in,out
	RegTrueName         int `json:"reg_true_name"` //in,out
}

func QueryMemberRegistration(db *gorm.DB, mid int) (*MemberRegistration, error) {
	var mr MemberRegistration
	if err := db.Table("merchant_website").Where("merchant_id = ?", mid).Find(&mr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &mr, nil
}

func ModifyMemberRegistration(db *gorm.DB, req *MemberRegistration, mid int) error {
	if err := db.Table("merchant_website").Where("merchant_id = ?", mid).Updates(*req).Error; err != nil {
		return err
	}
	return nil
}
