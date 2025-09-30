// Package model contains data models for the application.
package model

type User struct {
	username  string
	name      string
	password  string
	avatarUlr string
	links     []Link
}
