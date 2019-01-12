package status

import "fmt"

const (
	OK                  = 200
	ErrParamError       = 400
	ErrUnauthorized     = 401
	ErrTokenExpired     = 402
	ErrNotFound         = 404
	ErrUserOrPassError  = 408
	ErrUserExist        = 409
	ErrUserPasswodError = 410
	ErrNoEnoughMoney    = 406
	ErrPhoneExist       = 411
	ErrPassError        = 412
	ErrBankNoAdd        = 413
	ErrIpBanned         = 414
	ErrNoTrueName       = 415
	ErrPayPassError     = 416
	ErrAnswerError      = 417
	ErrAgentCodeError   = 418
	ErrUserNotExist     = 419
	ErrSecurityExist    = 420
	ErrNoPayPass       = 421
	ErrPayPassHadSet       = 422
	ErrServerError      = 500
)

var statusDesc = map[int]string{
	OK:                 "success",
	ErrParamError:      "parameter error",
	ErrTokenExpired:    "token expired",
	ErrUnauthorized:    "unauthorized",
	ErrNotFound:        "not found",
	ErrNoEnoughMoney:   "no enough money",
	ErrUserOrPassError: "user or password error",
	ErrPhoneExist:      "phone exist",
	ErrPassError:       "password error",
	ErrUserExist:       "user have existed",
	ErrBankNoAdd:       "user no add bank",
	ErrNoTrueName:      "user no truename validate",
	ErrPayPassError:    "user paypass error",
	ErrAnswerError:     "answer error",
	ErrAgentCodeError:  "angent code error",
	ErrUserNotExist:    "user not exist",
	ErrSecurityExist:   "security exist",
	ErrNoPayPass: "no pay password",
	ErrPayPassHadSet:"pay password had set",
	ErrServerError:     "server error",
	ErrIpBanned:        "ip banned",
}

func Description(code int) string {
	if msg, ok := statusDesc[code]; ok {
		return msg
	}

	return fmt.Sprintf("unknown error[%v]", code)
}
