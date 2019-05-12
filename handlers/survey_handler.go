package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/heroku/hotpotbot9/models"
)

var (
	//SurveyRegisterLogic ...
	SurveyRegisterLogic models.SurveyRegisterLogic
)

// SurveyEdit ...
func SurveyEdit(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("this is formdata")
		fmt.Println(r.Form)
		formData := map[string]string{
			"userID":       r.Form["user_id"][0],
			"satisfaction": r.Form["satisfaction"][0],
			"impression":   r.Form["impression"][0],
			"theme":        r.Form["theme"][0],
		}
		log.Print(formData)

		if e := surveyFormValidator(formData); e != nil {
			data["formErrors"] = e
		} else {
			err := SurveyRegisterLogic.Save(formData)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("location", "/survey-registered")
			w.WriteHeader(http.StatusSeeOther)
			return
		}
	}

	t := template.Must(template.ParseFiles("templates/survey.html"))
	if err := t.ExecuteTemplate(w, "survey.html", data); err != nil {
		log.Fatal(err)
	}
}

// SurveyRegistered ...
func SurveyRegistered(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/survey_registered.html"))
	if err := t.ExecuteTemplate(w, "survey_registered.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

func surveyFormValidator(formData map[string]string) map[string]error {
	e := make(map[string]error)

	userID := formData["userID"]
	count, err := SurveyRegisterLogic.CountByUserID(userID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(count)
	if count > 0 {
		e["userID"] = fmt.Errorf("回答済みです")
	}
	return e
}
