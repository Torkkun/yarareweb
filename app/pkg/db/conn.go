package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

// Driver名
const driverName = "postgres"

func init() {
	// ユーザ
	sql_url := os.Getenv("SQL_URL")

	var err error
	Conn, err = sql.Open(driverName, sql_url)
	if err != nil {
		log.Fatal(err)
	}

}
