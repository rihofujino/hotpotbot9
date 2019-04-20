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

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
