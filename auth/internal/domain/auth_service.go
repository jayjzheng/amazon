package domain

import (
	"fmt"
)

type AuthService struct {
	UserStore      UserStore
	TokenGenerator TokenGenerator
	Publisher      Publisher
}

type Publisher interface {
	PublishUserCreated(*User) error
	PublishTokenCreated(*User) error
	PublishPasswordChanged(*User) error
}

func (s *AuthService) CreateUser(u *User) error {
	if !u.Valid() {
		return ErrInvalidUser
	}

	if err := s.UserStore.CreateUser(u); err != nil {
		return fmt.Errorf("CreateUser %s: %w", u.Login, err)
	}

	if err := s.Publisher.PublishUserCreated(u); err != nil {
		return fmt.Errorf("PublishUserCreated %s: %w", u.Login, err)
	}

	return nil
}

func (s *AuthService) ChangePassword(login, old, new string) error {
	u, err := s.UserStore.FindUser(login)
	if err != nil {
		return fmt.Errorf("FindUser %s: %w", login, err)
	}

	if u == nil || !u.Authenticate(old) {
		return ErrUnauthenticated
	}

	newUser := User{
		Login:    u.Login,
		Password: new,
	}
	if !newUser.Valid() {
		return ErrInvalidPassword
	}

	if err := s.UserStore.UpdateUser(&newUser); err != nil {
		return fmt.Errorf("UpdateUser %s: %w", newUser.Login, err)
	}

	if err := s.Publisher.PublishPasswordChanged(&newUser); err != nil {
		return fmt.Errorf("PublishPasswordChanged %s: %w", newUser.Login, err)
	}

	return nil
}

func (s *AuthService) CreateToken(login, password string) (string, error) {
	u, err := s.UserStore.FindUser(login)
	if err != nil {
		return "", fmt.Errorf("FindUser %s: %w", login, err)
	}

	if u == nil {
		return "", ErrUserNotFound
	}

	if !u.Authenticate(password) {
		return "", ErrUnauthenticated
	}

	token, err := s.TokenGenerator.GenerateToken(u)
	if err != nil {
		return "", fmt.Errorf("GenerateToken %s: %w", u.Login, err)
	}

	if err := s.Publisher.PublishTokenCreated(u); err != nil {
		return "", fmt.Errorf("PublishTokenCreated %s: %w", u.Login, err)
	}

	return token, nil
}
