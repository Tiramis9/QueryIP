package action

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)

type (
	RebateLog struct {
		Id           int    `json:"id"`
		EffectiveBet int    `json:"effective_bet"` //有效投注
		Rebate       int    `json:"rebate"`        //反水金额
		UserName     string `json:"user_name"`     //用户账号
		CreateTime   int    `json:"create_time"`   //创建时间
		Status       int    `json:"status"`        //0.未操作；1.已反水; 2.反水驳回
		//UpdateTime   int    `json:"update_time"`   //更新时间
		//MerchantId   int   `json:"merchant_id"`   //商户id
		//Operator     int   `json:"operator"`      //操作员
	}

	RebateLogReq struct {
		StartTime int    `json:"start_time"` //开始时间
		EndTime   int    `json:"end_time"`   //结束时间
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
		UserName  string `json:"user_name"`  //用户账号，支持模糊查询
	}

	RebateLogListResp struct {
		List        []RebateLog `json:"list"`
		RebateTotal interface{} `json:"rebate_total"` //总反水数
		Total       interface{} `json:"total"`        //总数
	}
)

type (
	RebateSuccessReq struct {
		StartTime int `json:"start_time"` //开始时间
		EndTime   int `json:"end_time"`   //结束时间
		Page      int `json:"page"`       //页码
		PageCount int `json:"page_count"` //每页显示数量
	}

	RebateSuccessLog struct {
		Id           int    `json:"id"`
		EffectiveBet int    `json:"effective_bet"` //有效投注
		Rebate       int    `json:"rebate"`        //反水金额
		UserName     string `json:"user_name"`     //用户账号
		CreateTime   int    `json:"create_time"`   //创建时间
	}

	RebateSuccessList struct {
		List              []RebateSuccessLog `json:"detail_list"`
		RebateDate        int                `json:"rebate_date"`         //反水日期
		RebateMemberNum   int                `json:"rebate_member_num"`   //反水人数
		EffectiveBetTotal int                `json:"effective_bet_total"` //有效投注总额
		RebateTotal       int                `json:"rebate_total"`        //反水金额总额
		OperateDate       int                `json:"operate_date"`        //操作日期
	}

	RebateSuccessResp struct {
		List  []RebateSuccessList `json:"list"`
		Total int                 `json:"total"` //总数
	}
)

// 反水记录查询参数合法性检查
func RebateLogReqCheck(req *RebateLogReq) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if req.StartTime > 0 {
		m["start_time"] = req.StartTime
		if req.EndTime > 0 {
			if req.EndTime < req.StartTime {
				return nil, errors.New("start time less end time")
			}
			m["end_time"] = req.EndTime
		}
	} else if req.EndTime > 0 {
		m["end_time"] = req.EndTime
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)

	if req.UserName != "" {
		m["user_name"] = req.UserName
	}

	return m, nil
}

//反水记录查询
func GetRebateLogList(c *gin.Context) {
	var req RebateLogReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := RebateLogReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1

	list, count, sum, err := model.RebateLogList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	// 组装数据返回给前端显示
	resp := RebateLogListResp{
		List:        make([]RebateLog, 0),
		RebateTotal: sum.RebateTotal,
		Total:       count,
	}

	for i := range list {
		temp := RebateLog{
			Id:           list[i].Id,
			UserName:     list[i].UserName,
			EffectiveBet: list[i].EffectiveBet,
			Rebate:       list[i].Rebate,
			Status:       list[i].Status,
			CreateTime:   list[i].CreateTime,
			//UpdateTime:   list[i].UpdateTime,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)

}

func RebateSuccessReqCheck(req *RebateSuccessReq) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if req.StartTime > 0 {
		m["start_time"] = req.StartTime
		if req.EndTime > 0 {
			if req.EndTime < req.StartTime {
				return nil, errors.New("start time less end time")
			}
			m["end_time"] = req.EndTime
		}
	} else if req.EndTime > 0 {
		m["end_time"] = req.EndTime
	}
	return m, nil
}

//反水成功历史查询
func GetRebateLogSuccessList(c *gin.Context) {
	var req RebateSuccessReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := RebateSuccessReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1

	list, count, sum, err := model.RebateLogList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	// 组装数据返回给前端显示
	resp := RebateLogListResp{
		List:        make([]RebateLog, 0),
		RebateTotal: sum.RebateTotal,
		Total:       count,
	}

	for i := range list {
		temp := RebateLog{
			Id:           list[i].Id,
			UserName:     list[i].UserName,
			EffectiveBet: list[i].EffectiveBet,
			Rebate:       list[i].Rebate,
			Status:       list[i].Status,
			CreateTime:   list[i].CreateTime,
			//UpdateTime:   list[i].UpdateTime,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)

}
