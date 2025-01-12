package service

import (
	"errors"
	"goland-auth-service/internal/repository"
	"regexp"
	"strings"
)

var userRepository repository.UserRepository = &repository.UserRepositoryImpl{}

// AddUser добавляет пользователя с валидацией данных
func AddUser(email string, password string) error {
	// Регулярное выражение для проверки email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Проверка на пустую почту
	if strings.TrimSpace(email) == "" {
		return errors.New("почта не может быть пустой")
	}

	re := regexp.MustCompile(emailRegex)

	// Проверка соответствия регулярному выражению
	if !re.MatchString(email) {
		return errors.New("некорректный формат почты")
	}

	// Проверка на пустой пароль
	if strings.TrimSpace(password) == "" {
		return errors.New("пароль не может быть пустым")
	}

	// Минимальная длина пароля
	if len(password) < 8 {
		return errors.New("пароль должен быть не менее 8 символов")
	}

	// Попытка добавить пользователя в базу
	err := userRepository.AddUser(email, password)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
