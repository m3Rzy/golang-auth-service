package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // подключение PostgreSQL
)

var DataBase *sql.DB // соединение с БД

func InitDB() {
	var err error
	DataBase, err = sql.Open("postgres", "postgres://postgres:123321@localhost/golang_auth?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	pingErr := DataBase.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	} else {
		log.Println("Успешное подключение к базе данных!")
	}
}
