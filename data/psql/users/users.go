package users

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	port     = 5432
	username = "postgres"
	password = "Go-microservices"
	database = "users"
)

var (
	DB *sql.DB
)

func init() {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
	log.Println("Database connection successful")
}
