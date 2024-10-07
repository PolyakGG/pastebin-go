package repository

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"pastebin-go/internal/models"
)

type PasteRepository struct { // - Здесь мы организуем доступ к БД в приложении
	DB *sql.DB // - Используем указатель чтобы не копировать всего объекта базы данных при передаче его в функции или методы
} // - изменять соединение с базой данных в одном месте и видеть эти изменения во всем приложении

func NewPasteRepository(dbPath string) (*PasteRepository, error) {
	log.Printf("Connecting to database at path: %s", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.New("не удалось подключиться к базе данных")
	}
	return &PasteRepository{DB: db}, nil
}
func (r *PasteRepository) Create(paste *models.Paste) (int, error) {

	stmt, err := r.DB.Prepare("INSERT INTO pastes(title, content) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

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

func (r *PasteRepository) GetById(id int) (*models.Paste, error) { // С помощью запроса по id получаем данные
	paste := &models.Paste{}
	err := r.DB.QueryRow("SELECT id, title,content, created_at FROM pastes WHERE id=?", id). //[.]- это Метод-чейнинг(цепочка вызова методов)
													Scan(&paste.ID, &paste.Title, &paste.Content, &paste.CreatedAt) //  метод Scan автоматически заполняет соответствующие поля в структуре paste.
	if err != nil {
		return nil, err
	}
	return paste, nil
}

// Сервисный слой

func (r *PasteRepository) GetAll() ([]models.Paste, error) {
	rows, err := r.DB.Query("SELECT id, title, content, created_at FROM pastes")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var pastes []models.Paste
	for rows.Next() {
		var paste models.Paste
		if err := rows.Scan(&paste.ID, &paste.Title, &paste.Content, &paste.CreatedAt); err != nil {
			return nil, err
		}
		pastes = append(pastes, paste)
	}

	return pastes, nil
}

// Метод для поиска пасты по ID
func (r *PasteRepository) FindByID(id string) (models.Paste, error) {
	row := r.DB.QueryRow("SELECT id, title, content FROM pastes WHERE id = ?", id)

	var paste models.Paste
	err := row.Scan(&paste.ID, &paste.Title, &paste.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Paste{}, errors.New("paste not found")
		}
		return models.Paste{}, err
	}

	return paste, nil
}
