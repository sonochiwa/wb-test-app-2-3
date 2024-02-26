package repository

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	conn, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	// Проверяем подключение к базе данных
	if err := DB.Ping(); err != nil {
		DB.Close()
		log.Fatal(err)
	}
}
