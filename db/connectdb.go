package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connectdb() *sql.DB {

	godotenv.Load(".env.local")
	UNAMEDB := os.Getenv("DB_USERNAME")
	PASSDB := os.Getenv("DB_PASSWORD")
	HOSTDB := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	DB_PARAMS := os.Getenv("DB_PARAMS")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", UNAMEDB, PASSDB, HOSTDB, DBNAME, DB_PARAMS)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
