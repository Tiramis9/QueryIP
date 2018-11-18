package action

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang_game_merchant/global/status"
	"golang_game_merchant/model"
)

/**********************************
//http://xx.com/v1/global_report
new_register	int	新注册会员 .
recharge_member	int	充值会员 .
recharge_amount	int	充值 .
withdraw_amount	int	提现 .
rebate	int	反水 .
bonus_amount	int	红利 .
effective_bet	int	有效投注
win_lost_amount	int	总输赢
effective_member	int	有效会员
win_rate	int	胜率
win	int	输赢结果
tips_list	string	打赏小费列表
tips_num	int	打赏数额
game_name	string	游戏名或者直播名
game_active_list	string	游戏活动列表
award	int	派奖总金额
pool	int	奖池贡献金
******************************/
type (
	Reportrequest struct {
		Begintime string `json:"start_time"` // 开始时间
		EndTime   string `json:"end_time"`   // 结束时间
		//Token     string `json:"token"`
	}
	ReportInfo struct {
		New_register    interface{} `json:"new_register"`    // 新注册会员 .
		Bonus_amount    interface{} `json:"bonus_amount"`    // 红利 ..
		Rebate          interface{} `json:"rebate"`          // 反水 .
		Recharge_amount interface{} `json:"recharge_amount"` // 充值金额 .
		Recharge_member interface{} `json:"recharge_member"` // 充值会员 .
		Withdraw_amount interface{} `json:"withdraw_amount"` // 提现 .
		/**********************待完成*********************************/
		Effective_bet    interface{} `json:"effective_bet"` // 有效投注
		Win              interface{} `json:"win"`           // 输赢结果
		Win_lost_amount  interface{} `json:"win_lost_amount"`
		Game_active_list interface{} `json:"game_active_list"`
		Game_name        interface{} `json:"game_name"`
		Effective_member interface{} `json:"effective_member"`
		Win_rate         interface{} `json:"win_rate"`
		Tips_list        interface{} `json:"tips_list"`
		Tips_num         interface{} `json:"tips_num"`
		Award            interface{} `json:"award"`
		Pool             interface{} `json:"pool"`
	}
	ReportResponse struct {
		List ReportInfo `json:"data"`
	}
)

func globalReportCheck(request *Reportrequest) (map[string]interface{}, error) {
	msg := make(map[string]interface{})
	if request.EndTime == "" || request.Begintime == "" {
		return nil, errors.New("request  invalid of time")
	}
	if request.Begintime != "" {
		if request.EndTime != "" {
			if request.EndTime < request.Begintime {
				return nil, errors.New("start time less end time")
			}
			msg["end_time"] = request.EndTime
		}
		msg["start_time"] = request.Begintime
	}
	return msg, nil
}
func GlobalReport(c *gin.Context) {
	var request Reportrequest
	if err := c.Bind(&request); err != nil {
		RespParamErr(c)
		return
	}
	// 参数入参检查
	msg, err := globalReportCheck(&request)
	if err != nil {
		RespParamErr(c)
		return
	}
	dataList, err := model.GetMerchantAnnouncement(model.Db, msg)
	if err != nil {
		RespServerErr(c)
		return
	}
	resp := new(ReportResponse)
	resp.List.New_register = dataList.NewRegister
	resp.List.Bonus_amount = dataList.BonusAmount
	resp.List.Rebate = dataList.Rebate
	resp.List.Recharge_amount = dataList.RechargeAmount
	resp.List.Recharge_member = dataList.RechargeMember
	resp.List.Withdraw_amount = dataList.WithdrawAmount
	RespJson(c, status.OK, resp)
}
