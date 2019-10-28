package main

import "fmt"

// 声明一个map，用于存储每个数的阶乘
var myRes = make(map[int]int, 20)

func main() {
	chs:=make([]chan int,20)
	// 启用20个goroutine同时求1~20以内各个数的阶乘
	for i := 1; i <= 20; i++ {
		chs[i-1]=make(chan int)
		go factorial(chs[i-1],i)
	}
	for _,ch:=range(chs){
		<-ch
	}
	for v,i:=range myRes{
		fmt.Println(v," ",i)
	}
}

// 求n的阶乘，并将结果写进myRes
func factorial(ch chan int,n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myRes[n] = res
	ch<-1
}