package status

import "fmt"

const (
	OK                       = 200
	ErrParamError            = 400
	ErrUnauthorized          = 401
	ErrTokenExpired          = 402
	ErrForbidden             = 403
	ErrNotFound              = 404
	ErrGameAccountNotFound   = 405
	ErrNoEnoughMoney         = 406
	ErrUserOrPassError       = 408
	ErrUserExist             = 409
	ErrRoleExists            = 410
	ErrPhoneExist            = 411
	ErrPassError             = 412
	ErrNoWebSiteInfo         = 413
	ErrServerError           = 500
	ErrMerchantBailNotEnough = 414
	ErrGameTrans             = 415
	ErrBalanceNotEnough      = 416
	MerchantBailNotEnough    = 417
	AlreadyGeneratedRebate   = 419
	ErrWebSiteMaintenance    = 508
	ErrGeneratedRebate       = 420
)

var statusDesc = map[int]string{
	OK:                     "success",
	ErrParamError:          "parameter error",
	ErrTokenExpired:        "token expired",
	ErrUnauthorized:        "unauthorized",
	ErrForbidden:           "forbidden",
	ErrNotFound:            "not found",
	ErrNoEnoughMoney:       "no enough money",
	ErrUserOrPassError:     "user or password error",
	ErrPhoneExist:          "phone exist",
	ErrPassError:           "password error",
	ErrUserExist:           "user have existed",
	ErrRoleExists:          "role already exists",
	ErrNoWebSiteInfo:       "no website info",
	ErrServerError:         "server error",
	ErrWebSiteMaintenance:  "system under maintenance",
	MerchantBailNotEnough:  "merchant bail not enough",
	AlreadyGeneratedRebate: "already generated rebate",
	ErrGameAccountNotFound: "The game account does not exist",
	ErrGeneratedRebate:     "parameter error,Please check the time, the backwater generation date cannot be today and can only be operated after 14:00",
}

func Description(code int) string {
	if msg, ok := statusDesc[code]; ok {
		return msg
	}

	return fmt.Sprintf("unknown error[%v]", code)
}
