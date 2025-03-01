package service

import (
	"errors"
	"fmt"
	"log"

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

	activeNote, err := service.repository.FindActiveNote()
	if err != nil {
		return nil, err
	}
	if activeNote != nil {
		err = service.SetActive(activeNote.NoteId, false)
		if err != nil {
			return nil, err
		}
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
	if active {
		activeNote, err := service.repository.FindActiveNote()
		if err != nil {
			return err
		}
		if activeNote != nil {
			err = service.SetActive(activeNote.NoteId, false)
			if err != nil {
				return err
			}
		}
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

func (service *OpenNoteService) GetAllOpenNotes() ([]models.OpenNote, error) {
	openNotes, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return openNotes, nil
}

func (service *OpenNoteService) DeleteOpenNote(noteId int64) error {
	openNote, err := service.repository.FindByNoteId(noteId)
	if err != nil {
		return err
	}
	if openNote == nil {
		log.Printf("Open note not exists for note %d", noteId)
		return nil
	}
	err = service.repository.DeleteByNoteId(openNote.NoteId)
	if err != nil {
		return err
	}
	if !openNote.Active {
		return nil
	}
	openNotes, err := service.GetAllOpenNotes()
	if err != nil {
		return err
	}
	if len(openNotes) > 0 {
		return service.SetActive(openNotes[0].NoteId, true)
	}
	return nil
}
