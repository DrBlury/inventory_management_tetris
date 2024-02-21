package domain

type domain interface {
	// Inventories
	GetAllInventories() ([]*Inventory, error)
	AddInventory(inventory *Inventory) error
	DeleteInventoryById(inventoryId int) error
	GetInventoryById(inventoryId int) (*Inventory, error)
	AddItemInInventory(inventoryId int, item *Item) error
	AddItemInInventoryAtPosition(inventoryId int, item *Item, position Position) error
	DeleteItemFromInventory(inventoryId int, itemId int, position Position, amount int) error

	// Items
	GetAllItems() ([]*Item, error)
	AddItem(item *Item) error
	DeleteItemById(itemId int) error
	GetItemById(itemId int) (*Item, error)

	// Users
	GetUserById(userId int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	AddUser(user *User) error
	DeleteUserById(userId int) error
}
