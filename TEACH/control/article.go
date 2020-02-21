package control

import (
	"GO_WEB/TEACH/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func ApiArticleAdd(w http.ResponseWriter,r *http.Request){
//	r.ParseForm()
//	mod:=&model.Article{}
//	mod.Title=r.Form.Get("title")
//	mod.Author=r.Form["author"][0]
//	mod.Content=r.FormValue("content")
//	mod.Hits,_=strconv.Atoi(r.Form.Get("hits"))
//	err:=model.ArcticleAdd(mod)
//	if err==nil{
//		Succ(w,"添加成功")
//		return
//	}
//	Fail(w,"添加失败"+err.Error())
//	return
//}

func ApiArticleAdd(w http.ResponseWriter,r *http.Request){
	mod:=&model.Article{}
	err:=json.NewDecoder(r.Body).Decode(mod)
	if err!=nil{
		Fail(w,"输入数据有误",err.Error())
		return
	}
	err=model.ArcticleAdd(mod)
	if err!=nil{
	Fail(w,"添加失败"+err.Error())
	return
	}
	Succ(w,"添加成功")
	return
}

func ApiArticleEdit(w http.ResponseWriter,r *http.Request){
	mod:=&model.Article{}
	err:=json.NewDecoder(r.Body).Decode(mod)
	if err!=nil{
		Fail(w,"输入数据有误",err.Error())
		return
	}
	err=model.ArcticleEdit(mod)
	if err!=nil{
		Fail(w,"修改失败"+err.Error())
		return
	}
	Succ(w,"修改成功")
	return
}

//主页

//主页数据
func IndexData(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	w.Header().Set("Content-Type","application/json")
	idStr:=r.Form.Get("id")
	id,_:=strconv.ParseInt(idStr,10,64)
	mod,err:=model.ArticleGet(id)
	if err!=nil{
		Fail(w,err.Error())
		return
	}
	buf,_:=json.Marshal(mod)
	w.Write(buf)
}
//列表页
//func IndexLoginData(c *gin.Context){
//	username:=c.PostForm("username")
//	password:=c.PostForm("password")
//	if model.IndexLogin(username,password){
//		c.JSON(200,gin.H{"恭喜":"登录成功"})
//	}else{
//		c.JSON(400,gin.H{"false":"密码或用户名错误"})
//	}
//
//}
//注册数据获取
func Registe(gin *gin.Context) bool {
	username := gin.PostForm("username")
	password := gin.PostForm("password")
	registeRes := model.Registe(username, password)
	return registeRes
}

//登录数据获取
func Login(gin *gin.Context) bool {
	username := gin.PostForm("username")
	password := gin.PostForm("password")
	gin.SetCookie("username", username, 100, "/", "localhost", false, true)
	loginRes := model.Login(username, password)
	return loginRes
}

//列表页数据
func ListData(w http.ResponseWriter,r *http.Request){
	mods,err:=model.ArticleList()
	if err!=nil{
		Fail(w,err.Error())
		return
	}
	Succ(w,"列表",mods)
}

func ListDel(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	//w.Header().Set("Content-Type","application/json")
	idStr:=r.Form.Get("id")
	id,_:=strconv.ParseInt(idStr,10,64)
	if model.ArticleDel(id){
		Succ(w,"删除成功")
		return
	}
	Fail(w,"删除失败")
	return
}

func Fail(w http.ResponseWriter,msg string,data ...interface{}){
	mod:=Reply{
		Code: 300,
		Msg:  msg,
		//Data: data,
	}
	if len(data)>0{
		mod.Data=data[0]
	}
	buf,_:=json.Marshal(mod)
	w.Header().Set("Content-Type","application/json")
	w.Write(buf)

}

func Succ(w http.ResponseWriter,msg string,data ...interface{}){
	mod:=Reply{
		Code: 200,
		Msg:  msg,
		//Data: data,
	}
	if len(data)>0{
		mod.Data=data[0]
	}
	buf,_:=json.Marshal(mod)
	w.Header().Set("Content-Type","application/json")
	w.Write(buf)

}

type Reply struct {
	Code int  `json:"code"`//作为标识码 200 成功 300 失败
	Msg string	`json:"msg"`	//给用户提示
	Data interface{}	`json:"data"`//返回数据
}

