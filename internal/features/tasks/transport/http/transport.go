package tasks_transport_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

type TaskService interface {
	CreateTask(ctx context.Context, title string, description *string, authorUserID uuid.UUID) (domain.Task, error)
	GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error)
	GetTasks(ctx context.Context, userID *uuid.UUID, limit *int, offset *int) ([]domain.Task, error)
	PatchTask(ctx context.Context, id uuid.UUID, patch domain.TaskPatch) (domain.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

type TasksHTTPHandler struct {
	service TaskService
}

func NewTasksHTTPHandler(service TaskService) *TasksHTTPHandler {
	return &TasksHTTPHandler{service: service}
}

type TaskDTOResponse struct {
	ID           uuid.UUID  `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID uuid.UUID  `json:"author_user_id"`
}

func taskToDTO(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
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
