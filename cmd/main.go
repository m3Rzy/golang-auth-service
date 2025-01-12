package main

import (
	"fmt"
	"goland-auth-service/internal/handlers"
	_ "goland-auth-service/internal/handlers"
	"goland-auth-service/internal/utils"
	"log"
	"net/http"
)

const PORT string = ":8080"

func main() {
	// подключение к БД
	utils.InitDB()

	// регистрация хендлеров
	registerHandlers()

	// запуск сервера
	fmt.Printf("Сервер запущен http://localhost%s\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln("Ошибка при запуске сервера:", err)
	}
	
	defer utils.DataBase.Close()
}

func registerHandlers() {
	http.HandleFunc("/register", handlers.RegisterHandler)
}
