package statistics_transport_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

type StatisticsService interface {
	GetStatistics(ctx context.Context, userID *uuid.UUID, from *time.Time, to *time.Time) (domain.Statistics, error)
}

type StatisticsHTTPHandler struct {
	service StatisticsService
}

func NewStatisticsHTTPHandler(service StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{service: service}
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
