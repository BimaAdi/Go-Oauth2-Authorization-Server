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

	now := time.Now()
	session, err := repository.GenerateClientIdAndClientSecret(models.DBConn, requestUser, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, schemas.ClientRegiterResponse{
		ClientId:     session.ClientID,
		ClientSecret: session.ClientSecret,
	})
}
