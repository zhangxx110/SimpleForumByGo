package dao

import (
	"utils/dbutil"
	"utils/logger"
)

var user string = "user_dao"

/**
*向user表中增加一条记录
 */
func AddUser(name string, password string) bool {
	if len(name) < 6 || len(password) < 6 {
		logger.Println(user, " name or password less 6 letter")
		return false
	}
	var property = []string{"name", "password"}
	values := []interface{}{name, password}
	state := dbutil.Insert("user_pwd", property, values)
	return state
}

/**
*检查用户名密码是否正确
 */
func CheckUserPwd(uname string, pwd string) bool {
	var property []string = []string{"name", "password"}
	rows, _ := dbutil.Query("user_pwd", property, "name = "+"'"+uname+"'", "", "")
	if rows == nil {
		return false
	}
	for rows.Next() {
		var name string
		var password string
		err := rows.Scan(&name, &password)
		if err != nil {
			logger.Println(err)
			continue
		}
		logger.Println(name, password)
		if uname == name && password == pwd {
			return true
		}
	}
	return false
}

/**
*检查用户名是否存在
 */
func CheckUserName(uname string) bool {
	var property []string = []string{"name"}
	rows, _ := dbutil.Query("user_pwd", property, "name = "+"'"+uname+"'", "", "")
	if rows == nil {
		return false
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			logger.Println(err)
			continue
		}
		logger.Println(name)
		if uname == name {
			return true
		}
	}
	return false
}
