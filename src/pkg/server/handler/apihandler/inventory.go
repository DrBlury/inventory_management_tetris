package apihandler

import (
	"encoding/json"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	server "linuxcode/inventory_manager/pkg/server/generated"
	handler "linuxcode/inventory_manager/pkg/server/handler"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

// Get all inventories
// (GET /api/inventories)
func (a APIHandler) GetAllInventories(w http.ResponseWriter, r *http.Request) {
	// call domain layer
	inventories, err := a.AppLogic.GetAllInventories(r.Context())
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// map domain model to dto
	// inventoriesDTO := make([]server.Inventory, 0, len(inventories))
	// for _, inv := range inventories {
	// 	inventoriesDTO = append(inventoriesDTO, transform.DTOInventoryFromDomain(&inv))
	// }

	// return response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, inventories)
}

// Add new inventory
// (POST /api/inventories)
func (a APIHandler) AddInventory(w http.ResponseWriter, r *http.Request) {
	a.log.With(zap.String("request", r.RequestURI)).Info("adding inventory: ")
	// read dto inventory from request using unmarshal
	var dtoInventory server.InventoryPostRequest
	// read request body into bytes
	bodyBytes := make([]byte, r.ContentLength)

	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		a.log.With(zap.Error(err)).Error("error reading request body")
		return
	}
	// unmarshal bytes into dto
	err = json.Unmarshal(bodyBytes, &dtoInventory)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		a.log.With(zap.Error(err)).Error("error unmarshalling request body")
		return
	}

	createInventoryParams := &domain.CreateInventoryParams{
		UserId:    dtoInventory.UserId,
		Name:      dtoInventory.Name,
		MaxWeight: dtoInventory.MaxWeight,
		Width:     dtoInventory.Volume.Width,
		Height:    dtoInventory.Volume.Height,
	}

	// call domain layer
	addedInventory, err := a.AppLogic.AddInventory(r.Context(), createInventoryParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		a.log.With(zap.Error(err)).Error("error adding inventory")
		return
	}

	// TODO map to dto

	// return response
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, addedInventory)
}

// Delete inventory by ID
// (DELETE /api/inventories/{inventoryId})
func (a APIHandler) DeleteInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	// call domain layer
	err := a.AppLogic.DeleteInventoryById(r.Context(), inventoryId)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, struct{ Message string }{"Inventory deleted"})
}

// Get inventory by ID
// (GET /api/inventories/{inventoryId})
func (a APIHandler) GetInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	// call domain layer
	inventory, err := a.AppLogic.GetInventoryById(r.Context(), inventoryId)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// map domain model to dto

	// return response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, inventory)
}

// Update inventory by ID
// (PUT /api/inventories/{inventoryId})
func (a APIHandler) UpdateInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	// read dto inventory from request using unmarshal
	var dtoInventory server.InventoryPostRequest

	// read request body into bytes
	bodyBytes := make([]byte, r.ContentLength)
	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// unmarshal bytes into dto
	err = json.Unmarshal(bodyBytes, &dtoInventory)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	updateInventoryParams := &domain.UpdateInventoryParams{
		UserId:    dtoInventory.UserId,
		Name:      dtoInventory.Name,
		MaxWeight: dtoInventory.MaxWeight,
		Width:     dtoInventory.Volume.Width,
		Height:    dtoInventory.Volume.Height,
	}

	// call domain layer
	updatedInventory, err := a.AppLogic.UpdateInventory(r.Context(), inventoryId, updateInventoryParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// return response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, updatedInventory)
}

// Add a new item to the inventory at the first possible position
// (POST /api/inventories/{inventoryId}/add)
func (a APIHandler) AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
}

// Move an item already in the inventory
// (POST /api/inventories/{inventoryId}/move)
func (a APIHandler) MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
}
