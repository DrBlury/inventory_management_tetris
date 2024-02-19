// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

// Defines values for ItemType.
const (
	Armor            ItemType = "armor"
	Consumable       ItemType = "consumable"
	ConsumableWeapon ItemType = "consumableWeapon"
	MeleeWeapon      ItemType = "meleeWeapon"
	Quest            ItemType = "quest"
	RangedWeapon     ItemType = "rangedWeapon"
	Resource         ItemType = "resource"
)

// AddItemToInventoryRequest defines model for AddItemToInventoryRequest.
type AddItemToInventoryRequest struct {
	DurabilityLeft float32            `json:"durability_left"`
	Item           Item               `json:"item"`
	Position       *InventoryPosition `json:"position,omitempty"`
	Quantity       int                `json:"quantity"`
}

// Inventory defines model for Inventory.
type Inventory struct {
	Id        string  `json:"id"`
	MaxWeight float32 `json:"max_weight"`
	Name      string  `json:"name"`
	User      User    `json:"user"`
	Volume    Volume  `json:"volume"`
}

// InventoryItem defines model for InventoryItem.
type InventoryItem struct {
	DurabilityLeft float32           `json:"durability_left"`
	Item           Item              `json:"item"`
	Position       InventoryPosition `json:"position"`
	Quantity       int               `json:"quantity"`
}

// InventoryListResponse defines model for InventoryListResponse.
type InventoryListResponse struct {
	Inventories []Inventory `json:"inventories"`
	Pagination  Pagination  `json:"pagination"`
}

// InventoryPosition defines model for InventoryPosition.
type InventoryPosition struct {
	Rotation int `json:"rotation"`
	X        int `json:"x"`
	Y        int `json:"y"`
}

// InventoryPostRequest defines model for InventoryPostRequest.
type InventoryPostRequest struct {
	MaxWeight float32 `json:"max_weight"`
	Name      string  `json:"name"`
	User      User    `json:"user"`
	Volume    Volume  `json:"volume"`
}

// Item defines model for Item.
type Item struct {
	BuyValue    float32   `json:"buy_value"`
	Description string    `json:"description"`
	Durability  float32   `json:"durability"`
	Id          string    `json:"id"`
	MaxStack    int       `json:"max_stack"`
	Name        string    `json:"name"`
	SellValue   float32   `json:"sell_value"`
	Shape       ItemShape `json:"shape"`
	Type        ItemType  `json:"type"`
	Variant     string    `json:"variant"`
	Weight      float32   `json:"weight"`
}

// ItemListResponse defines model for ItemListResponse.
type ItemListResponse struct {
	Items      []Item     `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// ItemPostRequest defines model for ItemPostRequest.
type ItemPostRequest struct {
	BuyValue    float32   `json:"buy_value"`
	Description string    `json:"description"`
	Durability  float32   `json:"durability"`
	MaxStack    int       `json:"max_stack"`
	Name        string    `json:"name"`
	SellValue   float32   `json:"sell_value"`
	Shape       ItemShape `json:"shape"`
	Type        ItemType  `json:"type"`
	Variant     string    `json:"variant"`
	Weight      float32   `json:"weight"`
}

// ItemShape defines model for ItemShape.
type ItemShape struct {
	Rawshape string `json:"rawshape"`
	SizeH    int    `json:"size_h"`
	SizeV    int    `json:"size_v"`
}

// ItemType defines model for ItemType.
type ItemType string

// MoveItemRequest defines model for MoveItemRequest.
type MoveItemRequest struct {
	NewPosition      InventoryPosition `json:"new_position"`
	OriginalPosition InventoryPosition `json:"original_position"`
	Quantity         int               `json:"quantity"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
	Total  int `json:"total"`
}

// User defines model for User.
type User struct {
	Id          string      `json:"id"`
	Inventories []Inventory `json:"inventories"`
	Username    string      `json:"username"`
}

// UserListResponse defines model for UserListResponse.
type UserListResponse struct {
	Pagination Pagination `json:"pagination"`
	Users      []User     `json:"users"`
}

// UserPostRequest defines model for UserPostRequest.
type UserPostRequest struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Volume defines model for Volume.
type Volume struct {
	SizeH float32 `json:"size_h"`
	SizeV float32 `json:"size_v"`
}

// AddInventoryJSONRequestBody defines body for AddInventory for application/json ContentType.
type AddInventoryJSONRequestBody = InventoryPostRequest

// AddItemInInventoryJSONRequestBody defines body for AddItemInInventory for application/json ContentType.
type AddItemInInventoryJSONRequestBody = AddItemToInventoryRequest

// MoveItemInInventoryJSONRequestBody defines body for MoveItemInInventory for application/json ContentType.
type MoveItemInInventoryJSONRequestBody = MoveItemRequest

// AddItemJSONRequestBody defines body for AddItem for application/json ContentType.
type AddItemJSONRequestBody = ItemPostRequest

// UpdateItemByIdJSONRequestBody defines body for UpdateItemById for application/json ContentType.
type UpdateItemByIdJSONRequestBody = Item

// AddUserJSONRequestBody defines body for AddUser for application/json ContentType.
type AddUserJSONRequestBody = UserPostRequest

// UpdateUserByIdJSONRequestBody defines body for UpdateUserById for application/json ContentType.
type UpdateUserByIdJSONRequestBody = User
