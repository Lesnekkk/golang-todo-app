package statistics_transport_http

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type StatisticsResponse struct {
	TasksCreated               int      `json:"tasks_created"`
	TasksCompleted             int      `json:"tasks_completed"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	var userID *uuid.UUID
	if raw := r.URL.Query().Get("user_id"); raw != "" {
		id, err := uuid.Parse(raw)
		if err != nil {
			http.Error(w, "invalid user_id", http.StatusBadRequest)
			return
		}
		userID = &id
	}

	var from *time.Time
	if raw := r.URL.Query().Get("from"); raw != "" {
		t, err := time.Parse("2006-01-02", raw)
		if err != nil {
			http.Error(w, "invalid from date, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		from = &t
	}

	var to *time.Time
	if raw := r.URL.Query().Get("to"); raw != "" {
		t, err := time.Parse("2006-01-02", raw)
		if err != nil {
			http.Error(w, "invalid to date, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		to = &t
	}

	stats, err := h.service.GetStatistics(r.Context(), userID, from, to)
	if err != nil {
		writeError(w, err)
		return
	}

	var avgTime *string
	if stats.TasksAverageCompletionTime != nil {
		s := stats.TasksAverageCompletionTime.String()
		avgTime = &s
	}

	writeJSON(w, http.StatusOK, StatisticsResponse{
		TasksCreated:               stats.TasksCreated,
		TasksCompleted:             stats.TasksCompleted,
		TasksCompletedRate:         stats.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	})
}
