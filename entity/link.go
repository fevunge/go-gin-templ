// Package entity defines the Link struct used to represent a hyperlink with its URL and display text.
package entity

type Link struct {
	Title       string
	URL         string
	Description string
	IsPublic    bool
	Guests      []User
}
