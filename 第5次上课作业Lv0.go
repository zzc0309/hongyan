package main

import (

	"github.com/gin-gonic/gin"

)


func main() {
	Usermap:=map[string]string{
		"zzc0309":"1498155548",
	}

	var user User
	engin:=gin.Default()
	//注册
	routerGroup:=engin.Group("/user")
	routerGroup.POST("/register", func(context *gin.Context) {

		context.ShouldBind(&user)

		if _,ok:=Usermap[user.Username];!ok{
			context.Writer.WriteString("该用户名未被使用,注册成功")
			Usermap[user.Username]=user.Password
		}else{
			context.Writer.WriteString("该用户名已经被使用,注册失败")
		}
	})
	//登录

	routerGroup.POST("/login", func(context *gin.Context) {
		context.ShouldBind(&user)
		if _,ok:=Usermap[user.Username];ok{
			if user.Password==Usermap[user.Username]{
				context.Writer.WriteString("恭喜,登录成功")

			}else{
				context.Writer.WriteString("很遗憾,密码或用户名错误")
			}
		}else{
			context.Writer.WriteString("该用户不存在")
		}

	})

	engin.Run()
}
type User struct{
	Username string `form:"username"`
	Password string `form:"password"`
}
