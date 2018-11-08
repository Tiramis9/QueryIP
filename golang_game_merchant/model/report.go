package model

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

type BizReport struct {
	Report
}
type Report struct {
	New_register     int            `json:"new_register"`
	Recharge_member  int            `json:"recharge_member"`
	Effective_bet    int            `json:"effective_bet"`
	Win_lost_amount  int            `json:"win_lost_amount"`
	Recharge_amount  int            `json:"recharge_amount"`
	Withdraw_amount  int            `json:"withdraw_amount"`
	Rebate           int            `json:"rebate"`
	Bonus_amount     int            `json:"bonus_amount"`
	Effective_member string         `json:"effective_member"`
	Win_rate         int            `json:"win_rate"`
	Win              int            `json:"win"`
	Tips_list        map[string]int `json:"tips_list"`
	Tips_num         int            `json:"tips_num"`
	Game_name        string         `json:"game_name"`
	Game_active_list string         `json:"game_active_list"`
	Award            int            `json:"award"`
	Pool             int            `json:"pool"`
}

var (
	new_register    int // 新注册会员
	recharge_member int // 充值会员
	recharge_amount int // 充值金额
	withdraw_amount int // 提现
	bonus_amount    int // 红利
	rebate          int // 反水
	effective_bet   int // 有效投注
	win             int // 输赢结果
)

/*
new_register	int	新注册会员
recharge_member	int	充值会员
effective_bet	int	有效投注
win_lost_amount	int	总输赢
recharge_amount	int	充值
withdraw_amount	int	提现
rebate	int	反水  fy
bonus_amount	int	红利
effective_member	int	有效会员
win_rate	int	胜率
win	int	输赢结果
tips_list	string	打赏小费列表
tips_num	int	打赏数额
game_name	string	游戏名或者直播名
game_active_list	string	游戏活动列表
award	int	派奖总金额
pool	int	奖池贡献金
*/

func StringtoInt(message string) int {
	date, err := strconv.Atoi(message)
	if err != nil {
		panic(err)
	}
	return date
}
func GetMerchantAnnouncement(db *gorm.DB, date map[string]int) (*BizReport, error) {
	fmt.Println("weclone to GET SQL")
	start_time := date["start_time"]
	end_time := date["end_time"]

	bizReport := new(BizReport)
	//	if err := db.Table(`user`).Where("reg_time>= ? AND reg_time <=?", start_time, end_time).Find(&user).Error; err != nil {
	if err := db.Table(`user`).Where("reg_time>= ? AND reg_time <=?", start_time, end_time).Count(&new_register).Error; err != nil {
		return nil, err
	}
	if err := db.Table(`user_bill`).Where("create_time>= ? AND create_time <=?", start_time, end_time).Count(&recharge_member).Error; err != nil {
		return nil, err
	}
	if err := db.Table(`user_bill`).Where("create_time>= ? AND create_time <=?", start_time, end_time).Count(&recharge_amount).Error; err != nil { // SQL语句未完成
		return nil, err
	}
	//SELECT * FROM `user_bill `  WHERE (merchant_id='{id}')
	if err := db.Table(`user_bill`).Where("create_time>= ? AND create_time <=?", start_time, end_time).Count(&rebate).Error; err != nil { //
		return nil, err
	}
	bizReport.New_register = new_register
	bizReport.Recharge_member = recharge_member
	bizReport.Recharge_member = recharge_amount
	bizReport.Recharge_member = rebate
	return bizReport, nil
}

/*
new_register	int	新注册会员
recharge_member	int	充值会员
effective_bet	int	有效投注
win_lost_amount	int	总输赢
recharge_amount	int	充值
withdraw_amount	int	提现
rebate	int	反水
bonus_amount	int	红利
effective_member	int	有效会员
win_rate	int	胜率
win	int	输赢结果
tips_list	string	打赏小费列表
tips_num	int	打赏数额
game_name	string	游戏名或者直播名
game_active_list	string	游戏活动列表
award	int	派奖总金额
pool	int	奖池贡献金
*/
