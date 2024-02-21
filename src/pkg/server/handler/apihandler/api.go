package apihandler

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
	"net/http"

	transform "linuxcode/inventory_manager/pkg/server/transform"

	"github.com/go-chi/render"
)
type APIHandler struct{
	AppLogic domain.AppLogic
}

func NewAPIHandler(appLogic domain.AppLogic) *APIHandler {
	return &APIHandler{
		AppLogic: appLogic,
	}
}

// Get all inventories
// (GET /api/inventories)
func (a APIHandler) GetAllInventories(w http.ResponseWriter, r *http.Request) {
	// call domain layer
	inventories, err := a.AppLogic.GetAllInventories()
	if err != nil {
		// handle error
	}

	// map domain model to dto
	inventoriesDTO := make([]server.Inventory, 0, len(inventories))
	for _, inv := range inventories {
		inventoriesDTO = append(inventoriesDTO, transform.DTOInventoryFromDomain(inv))
	}

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, inventoriesDTO)
}

// Add new inventory
// (POST /api/inventories)
func (a APIHandler) AddInventory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete inventory by ID
// (DELETE /api/inventories/{inventoryId})
func (a APIHandler) DeleteInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get inventory by ID
// (GET /api/inventories/{inventoryId})
func (a APIHandler) GetInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add a new item to the inventory at the first possible position
// (POST /api/inventories/{inventoryId}/add)
func (a APIHandler) AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Move an item already in the inventory
// (POST /api/inventories/{inventoryId}/move)
func (a APIHandler) MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all items
// (GET /api/items)
func (a APIHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new item
// (POST /api/items)
func (a APIHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete item by ID
// (DELETE /api/items/{itemId})
func (a APIHandler) DeleteItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get item by ID
// (GET /api/items/{itemId})
func (a APIHandler) GetItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an item
// (PUT /api/items/{itemId})
func (a APIHandler) UpdateItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all users
// (GET /api/users)
func (a APIHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new user
// (POST /api/users)
func (a APIHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete user by ID
// (DELETE /api/users/{userId})
func (a APIHandler) DeleteUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get user by ID
// (GET /api/users/{userId})
func (a APIHandler) GetUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an user
// (PUT /api/users/{userId})
func (a APIHandler) UpdateUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get metrics
// (GET /info/metrics)
func (a APIHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get status
// (GET /info/status)
func (a APIHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get version
// (GET /info/version)
func (a APIHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
