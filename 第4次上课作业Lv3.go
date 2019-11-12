//爬的是2019级的

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"全新文件夹/github.com/PuerkitoBio/goquery"
)
//这个函数是把抓下来的字符串按字节分开,然后得到需要的东西
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
//主函数
func main(){
	//声明和创建一个Map来装新生的名字
	var mymap map[int]string
	mymap=make(map[int]string)
	//运用For循环访问每个学生的相应页面
	for  i:=0 ;i<=5000;i++ {
		var b string//接受网页爬取的字符串
		var a uint64//定义一个uint型变量a
		a=2019210001+uint64(i)
		str:=strconv.FormatUint(a,10)//将uint型变量转换为字符串型
		url := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" +str
		resp, _ := http.Get(url)//访问页面
		defer resp.Body.Close()//关闭文件
		dom, _ := goquery.NewDocumentFromReader(resp.Body)
		//查找筛选需要的信息
		dom.Find("body").Find("#head").Find("div").Find("ul").Find("li~li").Each(func(i int, selection *goquery.Selection) {
			str := selection.Text()
			b = str
		})
		b = SubstrByByte(b, 53)
		//装进map
		mymap[i]=b
		//fmt.Println(mymap[i])
	}
	//开始筛选了
	//定义定义一个人计数器和一个Map,利用Map的键值和键相对应的特点,把名字和重复的次数相对应
	var count int
	var mymap1=make(map[int]string)
	for ij:=0;ij<=5000;ij++{
		if ij==0 {
		}else if count>=5&&count<=20 {
			mymap1[count] = mymap[ij-1]
		}
		if count==5{break}
		count=0
		for ik:=0;ik<=5000;ik++{
			if   strings.EqualFold(mymap[ij],mymap[ik]){
				count++
			}
		}
	}
	//ok了
	for key,value:=range mymap1{
		fmt.Println(key," ",value)
	}
}