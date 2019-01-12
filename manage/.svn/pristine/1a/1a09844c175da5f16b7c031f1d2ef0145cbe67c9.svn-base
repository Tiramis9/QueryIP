package global

import "fmt"

const (
	OK               = 200
	ErrParamError    = 400
	ErrNotFound      = 404
	ErrNoEnoughMoney = 406
	ErrServerError   = 500
)

var statusDesc = map[int]string{
	OK:               "成功",
	ErrParamError:    "参数错误",
	ErrNotFound:      "请求找不到",
	ErrNoEnoughMoney: "余额不足",
	ErrServerError:   "服务器错误",
}

func StatusDesc(code int) string {
	if msg, ok := statusDesc[code]; ok {
		return msg
	}

	return fmt.Sprintf("未知错误：%v", code)
}
