package main

import "fmt"

func select_price(a int,b int)bool{
	db:=Conn()
	rows,err:= db.Query("select price from question_1 where id=? or id=?",a,b)
	defer rows.Close()
	if err!=nil{
		fmt.Println("SQL is wrong:",err)
	}
	var price  int
	for rows.Next(){
		var temp int
		err := rows.Scan(&temp)
		if err != nil {
			fmt.Println("rows.Scan is wrong:", err)
		}
		price += temp
	}
if price<=20{
	return true
}
return false
}