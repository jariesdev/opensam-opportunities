package database

import (
	"fmt"
	"log"

	// requires CGO_ENABLED=1
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

type Opportunity struct {
	gorm.Model
	NoticeID                  string           `json:"noticeID"`
	Title                     string           `json:"title"`
	SolicitationNumber        string           `json:"solicitationNumber"`
	FullParentPathName        string           `json:"fullParentPathName"`
	FullParentPathCode        string           `json:"fullParentPathCode"`
	PostedDate                string           `json:"postedDate"`
	Type                      string           `json:"type"`
	BaseType                  string           `json:"baseType"`
	ArchiveType               string           `json:"archiveType"`
	ArchiveDate               string           `json:"archiveDate"`
	TypeOfSetAsideDescription string           `json:"typeOfSetAsideDescription"`
	TypeOfSetAside            string           `json:"typeOfSetAside"`
	ResponseDeadLine          string           `json:"responseDeadLine"`
	NaicsCode                 string           `json:"naicsCode"`
	NaicsCodes                string           `json:"naicsCodes"`
	ClassificationCode        string           `json:"classificationCode"`
	Active                    string           `json:"active"`
	PointOfContact            []PointOfContact `json:"pointOfContact"`
	Description               string           `json:"description"`
	OrganizationType          string           `json:"organizationType"`
	OfficeAddress             string           `json:"officeAddress"`
	AdditionalInfoLink        string           `json:"additionalInfoLink"`
	UILink                    string           `json:"uiLink"`
	Links                     []Link           `json:"links"`
	ResourceLinks             string           `json:"resourceLinks"`
}

type Link struct {
	gorm.Model
	Rel           string `json:"rel"`
	Href          string `json:"href"`
	OpportunityID uint
}

type PointOfContact struct {
	gorm.Model
	Fax           string `json:"fax"`
	Type          string `json:"type"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Title         string `json:"title"`
	FullName      string `json:"fullName"`
	OpportunityID uint
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
	path, pathErr := os.Getwd()
	if pathErr != nil {
		panic(pathErr)
	}

	//path, pathErr := os.Executable()
	//if pathErr != nil {
	//	log.Println(pathErr)
	//}
	//
	//// If the file doesn't exist, create it, or append to the file
	_, errX := os.OpenFile("database.sqlite", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errX != nil {
		panic(errX)
	}
	//f.Close()

	fmt.Printf("CWD: %s ", path)

	dbName := os.Getenv("OPENSAM_DB")
	if dbName == "" {
		dbName = "database.sqlite"
	}

	db, dbErr := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if dbErr != nil {
		fmt.Print(dbErr)
		panic(dbErr)
	}

	// Migrate the schema
	err := db.AutoMigrate(&User{}, &Opportunity{}, &PointOfContact{}, &Link{}, &Setting{})
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
