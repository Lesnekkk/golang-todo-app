package web_transport_http

import (
	"errors"
	"net/http"

	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

func (h *WebHTTPHandler) GetMainPage(w http.ResponseWriter, r *http.Request) {
	file, err := h.service.GetMainPage()
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(file.Buffer())
}
