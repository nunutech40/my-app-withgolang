package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func ConnectDb() *sql.DB {
	if db != nil {
		return db // Return the existing connection if already established
	}

	var errMySql error
	db, errMySql = sql.Open("mysql", "root:myadmin12345678@tcp(localhost:3306)/myappdb")
	if errMySql != nil {
		log.Fatalf("Error opening database: %v", errMySql)
	}

	// check connection to DB
	errDb := db.Ping()
	if errDb != nil {
		log.Fatal("Error connection to db: ", errDb)
	}
	log.Println("Connected to the database successfully")
	return db
}
