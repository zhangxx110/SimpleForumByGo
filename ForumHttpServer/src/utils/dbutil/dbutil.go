package dbutil

import (
	"database/sql"
	"database/sql/driver"
	"log"
	_ "utils/github.com/go-sql-driver/mysql"
)

var tag_DBUtil string = "DBUtil"
var db *sql.DB = nil

/**
* 初始化数据库连接
* driverName：驱动名称，如"mysql"
* dataSourceName:它是go-sql-driver定义的一些数据库链接和配置信息。它支持如下格式：
* user@unix(/path/to/socket)/dbname?charset=utf8
* user:password@tcp(localhost:5555)/dbname?charset=utf8
* user:password@/dbname
* user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
**/
func Init(driverName string, dataSourceName string) error {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Println(tag_DBUtil, err)
		return err
	}
	return nil
}

/*
* 释放数据库连接
**/
func Relase() {
	if db == nil {
		return
	}
	db.Close()
}

/**
*向表插入一条记录
* tableName：数据库表名
* property：属性名
* values：属性值
* bool：插入成功返回true,否则范湖false
**/
func Insert(tableName string, property []string, values []interface{}) bool {
	if db == nil {
		log.Println(tag_DBUtil, "insert, db is nil")
		return false
	}
	log.Println(tag_DBUtil, property, values)
	if len(property) < 1 || len(property) != len(values) {
		log.Println(tag_DBUtil, "insert, property length is not equal values")
		return false
	}
	var mysql string = "insert into " + tableName + "("
	for i := 0; i < len(property); i++ {
		if i < len(property)-1 {
			mysql += property[i] + ","
		} else {
			mysql += property[i] + ")"
		}
	}
	mysql += " values("
	for i := 0; i < len(values); i++ {
		if i < len(values)-1 {
			mysql += "?,"
		} else {
			mysql += "?)"
		}
	}
	log.Println(tag_DBUtil, mysql)
	stmt, err := db.Prepare(mysql)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return false
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return false
	}
	return true
}

/**
* update tableName set property = values where whereProperty whereOpt whereValue
**/
func Update(tableName string, property []string, values []interface{}, whereProperty string, whereOpt string, whereValue interface{}) bool {
	if db == nil {
		log.Println(tag_DBUtil, "Update, db is nil")
		return false
	}
	log.Println(tag_DBUtil, property, values, whereProperty, whereValue)
	if len(property) < 1 || len(property) != len(values) {
		log.Println(tag_DBUtil, "Update, property length is not equal values")
		return false
	}
	var mysql string = "update " + tableName + " set"
	for i := 0; i < len(property); i++ {
		if i == 0 {
			mysql += " " + property[i] + " = ?"
		} else {
			mysql += " ," + property[i] + " = ?"
		}
	}
	//where
	if whereProperty != "" && whereOpt != "" {
		mysql += " where " + whereProperty + " " + whereOpt + " ?"
	}
	log.Println(tag_DBUtil, mysql)
	stmt, err := db.Prepare(mysql)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return false
	}
	if len(whereProperty) < 1 {
		_, err = stmt.Exec(values...)
	} else {
		newValue := append(values, whereValue)
		_, err = stmt.Exec(newValue...)
	}
	if err != nil {
		log.Println(tag_DBUtil, err)
		return false
	}
	return true
}

/**
* 删除表的某行（多行）记录
* delete from tableName where whereState
**/
func Delete(tableName string, whereState string) bool {
	if tableName == "" {
		log.Println(tag_DBUtil, "Delete tableName is nil")
		return false
	}
	if db == nil {
		log.Println(tag_DBUtil, "Delete, db is nil")
		return false
	}
	var mysql string = "delete from " + tableName
	if whereState != "" {
		mysql += " where " + whereState
	}
	stmt, err := db.Prepare(mysql)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return false
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(tag_DBUtil, err)
		return false
	}
	return true
}

func Query(tableName string, propertys []string, whereState string, orderByState string, groupByState string) (*sql.Rows, bool) {
	if tableName == "" {
		log.Println(tag_DBUtil, "Query tableName is nil")
		return nil, false
	}
	if db == nil {
		log.Println(tag_DBUtil, "Query, db is nil")
		return nil, false
	}
	var mysql string
	if len(propertys) < 1 {
		mysql = "select * from " + tableName
	} else {
		mysql = "select "
		for i := 0; i < len(propertys); i++ {
			if i < len(propertys)-1 {
				mysql += propertys[i] + ","
			} else {
				mysql += propertys[i] + " from " + tableName
			}
		}
	}
	//where
	if whereState != "" {
		mysql += " where " + whereState
	}
	// order by
	if orderByState != "" {
		mysql += " order by " + orderByState
	}
	//group by
	if groupByState != "" {
		mysql += " group by " + groupByState
	}
	log.Println(tag_DBUtil, mysql)
	rows, err := db.Query(mysql)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return nil, false
	}
	return rows, true
}

/*
*执行一个sql语句，sql中的变量用？代替，用？代替的值全部放在value数组中
**/
func ExcuteSql(sql string, value []interface{}) (driver.Result, bool) {
	if sql == "" {
		log.Println(tag_DBUtil, "ExcuteSql sql is nil")
		return nil, false
	}
	if db == nil {
		log.Println(tag_DBUtil, "ExcuteSql, db is nil")
		return nil, false
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println(tag_DBUtil, err)
		return nil, false
	}
	var result driver.Result
	if len(value) < 1 {
		result, err = stmt.Exec()
	} else {
		result, err = stmt.Exec(value...)
	}
	if err != nil {
		log.Println(tag_DBUtil, err)
		return nil, false
	}
	return result, true
}
