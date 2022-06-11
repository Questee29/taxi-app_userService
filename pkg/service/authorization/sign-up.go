package authorization

import (
	"log"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func (service *authService) IsPasswordValid(password string) bool {
	// upp: at least one upper case letter.
	// low: at least one lower case letter.
	// num: at least one digit.
	// sym: at least one special character.
	// tot: at least eight characters long.
	// No empty string or whitespace.
	var (
		upp, low, num bool
		tot           uint8
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			return false
		default:
			return false
		}
	}
	if !upp || !low || !num || tot < 8 {
		return false
	}
	return true
}
func (service *authService) IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)

}
func (service *authService) IsNumberValid(number string) bool {
	numberRegex := regexp.MustCompile(`^(\+375|80)(29|25|44|33)(\d{3})(\d{2})(\d{2})$`)
	log.Println(numberRegex.MatchString(number))

	return numberRegex.MatchString(number)
}
func (service *authService) IsRegistred(email, number string) (bool, error) {
	return service.repository.FindCopies(email, number)
}

func (service *authService) GeneratePasswordHash(password string) (string, error) {
	//hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	hashPass := string(bytes)
	return hashPass, nil
}

func (service *authService) RegisterUser(name, phone, email, password string) error {
	//Hash
	hashPass, err := service.GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	//insert emal,pass into db table
	return service.repository.CreateUser(name, phone, email, hashPass)
}
