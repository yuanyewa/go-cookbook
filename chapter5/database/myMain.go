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

func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@/gocookbook?parseTime=true",
			os.Getenv("MYSQLUSERNAME"),
			os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create(db *sql.DB) error {
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}
	return nil
}

func Query(db *sql.DB) error {
	name := "Aaron"
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("result:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}

func Exec(db *sql.DB) error {
	defer db.Exec("DROP TABLE example")
	if err := Create(db); err != nil {
		return err
	}
	if err := Query(db); err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := Setup()
	if err != nil {
		panic(err)
	}
	if err := Exec(db); err != nil {
		panic(err)
	}
}
