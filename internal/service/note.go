package service

import (
	"github.com/vigneshwaran-48/echo-safe/internal/models"
	"github.com/vigneshwaran-48/echo-safe/internal/repository"
)

type NoteService struct {
	repository *repository.NoteRepository
}

func CreateNoteService(repository *repository.NoteRepository) *NoteService {
	return &NoteService{repository}
}

func (service *NoteService) CreateNote(title string, content string) (*models.Note, error) {
	note := &models.Note{
		Title:   title,
		Content: content,
	}
	err := service.repository.Insert(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (service *NoteService) List() ([]*models.Note, error) {
	return service.repository.FindAll()
}
