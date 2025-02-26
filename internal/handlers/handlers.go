package handlers

import (
	"net/http"

	"github.com/vigneshwaran-48/echo-safe/internal/models"
	"github.com/vigneshwaran-48/echo-safe/internal/service"
)

func IsHxRequest(r *http.Request) bool {
	return r.Header.Get("Hx-Request") == "true"
}

func fillNotesName(openNotes []models.OpenNote, notesService *service.NoteService) error {
	for i := range openNotes {
		openNote := &openNotes[i]
		note, err := notesService.GetById(openNote.NoteId)
		if err != nil {
			return err
		}
		openNote.Title = note.Title
	}
	return nil
}
