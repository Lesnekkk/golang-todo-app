package users_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_http_types "github.com/Lesnekkk/golang-todo-app/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUUIDPathVar(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req PatchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	patch := domain.NewUserPatch(req.FullName.ToDomain(), req.PhoneNumber.ToDomain())

	user, err := h.service.PatchUser(r.Context(), id, patch)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, userToDTO(user))
}
