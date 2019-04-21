package models

import (
	"fmt"
	"log"

	"github.com/heroku/gyozabot/db"
)

//Member ...
type Member struct {
	ID      int
	Name    string
	Company string
	JobType int
}

//Save ...
func Save(formData map[string]string) error {
	db, err := db.OpenMysql()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := formData["name"]
	company := formData["company"]
	jobType := formData["jobType"]

	query := fmt.Sprintf("insert into gyozabot_db.member (name, company, jobtype) values ('%s', '%s', %s)", name, company, jobType)
	log.Print(query)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
