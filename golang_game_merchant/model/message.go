package model

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	Id         int    `json:"id"`
	MsgId      int    `json:"msg_id"`
	HaveRead   string `json:"have_read"`
	CreateTime string `json:"create_time"`
	ReadTime   string `json:"read_time"`
	UserId     string `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

/*
SELECT
	um.id,um.msg_id,um.have_read,um.create_time,um.read_time,um.user_id,sm.title,sm.content
FROM
	user_message AS um
LEFT JOIN
	sys_message AS sm ON sm.id=um.msg_id
WHERE
	(um.user_id='{id}')
LIMIT {limit} OFFSET {offset}
*/
func GetMessageList(db *gorm.DB, userId int, page int, pageCount int) ([]Message, error) {
	var msgList []Message
	if err := db.Table(`user_message AS um`).Joins(`
		LEFT JOIN sys_message AS sm ON sm.id=um.msg_id`).Select(`
		um.id,um.msg_id,um.have_read,um.create_time,um.read_time,um.user_id,sm.title,sm.content
	`).Where(`um.user_id=?`, userId).Offset((page - 1) * pageCount).Limit(pageCount).Find(&msgList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return msgList, nil
}

// SELECT count(*) FROM `user_message`  WHERE (user_id='{id}')
func GetMessageCount(db *gorm.DB, userId int) (int, error) {
	var count int
	if err := db.Table(`user_message`).Where(`user_id=?`, userId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

/*
UPDATE
	`user_message`
SET
	`have_read` = '1', `read_time` = '1541574995'
WHERE
	(`user_message`.`id` = '{id}') AND (`user_message`.`have_read` = '0')
*/
func ReadMessage(db *gorm.DB, id int, time int64) (bool, error) {
	if err := db.Table(`user_message`).Where(map[string]interface{}{
		"id":        id,
		"have_read": 0,
	}).Updates(map[string]interface{}{
		"have_read": 1,
		"read_time": time,
	}).Error; err != nil {
		return false, err
	}

	return true, nil
}
