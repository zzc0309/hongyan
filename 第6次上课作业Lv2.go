package main

import (
	"database/sql"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)
func insertDB(db *sql.DB,s string)  {
	stmt, err := db.Prepare("insert into stuinfo(name1) values(?)")
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec(s)
}

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
	db, err := sql.Open("mysql", "root:687211@tcp(127.0.0.1:3306)/test1")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	//声明和创建一个slice1来装同学的名字
	slice1:=make([]string,29)
	//运用For循环访问每个学生的相应页面
	for  i:=0 ;i<=28;i++ {
		var b string//接受网页爬取的字符串
		var a uint64//定义一个uint型变量a
		a=2019214104+uint64(i)
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
		//装进slice1
		slice1[i]=b
	}
//装到表CLASS6
	for i:=0;i<=27;i++{
		insertDB(db,slice1[i])
	}
}