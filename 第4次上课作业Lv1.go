package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//创建一个装输入时间戳的Map
	iotime:=make(map[int]string)
	//建立一个循环来向Map中输入值,输入的值为字符串,键值为int
	for i:=0;i<=1000;i++{
		var str string  //定义一个字符型变量来接受value的值
		fmt.Scanf("%s",&str)
		//判断输入是否结束
		if strings.EqualFold(str,"result"){
			fmt.Println("the result are:")
			break
		}else{
			iotime[i]=str
			fmt.Println("input ok!")
		}

}

	//由于时间戳转化为标准时间时,需要Int64类型,故下面的代码用于将Map中的字符串转换为Int64
	n:=len(iotime)//取切片的长度
iotime1:=make(map[int]int64)//创建一个新的Map来装Int64的值
//运用循环再向里面输入值
	for i:=0;i<n;i++{

	a,err:=strconv.ParseInt(iotime[i],10,64)
	if err!=nil{
		fmt.Println("有错")
	}
	iotime1[i]=a
	}
//最后打印出来
	for _,value:=range iotime1{
		i:=0
		i++
		fmt.Println(time.Unix(value,0))
	}
}