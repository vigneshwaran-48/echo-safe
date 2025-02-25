package service

import (
	"errors"
	"fmt"

	"github.com/vigneshwaran-48/echo-safe/internal/models"
	"github.com/vigneshwaran-48/echo-safe/internal/repository"
)

type OpenNoteService struct {
	repository *repository.OpenNotesRepository
}

func CreateOpenNotesService(repository *repository.OpenNotesRepository) *OpenNoteService {
	return &OpenNoteService{repository}
}

func (service *OpenNoteService) AddOpenNote(id int64) (*models.OpenNote, error) {
	existingNote, err := service.repository.FindByNoteId(id)
	if err != nil {
		return nil, err
	}
	if existingNote != nil {
		return nil, errors.New("Already note is open")
	}

	openNote := &models.OpenNote{NoteId: id, Active: true}
	err = service.repository.Save(openNote)
	if err != nil {
		return nil, err
	}
	return openNote, nil
}

func (service *OpenNoteService) SetActive(id int64, active bool) error {
	openNote, err := service.repository.FindByNoteId(id)
	if err != nil {
		return err
	}
	if openNote == nil {
		return fmt.Errorf("Note %d is not open", id)
	}
	openNote.Active = active
	return service.repository.Update(openNote)
}

func (service *OpenNoteService) GetOpenNote(id int64) (*models.OpenNote, error) {
	openNote, err := service.repository.FindByNoteId(id)
	if err != nil {
		return nil, err
	}
	return openNote, err
}
