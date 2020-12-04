package user

import (
	"commons"
	"fmt"
)

//返回nil查询失败
func SelectByUnPwdDao(un, pwd string) *TbUser {
	sql := "select * from tb_user where username =? and password =? or email =? and password =? or phone =? and password =?"
	rows, err := commons.Dql(sql, un, pwd, un, pwd, un, pwd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {
		user := new(TbUser)
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Created, &user.Updated)
		commons.CloseConn()
		return user
	}
	return nil
}
