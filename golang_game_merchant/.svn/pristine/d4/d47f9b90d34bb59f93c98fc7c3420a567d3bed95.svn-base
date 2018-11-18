package action

import (
	"errors"
	"golang_game_merchant/global/status"
	"golang_game_merchant/lib/utils"
	"golang_game_merchant/model"
	"time"

	"golang_game_merchant/lib/designpattern/builder"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	// （财务管理）帐变记录
	MemberBillReq struct {
		StartTime int64 `json:"start_time"` //开始时间时间戳
		EndTime   int64 `json:"end_time"`   //结束时间时间戳
		UserType  int   `json:"user_type"`  //用户类型，1-正常会员，2-代理，零值表示不考虑该字段，下同
		Type      int   `json:"type"`       //1-入款，2-出款
		Code      int   `json:"code"`       //100，充值; 200,提现; 301转入; 302 转出; 400 红利; 500 返利; 600 丢失补款; 700 多出扣款
		Page      int   `json:"page"`       //页码
		PageCount int   `json:"page_count"` //每页显示的数量
	}
	MemberBillInfo struct {
		Id         interface{} `json:"id"`          //记录id
		OrderSn    interface{} `json:"order_sn"`    //账变编号
		UserName   interface{} `json:"user_name"`   //会员账号
		TrueName   interface{} `json:"true_name"`   //真实姓名
		Type       interface{} `json:"type"`        //账变类型 1：加，-1：减
		Code       interface{} `json:"code"`        //交易类型 100，充值; 200,提现; 300转移; 400 红利; 500 返利
		OldBalance interface{} `json:"old_balance"` //变动前余额
		SettAmt    interface{} `json:"sett_amt"`    //变动金额
		Balance    interface{} `json:"balance"`     //变动后余额
		Memo       interface{} `json:"memo"`        //变动说明
		CreateTime interface{} `json:"create_time"` //变动时间
	}
	MemberBillResp struct {
		List  []MemberBillInfo `json:"list"`  //帐变数据列表
		Total interface{}      `json:"total"` //总数
	}

	// （财务管理）充值记录-在线支付
	RechargeBillReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		Platform  string `json:"platform"`   //支付平台
		PayType   string `json:"pay_type"`   //支付类型
		UserName  string `json:"user_name"`  //用户账号，支持模糊查询
		OrderSn   string `json:"order_sn"`   //订单账号，模糊查询
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
	}
	RechargeBillInfo struct {
		Id          interface{} `json:"id"`           //记录id
		OrderSn     interface{} `json:"order_sn"`     //流水单号
		UserName    interface{} `json:"user_name"`    //会员账号
		TrueName    interface{} `json:"true_name"`    //真实姓名
		PayMoney    interface{} `json:"pay_money"`    //充值金额
		SysPayType  interface{} `json:"sys_pay_type"` //支付类型
		Platform    interface{} `json:"platform"`     //支付平台
		Interface   interface{} `json:"interface"`    //接口名称
		CreateTime  interface{} `json:"create_time"`  //创建时间
		SuccessTime interface{} `json:"success_time"` //成功时间
		Status      interface{} `json:"status"`       //0-失败，1-待支付，2-支付成功
	}
	RechargeBillResp struct {
		List               []RechargeBillInfo `json:"list"`                 //充值记录数据列表
		Total              interface{}        `json:"total"`                //总数
		RechargeNum        interface{}        `json:"recharge_num"`         //充值数
		RechargeSuccessNum interface{}        `json:"recharge_success_num"` //成功数
		RechargeSuccessSum interface{}        `json:"recharge_success_sum"` //成功金额
	}

	// （财务管理）充值记录-转账汇款
	RechargeBillTransferReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		Type      string `json:"type"`       //转账类型
		UserName  string `json:"user_name"`  //用户账号，支持模糊查询
		OrderSn   string `json:"order_sn"`   //订单账号，模糊查询
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
	}
	RechargeBillTransferInfo struct {
		Id             interface{} `json:"id"`
		OrderSn        interface{} `json:"order_sn"`
		UserName       interface{} `json:"user_name"`
		TrueName       interface{} `json:"true_name"`
		TransAccount   interface{} `json:"trans_account"`
		PayMoney       interface{} `json:"pay_money"`
		SysPayType     interface{} `json:"sys_pay_type"`
		TransToAccount interface{} `json:"trans_to_account"`
		TransTime      interface{} `json:"trans_time"`
		CheckTime      interface{} `json:"check_time"`
		Status         interface{} `json:"status"`
	}
	RechargeBillTransferResp struct {
		List               []RechargeBillTransferInfo `json:"list"`
		Total              interface{}                `json:"total"`
		TransferNum        interface{}                `json:"transfer_num"`
		TransferSuccessNum interface{}                `json:"transfer_success_num"`
		TransferSuccessSum interface{}                `json:"transfer_success_sum"`
	}

	// （财务管理）充值记录-后台加款
	RechargeBillBackAddReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		UserName  string `json:"user_name"`  //用户账号，支持模糊查询
		OrderSn   string `json:"order_sn"`   //订单账号，模糊查询
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
	}
	RechargeBillBackAddInfo struct {
		Id         interface{} `json:"id"`          //记录id
		OrderSn    interface{} `json:"order_sn"`    //流水单号
		UserName   interface{} `json:"user_name"`   //会员账号
		TrueName   interface{} `json:"true_name"`   //真实姓名
		Money      interface{} `json:"money"`       //金额
		Memo       interface{} `json:"memo"`        //备注
		Operator   interface{} `json:"operator"`    //操作员
		CreateTime interface{} `json:"create_time"` //操作时间
	}
	RechargeBillBackAddResp struct {
		List              []RechargeBillBackAddInfo `json:"list"`
		Total             interface{}               `json:"total"`                //总数
		BackAddSuccessNum interface{}               `json:"back_add_success_num"` //后台操作数
		BackAddSuccessSum interface{}               `json:"back_add_success_sum"` //成功金额
	}

	// （财务管理）提现记录会员提现
	WithdrawBillReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		UserName  string `json:"user_name"`  //用户账号
		OrderSn   string `json:"order_sn"`   //流水单号
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
	}
	WithdrawBillInfo struct {
		Id          interface{} `json:"id"`           //记录id
		OrderSn     interface{} `json:"order_sn"`     //流水单号
		UserName    interface{} `json:"user_name"`    //用户名
		ClassName   interface{} `json:"class_name"`   //层级名称
		TrueName    interface{} `json:"true_name"`    //真实姓名
		BankName    interface{} `json:"bank_name"`    //银行名称
		CardNo      interface{} `json:"card_no"`      //卡号
		Money       interface{} `json:"money"`        //金额
		CreateTime  interface{} `json:"create_time"`  //申请时间
		ApproveTime interface{} `json:"approve_time"` //审核时间
		Memo        interface{} `json:"memo"`         //备注
		Status      interface{} `json:"status"`       //状态 1-审核中，2-审核成功，3-审核失败，4-打款成功，5-打款失败
	}
	WithdrawBillResp struct {
		List               []WithdrawBillInfo `json:"list"`                 //提现数据列表
		Total              interface{}        `json:"total"`                //总数
		WithdrawNum        interface{}        `json:"withdraw_num"`         //提现数
		WithdrawSuccessNum interface{}        `json:"withdraw_success_num"` //成功数
		WithdrawSuccessSum interface{}        `json:"withdraw_success_sum"` //成功金额
	}

	//（财务管理）提现记录后台扣款
	WithdrawBillBackReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		UserName  string `json:"user_name"`  //用户账号
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
	}
	WithdrawBillBackInfo struct {
		Id         interface{} `json:"id"`          //记录id
		OrderSn    interface{} `json:"order_sn"`    //流水单号
		UserName   interface{} `json:"user_name"`   //用户账号
		TrueName   interface{} `json:"true_name"`   //真实姓名
		Amount     interface{} `json:"amount"`      //金额
		Memo       interface{} `json:"memo"`        //备注
		Operator   interface{} `json:"operator"`    //操作员账号
		CreateTime interface{} `json:"create_time"` //操作时间
	}
	WithdrawBillBackResp struct {
		List               []WithdrawBillBackInfo `json:"list"`                 //后台扣款数据列表
		Total              interface{}            `json:"total"`                //总数
		WithdrawNum        interface{}            `json:"withdraw_num"`         //提现数
		WithdrawSuccessNum interface{}            `json:"withdraw_success_num"` //成功数
		WithdrawSuccessSum interface{}            `json:"withdraw_success_sum"` //成功金额
	}

	// （财务管理）帐变记录(转账)
	RechargeTransBillReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
		UserName  string `json:"user_name"`  //用户账号
	}
	RechargeTransBillInfo struct {
		CheckTime      interface{} `json:"check_time"`       //审核时间
		CreateTime     interface{} `json:"create_time"`      //申请时间
		BillNo         interface{} `json:"bill_no"`          //流水单号
		Amount         interface{} `json:"amount"`           //金额
		ContactName    interface{} `json:"contact_name"`     //真实姓名
		MerchantName   interface{} `json:"merchant_name"`    //会员账号
		Status         interface{} `json:"status"`           //0-失败,1-待支付,2-支付成功
		TransAccount   interface{} `json:"trans_account"`    //会员转账户名
		TransToAccount interface{} `json:"trans_to_account"` //转入账户
		TrueName       interface{} `json:"true_name"`        //真实姓名
	}
	RechargeTransBillResp struct {
		List               []RechargeTransBillInfo `json:"list"`                 //转账数据列表
		Total              interface{}             `json:"total"`                //总数
		RechargeNum        interface{}             `json:"recharge_num"`         //充值数
		RechargeSuccessNum interface{}             `json:"recharge_success_num"` //成功数
		RechargeSuccessSum interface{}             `json:"recharge_success_sum"` //成功金额
	}

	//（财务管理）游戏转账
	GameTransBillReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		UserName  string `json:"user_name"`  //用户账号
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
	}
	GameTransBillInfo struct {
		Id         interface{} `json:"id"`          //记录id
		BillNo     interface{} `json:"bill_no"`     //流水单号
		UserName   interface{} `json:"user_name"`   //会员账号
		TrueName   interface{} `json:"true_name"`   //真实姓名
		Type       interface{} `json:"type"`        //出入帐方向 1.中心账户向游戏账户; 2.游戏账户向中心账户
		Money      interface{} `json:"money"`       //金额
		CreateTime interface{} `json:"create_time"` //转账时间
		Ok         interface{} `json:"ok"`          //状态 1成功;2失败
	}
	GameTransBillResp struct {
		List  []GameTransBillInfo `json:"list"`  //游戏转账数据列表
		Total interface{}         `json:"total"` //总数
	}

	QueryUserBalanceReq struct {
		UserName string `json:"user_name" binding:"required"` //会员账号
	}
	QueryUserBalanceResp struct {
		Id       interface{} `json:"id"`
		UserName interface{} `json:"user_name"` //会员账号
		Balance  interface{} `json:"balance"`   //账户余额
	}

	QueryGameBalanceReq struct {
		GameId int `json:"game_id" binding:"required"` //游戏id
	}
	QueryGameBalanceResp struct {
		GameId  interface{} `json:"game_id"` //游戏id
		Balance interface{} `json:"balance"` //余额
	}

	OnlinePaymentsListReq struct {
		PayTag    int `json:"pay_tag" binding:"required"` //1-在线支付，2-转账汇款
		PayTypeId int `json:"pay_type_id"`                //支付类型id
		Page      int `json:"page"`                       //页码
		PageCount int `json:"page_count"`                 //每页显示数量
	}
	OnlinePaymentsInfo struct {
		Id         interface{} `json:"id"`           //记录id
		Code       interface{} `json:"code"`         //接口名称
		DayStopMax interface{} `json:"day_stop_max"` //单日停用上限
		SimpleMax  interface{} `json:"single_max"`   //单笔最高充值
		SimpleMin  interface{} `json:"single_min"`   //单笔最低充值
		Sort       interface{} `json:"sort"`         //排序值
		Status     interface{} `json:"status"`       //状态1.启用；2.禁用
		PayTypeId  interface{} `json:"pay_type_id"`  //支付类型id
		MerchNo    interface{} `json:"merch_no"`     //商户编号
		CreateTime interface{} `json:"create_time"`  //创建时间
		UpdateTime interface{} `json:"update_time"`  //更新时间
		Remark     interface{} `json:"remark"`       //备注
		Account    interface{} `json:"account"`      //账户名称
		Qrcode     interface{} `json:"qrcode"`       //二维码
		Url        interface{} `json:"url"`          //支付接口域名
		PayType    interface{} `json:"pay_type"`     //支付类型
		Platform   interface{} `json:"platform"`     //支付平台
	}
	OnlinePaymentsListResp struct {
		List  []OnlinePaymentsInfo `json:"list"`  //在线支付接口数据列表
		Total interface{}          `json:"total"` //总数
	}

	FinanceListReq struct {
		StartTime int64  `json:"start_time"` //开始时间
		EndTime   int64  `json:"end_time"`   //结束时间
		Page      int    `json:"page"`       //页码
		PageCount int    `json:"page_count"` //每页显示数量
		UserName  string `json:"user_name"`  //会员账号
		SortBy    int    `json:"sort_by"`    //排序，1-输赢金额，2-充值金额，3-提现金额，4-红利金额，5-反水金额，6-手续费
		Order     int    `json:"order"`      //0-或者不传降序，非0升序
	}
	FinanceInfo struct {
		UserId         interface{} `json:"user_id"`         //用户id
		UserName       interface{} `json:"user_name"`       //会员账号
		TrueName       interface{} `json:"true_name"`       //真实姓名
		EffectBet      interface{} `json:"effect_bet"`      //有效投注
		WinLostAmount  interface{} `json:"win_lost_amount"` //输赢金额
		RechargeAmount interface{} `json:"recharge_amount"` //充值金额
		WithdrawAmount interface{} `json:"withdraw_amount"` //提现金额
		BonusAmount    interface{} `json:"bonus_amount"`    //红利金额
		Rebate         interface{} `json:"rebate"`          //反水
		Fee            interface{} `json:"fee"`             //手续费
	}
	FinanceListResp struct {
		List              []FinanceInfo `json:"list"`
		Total             interface{}   `json:"total"`
		MemberSum         interface{}   `json:"member_sum"`
		EffectBetSum      interface{}   `json:"effect_bet_sum"`
		WinLostAmountSum  interface{}   `json:"win_lost_amount_sum"`
		RechargeAmountSum interface{}   `json:"recharge_amount_sum"`
		WithdrawAmountSum interface{}   `json:"withdraw_amount_sum"`
		BonusAmountSum    interface{}   `json:"bonus_amount_sum"`
		RebateSum         interface{}   `json:"rebate_sum"`
		FeeSum            interface{}   `json:"fee_sum"`
	}

	CenterAccountBalanceSwitchReq struct {
		UserId int     `json:"user_id" binding:"required"` //会员id
		Code   int     `json:"code" binding:"required"`    //100充值; 200提现; 301转入; 302转出; 400红利; 500返利; 600丢失补款; 700多出扣款
		Amount float64 `json:"amount" binding:"required"`  //操作金额
		Memo   string  `json:"memo" binding:"required"`    //操作原因
	}

	ThirdAccountBalanceSwitchReq struct {
		UserId int     `json:"user_id" binding:"required"` //会员id
		GameId int     `json:"game_id" binding:"required"` //游戏id
		Type   int     `json:"type" binding:"required"`    //1.第三方加款(不入帐变) 2.第三方扣款(不入帐变) 3.第三方账户->中心账户 4.中心账户->第三方账户
		Amount float64 `json:"amount" binding:"required"`  //操作金额
		Memo   string  `json:"memo" binding:"required"`    //操作原因
	}

	OnlinePaymentsAddReq struct {
		PayTag     int     `json:"pay_tag" binding:"required"`
		PayTypeId  int     `json:"pay_type_id" binding:"required"`
		Sort       int     `json:"sort" binding:"required"`
		Code       string  `json:"code" binding:"required"`
		MerchNo    string  `json:"merch_no"`   //在线支付参数
		Md5Key     string  `json:"md5_key"`    //在线支付参数
		PublicKey  string  `json:"public_key"` //在线支付参数
		SecretKey  string  `json:"secret_key"` //在线支付参数
		SimpleMin  float64 `json:"single_min" binding:"required"`
		SimpleMax  float64 `json:"single_max" binding:"required"`
		DayStopMax float64 `json:"day_stop_max" binding:"required"`
		Status     int     `json:"status" binding:"required"`
		Url        string  `json:"url"`     //在线支付参数
		Account    string  `json:"account"` //转账汇款参数
		Remark     string  `json:"remark"`  //备注
	}

	OnlinePaymentsEditReq struct {
		Id         int     `json:"id" binding:"required"`
		PayTag     int     `json:"pay_tag" binding:"required"`
		PayTypeId  int     `json:"pay_type_id" binding:"required"`
		Sort       int     `json:"sort" binding:"required"`
		Code       string  `json:"code" binding:"required"`
		MerchNo    string  `json:"merch_no"`   //在线支付参数
		Md5Key     string  `json:"md5_key"`    //在线支付参数
		PublicKey  string  `json:"public_key"` //在线支付参数
		SecretKey  string  `json:"secret_key"` //在线支付参数
		SimpleMin  float64 `json:"single_min" binding:"required"`
		SimpleMax  float64 `json:"single_max" binding:"required"`
		DayStopMax float64 `json:"day_stop_max" binding:"required"`
		Status     int     `json:"status" binding:"required"`
		Url        string  `json:"url"`     //在线支付参数
		Account    string  `json:"account"` //转账汇款参数
		Remark     string  `json:"remark"`
	}

	CreditLimitListReq struct {
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
		Page      int   `json:"page"`
		PageCount int   `json:"page_count"`
	}
	CreditLimitInfo struct {
		BillNo  interface{} `json:"bill_no"`
		Type    interface{} `json:"type"`
		Amount  interface{} `json:"amount"`
		OldBail interface{} `json:"old_bail"`
		NewBail interface{} `json:"new_bail"`
		Remark  interface{} `json:"remark"`
	}
	CreditLimitListResp struct {
		List  []CreditLimitInfo `json:"list"`
		Total interface{}       `json:"total"`
	}

	CreditLimitTransferListReq struct {
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
		Page      int   `json:"page"`
		PageCount int   `json:"page_count"`
	}
	CreditLimitTransferInfo struct {
		Id             interface{} `json:"id"`
		Amount         interface{} `json:"amount"`
		PayChannel     interface{} `json:"pay_channel"`
		PayType        interface{} `json:"pay_type"`
		OldBalance     interface{} `json:"old_balance"`
		NewBalance     interface{} `json:"new_balance"`
		TransAccount   interface{} `json:"trans_account"`
		TransToAccount interface{} `json:"trans_to_account"`
		CreateTime     interface{} `json:"create_time"`
		Remark         interface{} `json:"remark"`
		Status         interface{} `json:"status"`
		BillNo         interface{} `json:"bill_no"`
	}
	CreditLimitTransferListResp struct {
		List  []CreditLimitTransferInfo `json:"list"`
		Total interface{}               `json:"total"`
	}
)

func memberBillReqCheck(req *MemberBillReq) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if req.UserType != 0 && req.UserType != 1 && req.UserType != 2 {
		return nil, errors.New("user type error")
	} else if req.UserType == 1 || req.UserType == 2 {
		m["user_type"] = req.UserType
	}
	if req.Type != 0 && req.Type != 1 && req.Type != 2 {
		return nil, errors.New("type error")
	} else if req.Type == 1 || req.Type == 2 {
		m["type"] = req.Type
	}
	if req.Code != 0 {
		if !utils.IsIntContains(MemBillTypeList, req.Code) {
			return nil, errors.New("code error")
		} else {
			m["code"] = req.Code
		}
	}
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
	return m, nil
}

// （财务管理）帐变记录接口
func MemberBill(c *gin.Context) {
	var req MemberBillReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := memberBillReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 数据库查询数据
	//todo: get merchantId from token
	merchantId := 1
	list, count, err := model.GetMemberBillList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	// 组装数据返回给前端显示
	resp := MemberBillResp{
		List:  make([]MemberBillInfo, 0),
		Total: count,
	}

	for i := range list {
		temp := MemberBillInfo{
			Id:         list[i].Id,
			OrderSn:    list[i].OrderSn,
			UserName:   list[i].UserName,
			TrueName:   list[i].TrueName,
			Type:       list[i].Type,
			Code:       list[i].Code,
			OldBalance: list[i].OldBalance,
			SettAmt:    list[i].SettAmt,
			Balance:    list[i].Balance,
			Memo:       list[i].Memo,
			CreateTime: list[i].CreateTime,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)

}

func rechargeBillReqCheck(req *RechargeBillReq) (map[string]interface{}, error) {
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
	if req.Platform != "" {
		m["platform"] = req.Platform
	}
	if req.PayType != "" {
		m["pay_type"] = req.PayType
	}
	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if req.OrderSn != "" {
		m["order_sn"] = req.OrderSn
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)
	return m, nil
}

// （财务管理）充值记录（在线支付）接口
func RechargeBillOnlinePay(c *gin.Context) {
	var req RechargeBillReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := rechargeBillReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, rsn, rss, err := model.GetRechargeBillOnlinePayList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	resp := RechargeBillResp{
		List:               make([]RechargeBillInfo, 0),
		Total:              count,
		RechargeNum:        count,
		RechargeSuccessNum: rsn,
		RechargeSuccessSum: rss,
	}
	for i := range list {
		temp := RechargeBillInfo{
			Id:          list[i].Id,
			OrderSn:     list[i].OrderSn,
			UserName:    list[i].UserName,
			TrueName:    list[i].TrueName,
			PayMoney:    list[i].PayMoney,
			SysPayType:  list[i].SysPayType,
			Platform:    list[i].Platform,
			Interface:   list[i].Interface,
			CreateTime:  list[i].CreateTime,
			SuccessTime: list[i].SuccessTime,
			Status:      list[i].Status,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func rechargeBillTransferReqCheck(req *RechargeBillTransferReq) (map[string]interface{}, error) {
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
	if req.Type != "" {
		m["type"] = req.Type
	}
	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if req.OrderSn != "" {
		m["order_sn"] = req.OrderSn
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)
	return m, nil
}

// （财务管理）充值记录（转账汇款）接口
func RechargeBillTransfer(c *gin.Context) {
	var req RechargeBillTransferReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := rechargeBillTransferReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, rsn, rss, err := model.GetRechargeBillTransferList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	resp := RechargeBillTransferResp{
		List:               make([]RechargeBillTransferInfo, 0),
		Total:              count,
		TransferNum:        count,
		TransferSuccessNum: rsn,
		TransferSuccessSum: rss,
	}
	for i := range list {
		temp := RechargeBillTransferInfo{
			Id:             list[i].Id,
			OrderSn:        list[i].BillNo,
			UserName:       list[i].UserName,
			TrueName:       list[i].TrueName,
			TransAccount:   list[i].TransAccount,
			PayMoney:       list[i].Money,
			SysPayType:     list[i].SysPayType,
			TransToAccount: list[i].TransToAccount,
			TransTime:      list[i].TransTime,
			CheckTime:      list[i].CheckTime,
			Status:         list[i].Status,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func rechargeBillBackAddReqCheck(req *RechargeBillBackAddReq) (map[string]interface{}, error) {
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
	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if req.OrderSn != "" {
		m["order_sn"] = req.OrderSn
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)
	return m, nil
}

// （财务管理）充值记录（后台加款）接口
func RechargeBillBackAdd(c *gin.Context) {
	var req RechargeBillBackAddReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	// 参数合法性检查
	m, err := rechargeBillBackAddReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, basn, bass, err := model.GetRechargeBillBackAddList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	resp := RechargeBillBackAddResp{
		List:              make([]RechargeBillBackAddInfo, 0),
		Total:             count,
		BackAddSuccessNum: basn,
		BackAddSuccessSum: bass,
	}
	for i := range list {
		temp := RechargeBillBackAddInfo{
			Id:         list[i].Id,
			OrderSn:    list[i].OrderSn,
			UserName:   list[i].UserName,
			TrueName:   list[i].TrueName,
			Money:      list[i].SettAmt,
			Memo:       list[i].Memo,
			Operator:   list[i].Operator,
			CreateTime: list[i].CreateTime,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func withdrawBillReqCheck(req *WithdrawBillReq) (map[string]interface{}, error) {
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
	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if req.OrderSn != "" {
		m["order_sn"] = req.OrderSn
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)
	return m, nil
}

// （财务管理）提现记录会员提现接口
func WithdrawBill(c *gin.Context) {
	var req WithdrawBillReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	m, err := withdrawBillReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, wsn, wss, err := model.GetWithdrawBillList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := WithdrawBillResp{
		List:               make([]WithdrawBillInfo, 0),
		Total:              count,
		WithdrawNum:        count,
		WithdrawSuccessNum: wsn,
		WithdrawSuccessSum: wss,
	}
	for i := range list {
		temp := WithdrawBillInfo{
			Id:          list[i].Id,
			OrderSn:     list[i].OrderSn,
			UserName:    list[i].UserName,
			ClassName:   list[i].ClassName,
			TrueName:    list[i].TrueName,
			BankName:    list[i].BankName,
			CardNo:      list[i].CardNo,
			Money:       list[i].Money,
			CreateTime:  list[i].CreateTime,
			ApproveTime: list[i].ApproveTime,
			Memo:        list[i].Memo,
			Status:      list[i].Status,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func withdrawBillBackReqCheck(req *WithdrawBillBackReq) (map[string]interface{}, error) {
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

// （财务管理）提现记录后台扣款接口
func WithdrawBillBack(c *gin.Context) {
	var req WithdrawBillBackReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	m, err := withdrawBillBackReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, wsn, wss, err := model.GetWithdrawBillBackList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := WithdrawBillBackResp{
		List:               make([]WithdrawBillBackInfo, 0),
		Total:              count,
		WithdrawNum:        count,
		WithdrawSuccessNum: wsn,
		WithdrawSuccessSum: wss,
	}
	for i := range list {
		temp := WithdrawBillBackInfo{
			Id:         list[i].Id,
			OrderSn:    list[i].OrderSn,
			UserName:   list[i].UserName,
			TrueName:   list[i].TrueName,
			Amount:     list[i].SettAmt,
			Memo:       list[i].Memo,
			Operator:   list[i].Operator,
			CreateTime: list[i].CreateTime,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func rechargeTransBillReqCheck(req *RechargeBillReq) (map[string]interface{}, error) {
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

// （财务管理）帐变记录(转账)接口
func RechargeTransBill(c *gin.Context) {
	var req RechargeBillReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	m, err := rechargeTransBillReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, rsn, rss, err := model.GetRechargeTransBillList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := RechargeTransBillResp{
		List:               make([]RechargeTransBillInfo, 0),
		Total:              count,
		RechargeNum:        count,
		RechargeSuccessNum: rsn,
		RechargeSuccessSum: rss,
	}
	for i := range list {
		temp := RechargeTransBillInfo{
			CheckTime:      list[i].CheckTime,
			CreateTime:     list[i].TransTime,
			BillNo:         list[i].BillNo,
			Amount:         list[i].Money,
			ContactName:    list[i].ContactName,
			MerchantName:   list[i].MerchantName,
			Status:         list[i].Status,
			TransAccount:   list[i].TransAccount,
			TransToAccount: list[i].TransToAccount,
			TrueName:       list[i].TrueName,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func gameTransBillReqCheck(req *GameTransBillReq) (map[string]interface{}, error) {
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

// （财务管理）游戏转账接口
func GameTransBill(c *gin.Context) {
	var req GameTransBillReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	m, err := gameTransBillReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, err := model.GetGameTransBillList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := GameTransBillResp{
		List:  make([]GameTransBillInfo, 0),
		Total: count,
	}
	for i := range list {
		temp := GameTransBillInfo{
			Id:         list[i].Id,
			BillNo:     list[i].BillNo,
			UserName:   list[i].UserName,
			TrueName:   list[i].TrueName,
			Type:       list[i].Type,
			Money:      list[i].Money,
			CreateTime: list[i].CreateTime,
			Ok:         list[i].Ok,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

// （财务管理）根据用户名查询会员(加扣款前访问这个)接口
func QueryUserBalance(c *gin.Context) {
	var req QueryUserBalanceReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	info, err := model.GetUserBalanceByUserName(model.Db, merchantId, req.UserName)
	if err != nil {
		if err == model.ErrRecordNotFound {
			logrus.Infof("not found merchantId[%v], userName[%v]", merchantId, req.UserName)
			RespNotFoundErr(c)
			return
		}

		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := QueryUserBalanceResp{
		Id:       info.Id,
		UserName: info.UserName,
		Balance:  info.Balance,
	}

	RespJson(c, status.OK, resp)
}

func centerAccountBalanceSwitchReqCheck(req CenterAccountBalanceSwitchReq) error {
	if req.UserId <= 0 {
		return errors.New("user_id error")
	}
	if !utils.IsIntContains(MemBillTypeList, req.Code) {
		return errors.New("code error")
	}
	if req.Amount < 0 {
		return errors.New("amount error")
	}

	return nil
}

// 中心账户加扣款
func CenterAccountBalanceSwitch(c *gin.Context) {
	var req CenterAccountBalanceSwitchReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	if err := centerAccountBalanceSwitchReqCheck(req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get mUserId from token
	//todo: get merchantId from token
	merchantId := 1

	tx := model.TxBegin()
	// step0: check user exist
	exist, err := model.IsExistWithMerchantIdAndUserId(tx, merchantId, req.UserId)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	if !exist {
		tx.Rollback()
		RespNotFoundErr(c)
		return
	}

	// step1: update table `user`
	user := model.User{
		MerchantId: merchantId,
		Id:         req.UserId,
	}
	//丢失补款和多出扣款，减
	if req.Code == MemBillSupplement || req.Code == MemBillDeduction {
		req.Amount = -req.Amount
	}
	if err := user.MerchantUpdateUserBalance(tx, req.Amount); err != nil {
		tx.Rollback()
		logrus.Error(err)
		if err == model.ErrNoEnoughMoney {
			RespJson(c, status.ErrNoEnoughMoney, nil)
		} else {
			RespServerErr(c)
		}
		return
	}

	// step2: update table `user_backadd_bill`
	u, err := model.GetUserByMerchantIdAndUserId(tx, req.UserId, merchantId)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		if err == model.ErrRecordNotFound {
			RespNotFoundErr(c)
		} else {
			RespServerErr(c)
		}
		return
	}
	now := time.Now().Unix()
	orderSn := utils.CreateOrderNo(req.UserId)
	ubb := &model.UserBackaddBill{
		UserId:     req.UserId,
		SettAmt:    req.Amount,
		Memo:       req.Memo,
		Tips:       req.Memo,
		Balance:    u.Balance,
		OldBalance: u.Balance - req.Amount,
		OrderSn:    orderSn,
		CreateTime: now,
		MerchantId: merchantId,
		Operator:   merchantId,
	}
	if req.Code == MemBillSupplement || req.Code == MemBillDeduction {
		ubb.Type = UserBackAddBillSub
	} else {
		ubb.Type = UserBackAddBillAdd
	}

	if err := ubb.NewRecord(tx); err != nil {
		tx.Rollback()
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	//step3: update table `user_bill`
	ub := &model.UserBill{
		UserId:     req.UserId,
		MerchantId: merchantId,
		Type:       ubb.Type,
		SettAmt:    req.Amount,
		Memo:       req.Memo,
		Balance:    u.Balance,
		OldBalance: u.Balance - req.Amount,
		OrderSn:    orderSn,
		Code:       req.Code,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := ub.NewRecord(tx); err != nil {
		tx.Rollback()
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	tx.Commit()

	RespJson(c, status.OK, u.Balance)
}

// 查询游戏余额
func QueryGameBalance(c *gin.Context) {
	var req QueryGameBalanceReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	//todo: get userId from token
	userId := 1
	info, err := model.GetGameBalanceByGameId(model.Db, userId, req.GameId)
	if err != nil {
		if err == model.ErrRecordNotFound {
			RespNotFoundErr(c)
			return
		}

		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := QueryGameBalanceResp{
		GameId:  info.Id,
		Balance: info.GameBalance,
	}

	RespJson(c, status.OK, resp)
}

// 在线支付接口列表接口
func OnlinePaymentsList(c *gin.Context) {
	var req OnlinePaymentsListReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	if req.PayTag != PayTagOnlinePay && req.PayTag != PayTagTransfer {
		logrus.Error("pay_tag error")
		RespParamErr(c)
		return
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)

	m := make(map[string]interface{})
	m["pay_tag"] = req.PayTag
	if req.PayTypeId != 0 {
		m["pay_type_id"] = req.PayTypeId
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, err := model.GetOnlinePaymentsList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := OnlinePaymentsListResp{
		List:  make([]OnlinePaymentsInfo, 0),
		Total: count,
	}
	for i := range list {
		temp := OnlinePaymentsInfo{
			Id:         list[i].Id,
			Code:       list[i].Code,
			DayStopMax: list[i].DayStopMax,
			PayType:    list[i].PayType,
			Platform:   list[i].Platform,
			SimpleMax:  list[i].SimpleMax,
			SimpleMin:  list[i].SimpleMin,
			Sort:       list[i].Sort,
			Status:     list[i].Status,
			PayTypeId:  list[i].SysPayTypeId,
			MerchNo:    list[i].MerchNo,
			CreateTime: list[i].CreateTime,
			UpdateTime: list[i].UpdateTime,
			Remark:     list[i].Remark,
			Account:    list[i].Account,
			Qrcode:     list[i].Account,
			Url:        list[i].Url,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

func thirdAccountBalanceSwitchReqCheck(req ThirdAccountBalanceSwitchReq) error {
	if req.UserId <= 0 {
		return errors.New("user_id error")
	}
	if req.GameId <= 0 {
		return errors.New("game_id error")
	}
	if req.Type < ThirdAccountAddBalance || req.Type > ThirdAccountAG2Third {
		return errors.New("type error")
	}
	if req.Amount < 0 {
		return errors.New("amount error")
	}
	return nil
}

// 第三方账户加扣款
func ThirdAccountBalanceSwitch(c *gin.Context) {
	var req ThirdAccountBalanceSwitchReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	if err := thirdAccountBalanceSwitchReqCheck(req); err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	isMatch, err := model.IsMidGidUidMatch(model.Db, merchantId, req.GameId, req.UserId)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	if !isMatch {
		logrus.Error("merchant_id, game_id, user_id not match")
		RespParamErr(c)
		return
	}

	switch req.Type {
	case ThirdAccountAddBalance:
	case ThirdAccountSubBalance:
	case ThirdAccountThird2AG:
	case ThirdAccountAG2Third:
	}

	RespSuccess(c)
}

func financeListReqCheck(req *FinanceListReq) (map[string]interface{}, error) {
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
	if req.UserName != "" {
		m["user_name"] = req.UserName
	}
	if 0 < req.SortBy && req.SortBy <= 6 {
		m["order_by"] = FinanceListSortBy[req.SortBy]
	}
	m["order"] = req.Order
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)
	return m, nil
}

// 财务报表接口
func FinanceList(c *gin.Context) {
	var req FinanceListReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	m, err := financeListReqCheck(&req)
	if err != nil {
		logrus.Error(err)
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	list, count, err := model.GetFinanceList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := FinanceListResp{
		List:              make([]FinanceInfo, 0),
		Total:             count.MemberSum,
		MemberSum:         count.MemberSum,
		EffectBetSum:      count.EffectiveBetSum,
		WinLostAmountSum:  count.WinLostAmountSum,
		RechargeAmountSum: count.RechargeSum,
		WithdrawAmountSum: count.WithdrawSum,
		BonusAmountSum:    count.BonusSum,
		RebateSum:         count.RebateSum,
		FeeSum:            count.FeeSum,
	}

	for i := range list {
		temp := FinanceInfo{
			UserId:         list[i].UserId,
			UserName:       list[i].UserName,
			TrueName:       list[i].TrueName,
			EffectBet:      list[i].EffectBet,
			WinLostAmount:  list[i].WinLostAmount,
			RechargeAmount: list[i].RechargeAmount,
			WithdrawAmount: list[i].WithdrawAmount,
			BonusAmount:    list[i].BonusAmount,
			Rebate:         list[i].Rebate,
			Fee:            list[i].Fee,
		}
		resp.List = append(resp.List, temp)
	}

	RespJson(c, status.OK, resp)
}

// 导出充值记录excel
func ExportRechargeBillList(c *gin.Context) {
	var req RechargeBillReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	rechargeBill := NewRechargeBillExport(c, merchantId, req)
	director := builder.NewFileDirector()
	director.ExportFile(rechargeBill)
}

// 增加在线支付接口
func OnlinePaymentsAdd(c *gin.Context) {
	var req OnlinePaymentsAddReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	if req.PayTag != PayTagOnlinePay && req.PayTag != PayTagTransfer {
		logrus.Error("pay_tag error")
		RespParamErr(c)
		return
	}
	//检查pay_type_id
	match, err := model.CheckMPCIdAndPayTagMatch(model.Db, req.PayTypeId, req.PayTag)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	if !match {
		logrus.Error("pay_type_id and pay_tag not match")
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1

	now := time.Now().Unix()
	mpc := &model.MerchantPayConfig{
		SysPayTypeId: req.PayTypeId,
		MerchantId:   merchantId,
		Sort:         req.Sort,
		Code:         req.Code,
		SimpleMin:    req.SimpleMin,
		SimpleMax:    req.SimpleMax,
		DayStopMax:   req.DayStopMax,
		Status:       req.Status,
		CreateTime:   now,
		UpdateTime:   now,
		Remark:       req.Remark,
	}
	if req.PayTag == PayTagOnlinePay {
		mpc.MerchNo = req.MerchNo
		mpc.Md5Key = req.Md5Key
		mpc.PublicKey = req.PublicKey
		mpc.SecretKey = req.SecretKey
		mpc.Url = req.Url
	} else {
		mpc.Account = req.Account
	}
	if err := mpc.NewRecord(model.Db); err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

// 编辑在线支付接口
func OnlinePaymentsEdit(c *gin.Context) {
	var req OnlinePaymentsEditReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	if req.PayTag != PayTagOnlinePay && req.PayTag != PayTagTransfer {
		logrus.Error("pay_tag error")
		RespParamErr(c)
		return
	}
	//检查pay_type_id
	match, err := model.CheckMPCIdAndPayTagMatch(model.Db, req.PayTypeId, req.PayTag)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	if !match {
		logrus.Error("pay_type_id and pay_tag not match")
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	now := time.Now().Unix()
	mpc := &model.MerchantPayConfig{
		Id:           req.Id,
		SysPayTypeId: req.PayTypeId,
		MerchantId:   merchantId,
		Sort:         req.Sort,
		Code:         req.Code,
		MerchNo:      req.MerchNo,
		Md5Key:       req.Md5Key,
		PublicKey:    req.PublicKey,
		SecretKey:    req.SecretKey,
		Url:          req.Url,
		SimpleMin:    req.SimpleMin,
		SimpleMax:    req.SimpleMax,
		DayStopMax:   req.DayStopMax,
		Status:       req.Status,
		UpdateTime:   now,
		Account:      req.Account,
		Remark:       req.Remark,
	}
	if err := mpc.Update(model.Db, req.PayTag); err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

// 删除在线支付接口
func OnlinePaymentsDel(c *gin.Context) {
	var req struct {
		PayTag int `json:"pay_tag" binding:"required"`
		Id     int `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}

	if req.PayTag != PayTagOnlinePay && req.PayTag != PayTagTransfer {
		logrus.Error("pay_tag error")
		RespParamErr(c)
		return
	}
	//检查pay_type_id
	match, err := model.CheckMPCIdAndPayTagMatch(model.Db, req.Id, req.PayTag)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	if !match {
		logrus.Error("pay_type_id and pay_tag not match")
		RespParamErr(c)
		return
	}

	//todo: get merchantId from token
	merchantId := 1
	mpc := &model.MerchantPayConfig{
		Id:         req.Id,
		MerchantId: merchantId,
	}
	if err := mpc.Delete(model.Db); err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}
	RespSuccess(c)
}

// 信用额度-额度帐变接口
func CreditLimitList(c *gin.Context) {
	var req CreditLimitListReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}
	m := make(map[string]interface{})
	if req.StartTime > 0 {
		m["start_time"] = req.StartTime
		if req.EndTime > 0 {
			if req.EndTime < req.StartTime {
				logrus.Error("start time less end time")
				RespParamErr(c)
				return
			}
			m["end_time"] = req.EndTime
		}
	} else if req.EndTime > 0 {
		m["end_time"] = req.EndTime
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)

	//todo: get merchantId from token
	merchantId := 1
	list, count, err := model.GetCreditLimitList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := CreditLimitListResp{
		List:  make([]CreditLimitInfo, 0),
		Total: count,
	}
	for i := range list {
		temp := CreditLimitInfo{
			BillNo:  list[i].BillNo,
			Type:    list[i].Type,
			Amount:  list[i].Amount,
			OldBail: list[i].OldBail,
			NewBail: list[i].NewBail,
			Remark:  list[i].Remark,
		}
		resp.List = append(resp.List, temp)
	}
	RespJson(c, status.OK, resp)
}

// 信用额度-充值记录接口
func CreditLimitTransferList(c *gin.Context) {
	var req CreditLimitTransferListReq
	if err := c.BindJSON(&req); err != nil {
		RespParamErr(c)
		return
	}
	m := make(map[string]interface{})
	if req.StartTime > 0 {
		m["start_time"] = req.StartTime
		if req.EndTime > 0 {
			if req.EndTime < req.StartTime {
				logrus.Error("start time less end time")
				RespParamErr(c)
				return
			}
			m["end_time"] = req.EndTime
		}
	} else if req.EndTime > 0 {
		m["end_time"] = req.EndTime
	}
	req.Page, req.PageCount = InitPage(req.Page, req.PageCount)

	//todo: get merchantId from token
	merchantId := 1
	list, count, err := model.GetCreditLimitTransferList(model.Db, merchantId, req.Page, req.PageCount, m)
	if err != nil {
		logrus.Error(err)
		RespServerErr(c)
		return
	}

	resp := CreditLimitTransferListResp{
		List:  make([]CreditLimitTransferInfo, 0),
		Total: count,
	}
	for i := range list {
		temp := CreditLimitTransferInfo{
			Id:             list[i].Id,
			Amount:         list[i].Amount,
			PayChannel:     list[i].PayChannel,
			PayType:        list[i].PayType,
			OldBalance:     list[i].OldBalance,
			NewBalance:     list[i].NewBalance,
			TransAccount:   list[i].TransAccount,
			TransToAccount: list[i].TransToAccount,
			CreateTime:     list[i].CreateTime,
			Remark:         list[i].Remark,
			Status:         list[i].Status,
			BillNo:         list[i].BillNo,
		}
		resp.List = append(resp.List, temp)
	}
	RespJson(c, status.OK, resp)
}

func TransferAdd(c *gin.Context) {
	RespSuccess(c)
}
