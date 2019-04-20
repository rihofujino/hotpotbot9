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
