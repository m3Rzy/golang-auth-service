package utils

import (
	"database/sql"
	_ "github.com/lib/pq" // подключение PostgreSQL
	"log"
)

var DataBase *sql.DB // соединение с БД

func InitDB() {
	var err error
	DataBase, err = sql.Open("postgres", "postgres://postgres:123321@localhost/golang-auth?sslmode=disable")
	if err != nil {
		panic(err)
	}
	log.Println("Успешное подключение к базе данных!")

	//defer DataBase.Close()
}
