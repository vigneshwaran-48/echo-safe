package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vigneshwaran-48/echo-safe/internal/models"
	"github.com/vigneshwaran-48/echo-safe/internal/service"
	"github.com/vigneshwaran-48/echo-safe/internal/templates"
	"github.com/vigneshwaran-48/echo-safe/internal/templates/index"
	notesidebar "github.com/vigneshwaran-48/echo-safe/internal/templates/oob/note-sidebar"
	opennotes "github.com/vigneshwaran-48/echo-safe/internal/templates/oob/open-notes"
	"github.com/vigneshwaran-48/echo-safe/internal/templates/pages"
)

type NotesHandler struct {
	service          *service.NoteService
	openNotesService *service.OpenNoteService
}

func CreateNotesHandler(service *service.NoteService, openNotesService *service.OpenNoteService) *NotesHandler {
	return &NotesHandler{service, openNotesService}
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
	_, err = handler.openNotesService.AddOpenNote(note.Id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
	openNotes, err := handler.openNotesService.GetAllOpenNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = fillNotesName(openNotes, handler.service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("HX-Trigger-After-Swap", fmt.Sprintf("{\"oncreatenote\": {\"id\": %d, \"title\": \"%s\"}}", note.Id, note.Title))
	err = notesidebar.NoteWithSidebar(note, openNotes).Render(r.Context(), w)
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if note == nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	openNote, err := handler.openNotesService.GetOpenNote(note.Id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
	if openNote == nil {
		_, err = handler.openNotesService.AddOpenNote(note.Id)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	} else if !openNote.Active {
		err = handler.openNotesService.SetActive(note.Id, true)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}

	openNotes, err := getOpenNotes(handler)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	if IsHxRequest(r) {
		// HTMX request hence partial render the page.
		err = pages.NotePage(note, openNotes).Render(r.Context(), w)
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
	err = templates.Layout(pages.NotePage(note, nil), note.Title, notes, note.Id, openNotes).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id should be a number", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	_, err = handler.service.UpdateNote(int64(noteId), title, content)
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
	err = handler.openNotesService.DeleteOpenNote(int64(noteId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.service.DeleteNote(int64(noteId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handleDeleteNoteResponse(handler, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *NotesHandler) DeleteOpenNote(w http.ResponseWriter, r *http.Request) {
	noteId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id should be a number", http.StatusBadRequest)
		return
	}
	err = handler.openNotesService.DeleteOpenNote(int64(noteId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = handleDeleteNoteResponse(handler, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getOpenNotes(handler *NotesHandler) ([]models.OpenNote, error) {
	openNotes, err := handler.openNotesService.GetAllOpenNotes()
	if err != nil {
		return nil, err
	}

	err = fillNotesName(openNotes, handler.service)
	if err != nil {
		return nil, err
	}
	return openNotes, nil
}

func handleDeleteNoteResponse(handler *NotesHandler, w http.ResponseWriter, r *http.Request) error {
	openNotes, err := getOpenNotes(handler)
	if err != nil {
		return err
	}
	var activeNote int64
	activeNote = 0
	for _, openNote := range openNotes {
		if openNote.Active {
			activeNote = openNote.NoteId
			break
		}
	}

	if activeNote != 0 {
		note, err := handler.service.GetById(activeNote)
		if err != nil {
			return err
		}
		w.Header().Add("HX-Push-Url", fmt.Sprintf("/notes/%d", activeNote))

		err = pages.NotePage(note, openNotes).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return nil
		}
	} else {
		w.Header().Add("HX-Push-Url", "/")

		err = index.Index().Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
		err = opennotes.OpenNotes(openNotes).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}
	notes, err := handler.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	err = notesidebar.Sidebar(notes, activeNote).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}
