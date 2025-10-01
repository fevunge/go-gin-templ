package entity

type UserToCreate struct {
	Username string
	Name     string
	password string
}

type User struct {
	Username  string
	Name      string
	Password  string
	AvatarUlr string
	Links     []Link
}
