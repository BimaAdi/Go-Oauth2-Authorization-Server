package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
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
	"gorm.io/gorm"
)

type MigrateOauth2TestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *MigrateOauth2TestSuite) SetupSuite() {
	settings.InitiateSettings("../.env")
	models.Initiate()
	migrations.MigrateUp("../.env", "file://../migrations/migrations_files/")
	router := gin.Default()
	suite.router = routes.GetRoutes(router)
}

func (suite *MigrateOauth2TestSuite) SetupTest() {
	models.ClearAllData()
}

func (suite *MigrateOauth2TestSuite) TestOauth2Authorize() {
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
	client_id := core.GenerateRandomString(30)
	client_secret := core.GenerateRandomString(30)
	oauth_2_session := models.Oauth2Session{
		UserId:       requestUser.ID,
		Name:         "test app",
		Description:  "for test app application",
		ClientID:     client_id,
		ClientSecret: client_secret,
		IsActive:     true,
		CreatedAt:    time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&oauth_2_session)
	state := core.GenerateRandomString(30)

	// When
	callback_url := "http://localhost:8000/callbacks"
	requestJson := schemas.OauthLoginJsonRequest{
		Username:     "test",
		Password:     "Fakepassword",
		ResponseType: "code",
		ClientId:     client_id,
		RedirectUri:  callback_url,
		Scope:        "oauth",
		State:        state,
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/oauth/authorize/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 302, w.Code)
	token := models.Oauth2Token{}
	if err := models.DBConn.Where("user_id = ?", requestUser.ID).First(&token).Error; err != nil {
		panic(err.Error())
	}
	// Check redirect uri
	assert.Equal(suite.T(), callback_url+"?state="+state+"&code="+token.Code, w.Header()["Location"][0])

}

func (suite *MigrateOauth2TestSuite) TestOaut2TokenSuccess() {
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
	client_id := core.GenerateRandomString(30)
	client_secret := core.GenerateRandomString(30)
	oauth_2_session := models.Oauth2Session{
		UserId:       requestUser.ID,
		Name:         "test app",
		Description:  "for test app application",
		ClientID:     client_id,
		ClientSecret: client_secret,
		IsActive:     true,
		CreatedAt:    time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&oauth_2_session)
	code := core.GenerateRandomString(30)
	oauth_2_token := models.Oauth2Token{
		ID:        client_id,
		UserId:    requestUser.ID,
		Code:      code,
		CreatedAt: time.Time{},
	}
	models.DBConn.Create(&oauth_2_token)

	// When
	requestJson := schemas.Oauth2TokenJsonRequest{
		GrantType:    "authorization_code",
		ClientId:     client_id,
		ClientSecret: client_secret,
		RedirectUri:  "http://localhost:8000/callbacks/",
		Code:         code,
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/oauth/token/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	// Expect
	suite.T().Log(w.Body)
	assert.Equal(suite.T(), 200, w.Code)
	existing_oauth_token := models.Oauth2Token{}
	err = models.DBConn.Where("code = ?", code).First(&existing_oauth_token).Error
	assert.True(suite.T(), errors.Is(err, gorm.ErrRecordNotFound))
}

func (suite *MigrateOauth2TestSuite) TestOaut2TokeninvalideCode() {
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
	client_id := core.GenerateRandomString(30)
	client_secret := core.GenerateRandomString(30)
	oauth_2_session := models.Oauth2Session{
		UserId:       requestUser.ID,
		Name:         "test app",
		Description:  "for test app application",
		ClientID:     client_id,
		ClientSecret: client_secret,
		IsActive:     true,
		CreatedAt:    time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
	}
	models.DBConn.Create(&oauth_2_session)
	code := core.GenerateRandomString(30)
	oauth_2_token := models.Oauth2Token{
		ID:        client_id,
		UserId:    requestUser.ID,
		Code:      code,
		CreatedAt: time.Time{},
	}
	models.DBConn.Create(&oauth_2_token)

	// When
	requestJson := schemas.Oauth2TokenJsonRequest{
		GrantType:    "authorization_code",
		ClientId:     client_id,
		ClientSecret: client_secret,
		RedirectUri:  "http://localhost:8000/callbacks/",
		Code:         "aaaaabbbbccccdddd",
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/oauth/token/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	// Expect
	suite.T().Log(w.Body)
	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *MigrateOauth2TestSuite) TearDownTest() {
	models.ClearAllData()
}

func TestMigrateOauth2TestSuite(t *testing.T) {
	suite.Run(t, new(MigrateOauth2TestSuite))
}
