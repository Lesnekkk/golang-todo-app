package users_transport_http

import (
	"net/http"
	"strconv"
)

func (h *UsersHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var limit, offset *int

	if raw := r.URL.Query().Get("limit"); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
		limit = &v
	}

	if raw := r.URL.Query().Get("offset"); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil {
			http.Error(w, "invalid offset", http.StatusBadRequest)
			return
		}
		offset = &v
	}

	users, err := h.service.GetUsers(r.Context(), limit, offset)
	if err != nil {
		writeError(w, err)
		return
	}

	dtos := make([]UserDTOResponse, len(users))
	for i, user := range users {
		dtos[i] = userToDTO(user)
	}

	writeJSON(w, http.StatusOK, dtos)
}
