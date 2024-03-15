package domain

type User struct {
	ID           int
	Username     string
	Inventories  []Inventory
	Salt         string
	PasswordHash string
	Email        string
}
