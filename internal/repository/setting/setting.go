package setting

import (
	"open-gsa/internal/database"
)

type SystemState struct {
	LastPull string `json:"lastPull"`
}

func AllSettings() []database.Setting {
	var settings []database.Setting
	db := database.GetDbInstance()
	db.Find(&settings)

	return settings
}

func SystemStatus() SystemState {
	var lastPull database.Setting
	db := database.GetDbInstance()
	db.Find(&lastPull, "key = ?", "last_pull")

	return SystemState{
		LastPull: lastPull.Value,
	}
}
