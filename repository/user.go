// Package repository handles data storage and retrieval for user entities.
package repository

import (
	"context"

	"github.com/fevunge/let-go/entity"
	"github.com/fevunge/let-go/prisma/db"
	"github.com/fevunge/let-go/service/database"
)

var prisma = database.PrismaClient()

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
