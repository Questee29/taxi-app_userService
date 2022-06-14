package main

import (
	"log"
	"net/http"

	config "github.com/Questee29/taxi-app_userService/configs"
	"github.com/Questee29/taxi-app_userService/database"
	"github.com/Questee29/taxi-app_userService/middleware"
	_ "github.com/Questee29/taxi-app_userService/migrations"
	logout "github.com/Questee29/taxi-app_userService/pkg/handlers/log-out"
	signin "github.com/Questee29/taxi-app_userService/pkg/handlers/sign-in"
	signup "github.com/Questee29/taxi-app_userService/pkg/handlers/sign-up"
	welcome "github.com/Questee29/taxi-app_userService/pkg/handlers/welcome"
	authRepository "github.com/Questee29/taxi-app_userService/pkg/repository/authorization"
	authService "github.com/Questee29/taxi-app_userService/pkg/service/authorization"
)

func main() {
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	db, err := database.New()
	if err != nil {
		log.Fatalln(err)
	}

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

	http.ListenAndServe(config.Server.Port, nil)
}
