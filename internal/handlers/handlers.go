package handlers

import "net/http"

func IsHxRequest(r *http.Request) bool {
	return r.Header.Get("Hx-Request") == "true"
}
