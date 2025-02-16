package handlers

import (
	"net/http"

	"github.com/vigneshwaran-48/echo-safe/internal/service"
	"github.com/vigneshwaran-48/echo-safe/internal/templates"
	"github.com/vigneshwaran-48/echo-safe/internal/templates/index"
)

type HomeHandler struct {
	notesService *service.NoteService
}

func CreateHomeHandler(notesService *service.NoteService) *HomeHandler {
	return &HomeHandler{notesService}
}

func (handler *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	index := index.Index()

	notes, err := handler.notesService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templates.Layout(index, "Echo Safe", notes).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
