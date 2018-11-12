package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:example@tcp(localhost:3306)/")
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}

	var rs string
	row := db.QueryRow(`SELECT VERSION()`)
	err = row.Scan(&rs)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	fmt.Printf("row = %s\n", rs)
	for {

	}
}
