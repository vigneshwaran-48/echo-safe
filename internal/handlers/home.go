package handlers

import (
	"net/http"

	"github.com/vigneshwaran-48/echo-safe/internal/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	index := templates.Index()

	err := templates.Layout(index, "Echo Safe").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
