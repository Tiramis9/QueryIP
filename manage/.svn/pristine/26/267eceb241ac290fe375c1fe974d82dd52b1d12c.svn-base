package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang_game_merchant/global/status"
)

func MemberClassList(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "merchant_id": 1, "class_name": "白银"},
		{"id": 2, "merchant_id": 1, "class_name": "黄金"},
	}
	RespJson(c, status.OK, data)
}

func MemberGroupList(c *gin.Context) {
	data := [...]map[string]interface{}{
		{"id": 1, "merchant_id": 1, "group_name": "等级1"},
		{"id": 2, "merchant_id": 1, "group_name": "等级2"},
	}
	RespJson(c, status.OK, data)
}

func Message2UserAdd(c *gin.Context) {
	RespSuccess(c)
}

func Message2AgentAdd(c *gin.Context) {
	RespSuccess(c)
}

func MemberBill(c *gin.Context) {
	data := [...]map[string]interface{}{{"id": 1, "order_sn": "123456789", "user_name": "www.buzhidao.com", "true_name": "王鑫", "type": 1, "code": 100, "old_balance": 3000, "sett_amt": 1000, "balance": 4000, "memo": "111", "create_time": 1538016637}, {"id": 2, "order_sn": "123456789", "user_name": "aaa", "true_name": "王鑫", "type": 1, "code": 100, "old_balance": 3000, "sett_amt": 1000, "balance": 4000, "memo": "111", "create_time": 1538016637}}

	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func RechargeBill(c *gin.Context) {
	type_code := c.PostForm("pay_type")

	data_all := map[string]interface{}{}
	switch type_code {
	//1.在线支付
	case "1":
		datas := [...]map[string]interface{}{{"id": 1, "order_sn": "123456789", "user_name": "www.buzhidao.com", "true_name": "王鑫", "pay_money": 1000, "sys_pay_type": "网银", "platform_code": "alipay", "create_time": 1538016637, "callback_time": 1538016637, "status": 1}, {"id": 2, "order_sn": "123456789", "user_name": "www.buzhidao.com", "true_name": "王鑫", "pay_money": 1000, "sys_pay_type": "网银APP", "platform_code": "alipay", "create_time": 1538016637, "callback_time": 1538016637, "status": 1}}
		data_all["data"] = datas
		data_all["recharge_num"] = 10
		data_all["recharge_success_num"] = 10
		data_all["recharge_success_sum"] = 100
		//3.后台加款
	case "3":
		datas := [...]map[string]interface{}{{"id": 1, "bill_no": "123456789", "merchant_name": "巨星国际", "contact_name": "王鑫", "amount": 1000, "remark": "111", "operator_id": "1", "create_time": 1538016637}, {"id": 2, "bill_no": "123456788", "merchant_name": "巨星国际", "contact_name": "王鑫", "amount": 1000, "remark": "111", "operator_id": "1", "create_time": 1538016637}}
		data_all["data"] = datas
		data_all["recharge_num"] = 20
		data_all["recharge_success_num"] = 10
		data_all["recharge_success_sum"] = 100
	}

	res := gin.H{"code": 1, "data": data_all, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func RechargeTransBill(c *gin.Context) {
	data_all := map[string]interface{}{}

	datas := [...]map[string]interface{}{{"id": 1, "bill_no": "123456789", "merchant_name": "巨星国际", "contact_name": "王忠杰", "trans_account": "432255484848", "amount": 1000, "pay_type": 2, "trans_to_account": "43111111111", "create_time": 1538016637, "check_time": 1538016638, "status": 1}, {"id": 2, "bill_no": "123456788", "merchant_name": "巨星国际", "contact_name": "王忠杰", "trans_account": "432255484848", "amount": 1000, "pay_type": 2, "trans_to_account": "43111111111", "create_time": 1538016637, "check_time": 1538016638, "status": 2}}
	data_all["data"] = datas
	data_all["recharge_num"] = 20
	data_all["recharge_success_num"] = 10
	data_all["recharge_success_sum"] = 100
	res := gin.H{"code": 1, "data": data_all, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func WithdrawBill(c *gin.Context) {
	data_all := map[string]interface{}{}
	datas := [...]map[string]interface{}{{"id": 1, "order_sn": "123456789", "user_name": "fff123", "class_name": "黄金", "true_name": "王鑫", "card_no": "32132113131312", "money": 100.00, "memo": "111", "status": 1, "create_time": 1538016637, "approve_time": 1538016637}, {"id": 2, "order_sn": "123456788", "user_name": "fff123", "class_name": "黄金", "true_name": "王鑫", "card_no": "32132113131312", "money": 100.00, "memo": "111", "status": 2, "create_time": 1538016637, "approve_time": 1538016637}}
	data_all["data"] = datas
	data_all["withdraw_num"] = 10
	data_all["withdraw_success_num"] = 10
	data_all["withdraw_success_sum"] = 100
	res := gin.H{"code": 1, "data": data_all, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func WithdrawBillBack(c *gin.Context) {
	data_all := map[string]interface{}{}
	datas := [...]map[string]interface{}{{"id": 1, "bill_no": "123456789", "merchant_name": "巨星国际", "contact_name": "王鑫", "amount": 1000, "remark": "111", "operator": "小明", "create_time": 1538016637}, {"id": 2, "bill_no": "123456788", "merchant_name": "巨星国际", "contact_name": "王鑫", "amount": 1000, "remark": "111", "operator": "小明", "create_time": 1538016637}}
	data_all["data"] = datas
	data_all["withdraw_num"] = 10
	data_all["withdraw_success_num"] = 10
	data_all["withdraw_success_sum"] = 100
	res := gin.H{"code": 1, "data": data_all, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func GameTransBill(c *gin.Context) {
	data := [...]map[string]interface{}{{"id": 1, "bill_no": "123456789", "user_name": "fff123", "true_name": "王鑫", "type": 1, "money": 100.00, "create_time": 1539601066, "ok": 1}, {"id": 2, "bill_no": "123456789", "user_name": "fff123", "true_name": "王鑫", "type": 2, "money": 100.00, "create_time": 1539601066, "ok": 2}}

	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func QueryUserBalance(c *gin.Context) {
	data := [...]map[string]interface{}{{"id": 1, "balance": 200.00}}
	RespJson(c, status.OK, data)
}

func CenterAccountBalanceSwitch(c *gin.Context) {
	RespSuccess(c)
}

func ThirdAccountBalanceSwitch(c *gin.Context) {
	RespSuccess(c)
}

func QueryGameBalance(c *gin.Context) {
	data := map[string]interface{}{"balance": 200.00}
	res := gin.H{"code": 1, "data": data, "msg": "ok"}
	c.JSON(http.StatusOK, res)
}

func OnlinePaymentsList(c *gin.Context) {
	data := [...]map[string]interface{}{{"id": 1, "pay_type": "alipay", "platform": "高通支付", "code": "5000", "sort": 1, "simple_min": 100.00, "simple_max": 100000, "day_stop_max": 1000, "status": 1}, {"id": 2, "pay_type": "alipay", "platform": "高通支付2", "code": "5000", "sort": 1, "simple_min": 100.00, "simple_max": 100000, "day_stop_max": 1000, "status": 1}}

	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func TransferList(c *gin.Context) {
	data := [...]map[string]interface{}{{"id": 1, "pay_type": "alipay", "remark": "fff", "account": "313213131311313", "sort": 1, "simple_min": 100.00, "day_stop_max": 1000, "income_times_today": 50, "total_transfer_today": 50, "status": 1}, {"id": 2, "pay_type": "alipay", "remark": "fff", "account": "313213131311313", "sort": 1, "simple_min": 100.00, "day_stop_max": 1000, "income_times_today": 50, "total_transfer_today": 50, "status": 1}}

	res := gin.H{"code": 1, "data": data, "msg": "ok", "total": 10, "next_page": 2}
	c.JSON(http.StatusOK, res)
}

func OnlinePaymentsAdd(c *gin.Context) {
	RespSuccess(c)
}

func TransferAdd(c *gin.Context) {
	RespSuccess(c)
}
