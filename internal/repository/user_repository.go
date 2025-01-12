package repository

import (
	"errors"
	"goland-auth-service/internal/utils"
)

type UserRepository interface {
	AddUser(email string, password string) error
}

type UserRepositoryImpl struct{}

func (UserRepositoryImpl) AddUser(email string, password string) error {
	_, err := utils.DataBase.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
