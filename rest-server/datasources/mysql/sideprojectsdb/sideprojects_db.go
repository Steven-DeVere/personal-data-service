package sideprojectsdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// driver for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Client represent articlesDB client
var (
	Client *sql.DB
)

func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	userName := os.Getenv("MYSQL_USER_NAME")
	password := os.Getenv("MYSQL_PASSWORD")
	hostName := os.Getenv("MYSQL_HOST_NAME")
	dbName := os.Getenv("MYSQL_SIDEPROJECTS_DB_NAME")

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		userName,
		password,
		hostName,
		dbName,
	)

	var dbSetupErr error
	Client, dbSetupErr = sql.Open("mysql", dataSourceName)
	if dbSetupErr != nil {
		panic(dbSetupErr)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("database successfully configured")
}
