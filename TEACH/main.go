package main

import (
	"GO_WEB/TEACH/control"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	http.Handle("/static",http.StripPrefix("/static",http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/",control.IndexView)
	http.HandleFunc("/bilibili",control.BIndexView)//首页
	http.HandleFunc("/list",control.ListView)//新闻列表
	http.HandleFunc("/edit",control.EditView)//新闻编辑
	http.HandleFunc("/detail",control.DetailView)//详情
	http.HandleFunc("/add",control.ViewArticleAdd)//新闻添加
	http.HandleFunc("/api/index/data",control.IndexData)
	http.HandleFunc("/api/list/data",control.ListData)
	http.HandleFunc("/api/list/del",control.ListDel)
	http.HandleFunc("/api/article/add",control.ApiArticleAdd)
	http.HandleFunc("/api/article/edit",control.ApiArticleEdit)
	router:=gin.Default()
	router.POST("/registe", func(ctx *gin.Context) {
		res := control.Registe(ctx)
		if res {
			ctx.JSON(200, gin.H{
				"恭喜注册成功！": ctx.PostForm("username"),
			})
		} else {
			ctx.JSON(400, gin.H{
				"注册失败": "!",
			})
		}
	}) //注册
	router.POST("/login", func(ctx *gin.Context) {
		res := control.Login(ctx)
		if res {
			ctx.JSON(200, gin.H{
				"恭喜登录成功！": ctx.PostForm("username"),
			})
		} else {
			ctx.JSON(400, gin.H{
				"账号或密码错误": "!",
			})
		}
	})  //登录
    //留言
	router.POST("/PostMessage", func(ctx *gin.Context) {
		res:=control.PostMessage(ctx)
		if res{
			ctx.JSON(200,gin.H{
				"发表成功":" ",
			})
		}else{
			ctx.JSON(400,gin.H{
				"发表失败":" ",
			})
		}
	})

	router.POST("/takefri", func(ctx *gin.Context) {
		res:=control.Takefri(ctx)
		if res{
			ctx.JSON(200,gin.H{
				"添加成功":" ",
			})
		}else{
			ctx.JSON(400,gin.H{
				"添加失败":" ",
			})
		}
	})

	router.POST("/findmessage", func(ctx *gin.Context) {
		Responce:=control.OnlyFriend(ctx)
		if Responce{
			res := control.FindMessage(ctx)
			json := control.JsonNested(res)
			ctx.JSON(200, json)
		}else{
			ctx.JSON(400,gin.H{
				"你跟他不是好友":"--->拒绝",
			})
		}
	})

	router.Run(":8080")
	http.ListenAndServe(":8080",nil)
}
