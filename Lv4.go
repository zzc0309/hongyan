package main

import (
	"fmt"
)
/*func qiushushu(n int) {

	for i := 3 + n - 100; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if j == i {
				fmt.Print(i, " ")
			}
			if (i % j) == 0 {
				break
			}
		}
	}
}*/
func Count(ch chan int,n int){
	for i := 3 + n - 100; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if j == i {
				fmt.Print(i, " ")
			}
			if (i % j) == 0 {
				break
			}
		}
	}
	ch<-1
	}
func main() {
	chs := make([]chan int, 100)
	for k := 0; k < 100; k++ {
		chs[k] = make(chan int)
		//go qiushushu(100 * (k + 1))
		go Count(chs[k],100 * (k + 1))
		}
		for _,ch:=range(chs){
			<-ch
		}
	}

