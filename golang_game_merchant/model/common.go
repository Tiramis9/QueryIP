package model

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoEnoughMoney  = errors.New("no enough money")
)
