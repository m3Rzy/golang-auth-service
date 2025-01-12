package handlers

import (
	"encoding/json"
	"fmt"
	"goland-auth-service/internal/models"
	"goland-auth-service/internal/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// RegisterHandler POST-запрос на регистрацию
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if r.Method != http.MethodPost {
		http.Error(w, "Данный эндпоинт поддерживает только POST-запрос!", http.StatusNotFound)
		return
	}

	// Десериализация объекта из JSON
	if decodeErr := json.NewDecoder(r.Body).Decode(&user); decodeErr != nil {
		http.Error(w, "Ошибка при десериализации объекта User", http.StatusBadRequest)
		return
	}

	// Хэширование пароля пользователя
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		http.Error(w, "Ошибка при хэшировании пароля", http.StatusInternalServerError)
		return
	}

	// Передача данных для валидации
	validateErr := service.AddUser(user.Email, string(hashedPassword))
	if validateErr != nil {
		http.Error(w, fmt.Sprintf("Ошибка: %s", validateErr), http.StatusBadRequest)
		return
	}

	// Успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь успешно добавлен"))
}
