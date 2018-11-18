package action

const (
	MemBillRecharge    = 100 //充值
	MemBillWithdraw    = 200 //提现
	MemBillTransferIn  = 301 //转入
	MemBillTransferOut = 302 //转出
	MemBillBonus       = 400 //红利
	MemBillRebate      = 500 //返利
	MemBillSupplement  = 600 //丢失补款
	MemBillDeduction   = 700 //多出扣款

	UserBackAddBillAdd = 1  //加
	UserBackAddBillSub = -1 //减

	ThirdAccountAddBalance = 1 //第三方加款(不入帐变)
	ThirdAccountSubBalance = 2 //第三方扣款(不入帐变)
	ThirdAccountThird2AG   = 3 //第三方账户->中心账户
	ThirdAccountAG2Third   = 4 //中心账户->第三方账户

	PayTagOnlinePay = 1 //在线支付
	PayTagTransfer  = 2 //转账汇款
)

var (
	MemBillTypeList = []int{
		MemBillRecharge,
		MemBillWithdraw,
		MemBillTransferIn,
		MemBillTransferOut,
		MemBillBonus,
		MemBillRebate,
		MemBillSupplement,
		MemBillDeduction,
	}

	FinanceListSortBy = []string{
		"",                //默认值
		"win_lost_amount", //输赢金额
		"recharge_amount", //充值金额
		"withdraw_amount", //提现金额
		"bonus_amount",    //红利金额
		"rebate",          //反水金额
		"fee",             //手续费
	}
)
