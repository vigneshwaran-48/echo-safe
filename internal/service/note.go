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

func (service *NoteService) List() ([]models.Note, error) {
	return service.repository.FindAll()
}

func (service *NoteService) GetById(id int64) (*models.Note, error) {
	return service.repository.FindById(id)
}

func (service *NoteService) UpdateNote(id int64, title string, content string) (*models.Note, error) {
	note, err := service.GetById(id)
	if err != nil {
		return nil, err
	}
	if title != "" {
		note.Title = title
	}
	if content != "" {
		note.Content = content
	}
	err = service.repository.Update(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (service *NoteService) DeleteNote(id int64) error {
	_, err := service.GetById(id)
	if err != nil {
		return err
	}
	err = service.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
