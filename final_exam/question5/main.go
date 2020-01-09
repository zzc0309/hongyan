package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
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
engine:=gin.Default()
engine.POST("/search_xuehao", func(c *gin.Context) {
	var name string
	Stuid:=c.DefaultQuery("Stuid","2019214104")
	url:="http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh="+Stuid
	resp, err:= http.Get(url) //访问页面
	if err!=nil{fmt.Println("http.Get is wrong:",err)}
	defer resp.Body.Close()  //关闭文件
	dom,err2:=goquery.NewDocumentFromReader(resp.Body)
	if err2!=nil{fmt.Println("goquery.New is wrong:",err2)}
	dom.Find("body").Find("#head").Find("div").Find("ul").Find("li~li").Each(func(i int, selection *goquery.Selection) {
		name= selection.Text()
	})
	name = SubstrByByte(name, 53)
	c.JSON(200,gin.H{"xuehao":Stuid,"name":name})
})
engine.Run()
}
