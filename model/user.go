// Package model contains data models for the application.
package model

import (
	"fmt"

	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name, username, password string) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("Hashed password:", string(hashedPassword))
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}
	user := entity.User{name, username, string(hashedPassword), "", []entity.Link{}}
	repository.SaveUser(user)
	return user, err
}
