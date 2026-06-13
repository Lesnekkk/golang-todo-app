package tasks_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	AuthorUserID uuid.UUID `json:"author_user_id"`
}

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(r.Context(), req.Title, req.Description, req.AuthorUserID)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, taskToDTO(task))
}
