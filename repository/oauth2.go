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

func GetDetailOauth2SessionByClientId(tx *gorm.DB, client_id string) (models.Oauth2Session, error) {
	find_oauth2_session := models.Oauth2Session{}
	if err := tx.Where("client_id = ? AND is_active = ?", client_id, true).First(&find_oauth2_session).Error; err != nil {
		return models.Oauth2Session{}, err
	}
	return find_oauth2_session, nil
}

func GetDetailOauth2SessionByClientIdAndClientSecret(tx *gorm.DB, client_id string, client_secret string) (models.Oauth2Session, error) {
	find_oauth2_session := models.Oauth2Session{}
	if err := tx.Where(
		"client_id = ? AND client_secret = ? AND is_active = ?", client_id, client_secret, true,
	).First(&find_oauth2_session).Error; err != nil {
		return models.Oauth2Session{}, err
	}
	return find_oauth2_session, nil
}

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

func GetOauthTokenByCode(tx *gorm.DB, code string) (models.Oauth2Token, error) {
	find_oauth2_token := models.Oauth2Token{}
	if err := tx.Where(
		"code = ?", code,
	).First(&find_oauth2_token).Error; err != nil {
		return find_oauth2_token, err
	}
	return find_oauth2_token, nil
}

func DeleteOauthToken(tx *gorm.DB, oauth2token *models.Oauth2Token) error {
	if err := tx.Delete(oauth2token).Error; err != nil {
		return err
	}
	return nil
}
