package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

// goland wont like this import because it is currently "unused" without additonal context
// import "github.com/go-sql-driver/mysql"
var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		// this user requires mysql native password authentication
		User:   "derp",
		Passwd: "perd",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
