package repository

import (
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"gorm.io/gorm"
)

func GenerateClientIdAndClientSecret(tx *gorm.DB, name string, description string, user models.User, createdAt time.Time) (models.Oauth2Session, error) {
	newOauth2Session := models.Oauth2Session{
		UserId:       user.ID,
		Name:         name,
		Description:  description,
		ClientID:     core.GenerateRandomString(30),
		ClientSecret: core.GenerateRandomString(30),
		CreatedAt:    createdAt,
	}
	if err := tx.Save(&newOauth2Session).Error; err != nil {
		return newOauth2Session, err
	}
	return newOauth2Session, nil
}

func GetUserSession(tx *gorm.DB, user models.User) ([]models.Oauth2Session, error) {
	sessions := []models.Oauth2Session{}
	if err := tx.Where("user_id = ? ", user.ID).Find(&sessions).Error; err != nil {
		return sessions, err
	}
	return sessions, nil
}

// I Don't Know why this work !!!
// func GetUserSession(tx *gorm.DB, user models.User) ([]models.Oauth2Session, error) {
// 	sessions := []models.Oauth2Session{}
// 	if err := tx.Model(&models.Oauth2Session{}).Preload("User").Find(&sessions).Error; err != nil {
// 		return sessions, err
// 	}
// 	return sessions, nil
// }
