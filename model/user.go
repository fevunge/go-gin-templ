// Package model contains data models for the application.
package model

import (
	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/repository"
)

func CreateUser(name, username, password string) entity.User {
	user := entity.User{name, username, password, "", []entity.Link{}}
	repository.SaveUser(user)
	return user
}
