package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/heroku/hotpotbot9/handlers"
	"github.com/heroku/hotpotbot9/models"
)

func main() {
	fmt.Println("Server Start...")

	// class initialize
	fmt.Println("Initialize classes ...")
	handlers.PersonalInfoLogic = models.NewPersonalInfoLogic()
	handlers.SurveyRegisterLogic = models.NewSurveyRegisterLogic()

	http.HandleFunc("/webhook", handlers.Main)
	http.HandleFunc("/personal-info", handlers.PersonalInfoEdit)
	http.HandleFunc("/entry", handlers.Entry)
	http.HandleFunc("/survey", handlers.SurveyEdit)
	http.HandleFunc("/survey-registered", handlers.SurveyRegistered)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
