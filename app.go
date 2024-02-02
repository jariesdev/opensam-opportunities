package main

import (
	"context"
	"fmt"
	"open-gsa/internal/api/opportunities"
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
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(username string, password string) LoginResponse {
	success := username == "admin" && password == "password"
	return LoginResponse{Message: fmt.Sprintf("Welcome %s!", username), Result: success}
}

func (a *App) SearchOpportunities(dateFrom string, dateTo string) *opportunities.SearchResult {
	result := opportunities.GetOpportunities2(dateFrom, dateTo)
	fmt.Print(result)
	return result
}
