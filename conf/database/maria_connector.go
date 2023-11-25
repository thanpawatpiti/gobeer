package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MariaDB *sql.DB

// InitMariaDB is a function that connect to MariaDB
func InitMariaDB() error {
	// Connect to MariaDB
	db, err := sql.Open("mysql", "test:#123456789$@tcp(mariadb-155554-0.cloudclusters.net:19700)/gobeer")
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	MariaDB = db

	fmt.Println("Connected to MariaDB!")

	return nil
}
