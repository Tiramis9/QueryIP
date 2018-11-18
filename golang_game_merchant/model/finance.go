package model

import (
	"github.com/jinzhu/gorm"
)

type (
	// 用户财务明细表
	UserBill struct {
		Id         int
		UserId     int
		MerchantId int
		Type       int
		SettAmt    float64
		Memo       string
		Tips       string
		Fee        float64
		About      string
		Balance    float64
		OldBalance float64
		Status     int
		OrderSn    string
		Code       int
		CodeSn     string
		CreateTime int64
		UpdateTime int64
		Fs         float64
		Fy         float64
	}

	// 用户充值明细表
	UserPay struct {
		Id            int
		UserId        int
		OrderSn       string
		Status        int
		PayMoney      float64
		Memo          string
		Tips          string
		MerchantPayId int
		CreateTime    int64
		CallbackTime  int64
		UpdateTime    int64
		Fee           float64
		PayType       int
		PlatformCode  string
		SysPayType    string
	}

	// 用户提现表
	UserWithdraw struct {
		Id           int
		UserId       int
		Money        float64
		Status       int
		Memo         string
		Addition     string
		CreateTime   int64
		ApproveTime  int64
		OrderSn      string
		Type         int
		CallbackTime int64
		Fee          float64
		OperatorId   int
		OperatorIp   string
		CardNo       string
	}

	// 用户后台操作帐变记录表
	UserBackaddBill struct {
		Id         int
		UserId     int
		Type       int
		SettAmt    float64
		Memo       string
		Tips       string
		Fee        float64
		Balance    float64
		OldBalance float64
		OrderSn    string
		CreateTime int64
		MerchantId int
		Operator   int
	}

	//用户转账表
	UserTransBill struct {
		Id             int
		UserId         int
		Money          float64
		TransAccount   string
		TransToAccount string
		TransTime      int64
		CheckTime      int64
		Remark         string
		Status         int
		BillNo         string
	}

	//用户游戏账户明细表
	UserAccountBill struct {
		Id          int
		AccountName string
		UserId      int
		Money       float64
		Ok          int
		OldBalance  float64
		NewBalance  float64
		BillNo      string
		Fee         float64
		CreateTime  int64
		UpdateTime  int64
		Type        int
		GameCode    string
	}

	UserAccount struct {
		Id           int
		Type         int
		UserId       int
		GameUserName string
		GameBalance  float64
		Status       int
		CreateTime   int64
		UpdateTime   int64
		GameName     string
	}

	MerchantBailBill struct {
		Id         int
		MerchantId int
		Amount     float64
		Type       int
		Code       int
		OldBail    float64
		NewBail    float64
		CreateTime int64
		Remark     string
		Status     int
		BillNo     string
	}

	MerchantPayBill struct {
		Id             int
		MerchantId     int
		Amount         float64
		PayChannel     string
		PayType        string
		OldBalance     float64
		NewBalance     float64
		TransAccount   string
		TransToAccount string
		TransTime      int64
		CheckTime      int64
		CreateTime     int64
		Remark         string
		OperatorId     int
		Status         int
		BillNo         string
	}
)
type (
	MemberBillInfo struct {
		UserBill
		TrueName string
		UserName string
	}

	RechargeBillInfo struct {
		UserPay
		UserName    string
		TrueName    string
		SysPayType  string
		Platform    string
		Interface   string
		SuccessTime int
	}
	RechargeBillSuccessInfo struct {
		RechargeSuccessNum int
		RechargeSuccessSum float64
	}

	RechargeBillTransferInfo struct {
		UserTransBill
		UserName   string
		TrueName   string
		SysPayType string
	}

	RechargeBillBackAddInfo struct {
		UserBackaddBill
		TrueName string
		UserName string
		Operator string
	}

	WithdrawBillInfo struct {
		UserWithdraw
		UserName  string
		ClassName string
		TrueName  string
		BankName  string
	}
	WithdrawBillSuccessInfo struct {
		WithdrawSuccessNum int
		WithdrawSuccessSum float64
	}

	WithdrawBillBackInfo struct {
		UserBackaddBill
		UserName string
		TrueName string
		Operator string
	}

	UserTransBillInfo struct {
		UserTransBill
		ContactName  string
		MerchantName string
		TrueName     string
	}
	UserTransBillSuccessInfo struct {
		TransSuccessNum int
		TransSuccessSum float64
	}

	UserAccountBillInfo struct {
		UserAccountBill
		TrueName string
		UserName string
	}

	OnlinePaymentsInfo struct {
		MerchantPayConfig
		Platform string
		PayType  string
	}

	FinanceInfo struct {
		EffectBet     float64 //暂不处理
		WinLostAmount float64 //暂不处理

		RechargeAmount    float64 //充值
		WithdrawAmount    float64 //提现
		TransferInAmount  float64 //转入
		TransferOutAmount float64 //转出
		BonusAmount       float64 //红利
		Fee               float64 //手续费
		SupplementAmount  float64 //丢失补款
		DeductionAmount   float64 //多出扣款
		Rebate            float64 //反水
		UserId            int     //用户id
		TrueName          string  //真是姓名
		UserName          string  //用户名称
	}
	FinanceCount struct {
		EffectiveBetSum  float64 //暂不处理
		WinLostAmountSum float64 //暂不处理
		MemberSum        int     //会员数
		RechargeSum      float64 //充值
		WithdrawSum      float64 //提现
		TransferInSum    float64 //转入
		TransferOutSum   float64 //转出
		BonusSum         float64 //红利
		FeeSum           float64 //手续费
		SupplementSum    float64 //丢失补款
		DeductionSum     float64 //多出扣款
		RebateSum        float64 //反水
	}
)

func GetMemberBillList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]MemberBillInfo, int, error) {
	whereStr := "ub.merchant_id=?"
	condition := []interface{}{merchantId}
	if v, ok := m["user_type"]; ok {
		whereStr += " AND u.type=?"
		condition = append(condition, v)
	}
	if v, ok := m["type"]; ok {
		whereStr += " AND ub.type=?"
		condition = append(condition, v)
	}
	if v, ok := m["code"]; ok {
		whereStr += " AND ub.code=?"
		condition = append(condition, v)
	}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=ub.create_time AND ub.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=ub.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ub.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var list []MemberBillInfo
	if err := db.Table(`user_bill AS ub`).Joins(`
		LEFT JOIN user AS u ON u.id=ub.user_id
	`).Select(`
		ub.*,
		u.true_name,
		u.user_name
	`).Where(whereStr, condition...).Order(`
		ub.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`user_bill AS ub`).Joins(`
		LEFT JOIN user AS u ON u.id=ub.user_id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

//todo: fix
func GetRechargeBillExportList(db *gorm.DB, merchantId int, m map[string]interface{}) ([]RechargeBillInfo, error) {
	whereStr := "ub.merchant_id=?"
	condition := []interface{}{merchantId}
	if v, ok := m["pay_type"]; ok {
		whereStr += " AND up.pay_type=?"
		condition = append(condition, v)
	} else {
		//如果没有该参数，只返回“在线支付”类型数据
		whereStr += " AND up.pay_type=1"
	}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=up.create_time AND up.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=up.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND up.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var list []RechargeBillInfo
	if err := db.Table(`user_pay AS up`).Joins(`
		LEFT JOIN user_bill AS ub ON ub.order_sn=up.order_sn
	`).Joins(`
		LEFT JOIN user AS u ON u.id=up.user_id
	`).Joins(`
		LEFT JOIN user_backadd_bill AS ubb ON ubb.order_sn=ub.order_sn
	`).Select(`
		up.*,
		ub.balance,
		ub.old_balance,
		ub.code,
		ub.type,
		u.true_name,
		u.user_name,
		ubb.operator
	`).Where(whereStr, condition...).Order(`
		up.create_time DESC
	`).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return list, nil
}

func GetRechargeBillOnlinePayList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]RechargeBillInfo, int, int, float64, error) {
	whereStr := "ub.merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=up.create_time AND up.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=up.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND up.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["platform"]; ok {
		whereStr += " AND spt.platform=?"
		condition = append(condition, v)
	}
	if v, ok := m["pay_type"]; ok {
		whereStr += " AND spt.pay_type=?"
		condition = append(condition, v)
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}
	if v, ok := m["order_sn"]; ok {
		whereStr += " AND up.order_sn LIKE ?"
		orderSn, _ := v.(string)
		condition = append(condition, "%"+orderSn+"%")
	}

	var list []RechargeBillInfo
	if err := db.Table(`user_pay AS up`).Joins(`
		LEFT JOIN user_bill AS ub ON ub.order_sn=up.order_sn
	`).Joins(`
		LEFT JOIN user AS u ON u.id=up.user_id
	`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.platform_code=up.platform_code AND spt.pay_type=up.sys_pay_type
	`).Joins(`
		LEFT JOIN merchant_pay_config AS mpc ON mpc.merchant_id=up.merchant_id AND mpc.sys_pay_type_id=spt.id
	`).Select(`
		up.*,
		u.user_name,
		u.true_name,
		spt.pay_type AS sys_pay_type,
		spt.platform,
		mpc.code AS interface,
		IF(up.status=2, up.callback_time, 0) AS success_time
	`).Where(whereStr, condition...).Order(`
		up.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, 0, 0, nil
		}
		return nil, 0, 0, 0, err
	}

	var count int
	if err := db.Table(`user_pay AS up`).Joins(`
		LEFT JOIN user_bill AS ub ON ub.order_sn=up.order_sn
	`).Joins(`
		LEFT JOIN user AS u ON u.id=up.user_id
	`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.platform_code=up.platform_code AND spt.pay_type=up.sys_pay_type
	`).Joins(`
		LEFT JOIN merchant_pay_config AS mpc ON mpc.merchant_id=up.merchant_id AND mpc.sys_pay_type_id=spt.id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	whereStr += " AND up.status='2'"
	var successInfo RechargeBillSuccessInfo
	if err := db.Table(`user_pay AS up`).Joins(`
		LEFT JOIN user_bill AS ub ON ub.order_sn=up.order_sn
	`).Joins(`
		LEFT JOIN user AS u ON u.id=up.user_id
	`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.platform_code=up.platform_code AND spt.pay_type=up.sys_pay_type
	`).Joins(`
		LEFT JOIN merchant_pay_config AS mpc ON mpc.merchant_id=up.merchant_id AND mpc.sys_pay_type_id=spt.id
	`).Select(`
		COUNT(*) AS recharge_success_num,
		SUM(up.pay_money) AS recharge_success_sum
	`).Where(whereStr, condition...).Find(&successInfo).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return list, count, successInfo.RechargeSuccessNum, successInfo.RechargeSuccessSum, nil
}

func GetRechargeBillTransferList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]RechargeBillTransferInfo, int, int, float64, error) {
	whereStr := "utb.merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=utb.trans_time AND utb.trans_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=utb.trans_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND utb.trans_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["type"]; ok {
		whereStr += " AND spt.pay_type=?"
		condition = append(condition, v)
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}
	if v, ok := m["order_sn"]; ok {
		whereStr += " AND utb.bill_no LIKE ?"
		orderSn, _ := v.(string)
		condition = append(condition, "%"+orderSn+"%")
	}

	var list []RechargeBillTransferInfo
	if err := db.Table(`user_trans_bill AS utb`).Joins(`
		LEFT JOIN user AS u ON u.merchant_id=utb.merchant_id AND utb.user_id=u.id
	`).Joins(`
		LEFT JOIN merchant_pay_config AS mpc ON utb.merchant_pay_id=mpc.id
	`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.id=mpc.sys_pay_type_id
	`).Select(`
		utb.*,
		u.true_name,
		u.user_name,
		spt.pay_type AS sys_pay_type
	`).Where(whereStr, condition...).Order(`
		utb.trans_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, 0, 0, nil
		}
		return nil, 0, 0, 0, err
	}

	var count int
	if err := db.Table(`user_trans_bill AS utb`).Joins(`
		LEFT JOIN user AS u ON u.merchant_id=utb.merchant_id AND utb.user_id=u.id
	`).Joins(`
		LEFT JOIN merchant_pay_config AS mpc ON utb.merchant_pay_id=mpc.id
	`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.id=mpc.sys_pay_type_id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	whereStr += " AND utb.status='1'"
	var successInfo RechargeBillSuccessInfo
	if err := db.Table(`user_trans_bill AS utb`).Joins(`
		LEFT JOIN user AS u ON u.merchant_id=utb.merchant_id AND utb.user_id=u.id
	`).Joins(`
		LEFT JOIN merchant_pay_config AS mpc ON utb.merchant_pay_id=mpc.id
	`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.id=mpc.sys_pay_type_id
	`).Select(`
		COUNT(*) AS recharge_success_num,
		SUM(utb.money) AS recharge_success_sum
	`).Where(whereStr, condition...).Find(&successInfo).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return list, count, successInfo.RechargeSuccessNum, successInfo.RechargeSuccessSum, nil
}

func GetRechargeBillBackAddList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]RechargeBillBackAddInfo, int, int, float64, error) {
	whereStr := "ubb.merchant_id=? AND ubb.type='1'"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=ubb.create_time AND ubb.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=ubb.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ubb.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}
	if v, ok := m["order_sn"]; ok {
		whereStr += " AND ubb.order_sn LIKE ?"
		orderSn, _ := v.(string)
		condition = append(condition, "%"+orderSn+"%")
	}

	var list []RechargeBillBackAddInfo
	if err := db.Debug().Table(`user_backadd_bill AS ubb`).Joins(`
		LEFT JOIN user AS u ON u.id=ubb.user_id
	`).Joins(`
		LEFT JOIN merchant AS m ON m.id=ubb.operator
	`).Select(`
		ubb.*,
		u.true_name,
		u.user_name,
		m.merchant_name AS operator
	`).Where(whereStr, condition...).Order(`
		ubb.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, 0, 0, nil
		}
		return nil, 0, 0, 0, err
	}

	count := struct {
		Count int
		Sum   float64
	}{}
	if err := db.Table(`user_backadd_bill AS ubb`).Joins(`
		LEFT JOIN user AS u ON ubb.user_id=u.id
	`).Joins(`
		LEFT JOIN merchant AS m ON m.id=ubb.operator
	`).Select(`
		COUNT(*) AS count,
		SUM(ubb.sett_amt) AS sum
	`).Where(whereStr, condition...).Find(&count).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return list, count.Count, count.Count, count.Sum, nil
}

func GetWithdrawBillList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]WithdrawBillInfo, int, int, float64, error) {
	whereStr := "uw.merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=uw.create_time AND uw.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=uw.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND uw.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}
	if v, ok := m["order_sn"]; ok {
		whereStr += " AND uw.order_sn LIKE ?"
		orderSn, _ := v.(string)
		condition = append(condition, "%"+orderSn+"%")
	}

	var list []WithdrawBillInfo
	if err := db.Table(`user_withdraw AS uw`).Joins(`
		LEFT JOIN user AS u ON u.id=uw.user_id
	`).Joins(`
		LEFT JOIN merchant_user_class AS muc ON muc.id=u.class_id
	`).Joins(`
		LEFT JOIN user_bank AS ub ON ub.user_id=uw.user_id AND ub.card_no=uw.card_no
	`).Select(`
		uw.*,
		u.user_name,
		muc.class_name,
		u.true_name,
		ub.bank_name
	`).Where(whereStr, condition...).Order(`
		uw.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, 0, 0, nil
		}
		return nil, 0, 0, 0, err
	}

	var count int
	if err := db.Table(`user_withdraw AS uw`).Joins(`
		LEFT JOIN user AS u ON u.id=uw.user_id
	`).Joins(`
		LEFT JOIN merchant_user_class AS muc ON muc.id=u.class_id
	`).Joins(`
		LEFT JOIN user_bank AS ub ON ub.user_id=uw.user_id AND ub.card_no=uw.card_no
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	whereStr += " AND uw.status=4"
	var successInfo WithdrawBillSuccessInfo
	if err := db.Table(`user_withdraw AS uw`).Joins(`
		LEFT JOIN user AS u ON u.id=uw.user_id
	`).Joins(`
		LEFT JOIN merchant_user_class AS muc ON muc.id=u.class_id
	`).Joins(`
		LEFT JOIN user_bank AS ub ON ub.user_id=uw.user_id AND ub.card_no=uw.card_no
	`).Select(`
		COUNT(*) AS withdraw_success_num,
		SUM(uw.money) AS withdraw_success_sum
	`).Where(whereStr, condition...).Find(&successInfo).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return list, count, successInfo.WithdrawSuccessNum, successInfo.WithdrawSuccessSum, nil
}

func GetWithdrawBillBackList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]WithdrawBillBackInfo, int, int, float64, error) {
	whereStr := "ubb.merchant_id=? AND ubb.type='-1'"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=ubb.create_time AND ubb.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=ubb.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ubb.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var list []WithdrawBillBackInfo
	if err := db.Table(`user_backadd_bill AS ubb`).Joins(`
		LEFT JOIN user AS u ON ubb.user_id=u.id
	`).Joins(`
		LEFT JOIN merchant AS m ON ubb.merchant_id=m.id
	`).Select(`
		ubb.*,
		u.user_name,
		u.true_name,
		m.merchant_name AS operator
	`).Where(whereStr, condition...).Order(`
		ubb.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, 0, 0, nil
		}
		return nil, 0, 0, 0, err
	}

	count := struct {
		Count int
		Sum   float64
	}{}
	if err := db.Table(`user_backadd_bill AS ubb`).Joins(`
		LEFT JOIN user AS u ON ubb.user_id=u.id
	`).Joins(`
		LEFT JOIN merchant AS m ON ubb.merchant_id=m.id
	`).Select(`
		COUNT(*) AS count,
		SUM(ubb.sett_amt) AS sum
	`).Where(whereStr, condition...).Find(&count).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return list, count.Count, count.Count, count.Sum, nil
}

func GetRechargeTransBillList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]UserTransBillInfo, int, int, float64, error) {
	whereStr := "u.merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=utb.trans_time AND utb.trans_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=utb.trans_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND utb.trans_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var list []UserTransBillInfo
	if err := db.Table(`user_trans_bill AS utb`).Joins(`
		LEFT JOIN user AS u ON u.id=utb.user_id
	`).Joins(`
		LEFT JOIN merchant AS m ON m.id=u.merchant_id
	`).Select(`
		utb.*,
		u.true_name,
		m.contact_name,
		m.merchant_name
	`).Where(whereStr, condition...).Order(`
		utb.trans_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, 0, 0, nil
		}
		return nil, 0, 0, 0, err
	}

	var count int
	if err := db.Table(`user_trans_bill AS utb`).Joins(`
		LEFT JOIN user AS u ON u.id=utb.user_id
	`).Joins(`
		LEFT JOIN merchant AS m ON m.id=u.merchant_id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	whereStr += " AND utb.status=1"
	var successInfo UserTransBillSuccessInfo
	if err := db.Table(`user_trans_bill AS utb`).Joins(`
		LEFT JOIN user AS u ON u.id=utb.user_id
	`).Joins(`
		LEFT JOIN merchant AS m ON m.id=u.merchant_id
	`).Select(`
		COUNT(*) AS trans_success_num,
		SUM(utb.money) AS trans_success_sum
	`).Where(whereStr, condition...).Find(&successInfo).Error; err != nil {
		return nil, 0, 0, 0, err
	}

	return list, count, successInfo.TransSuccessNum, successInfo.TransSuccessSum, nil
}

func GetGameTransBillList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]UserAccountBillInfo, int, error) {
	whereStr := "u.merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=uab.create_time AND uab.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=uab.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND uab.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}

	var list []UserAccountBillInfo
	if err := db.Table(`user_account_bill AS uab`).Joins(`
		LEFT JOIN user AS u ON uab.user_id=u.id
	`).Select(`
		uab.*,
		u.true_name,
		u.user_name
	`).Where(whereStr, condition...).Order(`
		uab.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`user_account_bill AS uab`).Joins(`
		LEFT JOIN user AS u ON uab.user_id=u.id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func GetUserBalanceByUserName(db *gorm.DB, merchantId int, userName string) (*User, error) {
	var user User
	if err := db.Table(`user`).Where(`merchant_id=? AND user_name=?`, merchantId, userName).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

func GetGameBalanceByGameId(db *gorm.DB, userId int, gameId int) (*UserAccount, error) {
	var account UserAccount
	if err := db.Table(`user_account AS ua`).Joins(`
		LEFT JOIN sys_game AS sg ON sg.game_name=ua.game_name
	`).Select(`
		ua.*
	`).Where(`ua.user_id=? AND sg.id=?`, userId, gameId).Find(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &account, nil
}

func GetOnlinePaymentsList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]OnlinePaymentsInfo, int, error) {
	whereStr := "mpc.merchant_id=?"
	condition := []interface{}{merchantId}
	if v, ok := m["pay_tag"]; ok {
		whereStr += "AND spt.pay_tag=?"
		condition = append(condition, v)
	} else {
		return nil, 0, nil
	}
	if v, ok := m["pay_type_id"]; ok {
		whereStr += " AND spt.id=?"
		condition = append(condition, v)
	}

	var list []OnlinePaymentsInfo
	if err := db.Table(`merchant_pay_config AS mpc`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.id=mpc.sys_pay_type_id
	`).Select(`
		mpc.*,
		spt.platform,
		spt.pay_type
	`).Where(whereStr, condition...).Order(`
		mpc.create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`merchant_pay_config AS mpc`).Joins(`
		LEFT JOIN sys_pay_type AS spt ON spt.id=mpc.sys_pay_type_id
	`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func (ubb *UserBackaddBill) NewRecord(db *gorm.DB) error {
	return db.Table(`user_backadd_bill`).Create(ubb).Error
}

func (ub *UserBill) NewRecord(db *gorm.DB) error {
	return db.Table(`user_bill`).Create(ub).Error
}

func GetCreditLimitList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]MerchantBailBill, int, error) {
	whereStr := "merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=create_time AND create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND create_time<=?"
			condition = append(condition, v2)
		}
	}

	var list []MerchantBailBill
	if err := db.Table(`merchant_bail_bill`).Where(whereStr, condition...).Order(`
		create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`merchant_bail_bill`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func GetCreditLimitTransferList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]MerchantPayBill, int, error) {
	whereStr := "merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=create_time AND create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND create_time<=?"
			condition = append(condition, v2)
		}
	}

	var list []MerchantPayBill
	if err := db.Table(`merchant_pay_bill`).Where(whereStr, condition...).Order(`
		create_time DESC
	`).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var count int
	if err := db.Table(`merchant_pay_bill`).Where(whereStr, condition...).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func GetFinanceList(db *gorm.DB, merchantId, page, pageCount int, m map[string]interface{}) ([]FinanceInfo, *FinanceCount, error) {
	whereStr := "ub.merchant_id=?"
	condition := []interface{}{merchantId}
	if v1, ok1 := m["start_time"]; ok1 {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ?<=ub.create_time AND ub.create_time<=?"
			condition = append(condition, v1, v2)
		} else {
			whereStr += " AND ?<=ub.create_time"
			condition = append(condition, v1)
		}
	} else {
		if v2, ok2 := m["end_time"]; ok2 {
			whereStr += " AND ub.create_time<=?"
			condition = append(condition, v2)
		}
	}
	if v, ok := m["user_name"]; ok {
		whereStr += " AND u.user_name LIKE ?"
		userName, _ := v.(string)
		condition = append(condition, "%"+userName+"%")
	}
	var orderStr string
	if v, ok := m["order_by"]; ok {
		switch v {
		case 1:
			orderStr = ""
		case 2:
			orderStr = "recharge_amount"
		case 3:
			orderStr = "withdraw_amount"
		case 4:
			orderStr = "bonus_amount"
		case 5:
			orderStr = "rebate"
		case 6:
			orderStr = "fee"
		}
	}
	if orderStr != "" {
		if v, ok := m["order"]; ok {
			order, _ := v.(int)
			if order == 0 {
				orderStr += " DESC, "
			} else {
				orderStr += " ASC, "
			}
		} else {
			orderStr += " DESC, "
		}
	}
	orderStr += "ub.user_id DESC"

	var list []FinanceInfo
	if err := db.Table(`user_bill AS ub`).Joins(`
		LEFT JOIN user AS u ON u.id=ub.user_id
	`).Select(`
		SUM(IF(ub.code='100', sett_amt, 0)) AS recharge_amount,
		SUM(IF(ub.code='200', sett_amt, 0)) AS withdraw_amount,
		SUM(IF(ub.code='301', sett_amt, 0)) AS transfer_in_amount,
		SUM(IF(ub.code='302', sett_amt, 0)) AS transfer_out_amount,
		SUM(IF(ub.code='400', sett_amt, 0)) AS bonus_amount,
		SUM(IF(ub.code='500', sett_amt, 0)) AS fee,
		SUM(IF(ub.code='600', sett_amt, 0)) AS supplement_amount,
		SUM(IF(ub.code='700', sett_amt, 0)) AS deduction_amount,
		SUM(IF(ub.code='800', sett_amt, 0)) AS rebate,
		u.id AS user_id,
		u.true_name,
		u.user_name
	`).Where(whereStr, condition...).Group(`
		ub.user_id
	`).Order(orderStr).Offset((page - 1) * pageCount).Limit(pageCount).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	var memberSum int
	if err := db.Table(`user_bill AS ub`).Joins(`
		LEFT JOIN user AS u ON u.id=ub.user_id
	`).Where(whereStr, condition...).Group(`
		ub.user_id
	`).Count(&memberSum).Error; err != nil {
		return nil, nil, err
	}

	var total FinanceCount
	if err := db.Table(`user_bill AS ub`).Joins(`
		LEFT JOIN user AS u ON u.id=ub.user_id
	`).Select(`
		SUM(IF(ub.code='100', sett_amt, 0)) AS recharge_sum,
		SUM(IF(ub.code='200', sett_amt, 0)) AS withdraw_sum,
		SUM(IF(ub.code='301', sett_amt, 0)) AS transfer_in_sum,
		SUM(IF(ub.code='302', sett_amt, 0)) AS transfer_out_sum,
		SUM(IF(ub.code='400', sett_amt, 0)) AS bonus_sum,
		SUM(IF(ub.code='500', sett_amt, 0)) AS fee_sum,
		SUM(IF(ub.code='600', sett_amt, 0)) AS supplement_sum,
		SUM(IF(ub.code='700', sett_amt, 0)) AS deduction_sum,
		SUM(IF(ub.code='800', sett_amt, 0)) AS rebate_sum
	`).Where(whereStr, condition...).Find(&total).Error; err != nil {
		return nil, nil, err
	}
	total.MemberSum = memberSum

	return list, &total, nil
}
