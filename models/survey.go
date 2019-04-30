package models

import (
	"fmt"
	"log"
	"time"

	"github.com/heroku/hotpotbot9/db"
)

type (
	//Survey ...
	Survey struct {
		Name         string
		Satisfaction int
		Impression   string
		ExpectTheme  string
		CreatedAt    time.Time
	}

	//PostSurvey ...
	PostSurvey interface {
		Save(formData map[string]string) error
	}

	//PostSurveyImpl ...
	PostSurveyImpl struct{}
)

//Save ...
func (p *PostSurveyImpl) Save(formData map[string]string) error {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := formData["name"]
	satisfaction := formData["satisfaction"]
	impression := formData["impression"]
	theme := formData["theme"]

	query := fmt.Sprintf("INSERT INTO survey (name, satisfaction, impression, theme, created_at) VALUES ('%s', '%s', '%s', '%s', CURRENT_TIMESTAMP)", name, satisfaction, impression, theme)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
