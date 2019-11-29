//用的13001806班的数据
//实现的功能:
//查询某选修课,13001806班上有多少同学选,这些同学是哪些人.
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func selectDB(db *sql.DB,s1 string) {
	rows, err := db.Query("SELECT student.stu_name '学生名称'," +
		"course.course '课程名称' FROM student,course, stu_course" +
		" WHERE student.sid=stu_course.sid" +
		" AND course.cid=stu_course.cid" +
		" AND course.course=" +
		"'"+s1+"'")   //查询语法
	if err != nil {
		log.Fatal(err)
	}
	var count int   //记录人次
	defer rows.Close()
	fmt.Println("name      ","course")
	for rows.Next() {
		var course string
		var name string
		//var name sql.NullString
		err := rows.Scan(&course, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(course,"    ",name)
		count++
	}
	fmt.Println("一共有",count,"人选修")
}
func main(){
	db, err := sql.Open("mysql", "root:687211@tcp(127.0.0.1:3306)/lv3")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("请输入一门选修课")
	var s1 string
	fmt.Scanf("%s",&s1)
	selectDB(db,s1)
}

