package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
)

func CheckAuthorizedBearer(next http.Handler, authService TokenParser) http.Handler {
	return &authorizationBearerCheck{
		next:        next,
		authService: authService,
	}
}

type ContextKey string

const (
	authorizationHeader            = "Authorization"
	ContextUserKey      ContextKey = "name"
)

type TokenParser interface {
	ParseToken(tokenString string) (string, error)
	GetName(phone string) (string, error)
}

type authorizationBearerCheck struct {
	next        http.Handler
	authService TokenParser
}

func (m *authorizationBearerCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get(authorizationHeader)
	if header == "" {
		http.Error(w, "empty auth header", http.StatusBadRequest)
		return
	}

	//split header
	headerSplited := strings.Split(header, " ")
	if len(headerSplited) != 2 || headerSplited[0] != "Bearer" {
		http.Error(w, "invalid auth header", http.StatusUnauthorized)
		return
	}

	if len(headerSplited[1]) == 0 {
		http.Error(w, "token is empty", http.StatusUnauthorized)
	}

	phone, err := m.authService.ParseToken(headerSplited[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	name, _ := m.authService.GetName(phone)
	log.Println("authorized for", name)

	m.next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ContextUserKey, name)))
}
