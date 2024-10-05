package service

import (
	"pastebin-go/internal/models"
	"pastebin-go/internal/repository"
)

type PasteService struct { // ссылаемся на файл repository/sqlite.go
	Repo *repository.PasteRepository
}

func NewPasteService(repo *repository.PasteRepository) *PasteService { // Указатель для того чтобы структура не попал в кучу
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

/*
Функция GetAll необходима для правильной организации кода и разделения ответственности.
Она предоставляет удобный интерфейс для получения всех паст из базы данных через сервисный слой,
позволяя контроллерам (хэндлерам) не заниматься непосредственно запросами к базе данных.
*/

func (s *PasteService) GetAll() ([]models.Paste, error) {
	return s.Repo.GetAll()
}
