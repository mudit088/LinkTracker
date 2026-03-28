package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var connStr string

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL != "" {
		// Railway / Production
		connStr = databaseURL + "?sslmode=require"
	} else {
		// Local DB (fallback)
		connStr = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB ❌", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB ❌", err)
	}

	DB = db
	log.Println("Database connected ✅")
}
}