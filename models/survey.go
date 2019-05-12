package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/heroku/hotpotbot9/db"
)

type (
	//SurveyRegisterLogic ...
	SurveyRegisterLogic interface {
		CountByUserID(userID string) (int64, error)
		Save(formData map[string]string) error
	}

	//surveyRegisterLogicImpl ...
	surveyRegisterLogicImpl struct{}

	//Survey ...
	Survey struct {
		UserID       string         `db:"user_id"`
		Satisfaction int64          `db:"satisfaction"`
		Impression   sql.NullString `db:"satisfaction"`
		Theme        sql.NullString `db:"expect_theme"`
		CreatedAt    *time.Time     `db:"created_at"`
	}
)

// NewSurveyRegisterLogic ...
func NewSurveyRegisterLogic() SurveyRegisterLogic {
	return &surveyRegisterLogicImpl{}
}

//CountByUserID ...
func (p *surveyRegisterLogicImpl) CountByUserID(userID string) (int64, error) {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int64
	query := fmt.Sprintf("select count(*) from survey where user_id = '%s';", userID)
	row := db.QueryRow(query)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//Save ...
func (p *surveyRegisterLogicImpl) Save(formData map[string]string) error {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userID := formData["userID"]
	satisfaction := formData["satisfaction"]
	impression := formData["impression"]
	theme := formData["theme"]

	query := fmt.Sprintf("INSERT INTO survey (user_id, satisfaction, impression, expect_theme, created_at) VALUES ('%s', '%s', '%s', '%s', CURRENT_TIMESTAMP);", userID, satisfaction, impression, theme)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
