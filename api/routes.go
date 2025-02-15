package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vigneshwaran-48/echo-safe/internal/handlers"
	"github.com/vigneshwaran-48/echo-safe/internal/repository"
	"github.com/vigneshwaran-48/echo-safe/internal/service"
)

func SetupRouter(db *sql.DB, r chi.Router) http.Handler {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	notesRespository := repository.CreateNoteRepository(db)
	notesService := service.CreateNoteService(notesRespository)
	notesHandler := handlers.CreateNotesHandler(notesService)

	r.Route("/notes", func(r chi.Router) {
		r.Post("/", notesHandler.CreateNoteHandler)
		r.Get("/", notesHandler.ListNotesHandler)
	})

	return r
}
