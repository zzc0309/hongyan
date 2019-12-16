package dao

import (
	"fmt"
	"log"
	"redRockTest/util"
)

func Login(username string, password string) bool {
	db := util.Conn()
	res, err := db.Query("select id from user where username=? and password=? ", username, password)
fmt.Println("ssssssssssssssssssss")
	if err != nil {
		fmt.Println("db.query is error login:", err)
	}

	for res.Next() {
		var id int
		err := res.Scan(&id)
		fmt.Println("id=======",id)
		if err != nil {
			fmt.Println("res.Scan is error:", err)
		}

		if id >= 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func Registe(username string, password string) bool {
	db := util.Conn()
	res, err := db.Query("select id from user where username=? ", username)

	if err != nil {
		fmt.Println("db.query is error login:", err)
	}

	for res.Next() {
		var id int
		err := res.Scan(&id)
		fmt.Println("id=======",id)
		if err != nil {
			fmt.Println("res.Scan is error:", err)
		}

		if id >= 0 {
			return false
		} else {
			break
		}
	}
	stmt,err:=db.Prepare("insert into user(username,password) value (?,?)")
	if err!=nil{
		fmt.Println("错误")
		log.Fatal(err)
		return false
	}
	stmt.Exec(username,password)
	return true
}

func Delete(dlmessage string)bool{
	db:=util.Conn()
	stmt,err:=db.Prepare("delete from messageboard where message=?")
	if err!=nil{
		fmt.Println("delete is error:",err)
		return false
	}
	stmt.Exec(dlmessage)
	return true
}

func PostMessage(message string,username string)bool{
	db:=util.Conn()
	stmt,err:=db.Prepare("insert into messageboard(username,message,pid,zan) value (?,?,?,?)")
	if err!=nil{
		fmt.Println("postmessage is error:",err)
		return false
	}
	stmt.Exec(username,message,0,0)
	return true
}

func Takefri(id_me string,id_fri string)bool{
	db:=util.Conn()
	stmt,err:=db.Prepare("insert into friend(id_me,id_fri) value (?,?)")
	if err!=nil{
		fmt.Println("Take friends is error:",err)
		return false
	}
	stmt.Exec(id_me,id_fri)
	return true
}

func OnlyFri(id_me int,id_fri int)bool{
	db:=util.Conn()
	rows,err:=db.Query("select id_fri from friend where id_me=?",id_me)
	if err!=nil{
		fmt.Println("Onlyfri SQL is error:",err)
	}
	for rows.Next(){
		var id int
		err:=rows.Scan(&id)
		if err!=nil{
			fmt.Println("Scan friend is error:",err)
		}
		if id==id_fri{return true}
	}
	return false
}