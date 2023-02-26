package tasks

import (
	"errors"
	"fmt"
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/repository"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"gorm.io/gorm"
)

func GenerateClientId(envPath string, emailOrUsername string, name string, description string) (string, string) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)

	// Initiate Database connection
	models.Initiate()

	user, err := repository.GetUserByUsernameOrEmail(models.DBConn, emailOrUsername)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User with email or password " + emailOrUsername + " not found")
		} else {
			fmt.Println("error: " + err.Error())
		}
		return "", ""
	}

	session, err := repository.GenerateClientIdAndClientSecret(models.DBConn, name, description, user, time.Now())
	if err != nil {
		fmt.Println("error: " + err.Error())
		return "", ""
	}

	return session.ClientID, session.ClientSecret
}
