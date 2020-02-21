package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
var Db *sqlx.DB
func init(){
	db,err:=sqlx.Open(`mysql`,`root:687211@tcp(127.0.0.1:3306)/teach?charset=utf8`)
	if err!=nil{
		log.Fatalln(err.Error())
	}
	Db=db
}