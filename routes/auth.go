package routes

import (
	"net/http"
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/repository"
	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/gin-gonic/gin"
)

func authRoutes(rq *gin.RouterGroup) {
	auths := rq.Group("/auth")

	auths.POST("/login", authLoginRoute)
	auths.POST("/register-client", registerClientRoute)
	auths.GET("/client-id-client-secret/", getClientIdClientSecret)
}

// Login
//
//	@Summary		Login
//	@Description	login
//	@Tags			Auth
//	@Produce		json
//	@Param			payload	formData	schemas.LoginFormRequest	true	"form data"
//	@Success		200		{object}	schemas.LoginResponse
//	@Failure		400		{object}	schemas.BadRequestResponse
//	@Failure		500		{object}	schemas.InternalServerErrorResponse
//	@Router			/auth/login [post]
func authLoginRoute(c *gin.Context) {
	// Get data from form
	formRequest := schemas.LoginFormRequest{}
	err := c.ShouldBind(&formRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	// Get User
	user, err := repository.GetUserByUsername(models.DBConn, formRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: "invalid credentials",
		})
		return
	}

	// Check Password
	if !core.CheckPasswordHash(formRequest.Password, user.Password) {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: "invalid credentials",
		})
		return
	}

	// Generate JWT token
	token, err := core.GenerateJWTTokenFromUser(models.DBConn, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, schemas.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	})
}

// Register Client
//
//	@Summary		register client
//	@Description	generate client_id and client_secret for request user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			register	body		schemas.ClientRegiterRequest	true	"register client"
//	@Success		201			{object}	schemas.ClientRegiterResponse
//	@Failure		403			{object}	schemas.ForbiddenResponse
//	@Failure		500			{object}	schemas.InternalServerErrorResponse
//	@Security		OAuth2Password
//	@Router			/auth/register-client [post]
func registerClientRoute(c *gin.Context) {
	// Authorize User
	requestUser, err := core.GetUserFromAuthorizationHeader(models.DBConn, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, schemas.UnauthorizedResponse{
			Message: "Invalid/Expired token",
		})
		return
	}

	if !requestUser.IsSuperuser {
		c.JSON(http.StatusForbidden, schemas.ForbiddenResponse{
			Message: "User not allowed to perform this action",
		})
		return
	}

	// Parse JSON
	var request schemas.ClientRegiterRequest
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	// Generate client_id and client_secret
	now := time.Now()
	session, err := repository.GenerateClientIdAndClientSecret(
		models.DBConn,
		request.Name,
		request.Description,
		requestUser,
		now,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, schemas.ClientRegiterResponse{
		Name:         session.Name,
		Description:  session.Description,
		ClientId:     session.ClientID,
		ClientSecret: session.ClientSecret,
	})
}

// Get All Session
//
//	@Summary		Get All client_id and client_secret for user
//	@Description	Get All client_id and client_secret for user
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object}	schemas.ArrayClientRegisterResponse
//	@Failure		400	{object}	schemas.BadRequestResponse
//	@Failure		401	{object}	schemas.UnauthorizedResponse
//	@Failure		500	{object}	schemas.InternalServerErrorResponse
//	@Security		OAuth2Password
//	@Router			/auth/client-id-client-secret/ [get]
func getClientIdClientSecret(c *gin.Context) {
	// Authorize User
	requestUser, err := core.GetUserFromAuthorizationHeader(models.DBConn, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, schemas.UnauthorizedResponse{
			Message: "Invalid/Expired token",
		})
		return
	}

	// Preload Oauth2Sessions
	models.DBConn.Preload("Oauth2Sessions").Find(&requestUser)

	// construct json
	arraySession := []schemas.ClientRegiterResponse{}
	for _, item := range requestUser.Oauth2Sessions {
		arraySession = append(arraySession, schemas.ClientRegiterResponse{
			Name:         item.Name,
			Description:  item.Description,
			ClientId:     item.ClientID,
			ClientSecret: item.ClientSecret,
		})
	}
	listSessionSchema := schemas.ArrayClientRegisterResponse{
		Data: arraySession,
	}

	c.JSON(http.StatusOK, listSessionSchema)
}
