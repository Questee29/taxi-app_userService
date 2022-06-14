package authorization

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	user "github.com/Questee29/taxi-app_userService/models/user"
)

type Repository interface {
	GetUser(phone, password string) (user.ResponseAuthDetails, error)
	GetName(phone string) (string, error)

	IsRegistred(email, number string) (bool, error)
	CreateUser(name, phone, email, hashPass string) error
}
type authService struct {
	repository Repository
}

func New(repository Repository) *authService {
	return &authService{
		repository: repository,
	}
}

// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type tokenClaims struct {
	jwt.StandardClaims
	Phone string `json:"phone"`
}

// Create the JWT key used to create the signature
const (
	jwtKey       = "very_secret_key"
	tokenExpires = 15 * time.Minute
)

func (service *authService) GenerateJWT(phone, password string) (string, error) {

	user, err := service.repository.GetUser(phone, password)
	if err != nil {
		return "", err
	}

	if err := service.MatchPass(password, user.HashPassword); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpires).Unix(),
			IssuedAt:  time.Now().Unix(),
		},

		user.Phone,
	})

	return token.SignedString([]byte(jwtKey))

}

func (service *authService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(`invalid signing method`)
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New(`token claims are incorrect type `)
	}

	return claims.Phone, nil
}
func (service *authService) DeleteToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(`invalid signing method`)
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return errors.New(`token claims are incorrect type `)
	}
	claims.ExpiresAt = time.Now().Add(-time.Hour).Unix()
	log.Println(claims.ExpiresAt)
	return nil
}

func (service *authService) GetName(phone string) (string, error) {
	return service.repository.GetName(phone)
}

func (service *authService) MatchPass(password, HashPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
