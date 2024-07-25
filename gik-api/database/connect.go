package database

import (
	"GIK_Web/env"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("sqlite3", env.SqliteURI)
	if err != nil {
		fmt.Println("Unable to connect to database: " + err.Error())
	} else {
		fmt.Printf("Connected to database: %s", env.SqliteURI)
	}
	Database = db
}
