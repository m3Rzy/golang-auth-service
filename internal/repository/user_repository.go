package repository

import (
	"errors"
	"goland-auth-service/internal/utils"
)

/*
интерфейс сделан, так как функция сохранения пользователя может быть другой
*/
type UserRepository interface {
	AddUser(email string, password string) error
}

type UserRepositoryImpl struct{}

func (UserRepositoryImpl) AddUser(email string, password string) error {
	_, err := utils.DataBase.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)
	if err != nil {
		//log.Fatalf("Ошибка при отправке запроса на регистрацию в базу данных: %s", err)
		return errors.New("ошибка при отправке запроса на регистрацию в базу данных")
	}
	return nil
}
