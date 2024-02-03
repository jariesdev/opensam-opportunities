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

type User struct {
	gorm.Model
	Username string
	Password string
}

type Setting struct {
	gorm.Model
	Key   string
	Value string
}

func GetDbInstance() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("database.sqlite"))

	// Migrate the schema
	err := db.AutoMigrate(&User{}, &Opportunity{}, &Setting{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func RunSeeder() {
	db := GetDbInstance()
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		db.Create(&User{Username: "admin", Password: "password"})
	}
}
