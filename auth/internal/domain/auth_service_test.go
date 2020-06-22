package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthServiceCreateUser(t *testing.T) {
	tests := map[string]func(*testing.T){
		"error": testAuthServiceCreateUser_error,
		"ok":    testAuthServiceCreateUser_ok,
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}

func TestAuthServiceChangePassword(t *testing.T) {
	tests := map[string]func(*testing.T){
		"findUserError":    testAuthServiceChangePassword_findUserError,
		"notAuthenticated": testAuthServiceChangePassword_notAuthenticated,
		"updateUserError":  testAuthServiceChangePassword_updateUserError,
		"ok":               testAuthServiceChangePassword_ok,
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}

func testAuthServiceCreateUser_error(t *testing.T) {
	someErr := errors.New("some error")
	var invoked bool

	m := mockUserStore{
		CreateUserFn: func(*User) error {
			invoked = true
			return someErr
		},
	}

	s := AuthService{UserStore: &m}
	err := s.CreateUser(&User{})

	assert.True(t, invoked)
	assert.Equal(t, someErr, err)
}

func testAuthServiceCreateUser_ok(t *testing.T) {
	var invoked bool

	m := mockUserStore{
		CreateUserFn: func(*User) error {
			invoked = true
			return nil
		},
	}

	s := AuthService{UserStore: &m}
	err := s.CreateUser(&User{})

	assert.True(t, invoked)
	assert.Nil(t, err)
}

func testAuthServiceChangePassword_findUserError(t *testing.T) {
	someErr := errors.New("some error")

	m := mockUserStore{
		FindUserFn: func(string) (*User, error) {
			return nil, someErr
		},
	}

	s := AuthService{UserStore: &m}
	err := s.ChangePassword("login", "old", "new")

	assert.Contains(t, err.Error(), "FindUser")
	assert.True(t, errors.Is(err, someErr))
}

func testAuthServiceChangePassword_notAuthenticated(t *testing.T) {
	m := mockUserStore{
		FindUserFn: func(string) (*User, error) {
			return &User{
				Login:    "login",
				Password: "password",
			}, nil
		},
	}

	s := AuthService{UserStore: &m}
	err := s.ChangePassword("login", "old", "new")

	assert.Equal(t, ErrUnauthenticated, err)
}

func testAuthServiceChangePassword_updateUserError(t *testing.T) {
	someErr := errors.New("some error")

	m := mockUserStore{
		FindUserFn: func(string) (*User, error) {
			return &User{
				Login:    "login",
				Password: "password",
			}, nil
		},
		UpdateUserFn: func(*User) error {
			return someErr
		},
	}

	s := AuthService{UserStore: &m}
	err := s.ChangePassword("login", "password", "new")

	assert.Contains(t, err.Error(), "UpdateUser")
	assert.True(t, errors.Is(err, someErr))
}

func testAuthServiceChangePassword_ok(t *testing.T) {
	u := User{
		Login:    "login",
		Password: "old",
	}

	m := mockUserStore{
		FindUserFn: func(string) (*User, error) {
			return &u, nil
		},
		UpdateUserFn: func(*User) error {
			return nil
		},
	}

	s := AuthService{UserStore: &m}
	err := s.ChangePassword("login", "old", "new")

	assert.Nil(t, err)
	assert.Equal(t, "new", u.Password)
}

type mockUserStore struct {
	CreateUserFn func(u *User) error
	FindUserFn   func(login string) (*User, error)
	UpdateUserFn func(u *User) error
}

func (m *mockUserStore) CreateUser(u *User) error {
	return m.CreateUserFn(u)
}

func (m *mockUserStore) FindUser(login string) (*User, error) {
	return m.FindUserFn(login)
}

func (m *mockUserStore) UpdateUser(u *User) error {
	return m.UpdateUserFn(u)
}
