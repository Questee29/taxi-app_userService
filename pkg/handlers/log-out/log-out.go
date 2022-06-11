package logout

import (
	"net/http"
)

type UsersAuthService interface {
	DeleteToken(email string) error
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

}
