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
		GetByUserID(userID string) (*PersonalInfo, error)
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

//GetByUserID ...
func (p *personalInfoLogicImpl) GetByUserID(userID string) (*PersonalInfo, error) {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("select id, user_id, name, company, jobtype from member where user_id='%s'", userID)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var pi PersonalInfo
	for rows.Next() {
		err := rows.Scan(&pi.ID, &pi.UserID, &pi.Name, &pi.Company, &pi.JobType, &pi.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &pi, nil
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
