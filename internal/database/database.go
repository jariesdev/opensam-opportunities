package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Opportunity struct {
	gorm.Model
	NoticeID           string
	Title              string
	SolicitationNumber string
	FullParentPathName string
	FullParentPathCode string
	PostedDate         string
	Type               string
	BaseType           string
	ArchiveType        string
	ArchiveDate        string
	ResponseDeadLine   string
	NaicsCode          string
	ClassificationCode string
	Active             string
	Description        string
	OrganizationType   string
	UILink             string
}

type Setting struct {
	gorm.Model
	Key   string
	Value string
}

func GetDbInstance() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("database.db"))

	// Migrate the schema
	err := db.AutoMigrate(&Opportunity{}, &Setting{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
