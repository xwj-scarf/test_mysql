package main 
import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB
var err error
func main() {

	db,err = sql.Open("mysql","root:080349@tcp(127.0.0.1:3306)/oj?charset=utf8")
	if err != nil {
		return
	}
	defer db.Close()
	insert()
	update()
	query()
	
}

func query() {
	rows, _ := db.Query("select * from submit_status");
	defer rows.Close()	
	for rows.Next() {
		var id int
		var pid int
		var status int
		
		if err := rows.Scan(&id,&pid,&status);err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
		fmt.Println(pid)
	}
	if err := rows.Err();err != nil {
		fmt.Println(err)
		return
	}	
}

func insert() {
	stmt, err := db.Prepare(`insert into submit_status (pid) values(?)`)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return 
	}
	res,err := stmt.Exec("1")
	if err != nil {
		fmt.Println(err)
		return 
	}

	id, _ := res.LastInsertId()
	fmt.Println(id)	
}

func update() {
	stmt,err := db.Prepare(`update submit_status set status=? where id = ?`)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}

	res,err := stmt.Exec("1","1")
	if err != nil {
		fmt.Println(err)
	}

	num,_ := res.RowsAffected()
	fmt.Println(num)
}

