package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "application"
	db_password = "application127"
	db_address  = "192.168.2.140"
	db_db       = "test"
)

type (
	Person struct {
		Id       int
		Name     string
		Age      int
		Location string
	}
)

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s3306)/%s", db_user, db_password, db_address, db_db)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)

	// close it safely
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

}
