package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/heroku/hotpotbot9/models"
)

var (
	//SurveyPostLogic ...
	SurveyPostLogic models.SurveyPostLogic
)

// SurveyEdit ...
func SurveyEdit(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/survey.html"))

	if err := t.ExecuteTemplate(w, "survey.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

// SurveyPost ...
func SurveyPost(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/postsurvey.html"))

	if r.Method == "POST" {
		r.ParseForm()
		formData := map[string]string{
			"name":         r.Form["name"][0],
			"satisfaction": r.Form["satisfaction"][0],
			"impression":   r.Form["impression"][0],
			"theme":        r.Form["theme"][0],
		}
		log.Print(formData)
		err := SurveyPostLogic.Save(formData)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.ExecuteTemplate(w, "postsurvey.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}
