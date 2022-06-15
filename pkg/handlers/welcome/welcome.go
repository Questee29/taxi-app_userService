package welcome

import (
	"fmt"
	"net/http"

	"github.com/Questee29/taxi-app_userService/middleware"
)

type UsersAuthService interface {
	ParseToken(tokenString string) (string, error)
}

type Handler struct {
	usersService UsersAuthService
}

func New(usersService UsersAuthService) *Handler {
	return &Handler{
		usersService: usersService,
	}
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value(middleware.ContextUserKey)
	w.Write([]byte(fmt.Sprintf("Welcome %s!", userEmail)))

}
