package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/heroku/hotpotbot9/handlers"
	"github.com/heroku/hotpotbot9/models"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	fmt.Println("Server Start...")

	// DAO initialize
	fmt.Println("Initialize dao classes ...")
	handlers.PersonalInfoLogic = models.NewPersonalInfoLogic()
	handlers.SurveyPostLogic = models.NewSurveyPostLogic()

	http.HandleFunc("/webhook", handlers.Main)
	http.HandleFunc("/personal-info", handlers.PersonalInfoEdit)
	http.HandleFunc("/entry", handlers.PersonalInfoPost)
	http.HandleFunc("/survey", handlers.SurveyEdit)
	http.HandleFunc("/post-survey", handlers.SurveyPost)

	http.HandleFunc("/", testHandler)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
