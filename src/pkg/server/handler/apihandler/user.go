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

// Get all users
// (GET /api/users)
func (a APIHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// call domain layer
	users, err := a.AppLogic.GetAllUsers(r.Context())
	if err != nil {

		handler.HandleInternalServerError(w, r, err)
		return
	}

	// TODO map domain model to dto
	// usersDTO := make([]server.User, 0, len(*users))
	// for _, user := range *users {
	// 	usersDTO = append(usersDTO, transform.UserDTOFromDomain(&user))
	// }

	// return response
	render.JSON(w, r, users)
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
		Username: dtoUser.Username,
		Email:    dtoUser.Email,
		Password: dtoUser.Password,
	}

	// call domain layer
	addedUser, err := a.AppLogic.AddUser(r.Context(), createUserParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		zap.L().Error("error adding user", zap.Error(err))
		return
	}

	// TODO map to dto

	// return response
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, addedUser)
}

// Delete user by ID
// (DELETE /api/users/{userId})
func (a APIHandler) DeleteUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	// call domain layer
	err := a.AppLogic.DeleteUserById(r.Context(), int(userId))
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// return response
	// TODO Also change api response to return new status code
	w.WriteHeader(http.StatusNoContent)
}

// Get user by ID
// (GET /api/users/{userId})
func (a APIHandler) GetUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	// call domain layer
	user, err := a.AppLogic.GetUserById(r.Context(), int(userId))
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		return
	}

	// TODO: map domain model to dto
	// userDTO := transform.UserDTOFromDomain(user)

	// return response
	render.JSON(w, r, user)
}

// Update an user
// (PUT /api/users/{userId})
func (a APIHandler) UpdateUserById(w http.ResponseWriter, r *http.Request, userId int64) {
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

	updateUserParams := domain.UpdateUserParams{
		Username: dtoUser.Username,
		Email:    dtoUser.Email,
		Password: dtoUser.Password,
	}

	// call domain layer
	updatedUser, err := a.AppLogic.UpdateUser(r.Context(), int(userId), updateUserParams)
	if err != nil {
		handler.HandleInternalServerError(w, r, err)
		// log error
		zap.L().Error("error updating user", zap.Error(err))
		return
	}

	// TODO map to dto

	// return response
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, updatedUser)
}
