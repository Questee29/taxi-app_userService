package signup

import (
	"encoding/json"
	"net/http"

	user "github.com/Questee29/taxi-app_userService/models/user"
	_ "github.com/lib/pq"
)

type UsersSignupService interface {
	IsRegistred(email, phone string) (bool, error)
	IsPasswordValid(password string) bool
	IsEmailValid(email string) bool
	IsNumberValid(number string) bool
	RegisterUser(name, phone, email, password string) error
}

type Handler struct {
	service UsersSignupService
}

func New(u UsersSignupService) *Handler {
	return &Handler{
		service: u,
	}
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u user.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !handler.service.IsPasswordValid(u.Password) {
		http.Error(w, "Bad password,try again. At least 8 chars(at least 1 upper,1 lower,1 num)", http.StatusBadRequest)
		return
	}

	if !handler.service.IsEmailValid(u.Email) {
		http.Error(w, "invalid email. Example : example@gmail.com", http.StatusBadRequest)
		return
	}

	if !handler.service.IsNumberValid(u.Phone) {
		http.Error(w, "invalid phone number. Only for belarus users", http.StatusBadRequest)
		return
	}

	isRegistred, err := handler.service.IsRegistred(u.Email, u.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isRegistred {
		http.Error(w, "Sorry, email or phone number already exists", http.StatusBadRequest)
		return
	}

	if err := handler.service.RegisterUser(u.Name, u.Phone, u.Email, u.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
