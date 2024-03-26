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
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, inventories)
}

// Add new inventory
// (POST /api/inventories)
func (a APIHandler) AddInventory(w http.ResponseWriter, r *http.Request) {
	a.log.Info("adding inventory: ", zap.String("request", r.RequestURI))
	// read dto inventory from request using unmarshal
	var dtoInventory server.InventoryPostRequest
	// read request body into bytes
	bodyBytes := make([]byte, r.ContentLength)

	// log request body
	a.log.Error("request body", zap.String("body", string(bodyBytes)))

	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		a.log.Error("error reading request body", zap.Error(err))
		return
	}
	// unmarshal bytes into dto
	err = json.Unmarshal(bodyBytes, &dtoInventory)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		a.log.Error("error unmarshalling request body", zap.Error(err))
		return
	}

	createInventoryParams := domain.CreateInventoryParams{
		UserID:    dtoInventory.UserId,
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
		a.log.Error("error adding inventory", zap.Error(err))
		return
	}

	// TODO map to dto

	// return response
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, addedInventory)
}

// Delete inventory by ID
// (DELETE /api/inventories/{inventoryId})
func (a APIHandler) DeleteInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	// call domain layer
	err := a.AppLogic.DeleteInventoryById(r.Context(), int(inventoryId))
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
	inventory, err := a.AppLogic.GetInventoryById(r.Context(), int(inventoryId))
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// map domain model to dto

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, inventory)
}

// Add a new item to the inventory at the first possible position
// (POST /api/inventories/{inventoryId}/add)
func (a APIHandler) AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
}

// Move an item already in the inventory
// (POST /api/inventories/{inventoryId}/move)
func (a APIHandler) MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
}
