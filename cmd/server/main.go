package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
