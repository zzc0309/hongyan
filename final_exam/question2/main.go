package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main(){
	data,err:=ioutil.ReadFile("Students.txt")
	if err!=nil{
		fmt.Println("open file is wrong:",err)
	}
	reg:=regexp.MustCompile(`"xh":"\d{7}573".*?{`)
fmt.Println(reg.FindAllString(string(data),-1))
}
