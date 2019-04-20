package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/heroku/gyozabot/handlers"
)

func main() {
	fmt.Println("Server Start...")

	http.HandleFunc("/webhook", handlers.Main)
	http.HandleFunc("/personal-info", handlers.PersonalInfo)
	http.HandleFunc("/entry", handlers.Entry)
	http.HandleFunc("/survey", handlers.Survey)
	http.HandleFunc("/post-survey", handlers.PostSurvey)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
