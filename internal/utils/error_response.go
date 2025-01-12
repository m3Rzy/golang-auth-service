package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse представляет структуру для JSON ответа с ошибкой.
type ErrorResponse struct {
	Message string `json:"message"`
}

// ErrorHandler - функция для обработки ошибок.
func ErrorHandler(w http.ResponseWriter, status int, message string) {
	// Создаем объект ошибки.
	errorResponse := ErrorResponse{
		Message: message,
	}

	// Преобразуем объект в JSON.
	errorJSON, err := json.Marshal(errorResponse)
	if err != nil {
		// Если ошибка при сериализации, логируем её и возвращаем 500.
		log.Printf("Ошибка сериализации JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Логируем ошибку в консоль.
	log.Printf("Ошибка: [Message: %s StatusCode: %d]", message, status)

	// Устанавливаем заголовок и статус HTTP.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Отправляем JSON-ответ.
	w.Write(errorJSON)
}