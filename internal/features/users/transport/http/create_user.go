package users_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

type UserDTOResponse struct {
	ID          uuid.UUID `json:"id"`
	Version     int       `json:"version"`
	FullName    string    `json:"full_name"`
	PhoneNumber *string   `json:"phone_number"`
}

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userToDTO(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(r.Context(), req.FullName, req.PhoneNumber)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, userToDTO(user))
}
