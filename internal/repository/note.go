package repository

import (
	"database/sql"

	"github.com/vigneshwaran-48/echo-safe/internal/models"
)

type NoteRepository struct {
	db *sql.DB
}

func CreateNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (repository *NoteRepository) Insert(note *models.Note) error {
	query := `
    INSERT INTO note (title, content) values (?, ?) RETURNING id
  `
	row := repository.db.QueryRow(query, note.Title, note.Content)
	return row.Scan(&note.Id)
}

func (repository *NoteRepository) FindById(id int64) (*models.Note, error) {
	note := &models.Note{}
	query := `SELECT id, title, content FROM note WHERE id = ?`
	err := repository.db.QueryRow(query, id).Scan(&note.Id, &note.Title, &note.Content)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (repository *NoteRepository) Update(note *models.Note) error {
	query := `
        UPDATE note 
        SET title = ?, content = ? 
        WHERE id = ?`

	_, err := repository.db.Exec(query, note.Title, note.Content)
	return err
}

func (repository *NoteRepository) Delete(id int64) error {
	query := `DELETE FROM note WHERE id = ?`
	_, err := repository.db.Exec(query, id)
	return err
}

func (repository *NoteRepository) FindAll() ([]models.Note, error) {
	query := `SELECT id, title, content FROM note`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		note := models.Note{}
		err := rows.Scan(&note.Id, &note.Title, &note.Content)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
