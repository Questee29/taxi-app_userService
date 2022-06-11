package user

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthDetails struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
