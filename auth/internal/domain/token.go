package domain

type TokenGenerator interface {
	GenerateToken(u *User) (string, error)
}
