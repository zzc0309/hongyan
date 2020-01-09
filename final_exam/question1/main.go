
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main(){
	db:=Conn()
	var random1 int
	var random2 int
for i:=0;i<1000;i++{
	random1=random_num()
	time.Sleep(1000000 * time.Microsecond)
	random2=random_num()
	if select_price(random1,random2){
		break
	}
}
	rows,err:= db.Query("select name,price from question_1 where id=? or id=?",random2,random1)
	defer rows.Close()
	if err!=nil{
		fmt.Println("SQL is wrong:",err)
	}
	var name string
	var price  int
	for rows.Next(){
		err:=rows.Scan(&name,&price)
		if err!=nil{
			fmt.Println("rows.Scan is wrong:",err)
		}
		fmt.Println(name,price,"元")
	}

}