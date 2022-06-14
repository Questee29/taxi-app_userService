package authorization

import (
	"database/sql"

	user "github.com/Questee29/taxi-app_userService/models/user"
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
	FROM users
	WHERE phone=$1 
	`
	row := repository.db.QueryRow(query, phone)
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}
func (repository *authorizationRepository) GetUser(phone, password string) (user.ResponseAuthDetails, error) {
	var user user.ResponseAuthDetails

	query := `SELECT password
	FROM users
	WHERE phone=$1 
	`
	row := repository.db.QueryRow(query, phone)
	if err := row.Scan(&user.HashPassword); err != nil {

		return user, err
	}
	user.Phone = phone
	return user, nil

}

func (repository *authorizationRepository) IsRegistred(email, phone string) (bool, error) {
	result, err := repository.db.Query("SELECT email FROM users WHERE email = $1 or phone = $2", email, phone)
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
	INSERT into users(name,phone,email,password) 
	VALUES ($1,$2,$3,$4)`
	if _, err := repository.db.Exec(query, name, phone, email, hashPass); err != nil {
		return err
	}
	return nil
}
