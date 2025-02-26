package repository

import (
	"database/sql"

	"github.com/vigneshwaran-48/echo-safe/internal/models"
)

type OpenNotesRepository struct {
	db *sql.DB
}

func CreateOpenNotesRepository(db *sql.DB) *OpenNotesRepository {
	return &OpenNotesRepository{db}
}

func (repository *OpenNotesRepository) Save(openNote *models.OpenNote) error {
	query := `
    INSERT INTO open_note (note_id, active) values (?, ?) RETURNING id
  `
	active := 0
	if openNote.Active {
		active = 1
	}
	row := repository.db.QueryRow(query, openNote.NoteId, active)
	return row.Scan(&openNote.Id)
}

func (repository *OpenNotesRepository) Update(openNote *models.OpenNote) error {
	query := `
        UPDATE open_note 
        SET active = ? 
        WHERE note_id = ?`

	_, err := repository.db.Exec(query, openNote.Active, openNote.NoteId)
	return err
}

func (repository *OpenNotesRepository) DeleteById(noteId int64) error {
	query := `DELETE FROM open_note WHERE note_id = ?`
	_, err := repository.db.Exec(query, noteId)
	return err
}

func (repository *OpenNotesRepository) FindAll() ([]models.OpenNote, error) {
	query := `SELECT id, note_id, active FROM open_note`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var openNotes []models.OpenNote
	for rows.Next() {
		openNote := models.OpenNote{}
		err := rows.Scan(&openNote.Id, &openNote.NoteId, &openNote.Active)
		if err != nil {
			return nil, err
		}
		openNotes = append(openNotes, openNote)
	}
	return openNotes, nil
}

func (repository *OpenNotesRepository) FindByNoteId(noteId int64) (*models.OpenNote, error) {
	query := `SELECT id, note_id, active from open_note WHERE note_id = ?`

	openNote := &models.OpenNote{}
	err := repository.db.QueryRow(query, noteId).Scan(&openNote.Id, &openNote.NoteId, &openNote.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return openNote, nil
}

func (repository *OpenNotesRepository) FindActiveNote() (*models.OpenNote, error) {
	query := `SELECT id, note_id, active from open_note WHERE active = 1`

	openNote := &models.OpenNote{}
	err := repository.db.QueryRow(query).Scan(&openNote.Id, &openNote.NoteId, &openNote.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return openNote, nil
}
