package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/heroku/hotpotbot9/models"
)

// PersonalInfo ...
func PersonalInfo(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/personal_info.html"))

	if err := t.ExecuteTemplate(w, "personal_info.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

// Entry ...
func Entry(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/entry.html"))

	if r.Method == "POST" {
		r.ParseForm()
		formData := map[string]string{
			"name":    r.Form["name"][0],
			"company": r.Form["company"][0],
			"jobType": r.Form["jobType"][0],
		}
		log.Print(formData)
		err := models.Save(formData)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.ExecuteTemplate(w, "entry.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}
