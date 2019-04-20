package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Survey ...
func Survey(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/survey.html"))

	if err := t.ExecuteTemplate(w, "survey.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

// PostSurvey ...
func PostSurvey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("debug1")

	t := template.Must(template.ParseFiles("templates/postsurvey.html"))

	if r.Method == "POST" {
		r.ParseForm()
		// formData := map[string]string{
		// 	"satisfaction": r.Form["satisfaction"][0],
		// 	"impression":   r.Form["impression"][0],
		// 	"theme":        r.Form["theme"][0],
		// }
		// log.Print(formData)
		if err := t.ExecuteTemplate(w, "postsurvey.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}
