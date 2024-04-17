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

// AddItemToInventoryRequest This object holds the request data for adding an item to the inventory.
type AddItemToInventoryRequest struct {
	DurabilityLeft int `json:"durability_left"`

	// Item This object holds the item data.
	Item Item `json:"item"`

	// Position This object holds the position data.
	Position *Position `json:"position,omitempty"`
	Quantity int       `json:"quantity"`
}

// Error This object holds the error response data.
type Error struct {
	// ErrorType The error type
	ErrorType ErrorType `json:"ErrorType"`

	// Code The error code
	Code int `json:"code"`

	// Error The error message
	Error string `json:"error"`

	// ErrorId The unique identifier for the error
	ErrorId string `json:"errorId"`

	// Timestamp The time the error occurred
	Timestamp time.Time `json:"timestamp"`
}

// ErrorType The error type
type ErrorType string

// Inventory This object holds the inventory meta data and the items in the inventory as an array of InventoryItem objects.
type Inventory struct {
	// InventoryMeta This object holds the inventory meta data (restrictions, owner, etc.)
	InventoryMeta InventoryMeta `json:"inventoryMeta"`

	// Items An array of InventoryItem objects.
	Items []InventoryItem `json:"items"`
}

// InventoryItem This object holds the inventory item data.
type InventoryItem struct {
	DurabilityLeft int `json:"durability_left"`

	// Item This object holds the item data.
	Item Item `json:"item"`

	// Position This object holds the position data.
	Position Position `json:"position"`
	Quantity int      `json:"quantity"`
}

// InventoryListResponse This object holds the inventory list response data. It includes the list of inventories and pagination data.
type InventoryListResponse struct {
	// Inventories The list of inventories.
	Inventories []Inventory `json:"inventories"`

	// Pagination This object holds the pagination data.
	Pagination Pagination `json:"pagination"`
}

// InventoryMeta This object holds the inventory meta data (restrictions, owner, etc.)
type InventoryMeta struct {
	Id        int    `json:"id"`
	MaxWeight int    `json:"max_weight"`
	Name      string `json:"name"`
	UserId    int    `json:"userId"`

	// Volume This object holds the volume data.
	Volume Volume `json:"volume"`
}

// InventoryPostRequest This object holds the request data for creating an inventory.
type InventoryPostRequest struct {
	MaxWeight int    `json:"max_weight"`
	Name      string `json:"name"`
	UserId    int    `json:"user_id"`

	// Volume This object holds the volume data.
	Volume Volume `json:"volume"`
}

// Item This object holds the item data.
type Item struct {
	BuyValue    int    `json:"buy_value"`
	Description string `json:"description"`
	Durability  int    `json:"durability"`
	Id          int    `json:"id"`
	MaxStack    int    `json:"max_stack"`
	Name        string `json:"name"`
	SellValue   int    `json:"sell_value"`

	// Shape This object holds the item shape data.
	Shape ItemShape `json:"shape"`

	// Type The type of the item.
	Type    ItemType `json:"type"`
	Variant string   `json:"variant"`
	Weight  int      `json:"weight"`
}

// ItemListResponse This object holds the item list response data. It includes the list of items and pagination data.
type ItemListResponse struct {
	// Items The list of items.
	Items []Item `json:"items"`

	// Pagination This object holds the pagination data.
	Pagination Pagination `json:"pagination"`
}

// ItemPostRequest This object holds the request data for creating a new item.
type ItemPostRequest struct {
	BuyValue    int    `json:"buy_value"`
	Description string `json:"description"`
	Durability  int    `json:"durability"`
	MaxStack    int    `json:"max_stack"`
	Name        string `json:"name"`
	SellValue   int    `json:"sell_value"`

	// Shape This object holds the item shape data.
	Shape ItemShape `json:"shape"`

	// Type The type of the item.
	Type    ItemType `json:"type"`
	Variant string   `json:"variant"`
	Weight  int      `json:"weight"`
}

// ItemShape This object holds the item shape data.
type ItemShape struct {
	Height   int    `json:"height"`
	Rawshape string `json:"rawshape"`
	Width    int    `json:"width"`
}

// ItemType The type of the item.
type ItemType string

// MoveItemRequest This object holds the request data for moving an item within the inventory.
type MoveItemRequest struct {
	// NewPosition This object holds the position data.
	NewPosition Position `json:"new_position"`

	// OriginalPosition This object holds the position data.
	OriginalPosition Position `json:"original_position"`
	Quantity         int      `json:"quantity"`
}

// Pagination This object holds the pagination data.
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
	Total  int `json:"total"`
}

// Position This object holds the position data.
type Position struct {
	Rotation int `json:"rotation"`
	X        int `json:"x"`
	Y        int `json:"y"`
}

// Status The status of the API
type Status string

// User This object holds the user data.
type User struct {
	Id          int             `json:"id"`
	Inventories []InventoryMeta `json:"inventories"`
	Username    string          `json:"username"`
}

// UserListResponse This object holds the user list response data. It includes the list of users and pagination data.
type UserListResponse struct {
	// Pagination This object holds the pagination data.
	Pagination Pagination `json:"pagination"`

	// Users The list of users.
	Users []User `json:"users"`
}

// UserPostRequest This object holds the request data for creating a new user.
type UserPostRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// Version This object holds the API version data.
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

// Volume This object holds the volume data.
type Volume struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

// GetOpenAPIHTMLParams defines parameters for GetOpenAPIHTML.
type GetOpenAPIHTMLParams struct {
	// Render OpenAPI html renderer
	Render *string `form:"render,omitempty" json:"render,omitempty"`
}

// AddInventoryJSONRequestBody defines body for AddInventory for application/json ContentType.
type AddInventoryJSONRequestBody = InventoryPostRequest

// UpdateInventoryByIdJSONRequestBody defines body for UpdateInventoryById for application/json ContentType.
type UpdateInventoryByIdJSONRequestBody = InventoryPostRequest

// AddItemInInventoryJSONRequestBody defines body for AddItemInInventory for application/json ContentType.
type AddItemInInventoryJSONRequestBody = AddItemToInventoryRequest

// MoveItemInInventoryJSONRequestBody defines body for MoveItemInInventory for application/json ContentType.
type MoveItemInInventoryJSONRequestBody = MoveItemRequest

// AddItemJSONRequestBody defines body for AddItem for application/json ContentType.
type AddItemJSONRequestBody = ItemPostRequest

// UpdateItemByIdJSONRequestBody defines body for UpdateItemById for application/json ContentType.
type UpdateItemByIdJSONRequestBody = ItemPostRequest

// AddUserJSONRequestBody defines body for AddUser for application/json ContentType.
type AddUserJSONRequestBody = UserPostRequest

// UpdateUserByIdJSONRequestBody defines body for UpdateUserById for application/json ContentType.
type UpdateUserByIdJSONRequestBody = UserPostRequest
