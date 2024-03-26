// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

import (
	"time"
)

// Defines values for ErrorType.
const (
	BadRequest          ErrorType = "Bad Request"
	InternalServerError ErrorType = "Internal Server Error"
	NotFound            ErrorType = "Not Found"
	NotImplemented      ErrorType = "Not Implemented"
)

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

// Defines values for Status.
const (
	HEALTHY   Status = "HEALTHY"
	UNHEALTHY Status = "UNHEALTHY"
	UNKNOWN   Status = "UNKNOWN"
)

// AddItemToInventoryRequest defines model for AddItemToInventoryRequest.
type AddItemToInventoryRequest struct {
	DurabilityLeft int       `json:"durability_left"`
	Item           Item      `json:"item"`
	Position       *Position `json:"position,omitempty"`
	Quantity       int       `json:"quantity"`
}

// Error defines model for Error.
type Error struct {
	// Code The error code
	Code int `json:"code"`

	// Error The error message
	Error string `json:"error"`

	// ErrorId The unique identifier for the error
	ErrorId string `json:"errorId"`

	// ErrorType The error type
	ErrorType ErrorType `json:"errorType"`

	// Timestamp The time the error occurred
	Timestamp time.Time `json:"timestamp"`
}

// ErrorType The error type
type ErrorType string

// Inventory defines model for Inventory.
type Inventory struct {
	InventoryMeta InventoryMeta   `json:"inventoryMeta"`
	Items         []InventoryItem `json:"items"`
}

// InventoryItem defines model for InventoryItem.
type InventoryItem struct {
	DurabilityLeft int      `json:"durability_left"`
	Item           Item     `json:"item"`
	Position       Position `json:"position"`
	Quantity       int      `json:"quantity"`
}

// InventoryListResponse defines model for InventoryListResponse.
type InventoryListResponse struct {
	Inventories []Inventory `json:"inventories"`
	Pagination  Pagination  `json:"pagination"`
}

// InventoryMeta defines model for InventoryMeta.
type InventoryMeta struct {
	Id        int    `json:"id"`
	MaxWeight int    `json:"max_weight"`
	Name      string `json:"name"`
	UserId    int    `json:"userId"`
	Volume    Volume `json:"volume"`
}

// InventoryPostRequest defines model for InventoryPostRequest.
type InventoryPostRequest struct {
	Id        *string `json:"id,omitempty"`
	MaxWeight int     `json:"max_weight"`
	Name      string  `json:"name"`
	UserId    int     `json:"user_id"`
	Volume    Volume  `json:"volume"`
}

// Item defines model for Item.
type Item struct {
	BuyValue    int       `json:"buy_value"`
	Description string    `json:"description"`
	Durability  int       `json:"durability"`
	Id          int       `json:"id"`
	MaxStack    int       `json:"max_stack"`
	Name        string    `json:"name"`
	SellValue   int       `json:"sell_value"`
	Shape       ItemShape `json:"shape"`
	Type        ItemType  `json:"type"`
	Variant     string    `json:"variant"`
	Weight      int       `json:"weight"`
}

// ItemListResponse defines model for ItemListResponse.
type ItemListResponse struct {
	Items      []Item     `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// ItemPostRequest defines model for ItemPostRequest.
type ItemPostRequest struct {
	BuyValue    int       `json:"buy_value"`
	Description string    `json:"description"`
	Durability  int       `json:"durability"`
	MaxStack    int       `json:"max_stack"`
	Name        string    `json:"name"`
	SellValue   int       `json:"sell_value"`
	Shape       ItemShape `json:"shape"`
	Type        ItemType  `json:"type"`
	Variant     string    `json:"variant"`
	Weight      int       `json:"weight"`
}

// ItemShape defines model for ItemShape.
type ItemShape struct {
	Height   int    `json:"height"`
	Rawshape string `json:"rawshape"`
	Width    int    `json:"width"`
}

// ItemType defines model for ItemType.
type ItemType string

// MoveItemRequest defines model for MoveItemRequest.
type MoveItemRequest struct {
	NewPosition      Position `json:"new_position"`
	OriginalPosition Position `json:"original_position"`
	Quantity         int      `json:"quantity"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
	Total  int `json:"total"`
}

// Position defines model for Position.
type Position struct {
	Rotation int `json:"rotation"`
	X        int `json:"x"`
	Y        int `json:"y"`
}

// Status The status of the API
type Status string

// User defines model for User.
type User struct {
	Id          int             `json:"id"`
	Inventories []InventoryMeta `json:"inventories"`
	Username    string          `json:"username"`
}

// UserListResponse defines model for UserListResponse.
type UserListResponse struct {
	Pagination Pagination `json:"pagination"`
	Users      []User     `json:"users"`
}

// UserPostRequest defines model for UserPostRequest.
type UserPostRequest struct {
	Email    string  `json:"email"`
	Id       *string `json:"id,omitempty"`
	Password string  `json:"password"`
	Username string  `json:"username"`
}

// Version defines model for Version.
type Version struct {
	// BuildDate The date the code was built
	BuildDate string `json:"buildDate"`

	// CommitDate The date of the commit
	CommitDate string `json:"commitDate"`

	// CommitHash The hash of the commit
	CommitHash string `json:"commitHash"`

	// Description A description of the API
	Description string `json:"description"`

	// Version The version of the API
	Version string `json:"version"`
}

// Volume defines model for Volume.
type Volume struct {
	Height int `json:"height"`
	Width  int `json:"width"`
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
