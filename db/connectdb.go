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

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
