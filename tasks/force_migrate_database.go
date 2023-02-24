package tasks

import (
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
)

func ForceMigrate(envPath string) {
	settings.InitiateSettings(envPath)
	models.Initiate()
	models.AutoMigrate()
	// 20230224015024
	// models.DBConn.Exec("SELECT ")
}

func ForceRollback(envPath string) {
	settings.InitiateSettings(envPath)
	models.Initiate()
	models.AutoRollback()
	models.DBConn.Exec("DELETE FROM public.schema_migrations")
}
