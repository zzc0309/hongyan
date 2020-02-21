package control

import (
	"io"
	"net/http"
	"os"
)

func ListView(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./views/list.html")
	io.Copy(w,f)
	f.Close()
}

func EditView(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./views/edit.html")
	io.Copy(w,f)
	f.Close()
}

func DetailView(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./views/detail.html")
	io.Copy(w,f)
	f.Close()
}
//查询首页
func IndexView(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./views/index.html")
	io.Copy(w,f)
	f.Close()
}
//B战首页
func BIndexView(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./views/bilibili.index.html")
	io.Copy(w,f)
	f.Close()
}

//添加页面
func ViewArticleAdd(w http.ResponseWriter,r *http.Request){
	f,_:=os.Open("./views/add.html")
	io.Copy(w,f)
	f.Close()
}
