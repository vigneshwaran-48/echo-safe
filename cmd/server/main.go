package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vigneshwaran-48/echo-safe/api"
	"github.com/vigneshwaran-48/echo-safe/internal/db"
)

func main() {
	db := db.CreateDB()
	router := chi.NewRouter()

	api.SetupRouter(db, router)

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	log.Println("Server started on port 8089")
	log.Fatal(http.ListenAndServe(":8089", router))
}
