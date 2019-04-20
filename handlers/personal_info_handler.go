package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// PersonalInfo ...
func PersonalInfo(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("templates/personal_info.html"))
	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "personal_info.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}

// Entry ...
func Entry(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("templates/entry.html"))
	// テンプレートを描画

	if r.Method == "POST" {
		r.ParseForm()
		formData := map[string]string{
			"name":    r.Form["name"][0],
			"company": r.Form["company"][0],
			"jobType": r.Form["jobType"][0],
		}
		log.Print(formData)
		if err := t.ExecuteTemplate(w, "entry.html", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}
