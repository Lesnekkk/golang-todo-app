package tasks_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_http_types "github.com/Lesnekkk/golang-todo-app/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	id, err := getUUIDPathVar(r, "id")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req PatchTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	patch := domain.NewTaskPatch(req.Title.ToDomain(), req.Description.ToDomain(), req.Completed.ToDomain())

	task, err := h.service.PatchTask(r.Context(), id, patch)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, taskToDTO(task))
}
