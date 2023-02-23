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

func (suite *MigrateAuthTestSuite) TearDownTest() {
	models.ClearAllData()
}

func TestMigrateAuthTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateAuthTestSuite))
}
