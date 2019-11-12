


//输入学号获得课表
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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
func main() {
	var a uint64
	var b int
	var name string
	     kebiao:=make([]string,10)
	     classroom:=make([]string,10)
		course:=make([]string,10)
	//var zifu string
	fmt.Println("请输入你的学号")
	fmt.Scanf("%d", &a)
	str := strconv.FormatUint(a, 10) //将uint型变量转换为字符串型
	url := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + str
	resp, _ := http.Get(url) //访问页面
	defer resp.Body.Close()  //关闭文件

	dom, _ := goquery.NewDocumentFromReader(resp.Body)
	dom.Find("body").Find("#head").Find("div").Find("ul").Find("li~li").Each(func(i int, selection *goquery.Selection) {
		 name= selection.Text()
		})
	name = SubstrByByte(name, 53)

	t := time.Now()
	day:=int64(t.Weekday())
	dayweek:="星期"+strconv.FormatInt(day,10)
	switch dayweek {
	case "星期5":dom.Find("#kbStuTabs-list").Find("tbody").Find("td:contains(星期5)").Each(func(i int, selection *goquery.Selection) {
		kebiao[i]=selection.Text()
		course[i]=selection.Prev().Prev().Prev().Prev().Prev().Text()
		if len(course[i])==0{
			course[i]=course[i-1]
		}
		classroom[i]=selection.Next().Text()
		b=i})
	case "星期4":dom.Find("#kbStuTabs-list").Find("tbody").Find("td:contains(星期4)").Each(func(i int, selection *goquery.Selection) {
			kebiao[i]=selection.Text()
			course[i]=selection.Prev().Prev().Prev().Prev().Prev().Text()
		if len(course[i])==0{
			course[i]=course[i-1]
		}
			classroom[i]=selection.Next().Text()
			b=i})
	case "星期3":dom.Find("#kbStuTabs-list").Find("tbody").Find("td:contains(星期3)").Each(func(i int, selection *goquery.Selection) {
		kebiao[i]=selection.Text()
		course[i]=selection.Prev().Prev().Prev().Prev().Prev().Text()
		if len(course[i])==0{
			course[i]=course[i-1]
		}
		classroom[i]=selection.Next().Text()
		b=i})
	case "星期2":dom.Find("#kbStuTabs-list").Find("tbody").Find("td:contains(星期2)").Each(func(i int, selection *goquery.Selection) {
		kebiao[i]=selection.Text()
		course[i]=selection.Prev().Prev().Prev().Prev().Prev().Text()
		if len(course[i])==0{
			course[i]=course[i-1]
		}
		classroom[i]=selection.Next().Text()
		b=i})
	case "星期1":dom.Find("#kbStuTabs-list").Find("tbody").Find("td:contains(星期1)").Each(func(i int, selection *goquery.Selection) {
		kebiao[i]=selection.Text()
		course[i]=selection.Prev().Prev().Prev().Prev().Prev().Text()
		if len(course[i])==0{
			course[i]=course[i-1]
		}
		classroom[i]=selection.Next().Text()
		b=i})
	case "星期6":dom.Find("#kbStuTabs-list").Find("tbody").Find("td:contains(星期6)").Each(func(i int, selection *goquery.Selection) {
		kebiao[i]=selection.Text()
		course[i]=selection.Prev().Prev().Prev().Prev().Prev().Text()
		if len(course[i])==0{
			course[i]=course[i-1]
		}
		classroom[i]=selection.Next().Text()
		b=i})

	}
	total_classcourse:=make([]string,b+1)
	for i:=0;i<=b;i++{
		total_classcourse[i]=kebiao[i]+" "+course[i]+" "+classroom[i]
	}
		fmt.Println(name,"你今天的课表为:")
	for _,v:=range total_classcourse{
		fmt.Println(v)
	}
}