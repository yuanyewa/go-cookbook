package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Example struct {
	Name    string
	Created *time.Time
}

func main() {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)")
	db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`)
	db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`)
	rows, _ := db.Query("SELECT name, created FROM example where name=?", "Aaron")
	for rows.Next() {
		var e Example
		rows.Scan(&e.Name, &e.Created)
		fmt.Printf("result:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	rows.Close()
	db.Exec("DROP TABLE example")
}
