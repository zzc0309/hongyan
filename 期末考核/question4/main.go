package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db,err:=sql.Open("mysql","root:687211@tcp(127.0.0.1:3306)/final_exam?charset=utf8")
	if err!=nil{fmt.Println("open mysql is wrong:",err)}
	engin:=gin.Default()
	engin.POST("/sign", func(c *gin.Context) {
	cookie,err:=c.Cookie("key_cookie")
	if err!=nil {
		cookie="not set"
		c.SetCookie("key_cookie","value_cookie",60,"/","localhost",false,true)
		user:=c.PostForm("user")
		rows,err:= db.Query("select point from question4 where user=?",user)
		if err!=nil{fmt.Println("select出错:",err)}
		defer rows.Close()
		var point int
		for rows.Next() {
			err:=rows.Scan(&point)
			if  err!=nil{fmt.Println("scan is wrong:",err)}
		}
		point+=5
		stmt,err:=db.Prepare("update question4 set point=? where user=?")
		if err!=nil {fmt.Println("prepare is wrong:",err)}
		defer stmt.Close()
		stmt.Exec(point,user)
		c.JSON(200,gin.H{user:"签到完成,积分+5"})
		}else
		{
			c.JSON(500,gin.H{"false":"每天只签到一次"})
		}
fmt.Println(cookie)
})
	engin.POST("/cash_prizes", func(c *gin.Context) {
		user:=c.PostForm("user")
		prize:=c.PostForm("prize")
		rows,err:= db.Query("select point from question4 where user=?",user)
		if err!=nil{fmt.Println("select出错:",err)}
		defer rows.Close()
		var point int
		for rows.Next() {
			err:=rows.Scan(&point)
			if  err!=nil{fmt.Println("scan is wrong:",err)}
		}
		fmt.Println(prize)
		rows2,err2:= db.Query("select point from question4_1 where prize=?",prize)

		if err!=nil{fmt.Println("select出错:",err2)}
		defer rows2.Close()
		var point2 int
		for rows2.Next() {
			err:=rows2.Scan(&point2)
			if  err!=nil{fmt.Println("scan is wrong:",err)}
		}
		if point>=point2{
			point-=point2
			stmt,err:=db.Prepare("update question4 set point=? where user=?")
			if err!=nil {fmt.Println("prepare is wrong:",err)}
			defer stmt.Close()
			stmt.Exec(point,user)
			c.Writer.WriteString("兑换成功!!!")
		}else{
			c.Writer.WriteString("积分不足!!")
		}
	})
	engin.POST("/charge_point", func(c *gin.Context) {
		admin:=c.PostForm("admin")
		user:=c.PostForm("user")
		res, err := db.Query("select id from question4_2 where admin=? ", admin)
		if err != nil {
			fmt.Println("db.query is error login:", err)
		}
		for res.Next() {
			var id int
			err := res.Scan(&id)
			if err != nil {
				fmt.Println("res.Scan is error:", err)
			}
			if id >= 0 {
				rows,err:= db.Query("select point from question4 where user=?",user)
				if err!=nil{fmt.Println("select出错:",err)}
				defer rows.Close()
				var point int
				for rows.Next() {
					err:=rows.Scan(&point)
					if  err!=nil{fmt.Println("scan is wrong:",err)}
				}
				c.JSON(200,gin.H{user:point})
				fmt.Println("你是管理员")
			} else {
				fmt.Println("你不是管理员")
			}
		}})
	engin.POST("/charge_prize", func(c *gin.Context) {
		admin:=c.PostForm("admin")
		prize:=c.PostForm("prize")
		point:=c.PostForm("point")
		res, err := db.Query("select id from question4_2 where admin=? ", admin)
		if err != nil {
			fmt.Println("db.query is error login:", err)
		}
		for res.Next() {
			var id int
			err := res.Scan(&id)
			if err != nil {
				fmt.Println("res.Scan is error:", err)
			}
			if id >= 0 {
				stmt,err:=db.Prepare("insert into question4_1(prize,point) values (?,?) ")
				if err!=nil{
					fmt.Println("insert is wrong:",err)
				}
				stmt.Exec(prize,point)
				fmt.Println("你是管理员")
			} else {
				fmt.Println("你不是管理员")
			}
		}})
	engin.Run()
}