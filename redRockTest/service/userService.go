package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redRockTest/dao"
	"redRockTest/util"
	"strconv"
)

func Login(gin *gin.Context) bool {
	username := gin.PostForm("username")
	password := gin.PostForm("password")
	gin.SetCookie("username", username, 100, "/", "localhost", false, true)
	loginRes := dao.Login(username, password)
	return loginRes
}

func Registe(gin *gin.Context) bool {
	username := gin.PostForm("username")
	password := gin.PostForm("password")
	registeRes := dao.Registe(username, password)
	return registeRes
}

func Delete(gin *gin.Context)bool{
	dlmessage:=gin.PostForm("dlmessage")
	deleteRes:=dao.Delete(dlmessage)
	return deleteRes
}

func FindMessage(gin *gin.Context)  []util.Message {
	pid:=gin.PostForm("pid")
	pidNew, _ :=strconv.Atoi(pid)
	res:=dao.FindMessageByPid(pidNew)
	return res
}

func PostMessage(gin *gin.Context)bool{
	username,err:=gin.Cookie("username")
	if err!=nil{
		fmt.Println("cookie is error:",err)
		return false
	}
	message:=gin.PostForm("message")
	res:=dao.PostMessage(message,username)
	return res
}

func Takefri(gin *gin.Context)bool{
	id_me:=gin.PostForm("id_me")
	id_fri:=gin.PostForm("id_fri")
	res:=dao.Takefri(id_me,id_fri)
	return res
}

func OnlyFriend(gin *gin.Context)bool{
	id_me:=gin.PostForm("id_me")
	id_fri:=gin.PostForm("id_fri")
	id_meNew,_:=strconv.Atoi(id_me)
	id_friNew,_:=strconv.Atoi(id_fri)
	res:=dao.OnlyFri(id_meNew,id_friNew)
	return res
}