package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test1")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	insertDB(db)//增
	deleteDB(db)//删
	updateDB(db)//改
	selectDB(db)//查
}



func deleteDB(db *sql.DB)  {
	stmt, err := db.Prepare("delete from stuinfo where name =? or name =?")
	if err != nil{
		log.Fatal(err);
	}
	stmt.Exec("rose","lilei");
}


func insertDB(db *sql.DB)  {
	stmt, err := db.Prepare("insert into stuinfo(id,name) values (?,?)")
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec(2019214106,"叶睿康")
	stmt.Exec(2019214107,"赵书立")
}

func updateDB(db *sql.DB)  {
	stmt, err := db.Prepare("UPDATE stuinfo SET name = ? WHERE name =? ")
	if err != nil{
		log.Fatal(err)
	}

	stmt.Exec("赵书立","12222");
}


func selectDB(db *sql.DB)  {
	rows, err := db.Query("select * from stuinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var id int
		var name string
		//var name sql.NullString
		err := rows.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name)
	}

}