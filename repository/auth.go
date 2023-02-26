package repository

import (
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"gorm.io/gorm"
)

func GenerateClientIdAndClientSecret(tx *gorm.DB, user models.User, createdAt time.Time) (models.Oauth2Session, error) {
	newOauth2Session := models.Oauth2Session{
		UserId:       user.ID,
		ClientID:     core.GenerateRandomString(30),
		ClientSecret: core.GenerateRandomString(30),
		CreatedAt:    createdAt,
	}
	if err := tx.Save(&newOauth2Session).Error; err != nil {
		return newOauth2Session, err
	}
	return newOauth2Session, nil
}
