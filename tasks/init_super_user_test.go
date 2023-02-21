package tasks_test

import (
	"testing"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/BimaAdi/Oauth2AuthorizationServer/tasks"
	"github.com/stretchr/testify/assert"
)

func TestCreateSuperUser(t *testing.T) {
	// Given
	settings.InitiateSettings("../.env")
	models.Initiate()
	models.AutoMigrate()
	models.ClearAllData()

	// When
	tasks.CreateSuperUser("../.env", "test@local.com", "test", "password")

	// Expect
	createdUser := models.User{}
	err := models.DBConn.Where("email = ? AND username = ?", "test@local.com", "test").First(&createdUser).Error
	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
	assert.True(t, core.CheckPasswordHash("password", createdUser.Password))
}
