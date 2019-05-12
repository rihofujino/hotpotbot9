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
	//PersonalInfoLogic ...
	PersonalInfoLogic models.PersonalInfoLogic
)

// PersonalInfoEdit ...
func PersonalInfoEdit(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("this is formdata")
		fmt.Println(r.Form)
		formData := map[string]string{
			"userID":  r.Form["user_id"][0],
			"name":    r.Form["name"][0],
			"company": r.Form["company"][0],
			"jobType": r.Form["jobType"][0],
		}
		log.Print(formData)

		if e := personalInfoFormValidator(formData); e != nil {
			data["formErrors"] = e
		} else {
			err := PersonalInfoLogic.Save(formData)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("location", "/entry")
			w.WriteHeader(http.StatusSeeOther)
			return
		}
	}

	t := template.Must(template.ParseFiles("templates/personal_info.html"))
	if err := t.ExecuteTemplate(w, "personal_info.html", data); err != nil {
		log.Fatal(err)
	}
}

// Entry ...
func Entry(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/entry.html"))
	if err := t.ExecuteTemplate(w, "entry.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

func personalInfoFormValidator(formData map[string]string) map[string]error {
	e := make(map[string]error)

	userID := formData["userID"]
	pi, err := PersonalInfoLogic.GetByUserID(userID)
	if err != nil {
		log.Fatal(err)
	}
	if pi != nil {
		e["userID"] = fmt.Errorf("出席登録済みです")
	}
	return e
}
