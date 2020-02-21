package model

import (
	"fmt"
	"log"
)

//func IndexLogin(username string,password string)bool{
//		//var username2 string
//		var password2 string
//		err:=Db.Unsafe().Get(password2,"select * from user where username=?",username)
//		if err!=nil{fmt.Println(err)}
//		if strings.EqualFold(password2,password){
//			return true
//		}
//	return false
//}
//注册数据库操作
func Registe(username string, password string) bool {
	db :=Db
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

func Login(username string, password string) bool {
	db :=Db
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
