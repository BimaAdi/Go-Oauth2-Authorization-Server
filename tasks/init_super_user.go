package tasks

import (
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/repository"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
)

func CreateSuperUser(envPath string, email string, username string, password string) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)

	// Initiate Database connection
	models.Initiate()

	now := time.Now()
	repository.CreateUser(username, email, password, true, true, now, &now)
}
