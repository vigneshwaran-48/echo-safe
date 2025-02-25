package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vigneshwaran-48/echo-safe/internal/handlers"
	m "github.com/vigneshwaran-48/echo-safe/internal/middleware"
	"github.com/vigneshwaran-48/echo-safe/internal/repository"
	"github.com/vigneshwaran-48/echo-safe/internal/service"
)

func SetupRouter(db *sql.DB, r chi.Router) http.Handler {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	notesRespository := repository.CreateNoteRepository(db)
	notesService := service.CreateNoteService(notesRespository)

	openNotesRepository := repository.CreateOpenNotesRepository(db)
	openNotesService := service.CreateOpenNotesService(openNotesRepository)

	notesHandler := handlers.CreateNotesHandler(notesService, openNotesService)
	homeHanlder := handlers.CreateHomeHandler(notesService)

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			m.CSPMiddleware,
		)
		r.Get("/", homeHanlder.Home)
		r.Route("/notes", func(r chi.Router) {
			r.Post("/", notesHandler.CreateNoteHandler)
			r.Get("/", notesHandler.ListNotesHandler)
			r.Get("/{id}", notesHandler.GetNote)
			r.Patch("/{id}", notesHandler.UpdateNote)
			r.Delete("/{id}", notesHandler.DeleteNote)
		})
	})

	return r
}
