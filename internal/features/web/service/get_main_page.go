package web_service

import (
	"fmt"
	"os"
	"path"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

func (s *WebService) GetMainPage() (domain.File, error) {
	htmlFilePath := path.Join(os.Getenv("PROJECT_ROOT"), "/public/index.html")

	file, err := s.repo.GetFile(htmlFilePath)
	if err != nil {
		return domain.File{}, fmt.Errorf("get main page: %w", err)
	}

	return file, nil
}
