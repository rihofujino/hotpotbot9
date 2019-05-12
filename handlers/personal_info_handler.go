package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/heroku/hotpotbot9/models"
)

var (
	//PersonalInfoLogic ...
	PersonalInfoLogic models.PersonalInfoLogic
)

// PersonalInfoEdit ...
func PersonalInfoEdit(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/personal_info.html"))

	if err := t.ExecuteTemplate(w, "personal_info.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

// PersonalInfoPost ...
func PersonalInfoPost(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/entry.html"))

	if r.Method == "POST" {
		r.ParseForm()
		formData := map[string]string{
			"userID":  r.Form["user_id"][0],
			"name":    r.Form["name"][0],
			"company": r.Form["company"][0],
			"jobType": r.Form["jobType"][0],
		}
		log.Print(formData)
		err := PersonalInfoLogic.Save(formData)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.ExecuteTemplate(w, "entry.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}
