package main
import (
"database/sql"
"fmt"
"log"
"github.com/gin-gonic/gin"
_ "github.com/go-sql-driver/mysql"
)
var message []string=make([]string,0)
var username []interface{}=make([]interface{},5)
type Reply struct {
	username string
	message string
}
type User struct {
	username string
	password string
}
func Login(username string,password string,db *sql.DB)bool {
	rows,_:=db.Query("select password from user where username=?",username)
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.password)
		if password == user.password {
			return true
		}else{
			return false
		}
	}
	return true
}
func Registe(username string,password string,db *sql.DB)bool{
	/*rows,_:=db.Query("select username form user")
	defer rows.Close()
	for rows.Next(){
		var name string
		rows.Scan(&name)
		fmt.Println(name)
		if name==username{
			return false
		}
	}*/
	stmt,err:=db.Prepare("insert into user(username,password) value (?,?)")
	if err!=nil{
		fmt.Println("错误")
		log.Fatal(err)
		return false
	}
	stmt.Exec(username,password)
	return true
}
func Send(username string,message string,db *sql.DB)bool{
	stmt,err:=db.Prepare("insert into messageboard(username,message,pid,zan) value (?,?,?,?)")
	if err!=nil{
		fmt.Println("错误")
		log.Fatal(err)
		return false
	}
	stmt.Exec(username,message,0,0)
	return true
}//发送
func findContentChild(message_louzhu string,db *sql.DB) {
	rows,_:= db.Query("select id from messageboard where message=?", message_louzhu)
	defer rows.Close()
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	pid:=id
	rows1,_:= db.Query("select message from messageboard where pid=?", pid)
	defer rows.Close()
	i:=-1
	//fmt.Println(len(message))
	y:=len(message)
		for rows1.Next(){
			i++
			var x string
			rows1.Scan(&x)
			message=append(message,x)
			/*if y==len(message){
				fmt.Println("成功")
				//var examine interface{}
				//examine=message[0]
				//fmt.Println(examine)
				return*/
			}
			if y==len(message){
				fmt.Println("成功")
				return}
	//findContentChild(message[0],db)
		}//递归没写好,退不出来
	//fmt.Println(len(message))
	/*for j:=0;j<=len(message)-1;j++{
		findContentChild(message[j],db)
	}*/
//1.没有登录将不会返回信息,直接断开连接  2.不能检查用户名,只能检查密码 3.点赞需要测试  4.用interface类型的切片存储从mysql扫描来的varchar类型,结果切片里面全是一些数字
func main(){
	engin:=gin.Default()
	db,_:= sql.Open("mysql", "root:687211@tcp(127.0.0.1:3306)/seven")
	engin.POST("/reply", func(c *gin.Context) {
		username,err:=c.Cookie("username")
		if err!=nil||username==" "{
			c.JSON(200,gin.H{"错误":"你还没有登录"})
			log.Fatal(err)
			return
		}
		message:=c.PostForm("message")
		louzhu_name:=c.PostForm("louzhu_name")
		rows,_:= db.Query("select id from messageboard where username=?", louzhu_name)
		defer rows.Close()
		var louzhu_id int
		for rows.Next() {
			rows.Scan(&louzhu_id)
		}
		stmt,_:=db.Prepare("insert into messageboard(username,message,pid,zan) value (?,?,?,?)")
		stmt.Exec(username,message,louzhu_id,0)
		/*rows1, _ := db.Query("select username,message from messageboard where pid=?",louzhu_id)
		defer rows1.Close()
		for rows1.Next() {
			line := Reply{}
			rows1.Scan(&line.username,&line.message)
			c.JSON(200, gin.H{
				"回复": line.message, "用户": line.username})
		}*/

	})         //回复功能:需要输入发送的信息和回复的人
	engin.POST("/registe", func(context *gin.Context) {
		username:=context.PostForm("username")
		password:=context.PostForm("password")
		if Registe(username,password,db){
			context.Writer.WriteString("注册成功")
		}else {
			context.Writer.WriteString("注册失败")
		}
	})//注册
	engin.POST("/login", func(context *gin.Context) {

		username:=context.PostForm("username")
		password:=context.PostForm("password")
		context.SetCookie("username", username, 100, "/", "localhost", false, true)
		if Login(username,password,db){
			context.Writer.WriteString("登录成功")
		}else {
			context.Writer.WriteString("密码,用户名错误")
		}

	})//登录
	engin.POST("/send", func(context *gin.Context) {
		username,err:=context.Cookie("username")
		if username==" "||err!=nil{
			context.Writer.WriteString("请登录")
			context.JSON(200,gin.H{"错误":"你还没有登录"})
			log.Fatal(err)//有问题!!!!!!!!!!!!!!!!!!!!!!!!!!!
			return
		}
		message:=context.PostForm("message")
		if Send(username,message,db){
			context.JSON(200,gin.H{"用户":username,"留言":message})
		}else{context.JSON(200,gin.H{"系统消息":"发送失败"})}

	})//发帖子 需要发送的信息
	engin.POST("/dianzan", func(c *gin.Context) {
		username,err:=c.Cookie("username")
		if err!=nil||username==" "{
			c.JSON(200,gin.H{"错误":"你还没有登录"})
			log.Fatal(err)
			return
		}
		username_zan:=c.PostForm("username_zan")
		fmt.Println(username_zan)
		rows,_:= db.Query("select zan from messageboard where username=? and pid=?",username_zan,0)
		defer rows.Close()
		var zan int
		for rows.Next() {
			rows.Scan(&zan)
		}
		fmt.Println(zan)
		zan=zan+1
		fmt.Println(zan)
		stmt,_:=db.Prepare("update messageboard set zan=? where username=? and pid=?")
		defer stmt.Close()
		stmt.Exec(zan,username_zan,0)

	})//点赞
	engin.POST("/redianzan", func(c *gin.Context) {
		username,err:=c.Cookie("username")
		if err!=nil||username==" "{
			c.JSON(200,gin.H{"错误":"你还没有登录"})
			log.Fatal(err)
			return
		}
		username_zan:=c.PostForm("username_zan")
		rows,_:= db.Query("select zan from messageboard where username=? and pid=?",username_zan,0)
		defer rows.Close()
		var zan int
		for rows.Next() {
			rows.Scan(&zan)
		}
		zan=zan-1
		stmt,_:=db.Prepare("update messageboard set zan=? where username=? and pid=?")
		defer stmt.Close()
		stmt.Exec(zan,username_zan,0)

	})//取消点赞
	engin.POST("/chakanliuyan", func(c *gin.Context) {
		message_louzhu:=c.PostForm("message_louzhu")
		findContentChild(message_louzhu,db)
		for i:=0;i<=len(message)-1;i++{
			c.JSON(200,gin.H{"消息":message[i]})
		}
	})//查看留言,输入想查看的帖子(不完善)
		engin.POST("/zhuxiao", func(c *gin.Context) {
		c.SetCookie("username"," ",100,"/","localhost",false,true)
	})//注销
	engin.Run()
}
