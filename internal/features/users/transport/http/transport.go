package users_transport_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

type UserService interface {
	CreateUser(ctx context.Context, fullName string, phoneNumber *string) (domain.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	PatchUser(ctx context.Context, id uuid.UUID, patch domain.UserPatch) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UsersHTTPHandler struct {
	service UserService
}

func NewUsersHTTPHandler(service UserService) *UsersHTTPHandler {
	return &UsersHTTPHandler{service: service}
}

func getUUIDPathVar(r *http.Request, key string) (uuid.UUID, error) {
	vars := mux.Vars(r)
	return uuid.Parse(vars[key])
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, core_errors.ErrConflict):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, core_errors.ErrInvalidArgument):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
