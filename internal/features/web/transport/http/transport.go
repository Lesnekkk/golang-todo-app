package web_transport_http

import (
	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

type WebService interface {
	GetMainPage() (domain.File, error)
}

type WebHTTPHandler struct {
	service WebService
}

func NewWebHTTPHandler(service WebService) *WebHTTPHandler {
	return &WebHTTPHandler{service: service}
}
