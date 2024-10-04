package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"pastebin-go/internal/models"
)

type PasteRepository struct {
	DB *sql.DB
}

func NewPasteRepository(dbPath string) (*PasteRepository, error) {
	log.Printf("Connecting to database at path: %s", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &PasteRepository{DB: db}, nil
}
func (r *PasteRepository) Create(paste *models.Paste) (int, error) {

	stmt, err := r.DB.Prepare("INSERT INTO pastes(title, content) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(paste.Title, paste.Content)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *PasteRepository) GetById(id int) (*models.Paste, error) {
	paste := &models.Paste{}
	err := r.DB.QueryRow("SELECT id, title,content, created_at FROM pastes WHERE id=?", id).
		Scan(&paste.ID, &paste.Title, &paste.Content, &paste.CreatedAt)
	if err != nil {
		return nil, err
	}
	return paste, nil
}
