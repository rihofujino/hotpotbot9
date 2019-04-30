package models

import (
	"fmt"
	"log"

	"github.com/heroku/hotpotbot9/db"
)

type (
	//PersonalInfoLogic ...
	PersonalInfoLogic interface {
		Save(formData map[string]string) error
	}

	//personalInfoLogicImpl ...
	personalInfoLogicImpl struct{}
)

// NewPersonalInfoLogic ...
func NewPersonalInfoLogic() PersonalInfoLogic {
	return &personalInfoLogicImpl{}
}

//Save ...
func (p *personalInfoLogicImpl) Save(formData map[string]string) error {
	db, err := db.OpenPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := formData["name"]
	company := formData["company"]
	jobType := formData["jobType"]

	query := fmt.Sprintf("INSERT INTO member (name, company, jobType, created_at) VALUES ('%s', '%s', %s, CURRENT_TIMESTAMP);", name, company, jobType)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
