package domain

import (
	"fmt"
)

type AuthService struct {
	UserStore      UserStore
	TokenGenerator TokenGenerator
}

func (s *AuthService) CreateUser(u *User) error {
	return (s.UserStore.CreateUser(u))
}

func (s *AuthService) ChangePassword(login, old, new string) error {
	u, err := s.UserStore.FindUser(login)
	if err != nil {
		return fmt.Errorf("FindUser %s: %w", login, err)
	}

	if !u.Authenticate(old) {
		return ErrUnauthenticated
	}

	u.Password = new
	if err := s.UserStore.UpdateUser(u); err != nil {
		return fmt.Errorf("UpdateUser %s: %w", login, err)
	}

	return nil
}

func (s *AuthService) CreateToken(u *User) (string, error) {
	u, err := s.UserStore.FindUser(u.Login)
	if err != nil {
		return "", fmt.Errorf("FindUser %s: %w", u.Login, err)
	}

	if !u.Authenticate(u.Password) {
		return "", ErrUnauthenticated
	}

	token, err := s.TokenGenerator.GenerateToken(u)
	if err != nil {
		return "", fmt.Errorf("GenerateToken %s: %w", u.Login, err)
	}

	return token, nil
}
