package dao

import (
	"fmt"
	"redRockTest/util"
)

func FindMessageByPid(pid int) []util.Message {
	db := util.Conn()

	res, err := db.Query("select id,message,username from messageboard where pid=?", pid)

	if err != nil {
		fmt.Println("du.query findMessage is error !")
	}

	var id int
	var messageSlice []util.Message
	for res.Next() {		//如果该pid没有评论  则不会进入for循环
		var messages util.Message
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
