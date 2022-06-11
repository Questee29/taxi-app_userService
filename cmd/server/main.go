package main

import (
	"net/http"

	"github.com/Questee29/taxi-app_userService/database"
	"github.com/Questee29/taxi-app_userService/middleware"
	logout "github.com/Questee29/taxi-app_userService/pkg/handlers/log-out"
	signin "github.com/Questee29/taxi-app_userService/pkg/handlers/sign-in"
	signup "github.com/Questee29/taxi-app_userService/pkg/handlers/sign-up"
	welcome "github.com/Questee29/taxi-app_userService/pkg/handlers/welcome"
	authRepository "github.com/Questee29/taxi-app_userService/pkg/repository/authorization"
	authService "github.com/Questee29/taxi-app_userService/pkg/service/authorization"
)

func main() {
	db := database.New()

	usersRepository := authRepository.New(db)
	authService := authService.New(usersRepository)
	handlerSignUp := signup.New(authService)
	handlerSignIn := signin.New(authService)
	handlerWelcome := welcome.New(authService)
	handlerLogout := logout.New(authService)

	http.Handle("/sign-up", middleware.SetContentTypeJSON(handlerSignUp))
	http.Handle("/sign-in", middleware.SetContentTypeJSON(handlerSignIn))
	http.Handle("/welcome", middleware.SetContentTypeJSON(middleware.CheckAuthorizedBearer(handlerWelcome, authService)))
	http.Handle("/logout", middleware.CheckAuthorizedBearer(handlerLogout, authService))

	http.ListenAndServe(":8080", nil)
}
