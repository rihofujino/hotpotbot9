package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/heroku/gyozabot/db"
)

//Member ...
type Member struct {
	ID      int
	Name    string
	Company string
	JobType int
}

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
		err := save(formData)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.ExecuteTemplate(w, "entry.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}

func save(formData map[string]string) error {
	db, err := db.OpenMysql()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := formData["name"]
	company := formData["company"]
	jobType := formData["jobtype"]

	query := fmt.Sprintf("insert into member values (%s, %s, %s)", name, company, jobType)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
