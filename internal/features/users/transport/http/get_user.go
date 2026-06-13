package users_transport_http

import (
	"net/http"
)

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUUIDPathVar(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, userToDTO(user))
}
