package domain

type User struct {
	ID           int
	Username     string
	Inventories  []InventoryMeta
	Salt         string
	PasswordHash string
	Email        string
}
