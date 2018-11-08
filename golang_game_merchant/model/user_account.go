package model

type UserAccount struct {
	Game_name    string  `json:"game_name,omitempty"`
	Account_name string  `json:"account_name,omitempty"`
	User_id      int     `json:"user_id,omitempty"`
	Money        float64 `json:"money,omitempty"`
}

func (u UserAccount) GetAccountListByUserId(user_id int) interface{} {
	return nil
	/*var accountlist []UserAccount
	stmt, err := Db.Prepare("SELECT sg.channel FROM merchant_game mg LEFT JOIN sys_game sg ON mg.game_id=sg.id " +
		"WHERE mg.merchant_id = ? GROUP BY sg.channel")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rows, err := stmt.Query(user_id)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var conv UserAccount
		rows.Scan(&conv.Game_name)
		accountlist = append(accountlist, conv)
	}
	return accountlist*/
}

func (u UserAccount) TransAccount(user_id int, from string, to string, amount float64) (bool, string) {
	return true, ""
	/*//从什么账户转至什么账户(开启事务)
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println(err)
		return false, "system error"
	}
	//判断中心余额是否足够
	stmt, err := conn.Prepare("SELECT balance FROM user WHERE id=?")
	defer stmt.Close()
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return false, "system error"
	}
	row := stmt.QueryRow(user_id)
	var balance float64
	row.Scan(&balance)
	//转出账户是中心账户
	if from == "0" {
		if balance < amount {
			return false, "中心账户余额不足"
		}

		//先减中心账户的金额
		stmt2, err := conn.Prepare("UPDATE user SET balance=balance-? WHERE id=?")
		defer stmt2.Close()
		if err != nil {
			conn.Rollback()
			fmt.Println(err)
			return false, "system error"
		}
		ret, err := stmt2.Exec(amount, user_id)
		if err != nil {
			conn.Rollback()
			fmt.Println(err)
			return false, "system error"
		}
		if _, err := ret.RowsAffected(); nil == err {
		} else {
			return false, "system error"
		}
		//调用第三方接口
	} else {
		//调用第三方接口
		//增加中心账户余额
		stmt2, err := conn.Prepare("UPDATE user SET balance=balance+? WHERE id=?")
		defer stmt2.Close()
		if err != nil {
			conn.Rollback()
			fmt.Println(err)
			return false, "system error"
		}
		ret, err := stmt2.Exec(amount, user_id)
		if err != nil {
			conn.Rollback()
			fmt.Println(err)
			return false, "system error"
		}
		if _, err := ret.RowsAffected(); nil == err {
		} else {
			return false, "system error"
		}
	}
	//转账记录
	old_balance := 100.00
	new_balance := old_balance + u.Money
	bill_no := utils.CreateOrderNo(user_id)
	//增加转账记录
	stmt2, err := conn.Prepare("INSERT INTO user_account_bill (account_name, user_id, money, ok, " +
		"old_balance, new_balance, bill_no, create_time, update_time) values(?,?,?,?,?,?,?,?,?)")
	defer stmt2.Close()
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return false, "system error"
	}
	timestamp := time.Now().Unix()
	ret, err := stmt2.Exec(u.Account_name, u.User_id, u.Money, 1, old_balance, new_balance, bill_no, timestamp, timestamp)
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return false, "system error"
	}
	if _, err := ret.RowsAffected(); nil == err {
	} else {
		conn.Rollback()
		return false, "system error"
	}
	//交易明细
	stmt2, err = conn.Prepare("INSERT INTO user_bill (user_id, type, sett_amt, about, " +
		"balance, old_balance, order_sn, code, code_sn, create_time, update_time) values(?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt2.Close()
	if err != nil {
		fmt.Println(err)
		return false, "system error"
	}
	trans_type := 1
	center_new_balance := balance
	if from == "0" { //中心账户转出
		trans_type = -1
		center_new_balance = center_new_balance - u.Money
	} else {
		center_new_balance = center_new_balance + u.Money
	}
	order_no := utils.CreateOrderNo(user_id)
	ret, err = stmt2.Exec(u.User_id, trans_type, u.Money, u.Account_name, center_new_balance, balance,
		order_no, 300, bill_no, timestamp, timestamp)
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return false, "system error"
	}
	if _, err := ret.RowsAffected(); nil == err {
	} else {
		conn.Rollback()
		return false, "system error"
	}
	conn.Commit()
	return true, ""*/
}
