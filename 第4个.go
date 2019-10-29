package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
	func main(){

		f, err := os.Create("proverb.txt")
		check(err)
		f.Write([]byte("don't communicate by sharing memory share memory by communicating"))

		data2, err := ioutil.ReadFile("proverb.txt")
		if err != nil {
			fmt.Printf("ioutil read file err : %v\n", err)
		}
		fmt.Printf("ioutil read file success.\n内容：\n%s\n", string(data2))

				defer f.Close()
	}
