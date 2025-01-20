package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			user, password, host, port, name)

		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatalf("Could not stabish connection to database: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Could not connect to database: %v", err)
		}
	})
	return db
}
