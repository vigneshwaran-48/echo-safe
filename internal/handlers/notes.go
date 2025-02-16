package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vigneshwaran-48/echo-safe/internal/service"
	"github.com/vigneshwaran-48/echo-safe/internal/templates"
	"github.com/vigneshwaran-48/echo-safe/internal/templates/pages"
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

func (handler *NotesHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id should be a number", http.StatusBadRequest)
		return
	}
	note, err := handler.service.GetById(int64(noteId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if IsHxRequest(r) {
		// HTMX request hence partial render the page.
		err = pages.NotePage(note).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	// Render the whole page
	notes, err := handler.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templates.Layout(pages.NotePage(note), note.Title, notes).Render(r.Context(), w)
}
