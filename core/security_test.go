package core_test

import (
	"testing"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestGenerateJWTToken(t *testing.T) {
	settings.InitiateSettings("../.env")
	token, err := core.GenerateJWTToken("aaaaa-bbbbb-ccccc", "bimaadi419@gmail.com")
	assert.Nil(t, err)
	id, email, err := core.GetPayloadFromJWTToken(token)
	assert.Equal(t, "aaaaa-bbbbb-ccccc", id)
	assert.Equal(t, "bimaadi419@gmail.com", email)
}

func TestExpiredJWTToken(t *testing.T) {
	settings.InitiateSettings("../.env")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJpbWFhZGk0MTlAZ21haWwuY29tIiwiZXhwIjoxNjc2MjYwMDg0LCJpYXQiOjE2NzYyNjE4ODQsImlkIjoiYWFhYWEtYmJiYmItY2NjY2MifQ.8VwyTsDIfUf8WT1A9MxFHMqWn-ZVi9RZACvA2y5K9WE"
	_, _, err := core.GetPayloadFromJWTToken(token)
	assert.NotNil(t, err)
}

type MigrateTestSuite struct {
	suite.Suite
}

func (suite *MigrateTestSuite) SetupSuite() {
	settings.InitiateSettings("../.env")
	models.Initiate()
	models.AutoMigrate()
}

func (suite *MigrateTestSuite) SetupTest() {
	models.ClearAllData()
}

// ==========================================

func (suite *MigrateTestSuite) TestGetUserFromJwtToken() {
	// TODO add function get user from jwt token
}

// ==========================================

func (suite *MigrateTestSuite) TearDownTest() {
	models.ClearAllData()
}

func TestMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateTestSuite))
}
