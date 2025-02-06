package dbconnection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// This function will initialize the database connection.
	// This function will be called from the main.go file.

	connectString := " host=localhost port=5432 user=postgres password=admin dbname=ordertracker sslmode=disable"
	// Open a database connection
	DB, err := sql.Open("postgres", connectString)

	if err != nil {
		// If there is an error opening the connection, handle it
		panic(err)
	}

	// Set the maximum number of open connections to 10
	DB.SetMaxOpenConns(10)
	// Set the maximum number of idle connections to 5
	DB.SetMaxIdleConns(5)
}
