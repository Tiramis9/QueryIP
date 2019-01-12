package status

import "fmt"

const (
	OK               = 200
	ErrParamError    = 400
	ErrUnauthorized  = 401
	ErrNotFound      = 404
	ErrNoEnoughMoney = 406
	ErrServerError   = 500
)

var statusDesc = map[int]string{
	OK:               "success",
	ErrParamError:    "parameter error",
	ErrUnauthorized:  "unauthorized",
	ErrNotFound:      "not found",
	ErrNoEnoughMoney: "no enough money",
	ErrServerError:   "server error",
}

func Description(code int) string {
	if msg, ok := statusDesc[code]; ok {
		return msg
	}

	return fmt.Sprintf("unknown error[%v]", code)
}
