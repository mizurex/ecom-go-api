package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	dsn := "root:asdf1434@tcp(127.0.0.1:3306)/college_project"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Failed To Connect", err)

	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database Not Reachable", err)
	}

	fmt.Println("Connected To Database")
	DB = db
	return db
}
