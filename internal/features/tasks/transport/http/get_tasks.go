package tasks_transport_http

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var userID *uuid.UUID
	if raw := r.URL.Query().Get("user_id"); raw != "" {
		id, err := uuid.Parse(raw)
		if err != nil {
			http.Error(w, "invalid user_id", http.StatusBadRequest)
			return
		}
		userID = &id
	}

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

	tasks, err := h.service.GetTasks(r.Context(), userID, limit, offset)
	if err != nil {
		writeError(w, err)
		return
	}

	dtos := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		dtos[i] = taskToDTO(task)
	}

	writeJSON(w, http.StatusOK, dtos)
}
