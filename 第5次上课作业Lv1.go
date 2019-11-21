package main
import (

	"github.com/gin-gonic/gin"

)
func main(){
	var user2 User2
engin:=gin.Default()
engin.POST("/cookie", func(context *gin.Context) {
	context.BindJSON(&user2)
	cookie,err:=context.Cookie("key_cookie")

	c:="hello "+user2.Username
	context.SetCookie("key_cookie",c,60,
		"/","localhost",false,true)
	//判断该用户有没有COOKIE
	//没有就"hello guest"
	//有就"hello 用户"
	if err!=nil{
		resp:=Response{Code:200,Message:"hello guest"}
		context.JSON(200,&resp)
	}else{
		resp:=Response{200,cookie}
		context.JSON(200,&resp)

	}
})
engin.Run(":8080")
}
type User2 struct {
	Username string `form:"username"`
}
type Response struct {
	Code int
	Message string
}