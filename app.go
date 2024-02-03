package main

import (
	"context"
	"fmt"
	"open-gsa/internal/api/opportunities"
	"open-gsa/internal/database"
	"open-gsa/internal/repository"
)

// App struct
type App struct {
	ctx context.Context
}

type LoginResponse struct {
	Message string
	Result  bool
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

func (a *App) SearchOpportunities(dateFrom string, dateTo string) opportunities.SearchResult {
	result := opportunities.GetOpportunities(dateFrom, dateTo)
	return result
}

func (a *App) PullLatest() {
	repository.PullLatest()
}
