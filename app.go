package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"open-gsa/internal/database"
	opportunityRepo "open-gsa/internal/repository/opportunity"
	"open-gsa/internal/repository/setting"
	"os"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type SettingRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SettingResponse struct {
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

func (a *App) SearchOpportunities(keyword string, filters opportunityRepo.OpportunityFilter, saveAs bool) opportunityRepo.PaginatedResult {
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
		downloadReport(result.Data, selection)
	}

	return result
}

func (a *App) PullLatest() Response {
	msg, err := opportunityRepo.PullLatest()
	success := true
	if err != nil {
		msg = err.Error()
		success = false
	}
	return Response{Message: msg, Success: success}
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

func (a *App) GetOpportunityTypes() []string {
	return opportunityRepo.GetTypes()
}

func (a *App) GetOpportunityNaicsCodes() []string {
	return opportunityRepo.GetNaicsCodes()
}

func (a *App) GetAllSettings() []database.Setting {
	return setting.AllSettings()
}

func (a *App) GetSystemStatus() setting.SystemState {
	return setting.SystemStatus()
}

func (a *App) UpdateSettings(settings []SettingRequest) SettingResponse {
	db := database.GetDbInstance()

	for i := 0; i < len(settings); i++ {
		rSetting := settings[i]

		settingModel := database.Setting{Value: rSetting.Value}
		result := db.FirstOrCreate(&settingModel, database.Setting{Key: rSetting.Key})

		if result.RowsAffected == 0 {
			settingModel.Value = rSetting.Value
			db.Save(&settingModel)
		}
	}

	_, err1 := runtime.MessageDialog(
		a.ctx,
		runtime.MessageDialogOptions{
			Title:   "Update settings",
			Message: "System setting updated",
		},
	)
	if err1 != nil {
		log.Fatal(err1)
	}

	err2 := beeep.Alert("OpenGsaApp", "System setting updated.", "static/favicon.png")
	if err2 != nil {
		log.Fatal(err2)
	}

	return SettingResponse{Message: "Settings updated.", Result: true}
}
