// Package model contains data models for the application.
package model

import (
	"github.com/labstack/gommon/log"

	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name, username, password string) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Debug("Hashed password:", string(hashedPassword))
	if err != nil {
		log.Debug("error hashing password:", err)
	}
	user := entity.User{name, username, string(hashedPassword), "", []entity.Link{}}
	repository.SaveUser(user)
	return user, err
}

func GetProfile(username string) (entity.User, error) {
	user, err := repository.FindUserByUsername(username)
	if err != nil {
		log.Debug("error finding user by username:", err)
	}
	return user, err
}