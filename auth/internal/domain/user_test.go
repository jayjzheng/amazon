package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValid(t *testing.T) {
	tests := map[string]func(*testing.T){
		"emptyLogin":    testUserValid_emptyLogin,
		"emptyPassword": testUserValid_emptyPassword,
		"ok":            testUserValid_ok,
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}

func TestUserAuthenticate(t *testing.T) {
	tests := map[string]func(*testing.T){
		"wrongPassword": testUserAuthenticate_wrongPassword,
		"ok":            testUserAuthenticate_ok,
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}

func testUserValid_emptyLogin(t *testing.T) {
	u := User{
		Login:    "  ",
		Password: "password",
	}

	assert.False(t, u.Valid())
}

func testUserValid_emptyPassword(t *testing.T) {
	u := User{
		Login:    "login",
		Password: "  ",
	}

	assert.False(t, u.Valid())
}

func testUserValid_ok(t *testing.T) {
	u := User{
		Login:    "login",
		Password: "password",
	}

	assert.True(t, u.Valid())
}

func testUserAuthenticate_wrongPassword(t *testing.T) {
	u := User{
		Login:    "login",
		Password: "password",
	}

	assert.False(t, u.Authenticate("wrongPassword"))
}

func testUserAuthenticate_ok(t *testing.T) {
	u := User{
		Login:    "login",
		Password: "password",
	}

	assert.True(t, u.Authenticate("password"))
}
