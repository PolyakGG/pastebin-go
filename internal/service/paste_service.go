package service

import (
	"pastebin-go/internal/models"
	"pastebin-go/internal/repository"
)

type PasteService struct {
	Repo *repository.PasteRepository
}

func NewPasteService(repo *repository.PasteRepository) *PasteService {
	return &PasteService{Repo: repo}
}

func (s *PasteService) CreatePaste(title, content string) (int, error) {
	paste := &models.Paste{
		Title:   title,
		Content: content,
	}
	return s.Repo.Create(paste)
}

func (s *PasteService) GetPaste(id int) (*models.Paste, error) {
	return s.Repo.GetById(id)
}
