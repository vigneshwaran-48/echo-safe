package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vigneshwaran-48/echo-safe/internal/service"
	"github.com/vigneshwaran-48/echo-safe/internal/templates"
	notesidebar "github.com/vigneshwaran-48/echo-safe/internal/templates/oob/note-sidebar"
	updatenote "github.com/vigneshwaran-48/echo-safe/internal/templates/oob/update-note"
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
	if title == "" {
		title = "Untitled"
	}
	note, err := handler.service.CreateNote(title, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("HX-Trigger-After-Swap", fmt.Sprintf("{\"oncreatenote\": {\"id\": %d, \"title\": \"%s\"}}", note.Id, note.Title))
	err = notesidebar.NoteWithSidebar(note).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	if note == nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	if IsHxRequest(r) {
		// HTMX request hence partial render the page.
		w.Header().Add("HX-Trigger", fmt.Sprintf("{\"onactivenote\": {\"id\": %d, \"title\": \"%s\"}}", note.Id, note.Title))
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
	err = templates.Layout(pages.NotePage(note), note.Title, notes, note.Id).Render(r.Context(), w)
}

func (handler *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id should be a number", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	note, err := handler.service.UpdateNote(int64(noteId), title, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Adding header for updating the title
	w.Header().Add("HX-Trigger-After-Swap", fmt.Sprintf("{\"onactivenote\": {\"id\": %d, \"title\": \"%s\"}}", note.Id, note.Title))
	err = updatenote.UpdateNote(note).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *NotesHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id should be a number", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.service.DeleteNote(int64(noteId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
