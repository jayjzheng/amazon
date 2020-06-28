package domain

import "errors"

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidUser     = errors.New("invalid user")
	ErrInvalidPassword = errors.New("invalid password")
)
