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

// Get all items
// (GET /api/items)
func (a APIHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	// call domain layer
	items, err := a.AppLogic.GetAllItems(r.Context())
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// TODO map domain model to dto
	// itemsDTO := make([]server.Item, 0, len(*items))
	// for _, item := range *items {
	// 	itemsDTO = append(itemsDTO, transform.ItemDTOfromDomain(&item))
	// }

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, items)
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

	// TODO map dto to domain model for createItemParams
	var createItemParams = domain.CreateItemParams{}

	// call domain layer
	/* //TODO maybe add function to add multiple items at
	 * //TODO once (...) and return list of items that could not be added
	 */
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
	// TODO fix this response in API spec
	w.WriteHeader(http.StatusNoContent)
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

	// TODO map domain model to dto
	// itemDTO := transform.ItemDTOFromDomain(item)

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, item)
}

// Update an item
// (PUT /api/items/{itemId})
func (a APIHandler) UpdateItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	// TODO use custom model for domain.UpdateItemParams,
	// TODO do not use domain.CreateItemParams
}
