package tasks_test

import (
	"testing"
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/migrations"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/BimaAdi/Oauth2AuthorizationServer/tasks"
	"github.com/stretchr/testify/assert"
)

func TestGenerateClientIDSuccess(t *testing.T) {
	// Given
	settings.InitiateSettings("../.env")
	models.Initiate()
	migrations.MigrateUp("../.env", "file://../migrations/migrations_files/")
	models.ClearAllData()
	timeZoneAsiaJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err.Error())
	}
	hashPasword, err := core.HashPassword("Fakepassword")
	if err != nil {
		panic(err.Error())
	}
	requestUser := models.User{
		Email:       "test@test.com",
		Username:    "test",
		Password:    hashPasword,
		IsActive:    true,
		IsSuperuser: false,
		CreatedAt:   time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&requestUser)

	// When
	clientId, clientSecret := tasks.GenerateClientId("../.env", "test")

	// Expect
	assert.NotEqual(t, "", clientId)
	assert.NotEqual(t, "", clientSecret)
	session := models.Oauth2Session{}
	if err := models.DBConn.Where("user_id = ?", requestUser.ID).First(&session).Error; err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, session.ClientID, clientId)
	assert.Equal(t, session.ClientSecret, clientSecret)
	assert.NotNil(t, session.CreatedAt)
}

func TestGenerateClientIDNotFound(t *testing.T) {
	// Given
	settings.InitiateSettings("../.env")
	models.Initiate()
	migrations.MigrateUp("../.env", "file://../migrations/migrations_files/")
	models.ClearAllData()

	// When
	clientId, clientSecret := tasks.GenerateClientId("../.env", "hello")

	// Expect
	assert.Equal(t, "", clientId)
	assert.Equal(t, "", clientSecret)
}
