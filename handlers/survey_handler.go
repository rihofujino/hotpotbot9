package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// Survey ...
func Survey(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("templates/survey.html"))
	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "survey.html", time.Now()); err != nil {
		log.Fatal(err)
	}
}
