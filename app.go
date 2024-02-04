package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"open-gsa/internal/database"
	opportunityRepo "open-gsa/internal/repository/opportunity"
	"os"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type LoginResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	database.RunSeeder()
	changeCwdDir()
	a.ctx = ctx
}

func changeCwdDir() {
	home, _ := os.UserHomeDir()
	newHomePath := filepath.Join(home, "OpenSamGov")
	_ = os.MkdirAll(newHomePath, os.ModePerm)
	_ = os.Chdir(newHomePath)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(username string, password string) LoginResponse {
	db := database.GetDbInstance()
	var user database.User
	result := db.First(&user, "username = ? AND password = ?", username, password)
	success := result.RowsAffected > 0

	return LoginResponse{Message: fmt.Sprintf("Welcome %s!", username), Result: success}
}

func (a *App) SearchOpportunities(keyword string, filters opportunityRepo.OpportunityFilter, saveAs bool) []database.Opportunity {
	result := opportunityRepo.Search(keyword, filters)

	if saveAs {
		var filters []runtime.FileFilter
		fileFilters := append(filters, runtime.FileFilter{
			DisplayName: "Comma-separate Values",
			Pattern:     "*.csv",
		})
		selection, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
			Title:                      "Save Report",
			DefaultFilename:            "Opportunities",
			Filters:                    fileFilters,
			ShowHiddenFiles:            false,
			CanCreateDirectories:       false,
			TreatPackagesAsDirectories: false,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("OpenFile: selection ", selection)
		downloadReport(result, selection)
	}

	return result
}

func (a *App) PullLatest() {
	opportunityRepo.PullLatest()
}

func downloadReport(data []database.Opportunity, filename string) {

	if !strings.HasSuffix(filename, ".csv") {
		filename = fmt.Sprintf("%s.csv", filename)
	}

	csvFile, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	_ = csvwriter.Write([]string{
		"NoticeID",
		"Title",
		"SolicitationNumber",
		"FullParentPathName",
		"FullParentPathCode",
		"PostedDate",
		"Type",
		"BaseType",
		"ArchiveType",
		"ArchiveDate",
		"ResponseDeadLine",
		"NaicsCode",
		"ClassificationCode",
		"Active",
		"Description",
		"OrganizationType",
		"UILink",
	})

	for _, op := range data {
		_ = csvwriter.Write([]string{
			op.NoticeID,
			op.Title,
			op.SolicitationNumber,
			op.FullParentPathName,
			op.FullParentPathCode,
			op.PostedDate,
			op.Type,
			op.BaseType,
			op.ArchiveType,
			op.ArchiveDate,
			op.ResponseDeadLine,
			op.NaicsCode,
			op.ClassificationCode,
			op.Active,
			op.Description,
			op.OrganizationType,
			op.UILink,
		})
	}
	csvwriter.Flush()
	csvFile.Close()

}

func (a *App) GetCwd() (string, error) {
	return os.Getwd()
}
