package logic

import (
	"game2/service"
)

func UserInfoByRedis(token string) map[string]interface{} {
	//用户信息
	mapInfo, err := service.RedisGetMap(token)
	if err != nil {
		return nil
	}
	return mapInfo
}
