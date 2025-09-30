// Package database implements the database connection and operations using Prisma ORM.
package database

import "github.com/fevunge/let-go/prisma/db"

func PrismaClient() *db.PrismaClient {
	client := db.NewClient()

	if err := client.Connect(); err != nil {
		panic(err)
	}

	return client
}
