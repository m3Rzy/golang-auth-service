package handlers

import (
	"encoding/json"
	"goland-auth-service/internal/models"
	"goland-auth-service/internal/service"
	"goland-auth-service/internal/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler POST-запрос на регистрацию
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if r.Method != http.MethodPost {
		utils.ErrorHandler(w, http.StatusNotFound, "Данный эндпоинт поддерживает только POST-запрос!")
		return
	}

	// Десериализация объекта из JSON
	if decodeErr := json.NewDecoder(r.Body).Decode(&user); decodeErr != nil {
		utils.ErrorHandler(w, http.StatusBadRequest, "Ошибка при десериализации объекта User!")
		return
	}

	// Хэширование пароля пользователя
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		utils.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при хэшировании пароля!")
		return
	}

	// Передача данных для валидации
	validateErr := service.AddUser(user.Email, string(hashedPassword))
	if validateErr != nil {
		utils.ErrorHandler(w, http.StatusBadRequest, validateErr.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь успешно добавлен"))
}
