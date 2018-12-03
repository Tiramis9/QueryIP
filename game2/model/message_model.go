package model

import "github.com/jinzhu/gorm"

type Message struct {
	Id         int    `json:"id"`
	MsgId      int    `json:"msg_id"`
	HaveRead   int    `json:"have_read"`
	CreateTime string `json:"create_time"`
	ReadTime   string `json:"read_time"`
	UserId     int    `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

type MessageInterface interface {
	GetMessageList()
	GetMessageCount()
	ReadMessage()
}

func GetMessageList(db *gorm.DB, userId int, page int, pageCount int) ([]Message, error) {
	var messList []Message
	if err := db.Table("user_message as um").Select("um.id,um.msg_id,um.have_read,um.create_time,um.read_time,um.user_id,sm.title,sm.content").
		Joins("LEFT JOIN sys_message sm ON sm.id=um.msg_id").Where("um.user_id=? ", userId).
		Offset((page - 1) * pageCount).Limit(pageCount).Find(&messList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return messList, nil
}

func GetMessageInfo(db *gorm.DB, id int, userId int) (*Message, error) {
	var m Message
	if err := db.Table("user_message um").Joins("LEFT JOIN sys_message sm ON sm.id=um.msg_id").Select("um.id,um.msg_id,um.have_read,"+
		"um.create_time,um.read_time,um.user_id,sm.title,sm.content").Where("um.id=? AND um.user_id=?", id, userId).Find(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func GetMessageCount(db *gorm.DB, userId int) (int,error) {
	var total int
	if err:=db.Table("user_message").Where("user_id=?",userId).Count(&total).Error;err!=nil{
		return 0,err
	}
	return total,nil
}

func ReadMessage(db *gorm.DB, id int, time int64) (bool,error) {
	//对于未阅读的消息更新
	fields := make(map[string]interface{})
	fields["have_read"] = 1
	fields["read_time"] = time
	if err:=db.Table("user_message").Where("id=? AND have_read=0",id).Update(fields).Error;err!=nil{
		return false, err
	}
	return true, nil
}
