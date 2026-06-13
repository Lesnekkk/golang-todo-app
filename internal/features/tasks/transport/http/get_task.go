package tasks_transport_http

import (
	"net/http"
)

func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getUUIDPathVar(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, taskToDTO(task))
}
