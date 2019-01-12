package dg

import (
	"encoding/json"
	"fmt"
	"game2/lib/utils"
)

type TryPlayLoginReq struct {
	Token  string `json:"token"`
	Random string `json:"random"`
	Lang   string `json:"lang"`
	Device int    `json:"device"`
}
type TryPlayLoginResp struct {
	CodeId int      `json:"codeId"`
	Token  string   `json:"token"`
	Random string   `json:"random"`
	List   []string `json:"list"`
}

// 会员试玩登入
func (g *GameDG) TryPlayLogin(req TryPlayLoginReq) (*TryPlayLoginResp, error) {
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(TryPlayLoginUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp TryPlayLoginResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type UpdateUserInfoMember struct {
	UserName string  `json:"username"`
	Password string  `json:"password"`
	WinLimit float64 `json:"winLimit"`
	Status   int     `json:"status"`
}
type UpdateUserInfoReq struct {
	Token  string               `json:"token"`
	Random string               `json:"random"`
	Member UpdateUserInfoMember `json:"member"`
}
type UpdateUserInfoResp struct {
	CodeId int    `json:"codeId"`
	Token  string `json:"token"`
	Random string `json:"random"`
}

// 修改会员信息
func (g *GameDG) UpdateUserInfo(req UpdateUserInfoReq) (*UpdateUserInfoResp, error) {
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(UpdateUserInfoUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp UpdateUserInfoResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type (
	UpdateUserLimitGroupMember struct {
		Username string `json:"username"`
	}
	UpdateUserLimitGroupReq struct {
		Token  string                     `json:"token"`
		Random string                     `json:"random"`
		Data   string                     `json:"data"`
		Member UpdateUserLimitGroupMember `json:"member"`
	}
	UpdateUserLimitGroupResp struct {
		CodeId string `json:"codeId"`
		Token  string `json:"token"`
		Random string `json:"random"`
	}
)

// 修改会员限红组
func (g *GameDG) UpdateUserLimitGroup(req UpdateUserLimitGroupReq) (*UpdateUserLimitGroupResp, error) {
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(UpdateUserLimitGroupUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp UpdateUserLimitGroupResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type (
	GetAgentUserListInfoReq struct {
		Token  string `json:"token"`
		Random string `json:"random"`
	}
	GetAgentUserListInfoResp struct {
		CodeId int           `json:"codeId"`
		Token  string        `json:"token"`
		Random string        `json:"random"`
		List   []interface{} `json:"list"`
	}
)

// 获取当前代理下在DG在线会员信息
func (g *GameDG) GetAgentUserListInfo(req GetAgentUserListInfoReq) (*GetAgentUserListInfoResp, error) {
	data, err := utils.HttpPostJson(g.Host+fmt.Sprintf(GetAgentUserListInfoUrl, g.Agent), req)
	if err != nil {
		return nil, err
	}

	var resp GetAgentUserListInfoResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
