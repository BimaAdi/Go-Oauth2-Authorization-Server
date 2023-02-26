package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/migrations"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/routes"
	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MigrateAuthTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *MigrateAuthTestSuite) SetupSuite() {

	settings.InitiateSettings("../.env")
	models.Initiate()
	migrations.MigrateUp("../.env", "file://../migrations/migrations_files/")
	router := gin.Default()
	suite.router = routes.GetRoutes(router)
}

func (suite *MigrateAuthTestSuite) SetupTest() {
	models.ClearAllData()

}

func (suite *MigrateAuthTestSuite) TestLoginSuccess() {
	// Given
	timeZoneAsiaJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err.Error())
	}
	hashPasword, err := core.HashPassword("Fakepassword")
	if err != nil {
		panic(err.Error())
	}
	user_login := models.User{
		Email:       "test@test.com",
		Username:    "test",
		Password:    hashPasword,
		IsActive:    true,
		IsSuperuser: true,
		CreatedAt:   time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&user_login)

	// When
	var param = url.Values{}
	param.Set("username", "test")
	param.Set("password", "Fakepassword")
	var payload = bytes.NewBufferString(param.Encode())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/login", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 200, w.Code)
	jsonResponse := schemas.LoginResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")
}

func (suite *MigrateAuthTestSuite) TestLoginFailed() {
	// Given
	timeZoneAsiaJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err.Error())
	}
	hashPasword, err := core.HashPassword("Fakepassword")
	if err != nil {
		panic(err.Error())
	}
	user_login := models.User{
		Email:       "test@test.com",
		Username:    "test",
		Password:    hashPasword,
		IsActive:    true,
		IsSuperuser: true,
		CreatedAt:   time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&user_login)

	// When
	var param = url.Values{}
	param.Set("username", "test")
	param.Set("password", "wrong password")
	var payload = bytes.NewBufferString(param.Encode())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/login", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *MigrateAuthTestSuite) TestRegisterClientSuccess() {
	// Given
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
		IsSuperuser: true,
		CreatedAt:   time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&requestUser)
	token, err := core.GenerateJWTTokenFromUser(models.DBConn, requestUser)
	if err != nil {
		panic(err.Error())
	}

	// When
	requestJson := schemas.ClientRegiterRequest{
		Name:        "app",
		Description: "for my app",
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/register-client", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 201, w.Code)
	jsonResponse := schemas.ClientRegiterResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")
	session := models.Oauth2Session{}
	if err := models.DBConn.Where("user_id = ?", requestUser.ID).First(&session).Error; err != nil {
		suite.T().Error(err.Error())
	}
	assert.Equal(suite.T(), session.Name, "app")
	assert.Equal(suite.T(), session.Description, "for my app")
	assert.Equal(suite.T(), session.ClientID, jsonResponse.ClientId)
	assert.Equal(suite.T(), session.ClientSecret, jsonResponse.ClientSecret)
	assert.NotNil(suite.T(), session.CreatedAt)
}

func (suite *MigrateAuthTestSuite) TestRegisterClientForbidden() {
	// Given
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
	token, err := core.GenerateJWTTokenFromUser(models.DBConn, requestUser)
	if err != nil {
		panic(err.Error())
	}

	// When
	requestJson := schemas.ClientRegiterRequest{
		Name:        "app",
		Description: "for my app",
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/register-client", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 403, w.Code)
}

func (suite *MigrateAuthTestSuite) TestGetClientIdClientSecret() {
	// Given
	timeZoneAsiaJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err.Error())
	}
	hashPasword, err := core.HashPassword("Fakepassword")
	if err != nil {
		panic(err.Error())
	}
	userSessions := []models.Oauth2Session{
		{
			Name:         "app 1",
			Description:  "for app 1",
			ClientID:     "aaaabbbbccc",
			ClientSecret: "dddeeefff",
		},
		{
			Name:         "app 2",
			Description:  "for app 2",
			ClientID:     "dddeeefff",
			ClientSecret: "aaaabbbbccc",
		},
	}
	requestUser := models.User{
		Email:          "test@test.com",
		Username:       "test",
		Password:       hashPasword,
		IsActive:       true,
		IsSuperuser:    true,
		CreatedAt:      time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
		Oauth2Sessions: userSessions,
	}
	models.DBConn.Create(&requestUser)
	token, err := core.GenerateJWTTokenFromUser(models.DBConn, requestUser)
	if err != nil {
		panic(err.Error())
	}

	// When
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auth/client-id-client-secret/", nil)
	req.Header.Set("authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	// Expect
	suite.T().Log(w.Body)
	assert.Equal(suite.T(), 200, w.Code)
	jsonResponse := schemas.ArrayClientRegisterResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")
	for index, item := range userSessions {
		assert.Equal(suite.T(), item.Name, jsonResponse.Data[index].Name)
		assert.Equal(suite.T(), item.Description, jsonResponse.Data[index].Description)
		assert.Equal(suite.T(), item.ClientID, jsonResponse.Data[index].ClientId)
		assert.Equal(suite.T(), item.ClientSecret, jsonResponse.Data[index].ClientSecret)
	}
}

func (suite *MigrateAuthTestSuite) TearDownTest() {
	// models.ClearAllData()
}

func TestMigrateAuthTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateAuthTestSuite))
}
