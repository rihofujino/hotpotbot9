package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/heroku/gyozabot/handlers"
)

func main() {
	fmt.Println("Server Start...")

	r := mux.NewRouter()
	fmt.Println("debug1")
	r.HandleFunc("/webhook", handlers.Main)
	r.HandleFunc("/personal-info", handlers.PersonalInfo)

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
