package apihandler

import "net/http"

type APIHandler struct{}

// NewAPIHandler returns a new *APIHandler
func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

// Get all inventories
// (GET /inventories)
func (a APIHandler) GetAllInventories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllInventories"))
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new inventory
// (POST /inventories)
func (a APIHandler) AddInventory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete inventory by ID
// (DELETE /inventories/{inventoryId})
func (a APIHandler) DeleteInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get inventory by ID
// (GET /inventories/{inventoryId})
func (a APIHandler) GetInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add a new item to the inventory at the first possible position
// (POST /inventories/{inventoryId}/add)
func (a APIHandler) AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add a new item to the inventory at a specific position
// (POST /inventories/{inventoryId}/addAtPosition)
func (a APIHandler) AddItemInInventoryAtPosition(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Move an item already in the inventory
// (POST /inventories/{inventoryId}/move)
func (a APIHandler) MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all items
// (GET /items)
func (a APIHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new item
// (POST /items)
func (a APIHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete item by ID
// (DELETE /items/{itemId})
func (a APIHandler) DeleteItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get item by ID
// (GET /items/{itemId})
func (a APIHandler) GetItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an item
// (PUT /items/{itemId})
func (a APIHandler) UpdateItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all users
// (GET /users)
func (a APIHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new user
// (POST /users)
func (a APIHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete user by ID
// (DELETE /users/{userId})
func (a APIHandler) DeleteUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get user by ID
// (GET /users/{userId})
func (a APIHandler) GetUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an user
// (PUT /users/{userId})
func (a APIHandler) UpdateUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
