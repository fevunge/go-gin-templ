// Package model contains data models for the application.
package model

type Link struct {
	title       string
	url         string
	description string
	isPublic    bool
	guests      []User
}
