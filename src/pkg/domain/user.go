package domain

type User struct {
	ID       int
	Username string
	Salt    string
	PasswordHash string
	Email    string
}
