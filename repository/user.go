// Package repository handles data storage and retrieval for user entities.
package repository

import (
	"context"
	"errors"

	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/prisma/db"
	"github.com/fevunge/let-go/service/database"
)

var (
	prisma = database.PrismaClient()
	ctxt   = context.Background()
)

func GetUserByUsername(username string) (*db.UserModel, error) {
	user, err := prisma.User.FindUnique(
		db.User.Username.Equals(username),
	).Exec(ctxt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validUsername(username string) bool {
	if user, err := GetUserByUsername(username); err == nil {
		if user != nil {
			return false
		}
		return true
	}
	return true
}

func SaveUser(user entity.User) (*db.UserModel, error) {
	if !validUsername(user.Username) {
		return nil, errors.New("invalid username")
	}
	created, err := prisma.User.CreateOne(
		db.User.Username.Set(user.Username),
		db.User.Password.Set(user.Password),
		db.User.Name.Set(user.Name),
	).Exec(ctxt)

	if created != nil {
		return created, nil
	}
	return nil, err
}
