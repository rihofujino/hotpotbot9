package models

import (
	"fmt"
	"log"

	"github.com/heroku/hotpotbot9/db"
)

type (
	//SurveyPostLogic ...
	SurveyPostLogic interface {
		Save(formData map[string]string) error
	}

	//surveyPostLogicImpl ...
	surveyPostLogicImpl struct{}
)

// NewSurveyPostLogic ...
func NewSurveyPostLogic() SurveyPostLogic {
	return &surveyPostLogicImpl{}
}

//Save ...
func (p *surveyPostLogicImpl) Save(formData map[string]string) error {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := formData["name"]
	satisfaction := formData["satisfaction"]
	impression := formData["impression"]
	theme := formData["theme"]

	query := fmt.Sprintf("INSERT INTO survey (name, satisfaction, impression, expect_theme, created_at) VALUES ('%s', '%s', '%s', '%s', CURRENT_TIMESTAMP);", name, satisfaction, impression, theme)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
