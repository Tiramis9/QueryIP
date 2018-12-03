package ky

import (
	"fmt"
	"game2/lib/encrypt"
	"game2/lib/utils"
	"encoding/json"
)

type KickUserOffLineReq struct {
	BaseReq
	Account string
}

type KickUserOffLineResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code int `json:"code"`
	} `json:"d"`
}

// 踢玩家下线
func (ky *GameKY) KickUserOffLine(req KickUserOffLineReq) (*KickUserOffLineResp, error) {
	param := fmt.Sprintf("s=%v&account=%v", OperationTypeKickUserOffline, req.Account)
	req.Param = encrypt.AESEncrypt(param, ky.DesKey)
	req.Key = encrypt.MD5(generateMd5Str(ky.Agent, req.TimeStamp, ky.Md5Key))
	data, err := utils.HttpPostForm(channelHandleUrl, baseReq2Map(req.BaseReq))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var resp KickUserOffLineResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}


type GetOrderInfoReq struct {
	BaseReq
	Account string
	OrderId string
}

type GetOrderInfoResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code   int `json:"code"`
		Status int `json:"status"` //-1不存在，0成功，2失败
	} `json:"d"`
}

// 查询订单状态
func (ky *GameKY) GetOrderInfo(req GetOrderInfoReq) (*GetOrderInfoResp, error) {
	param := fmt.Sprintf("s=%v&orderid=%v",
		OperationTypeQueryOrder, req.OrderId)
	req.Param = encrypt.AESEncrypt(param, ky.DesKey)
	req.Key = encrypt.MD5(generateMd5Str(ky.Agent, req.TimeStamp, ky.Md5Key))
	data, err := utils.HttpPostForm(channelHandleUrl, baseReq2Map(req.BaseReq))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var resp GetOrderInfoResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type IsOnlineReq struct {
	BaseReq
	Account string
}

type IsOnlineResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		Code   int `json:"code"`
		Status int `json:"status"` //-1不存在，0不在线，1在线
	} `json:"d"`
}

// 查询在线状态
func (ky *GameKY) IsOnline(req IsOnlineReq) (*IsOnlineResp, error) {
	param := fmt.Sprintf("s=%v&account=%v",
		OperationTypeQueryIsOnline, req.Account)
	req.Param = encrypt.AESEncrypt(param, ky.DesKey)
	req.Key = encrypt.MD5(generateMd5Str(ky.Agent, req.TimeStamp, ky.Md5Key))
	data, err := utils.HttpPostForm(channelHandleUrl, baseReq2Map(req.BaseReq))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var resp IsOnlineResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type GetTotalScoreReq struct {
	BaseReq
	Account string
}

type GetTotalScoreResp struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		TotalMoney float64 `json:"totalMoney"`
		FreeMoney  float64 `json:"freeMoney"`
		Status     int     `json:"status"` //-1不存在，0不在线，1在线
		Code       int     `json:"code"`
	} `json:"d"`
}

// 查询用户总分
func (ky *GameKY) GetUserTotalScore(req GetTotalScoreReq) (*GetTotalScoreResp, error) {
	param := fmt.Sprintf("s=%v&account=%v", OperationTypeGetTotalScore, req.Account)
	req.Param = encrypt.AESEncrypt(param, ky.DesKey)
	req.Key = encrypt.MD5(generateMd5Str(ky.Agent, req.TimeStamp, ky.Md5Key))
	data, err := utils.HttpPostForm(channelHandleUrl, baseReq2Map(req.BaseReq))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var resp GetTotalScoreResp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
