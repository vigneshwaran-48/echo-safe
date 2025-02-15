package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vigneshwaran-48/echo-safe/internal/service"
)

type NotesHandler struct {
	service *service.NoteService
}

func CreateNotesHandler(service *service.NoteService) *NotesHandler {
	return &NotesHandler{service}
}

func (handler *NotesHandler) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	if _, err := handler.service.CreateNote(title, content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Note saved successfully!"))
}

func (handler *NotesHandler) ListNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := handler.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notes)
}
