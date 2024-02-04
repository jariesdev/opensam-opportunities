package database

import (
	// requires CGO_ENABLED=1
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

type Opportunity struct {
	gorm.Model
	NoticeID                  string `json:"noticeID,omitempty"`
	Title                     string `json:"title,omitempty"`
	SolicitationNumber        string `json:"solicitationNumber"`
	FullParentPathName        string `json:"fullParentPathName,omitempty"`
	FullParentPathCode        string `json:"fullParentPathCode,omitempty"`
	PostedDate                string `json:"postedDate,omitempty"`
	Type                      string `json:"type,omitempty"`
	BaseType                  string `json:"baseType,omitempty"`
	ArchiveType               string `json:"archiveType,omitempty"`
	ArchiveDate               string `json:"archiveDate,omitempty"`
	TypeOfSetAsideDescription string `json:"typeOfSetAsideDescription,omitempty"`
	TypeOfSetAside            string `json:"typeOfSetAside,omitempty"`
	ResponseDeadLine          string `json:"responseDeadLine,omitempty"`
	NaicsCode                 string `json:"naicsCode,omitempty"`
	ClassificationCode        string `json:"classificationCode,omitempty"`
	Active                    string `json:"active,omitempty"`
	Description               string `json:"description,omitempty"`
	OrganizationType          string `json:"organizationType,omitempty"`
	UILink                    string `json:"uiLink,omitempty"`
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
