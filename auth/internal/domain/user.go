package domain

import "strings"

type User struct {
	Login    string
	Password string
}

func (u User) Valid() bool {
	return strings.TrimSpace(u.Login) != "" &&
		strings.TrimSpace(u.Password) != ""
}

func (u User) Authenticate(pass string) bool {
	return u.Password == pass
}

type UserStore interface {
	CreateUser(u *User) error
	FindUser(login string) (*User, error)
	UpdateUser(u *User) error
}
