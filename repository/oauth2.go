package repository

import (
	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"gorm.io/gorm"
)

func GenerateCodeForUser(tx *gorm.DB, user models.User) (string, error) {
	code := core.GenerateRandomString(30)
	new_token := models.Oauth2Token{
		UserId: user.ID,
		Code:   code,
	}
	if err := tx.Create(&new_token).Error; err != nil {
		return "", err
	}
	return code, nil
}
