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

	fillNotesName(openNotes, handler.notesService)

	// TODO Need to change it to active note once the last active note changes have been done
	err = templates.Layout(index, "Echo Safe", notes, 0, openNotes).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
