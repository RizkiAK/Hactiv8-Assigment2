package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env data")
	}
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbdriver := os.Getenv("DBDRIVER")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbname := os.Getenv("DATABASE")
	DBURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)

	db, err := sql.Open(dbdriver, DBURL)

	if err != nil {
		log.Fatal("error connecting to database:", err.Error())
	}

	fmt.Println("Suksses terkoneksi ke database")

	return db
}
