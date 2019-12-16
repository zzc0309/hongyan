package api

import (
	"github.com/gin-gonic/gin"
	"redRockTest/service"
	"redRockTest/util"
)

func Entrance() {
	router := gin.Default()
	router.POST("/login", func(ctx *gin.Context) {
		res := service.Login(ctx)
		if res {
			ctx.JSON(200, gin.H{
				"恭喜登录成功！": ctx.PostForm("username"),
			})
		} else {
			ctx.JSON(400, gin.H{
				"账号或密码错误": "!",
			})
		}
	})

	router.POST("/registe", func(ctx *gin.Context) {
		res := service.Registe(ctx)
		if res {
			ctx.JSON(200, gin.H{
				"恭喜注册成功！": ctx.PostForm("username"),
			})
		} else {
			ctx.JSON(400, gin.H{
				"注册失败": "!",
			})
		}
	})

	router.POST("/delete", func(ctx *gin.Context) {
		res:=service.Delete(ctx)
		if res{
			ctx.JSON(200,gin.H{
				"删除成功":" ",
			})
		}else{
			ctx.JSON(400,gin.H{
				"删除失败":" ",
			})
		}
	})

	router.POST("/PostMessage", func(ctx *gin.Context) {
		res:=service.PostMessage(ctx)
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
		res:=service.Takefri(ctx)
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
		Responce:=service.OnlyFriend(ctx)
		if Responce{
			res := service.FindMessage(ctx)
			json := util.JsonNested(res)
			ctx.JSON(200, json)
		}else{
			ctx.JSON(400,gin.H{
				"你跟他不是好友":"--->拒绝",
			})
		}
	})
	router.Run(":8080")
}



