package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/heroku/hotpotbot9/db"
)

type (
	//PersonalInfoLogic ...
	PersonalInfoLogic interface {
		CountByUserID(userID string) (int64, error)
		Save(formData map[string]string) error
	}

	//personalInfoLogicImpl ...
	personalInfoLogicImpl struct{}

	//PersonalInfo ...
	PersonalInfo struct {
		ID        int64          `db:"id"`
		UserID    string         `db:"user_id"`
		Name      string         `db:"name"`
		Company   sql.NullString `db:"company"`
		JobType   int64          `db:"job_type"`
		CreatedAt *time.Time     `db:"created_at"`
	}
)

// NewPersonalInfoLogic ...
func NewPersonalInfoLogic() PersonalInfoLogic {
	return &personalInfoLogicImpl{}
}

//CountByUserID ...
func (p *personalInfoLogicImpl) CountByUserID(userID string) (int64, error) {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int64
	query := fmt.Sprintf("select count(*) from member where user_id = '%s';", userID)
	row := db.QueryRow(query)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

//Save ...
func (p *personalInfoLogicImpl) Save(formData map[string]string) error {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userID := formData["userID"]
	name := formData["name"]
	company := formData["company"]
	jobType := formData["jobType"]

	query := fmt.Sprintf("INSERT INTO member (user_id, name, company, jobType, created_at) VALUES ('%s', '%s', '%s', %s, CURRENT_TIMESTAMP);", userID, name, company, jobType)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
