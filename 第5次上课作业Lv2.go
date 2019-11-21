package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"全新文件夹/github.com/PuerkitoBio/goquery"
)
func SubstrByByte(str string, length int) string {
	bs := []byte(str)[length:]
	bl := 0
	for i:=len(bs)-1; i>=0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++;
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i] & 252 == 252:
				cl = 6
			case bs[i] & 248 == 248:
				cl = 5
			case bs[i] & 240 == 240:
				cl = 4
			case bs[i] & 224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}
func main(){
	//var Gender string
	var name string
	engin:=gin.Default()
	engin.GET("/search", func(context *gin.Context) {
		Stuid:=context.DefaultQuery("Stuid","2019214104")

		//在教务在线找名字
		url := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + Stuid
		resp, _ := http.Get(url) //访问页面
		defer resp.Body.Close()  //关闭文件
		dom, _ := goquery.NewDocumentFromReader(resp.Body)
		dom.Find("body").Find("#head").Find("div").Find("ul").Find("li~li").Each(func(i int, selection *goquery.Selection) {
			name= selection.Text()
		})
		name = SubstrByByte(name, 53)


		resp1:=Response{Code:200,Message:"success",Stuid:Stuid,Name:name}
		context.JSON(200,&resp1)
	})
	engin.Run()
}
type Response struct {
	Code int
	Message string
	Stuid string
	//Gender string
	Name string
}
