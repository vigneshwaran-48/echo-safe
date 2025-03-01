package handlers

import (
	"net/http"

	"github.com/vigneshwaran-48/echo-safe/internal/service"
	"github.com/vigneshwaran-48/echo-safe/internal/templates"
	"github.com/vigneshwaran-48/echo-safe/internal/templates/index"
)

type HomeHandler struct {
	notesService     *service.NoteService
	openNotesService *service.OpenNoteService
}

func CreateHomeHandler(notesService *service.NoteService, openNotesService *service.OpenNoteService) *HomeHandler {
	return &HomeHandler{notesService, openNotesService}
}

func (handler *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	index := index.Index()

	notes, err := handler.notesService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	openNotes, err := handler.openNotesService.GetAllOpenNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = fillNotesName(openNotes, handler.notesService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var activeOpenNote int64
	for _, openNote := range openNotes {
		if openNote.Active {
			activeOpenNote = openNote.NoteId
      break
		}
	}
	err = templates.Layout(index, "Echo Safe", notes, activeOpenNote, openNotes).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
