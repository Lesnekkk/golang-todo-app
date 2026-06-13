package tasks_transport_http

import (
	"net/http"
)

func (h *TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getUUIDPathVar(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
