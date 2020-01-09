package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	MysqlSource := "root:687211@tcp(127.0.0.1:3306)/final_exam?charset=utf8"
	var err error
	db, err = sql.Open("mysql", MysqlSource)
	if err != nil {
		fmt.Println("sql.open is error !", err)
	}
}

func Conn() *sql.DB {
	return db
}

