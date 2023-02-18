package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/routes"
	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MigrateTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *MigrateTestSuite) SetupSuite() {
	settings.InitiateSettings("../.env")
	models.Initiate()
	models.AutoMigrate()
	router := gin.Default()
	suite.router = routes.GetRoutes(router)
}

func (suite *MigrateTestSuite) SetupTest() {
	models.ClearAllData()
}

// ==========================================

func (suite *MigrateTestSuite) TestGetPaginateUser() {
	// Given
	timeZoneAsiaJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err.Error())
	}
	users := []models.User{
		{
			Email:       "a@test.com",
			Username:    "a",
			Password:    "Fakepassword",
			IsActive:    true,
			IsSuperuser: true,
			CreatedAt:   time.Date(2022, 10, 5, 10, 0, 0, 0, timeZoneAsiaJakarta),
		},
		{
			Email:       "b@test.com",
			Username:    "b",
			Password:    "Fakepassword",
			IsActive:    true,
			IsSuperuser: true,
			CreatedAt:   time.Date(2022, 10, 4, 10, 0, 0, 0, timeZoneAsiaJakarta),
		},
		{
			Email:       "c@test.com",
			Username:    "c",
			Password:    "Fakepassword",
			IsActive:    true,
			IsSuperuser: true,
			CreatedAt:   time.Date(2022, 10, 3, 10, 0, 0, 0, timeZoneAsiaJakarta),
		},
	}
	models.DBConn.Create(&users)
	request_user := users[0]
	token, err := core.GenerateJWTTokenFromUser(models.DBConn, request_user)
	if err != nil {
		panic(err.Error())
	}

	// When
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/?page=1&page_size=2", nil)
	req.Header.Set("authorization", "Bearer "+token)
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 200, w.Code)
	jsonResponse := schemas.UserPaginateResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")
	assert.Equal(suite.T(), 3, jsonResponse.Counts)
	assert.Equal(suite.T(), 2, jsonResponse.PageCount)
	assert.Equal(suite.T(), 1, jsonResponse.Page)
	assert.Equal(suite.T(), 2, jsonResponse.PageSize)
	assert.Len(suite.T(), jsonResponse.Results, 2)
	for i := 0; i < 2; i++ {
		assert.Equal(suite.T(), users[i].ID, jsonResponse.Results[i].Id)
		assert.Equal(suite.T(), users[i].Username, jsonResponse.Results[i].Username)
		assert.Equal(suite.T(), users[i].Email, jsonResponse.Results[i].Email)
		assert.Equal(suite.T(), users[i].IsActive, jsonResponse.Results[i].IsActive)
	}

	// When 2
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/user/?page=2&page_size=2", nil)
	req2.Header.Set("authorization", "Bearer "+token)
	suite.router.ServeHTTP(w2, req2)

	// Expect 2
	assert.Equal(suite.T(), 200, w2.Code)
	jsonResponse2 := schemas.UserPaginateResponse{}
	err = json.Unmarshal(w2.Body.Bytes(), &jsonResponse2)
	assert.Nil(suite.T(), err, "Invalid response json")
	assert.Equal(suite.T(), 3, jsonResponse2.Counts)
	assert.Equal(suite.T(), 2, jsonResponse2.PageCount)
	assert.Equal(suite.T(), 2, jsonResponse2.Page)
	assert.Equal(suite.T(), 2, jsonResponse2.PageSize)
	assert.Len(suite.T(), jsonResponse2.Results, 1)
	expect_2 := users[2:3]
	for i := 0; i < 1; i++ {
		assert.Equal(suite.T(), expect_2[i].ID, jsonResponse2.Results[i].Id)
		assert.Equal(suite.T(), expect_2[i].Username, jsonResponse2.Results[i].Username)
		assert.Equal(suite.T(), expect_2[i].Email, jsonResponse2.Results[i].Email)
		assert.Equal(suite.T(), expect_2[i].IsActive, jsonResponse2.Results[i].IsActive)
	}
}

func (suite *MigrateTestSuite) TestGetDetailUser() {
	// Given
	givenW := httptest.NewRecorder()
	requestJson := schemas.UserCreateRequest{
		Username:    "test",
		Password:    "testpassword",
		Email:       "test@example.com",
		IsActive:    true,
		IsSuperuser: true,
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(givenW, req)
	assert.Equal(suite.T(), 201, givenW.Code)
	givenJsonResponse := schemas.UserCreateResponse{}
	err := json.Unmarshal(givenW.Body.Bytes(), &givenJsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")

	// When 1
	w1 := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/"+givenJsonResponse.Id, nil)
	suite.router.ServeHTTP(w1, req)

	// Expect 1
	assert.Equal(suite.T(), 200, w1.Code)
	jsonResponse1 := schemas.UserDetailResponse{}
	err = json.Unmarshal(w1.Body.Bytes(), &jsonResponse1)
	assert.Nil(suite.T(), err, "Invalid response json")

	// When 2
	w2 := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/aaaa-bbbbb-ccccc-ddddd", nil)
	suite.router.ServeHTTP(w2, req)

	// Expect 2
	assert.Equal(suite.T(), 404, w2.Code)
	jsonResponse2 := schemas.UserDetailResponse{}
	err = json.Unmarshal(w2.Body.Bytes(), &jsonResponse2)
	assert.Nil(suite.T(), err, "Invalid response json")
}

func (suite *MigrateTestSuite) TestCreateUser() {
	// Given
	suite.T().Log("test create user")

	// When
	w := httptest.NewRecorder()
	requestJson := schemas.UserCreateRequest{
		Username:    "test",
		Password:    "testpassword",
		Email:       "test@example.com",
		IsActive:    true,
		IsSuperuser: true,
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	// Expect
	assert.Equal(suite.T(), 201, w.Code)
	jsonResponse := schemas.UserCreateResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")

	createdUser := models.User{}
	err = models.DBConn.Where("id = ?", jsonResponse.Id).First(&createdUser).Error
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), createdUser.ID)
	assert.Equal(suite.T(), requestJson.Email, createdUser.Email)
	assert.Equal(suite.T(), requestJson.Username, createdUser.Username)
	assert.Equal(suite.T(), requestJson.IsActive, createdUser.IsActive)
	assert.Equal(suite.T(), requestJson.IsSuperuser, createdUser.IsSuperuser)
	assert.NotNil(suite.T(), createdUser.CreatedAt)
}

func (suite *MigrateTestSuite) TestUpdateUser() {
	// Given
	givenW := httptest.NewRecorder()
	requestJson := schemas.UserCreateRequest{
		Username:    "test",
		Password:    "testpassword",
		Email:       "test@example.com",
		IsActive:    true,
		IsSuperuser: true,
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(givenW, req)
	assert.Equal(suite.T(), 201, givenW.Code)
	givenJsonResponse := schemas.UserCreateResponse{}
	err := json.Unmarshal(givenW.Body.Bytes(), &givenJsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")

	// When 1
	w1 := httptest.NewRecorder()
	password := "testpassword"
	requestJson1 := schemas.UserUpdateRequest{
		Username:    "test",
		Password:    &password,
		Email:       "test@example.com",
		IsActive:    true,
		IsSuperuser: true,
	}
	requestJsonByte1, _ := json.Marshal(requestJson1)
	req, _ = http.NewRequest("PUT", "/user/"+givenJsonResponse.Id, bytes.NewBuffer(requestJsonByte1))
	suite.router.ServeHTTP(w1, req)

	// Expect 1
	assert.Equal(suite.T(), 200, w1.Code)
	jsonResponse1 := schemas.UserUpdateResponse{}
	err = json.Unmarshal(w1.Body.Bytes(), &jsonResponse1)
	assert.Nil(suite.T(), err, "Invalid response json")

	// When 2
	w2 := httptest.NewRecorder()
	password2 := "testpassword2"
	requestJson2 := schemas.UserUpdateRequest{
		Username:    "test",
		Password:    &password2,
		Email:       "test@example.com",
		IsActive:    true,
		IsSuperuser: true,
	}
	requestJsonByte2, _ := json.Marshal(requestJson2)
	req, _ = http.NewRequest("PUT", "/user/aaaaa-bbbbb-ccccc", bytes.NewBuffer(requestJsonByte2))
	suite.router.ServeHTTP(w2, req)

	// Expect 2
	assert.Equal(suite.T(), 404, w2.Code)
	jsonResponse2 := schemas.NotFoundResponse{}
	err = json.Unmarshal(w1.Body.Bytes(), &jsonResponse2)
	assert.Nil(suite.T(), err, "Invalid response json")
}

func (suite *MigrateTestSuite) TestDeleteUser() {
	// Given
	w := httptest.NewRecorder()
	requestJson := schemas.UserCreateRequest{
		Username:    "test",
		Password:    "testpassword",
		Email:       "test@example.com",
		IsActive:    true,
		IsSuperuser: true,
	}
	requestJsonByte, _ := json.Marshal(requestJson)
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(requestJsonByte))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 201, w.Code)
	jsonResponse := schemas.UserCreateResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Nil(suite.T(), err, "Invalid response json")

	// When 1
	w1 := httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/user/"+jsonResponse.Id, nil)
	suite.router.ServeHTTP(w1, req)

	// Expect 1
	assert.Equal(suite.T(), 204, w1.Code)

	// When 2
	w2 := httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/user/"+jsonResponse.Id, nil)
	suite.router.ServeHTTP(w2, req)

	// Expect 2
	assert.Equal(suite.T(), 404, w2.Code)

	// When 3
	w3 := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/user/"+jsonResponse.Id, nil)
	suite.router.ServeHTTP(w3, req)

	// Expect 3
	assert.Equal(suite.T(), 404, w3.Code)
}

// ==========================================

func (suite *MigrateTestSuite) TearDownTest() {
	models.ClearAllData()
}

func TestMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateTestSuite))
}
