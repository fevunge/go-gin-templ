// Package model contains data models for the application.
package model

import (
	"errors"

	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name, username, password string) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		name,
		username,
		string(hashedPassword),
		"",
		[]entity.Link{},
	}
	repository.SaveUser(user)
	return user, nil
}

func GetUserByUsername(username string) (*entity.User, error) {
	return nil, errors.New("not implemented")
}
