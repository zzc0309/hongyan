package model

import (

	"fmt"
)

func PostMessage(message string,username string)bool{
	db:=Db
	stmt,err:=db.Prepare("insert into messageboard(username,message,pid,zan) value (?,?,?,?)")
	if err!=nil{
		fmt.Println("postmessage is error:",err)
		return false
	}
	stmt.Exec(username,message,0,0)
	return true
}

func Takefri(id_me string,id_fri string)bool{
	db:=Db
	stmt,err:=db.Prepare("insert into friend(id_me,id_fri) value (?,?)")
	if err!=nil{
		fmt.Println("Take friends is error:",err)
		return false
	}
	stmt.Exec(id_me,id_fri)
	return true
}

func OnlyFri(id_me int,id_fri int)bool{
	db:=Db
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

func FindMessageByPid(pid int) []Message {
	db :=Db

	res, err := db.Query("select id,message,username from messageboard where pid=?", pid)

	if err != nil {
		fmt.Println("du.query findMessage is error !")
	}

	var id int
	var messageSlice []Message
	for res.Next() {		//如果该pid没有评论  则不会进入for循环
		var messages Message
		err := res.Scan(&id, &messages.Message, &messages.Username)
		if err != nil {
			fmt.Println("res.scan id is error !", err)
		}
		child := FindMessageByPid(id)
		messages.ChildMessage = &child
		messageSlice = append(messageSlice, messages)
		//fmt.Println("数据库装载数据时我们的messages", messages, "----messageSlice:", messageSlice)
	}
	return messageSlice
}

type Message struct {
	Username string
	Message string
	ChildMessage *[]Message
}