package game

import (
	"fmt"
	"errors"
)

var ErrNoEnoughMoney = errors.New("no enough money")

type Game interface {
	Register(req map[string]interface{}) (resp interface{}, err error)
	Login(req map[string]interface{}) (resp interface{}, err error)
	GetBalance(req map[string]interface{}) (resp interface{}, err error)
	Account2GameTransfer(req map[string]interface{}) (resp interface{}, err error)
	Game2AccountTransfer(req map[string]interface{}) (resp interface{}, err error)
	QueryRecord(req map[string]interface{}) (resp interface{}, err error)
	GetPrefix()(string)
}

type Instance func() Game

var adapter = make(map[string]Instance)

func Register(name string, game Instance) {
	if _, ok := adapter[name]; ok {
		panic("game: Register called twice for adapter " + name)
	}
	adapter[name] = game
}

func NewGame(name string) (game Game, err error) {
	instanceFunc, ok := adapter[name]
	if !ok {
		err = fmt.Errorf("game: unknown adapter name %v (forgot to import?)", name)
		return
	}

	return instanceFunc(), nil
}
