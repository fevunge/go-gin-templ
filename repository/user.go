// Package repository handles data storage and retrieval for user entities.
package repository

import (
	"context"

	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/prisma/db"
	"github.com/fevunge/let-go/service/database"
)

var prisma = database.PrismaClient()

func FindUserByUsername(username string) (entity.User, error) {
	ctxt := context.Background()
	userModel, err := prisma.User.FindUnique(
		db.User.Username.Equals(username),
	).Exec(ctxt)
	if err != nil {
		return entity.User{}, err
	}
	if userModel == nil {
		return entity.User{}, nil
	}
	user := entity.User{
		Name:     userModel.Name,
		Username: userModel.Username,
		Password: userModel.Password,
	}
	return user, nil
}

func FindUserByName(name string) (entity.User, error) {
	ctxt := context.Background()
	userModel, err := prisma.User.FindUnique(
		db.User.Name.Equals(name),
	).Exec(ctxt)
	if err != nil {
		return entity.User{}, err
	}
	if userModel == nil {
		return entity.User{}, nil
	}
	user := entity.User{
		Name:     userModel.Name,
		Username: userModel.Username,
		Password: userModel.Password,
	}
	return user, nil
}

func SaveUser(user entity.User) (*db.UserModel, error) {
	ctxt := context.Background()

	created, err := prisma.User.CreateOne(
		db.User.Username.Set(user.Username),
		db.User.Password.Set(user.Password),
		db.User.Name.Set(user.Name),
	).Exec(ctxt)

	if created != nil {
		return created, err
	}
	return nil, err
}
