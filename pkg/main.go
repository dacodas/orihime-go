package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/subcommands"
)

var db, err = sql.Open("mysql", "root:dacodastrackoda@tcp(127.0.0.1:3306)/orihime")

func checkDatabase() {
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func main() {
	checkDatabase()
}
