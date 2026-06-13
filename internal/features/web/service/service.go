package web_service

import "github.com/Lesnekkk/golang-todo-app/internal/core/domain"

type WebRepository interface {
	GetFile(filePath string) (domain.File, error)
}

type WebService struct {
	repo WebRepository
}

func NewWebService(repo WebRepository) *WebService {
	return &WebService{repo: repo}
}
