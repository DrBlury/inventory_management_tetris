// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all inventories
	// (GET /api/inventories)
	GetAllInventories(w http.ResponseWriter, r *http.Request)
	// Add new inventory
	// (POST /api/inventories)
	AddInventory(w http.ResponseWriter, r *http.Request)
	// Delete an inventory
	// (DELETE /api/inventories/{inventoryId})
	DeleteInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64)
	// Get inventory by ID
	// (GET /api/inventories/{inventoryId})
	GetInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64)
	// Update an inventory
	// (PUT /api/inventories/{inventoryId})
	UpdateInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64)
	// Add item to inventory (with optional position)
	// (POST /api/inventories/{inventoryId}/add)
	AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64)
	// Move item in Inventory
	// (POST /api/inventories/{inventoryId}/move)
	MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64)
	// Get all items
	// (GET /api/items)
	GetAllItems(w http.ResponseWriter, r *http.Request)
	// Add a new item
	// (POST /api/items)
	AddItem(w http.ResponseWriter, r *http.Request)
	// Delete an item
	// (DELETE /api/items/{itemId})
	DeleteItemById(w http.ResponseWriter, r *http.Request, itemId int64)
	// Get item
	// (GET /api/items/{itemId})
	GetItemById(w http.ResponseWriter, r *http.Request, itemId int64)
	// Update an item
	// (PUT /api/items/{itemId})
	UpdateItemById(w http.ResponseWriter, r *http.Request, itemId int64)
	// Get all users
	// (GET /api/users)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	// Add a new user
	// (POST /api/users)
	AddUser(w http.ResponseWriter, r *http.Request)
	// Delete a user
	// (DELETE /api/users/{userId})
	DeleteUserById(w http.ResponseWriter, r *http.Request, userId int64)
	// Get a user
	// (GET /api/users/{userId})
	GetUserById(w http.ResponseWriter, r *http.Request, userId int64)
	// Update a user
	// (PUT /api/users/{userId})
	UpdateUserById(w http.ResponseWriter, r *http.Request, userId int64)
	// Get metrics for the service
	// (GET /info/metrics)
	GetMetrics(w http.ResponseWriter, r *http.Request)
	// Get describing html of openapi spec
	// (GET /info/openapi.html)
	GetOpenAPIHTML(w http.ResponseWriter, r *http.Request, params GetOpenAPIHTMLParams)
	// Get openapi spec as json
	// (GET /info/openapi.json)
	GetOpenAPIJSON(w http.ResponseWriter, r *http.Request)
	// Get status of the service
	// (GET /info/status)
	GetStatus(w http.ResponseWriter, r *http.Request)
	// Get version info of the service
	// (GET /info/version)
	GetVersion(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get all inventories
// (GET /api/inventories)
func (_ Unimplemented) GetAllInventories(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new inventory
// (POST /api/inventories)
func (_ Unimplemented) AddInventory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete an inventory
// (DELETE /api/inventories/{inventoryId})
func (_ Unimplemented) DeleteInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get inventory by ID
// (GET /api/inventories/{inventoryId})
func (_ Unimplemented) GetInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an inventory
// (PUT /api/inventories/{inventoryId})
func (_ Unimplemented) UpdateInventoryById(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add item to inventory (with optional position)
// (POST /api/inventories/{inventoryId}/add)
func (_ Unimplemented) AddItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Move item in Inventory
// (POST /api/inventories/{inventoryId}/move)
func (_ Unimplemented) MoveItemInInventory(w http.ResponseWriter, r *http.Request, inventoryId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all items
// (GET /api/items)
func (_ Unimplemented) GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add a new item
// (POST /api/items)
func (_ Unimplemented) AddItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete an item
// (DELETE /api/items/{itemId})
func (_ Unimplemented) DeleteItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get item
// (GET /api/items/{itemId})
func (_ Unimplemented) GetItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an item
// (PUT /api/items/{itemId})
func (_ Unimplemented) UpdateItemById(w http.ResponseWriter, r *http.Request, itemId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all users
// (GET /api/users)
func (_ Unimplemented) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add a new user
// (POST /api/users)
func (_ Unimplemented) AddUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a user
// (DELETE /api/users/{userId})
func (_ Unimplemented) DeleteUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a user
// (GET /api/users/{userId})
func (_ Unimplemented) GetUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a user
// (PUT /api/users/{userId})
func (_ Unimplemented) UpdateUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get metrics for the service
// (GET /info/metrics)
func (_ Unimplemented) GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get describing html of openapi spec
// (GET /info/openapi.html)
func (_ Unimplemented) GetOpenAPIHTML(w http.ResponseWriter, r *http.Request, params GetOpenAPIHTMLParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get openapi spec as json
// (GET /info/openapi.json)
func (_ Unimplemented) GetOpenAPIJSON(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get status of the service
// (GET /info/status)
func (_ Unimplemented) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get version info of the service
// (GET /info/version)
func (_ Unimplemented) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetAllInventories operation middleware
func (siw *ServerInterfaceWrapper) GetAllInventories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllInventories(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddInventory operation middleware
func (siw *ServerInterfaceWrapper) AddInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddInventory(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteInventoryById operation middleware
func (siw *ServerInterfaceWrapper) DeleteInventoryById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "inventoryId" -------------
	var inventoryId int64

	err = runtime.BindStyledParameterWithOptions("simple", "inventoryId", chi.URLParam(r, "inventoryId"), &inventoryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inventoryId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteInventoryById(w, r, inventoryId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetInventoryById operation middleware
func (siw *ServerInterfaceWrapper) GetInventoryById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "inventoryId" -------------
	var inventoryId int64

	err = runtime.BindStyledParameterWithOptions("simple", "inventoryId", chi.URLParam(r, "inventoryId"), &inventoryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inventoryId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInventoryById(w, r, inventoryId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateInventoryById operation middleware
func (siw *ServerInterfaceWrapper) UpdateInventoryById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "inventoryId" -------------
	var inventoryId int64

	err = runtime.BindStyledParameterWithOptions("simple", "inventoryId", chi.URLParam(r, "inventoryId"), &inventoryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inventoryId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateInventoryById(w, r, inventoryId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddItemInInventory operation middleware
func (siw *ServerInterfaceWrapper) AddItemInInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "inventoryId" -------------
	var inventoryId int64

	err = runtime.BindStyledParameterWithOptions("simple", "inventoryId", chi.URLParam(r, "inventoryId"), &inventoryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inventoryId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddItemInInventory(w, r, inventoryId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// MoveItemInInventory operation middleware
func (siw *ServerInterfaceWrapper) MoveItemInInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "inventoryId" -------------
	var inventoryId int64

	err = runtime.BindStyledParameterWithOptions("simple", "inventoryId", chi.URLParam(r, "inventoryId"), &inventoryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inventoryId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.MoveItemInInventory(w, r, inventoryId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAllItems operation middleware
func (siw *ServerInterfaceWrapper) GetAllItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllItems(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddItem operation middleware
func (siw *ServerInterfaceWrapper) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddItem(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteItemById operation middleware
func (siw *ServerInterfaceWrapper) DeleteItemById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "itemId" -------------
	var itemId int64

	err = runtime.BindStyledParameterWithOptions("simple", "itemId", chi.URLParam(r, "itemId"), &itemId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteItemById(w, r, itemId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetItemById operation middleware
func (siw *ServerInterfaceWrapper) GetItemById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "itemId" -------------
	var itemId int64

	err = runtime.BindStyledParameterWithOptions("simple", "itemId", chi.URLParam(r, "itemId"), &itemId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItemById(w, r, itemId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateItemById operation middleware
func (siw *ServerInterfaceWrapper) UpdateItemById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "itemId" -------------
	var itemId int64

	err = runtime.BindStyledParameterWithOptions("simple", "itemId", chi.URLParam(r, "itemId"), &itemId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateItemById(w, r, itemId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAllUsers operation middleware
func (siw *ServerInterfaceWrapper) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddUser operation middleware
func (siw *ServerInterfaceWrapper) AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUserById operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", chi.URLParam(r, "userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUserById(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserById operation middleware
func (siw *ServerInterfaceWrapper) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", chi.URLParam(r, "userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserById(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUserById operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", chi.URLParam(r, "userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUserById(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetMetrics operation middleware
func (siw *ServerInterfaceWrapper) GetMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMetrics(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetOpenAPIHTML operation middleware
func (siw *ServerInterfaceWrapper) GetOpenAPIHTML(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetOpenAPIHTMLParams

	// ------------- Optional query parameter "render" -------------

	err = runtime.BindQueryParameter("form", true, false, "render", r.URL.Query(), &params.Render)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "render", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetOpenAPIHTML(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetOpenAPIJSON operation middleware
func (siw *ServerInterfaceWrapper) GetOpenAPIJSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetOpenAPIJSON(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetStatus operation middleware
func (siw *ServerInterfaceWrapper) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetStatus(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetVersion operation middleware
func (siw *ServerInterfaceWrapper) GetVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVersion(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/inventories", wrapper.GetAllInventories)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/inventories", wrapper.AddInventory)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/inventories/{inventoryId}", wrapper.DeleteInventoryById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/inventories/{inventoryId}", wrapper.GetInventoryById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/inventories/{inventoryId}", wrapper.UpdateInventoryById)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/inventories/{inventoryId}/add", wrapper.AddItemInInventory)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/inventories/{inventoryId}/move", wrapper.MoveItemInInventory)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/items", wrapper.GetAllItems)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/items", wrapper.AddItem)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/items/{itemId}", wrapper.DeleteItemById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/items/{itemId}", wrapper.GetItemById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/items/{itemId}", wrapper.UpdateItemById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/users", wrapper.GetAllUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/users", wrapper.AddUser)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/users/{userId}", wrapper.DeleteUserById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/users/{userId}", wrapper.GetUserById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/users/{userId}", wrapper.UpdateUserById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/info/metrics", wrapper.GetMetrics)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/info/openapi.html", wrapper.GetOpenAPIHTML)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/info/openapi.json", wrapper.GetOpenAPIJSON)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/info/status", wrapper.GetStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/info/version", wrapper.GetVersion)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"//+QEhIsUJgAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
