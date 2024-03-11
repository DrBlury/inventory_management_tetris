package apihandler

import (
	"encoding/json"
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
	handler "linuxcode/inventory_manager/pkg/server/handler"
	transform "linuxcode/inventory_manager/pkg/server/transform"
	"net/http"

	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type APIHandler struct {
	AppLogic domain.AppLogic
	version  VersionInfo
	log      *zap.SugaredLogger
}

func NewAPIHandler(appLogic domain.AppLogic, versionInfo VersionInfo, logger *zap.SugaredLogger) *APIHandler {
	return &APIHandler{
		AppLogic: appLogic,
		version:  versionInfo,
		log:      logger,
	}
}

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
		Width:     dtoInventory.Volume.SizeH,
		Height:    dtoInventory.Volume.SizeV,
	}

	// call domain layer
	addedInventory, err := a.AppLogic.AddInventory(r.Context(), createInventoryParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		a.log.Error("error adding inventory", zap.Error(err))
		return
	}

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
	inventoryDTO := transform.DTOInventoryFromDomain(&inventory)

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, inventoryDTO)
}

// Add a new item to the inventory at the first possible position
// (POST /api/inventories/{inventoryId}/add)
func (a APIHandler) AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
}

// Move an item already in the inventory
// (POST /api/inventories/{inventoryId}/move)
func (a APIHandler) MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
}

// Get all items
// (GET /api/items)
func (a APIHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	// call domain layer
	items, err := a.AppLogic.GetAllItems(r.Context())
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// map domain model to dto
	itemsDTO := make([]server.Item, 0, len(items))
	for _, item := range items {
		itemsDTO = append(itemsDTO, transform.ItemDTOfromDomain(&item))
	}

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, itemsDTO)
}

// Add new item
// (POST /api/items)
func (a APIHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	// read dto item from request using unmarshal
	var dtoItem server.ItemPostRequest
	// read request body into bytes
	bodyBytes := make([]byte, r.ContentLength)

	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		zap.L().Error("error reading request body", zap.Error(err))
		return
	}
	// unmarshal bytes into dto
	err = json.Unmarshal(bodyBytes, &dtoItem)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		zap.L().Error("error unmarshalling request body", zap.Error(err))
		return
	}

	domainShape := transform.DomainShapeFromDTO(&dtoItem.Shape)

	domainType := transform.DomainItemTypeFromDTO(&dtoItem.Type)

	createItemParams := domain.CreateItemParams{
		Name:        dtoItem.Name,
		Variant:     dtoItem.Variant,
		Description: dtoItem.Description,
		BuyValue:    dtoItem.BuyValue,
		SellValue:   dtoItem.SellValue,
		Weight:      dtoItem.Weight,
		Durability:  dtoItem.Durability,
		MaxStack:    dtoItem.MaxStack,
		Type:        domainType,
		Shape:       domainShape,
	}

	// call domain layer
	err = a.AppLogic.AddItem(r.Context(), createItemParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// return response
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, struct{ Message string }{"Item added"})
}

// Delete item by ID
// (DELETE /api/items/{itemId})
func (a APIHandler) DeleteItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	// call domain layer
	err := a.AppLogic.DeleteItemById(r.Context(), int(itemId))
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, struct{ Message string }{"Item deleted"})
}

// Get item by ID
// (GET /api/items/{itemId})
func (a APIHandler) GetItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	// call domain layer
	item, err := a.AppLogic.GetItemById(r.Context(), int(itemId))
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// map domain model to dto
	itemDTO := transform.ItemDTOfromDomain(&item)

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, itemDTO)
}

// Update an item
// (PUT /api/items/{itemId})
func (a APIHandler) UpdateItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
}

// Get all users
// (GET /api/users)
func (a APIHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// call domain layer
	users, err := a.AppLogic.GetAllUsers(r.Context())
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// map domain model to dto
	usersDTO := make([]server.User, 0, len(users))
	for _, user := range users {
		usersDTO = append(usersDTO, transform.UserDTOFromDomain(&user))
	}

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, usersDTO)
}

// Add new user
// (POST /api/users)
func (a APIHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	// read dto user from request using unmarshal
	var dtoUser server.UserPostRequest
	// read request body into bytes
	bodyBytes := make([]byte, r.ContentLength)

	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		zap.L().Error("error reading request body", zap.Error(err))
		return
	}
	// unmarshal bytes into dto
	err = json.Unmarshal(bodyBytes, &dtoUser)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		zap.L().Error("error unmarshalling request body", zap.Error(err))
		return
	}

	createUserParams := domain.CreateUserParams{
		Username: dtoUser.Name,
		Email:    dtoUser.Email,
		Password: dtoUser.Password,
	}

	// call domain layer
	addedUser, err := a.AppLogic.AddUser(r.Context(), createUserParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		zap.L().Error("error adding user", zap.Error(err))
		return
	}

	// return response
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, addedUser)
}

// Delete user by ID
// (DELETE /api/users/{userId})
func (a APIHandler) DeleteUserById(w http.ResponseWriter, r *http.Request, userId int64) {
}

// Get user by ID
// (GET /api/users/{userId})
func (a APIHandler) GetUserById(w http.ResponseWriter, r *http.Request, userId int64) {
}

// Update an user
// (PUT /api/users/{userId})
func (a APIHandler) UpdateUserById(w http.ResponseWriter, r *http.Request, userId int64) {
}

// === Info ===

// Get metrics
// (GET /info/metrics)
func (a APIHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

// Get status
// (GET /info/status)
func (a APIHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, &struct{ Status string }{"HEALTHY"})
}

// Get version
// (GET /info/version)
func (a APIHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, a.version)
}
