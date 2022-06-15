package middleware

import (
	"net/http"
)

func SetContentTypeJSON(next http.Handler) http.Handler {
	return &contentTypeJSON{
		next: next,
	}
}

type contentTypeJSON struct {
	next http.Handler
}

func (m *contentTypeJSON) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m.next.ServeHTTP(w, r)
}
