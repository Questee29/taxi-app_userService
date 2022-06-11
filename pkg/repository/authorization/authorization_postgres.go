package authorization

import (
	"database/sql"
	"errors"

	user "github.com/Questee29/taxi-app_userService/models/user"
	"golang.org/x/crypto/bcrypt"
)

type authorizationRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *authorizationRepository {
	return &authorizationRepository{
		db: db,
	}
}
func (repository *authorizationRepository) GetName(phone string) (string, error) {
	var name string
	query := `SELECT name
	FROM uusers
	WHERE phone=$1 
	`
	row := repository.db.QueryRow(query, phone)
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}
func (repository *authorizationRepository) GetUser(phone, password string) (user.AuthDetails, error) {
	var user user.AuthDetails

	query := `SELECT password
	FROM uusers
	WHERE phone=$1 
	`
	row := repository.db.QueryRow(query, phone)
	if err := row.Scan(&user.Password); err != nil {

		return user, err
	}
	if repository.MatchPass(password, user.Password) {
		user.Password = password
		user.Phone = phone

		return user, nil
	}
	return user, errors.New(`invalid password`)
}

func (repository *authorizationRepository) MatchPass(password, hashPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false
	}
	return true

}
func (repository *authorizationRepository) FindCopies(email, phone string) (bool, error) {
	result, err := repository.db.Query("SELECT email FROM uusers WHERE email = $1 or phone = $2", email, phone)
	if err != nil {
		return false, err
	}
	if result.Next() {
		return true, nil
	}

	return false, nil
}
func (repository *authorizationRepository) CreateUser(name, phone, email, hashPass string) error {
	query := `
	INSERT into uusers(name,phone,email,password) 
	VALUES ($1,$2,$3,$4)`
	if _, err := repository.db.Exec(query, name, phone, email, hashPass); err != nil {
		return err
	}
	return nil
}
