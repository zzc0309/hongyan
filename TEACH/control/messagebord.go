package control

import (
	"GO_WEB/TEACH/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)



func FindMessage(gin *gin.Context)  []model.Message {
	pid:=gin.PostForm("pid")
	pidNew, _ :=strconv.Atoi(pid)
	res:=model.FindMessageByPid(pidNew)
	return res
}

func PostMessage(gin *gin.Context)bool{
	username,err:=gin.Cookie("username")
	if err!=nil{
		fmt.Println("cookie is error:",err)
		return false
	}
	message:=gin.PostForm("message")
	res:=model.PostMessage(message,username)
	return res
}

func Takefri(gin *gin.Context)bool{
	id_me:=gin.PostForm("id_me")
	id_fri:=gin.PostForm("id_fri")
	res:=model.Takefri(id_me,id_fri)
	return res
}

func OnlyFriend(gin *gin.Context)bool{
	id_me:=gin.PostForm("id_me")
	id_fri:=gin.PostForm("id_fri")
	id_meNew,_:=strconv.Atoi(id_me)
	id_friNew,_:=strconv.Atoi(id_fri)
	res:=model.OnlyFri(id_meNew,id_friNew)
	return res
}

func JsonNested(messageSlice []model.Message) []gin.H {
	//order++
	var messageJsons []gin.H
	//fmt.Printf("第%d层开始", order)
	//fmt.Println()
	var messageJson gin.H
	for _, messages := range messageSlice {
		//fmt.Println("分解过程", messages)
		message := *messages.ChildMessage
		//fmt.Println("分解过程的的子留言", message)
		if messages.ChildMessage != nil {
			messageJson = gin.H{
				"user_id":         messages.Username,
				"message":         messages.Message,
				"ChildrenMessage": JsonNested(message),
			}
		} else {
			messageJson = gin.H{
				"user_id": messages.Username,
				"message": messages.Message,
				"ChildrenMessage":"null",
			}
		}
		messageJsons = append(messageJsons, messageJson)
	}
	//fmt.Printf("第%d层结束。", order)
	//fmt.Println()
	//order--
	return messageJsons
}

