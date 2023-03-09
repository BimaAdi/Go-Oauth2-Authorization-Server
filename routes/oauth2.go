package routes

import (
	"net/http"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/repository"
	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/gin-gonic/gin"
)

func oauthRoutes(rq *gin.RouterGroup) {
	oauth := rq.Group("/oauth")

	oauth.GET("/authorize/", loginUiRoute)
	oauth.POST("/authorize/", oauthLoginRoute)
}

func loginUiRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "login-ui.html", nil)
}

// Oauth Login
//
//	@Summary		Authorize Oauth Login
//	@Description	Login for Oauth
//	@Tags			Oauth
//	@Accept			json
//	@Param			payload	body	schemas.OauthLoginJsonRequest	true	"login data"
//	@Success		302
//	@Failure		400	{object}	schemas.BadRequestResponse
//	@Failure		403	{object}	schemas.ForbiddenResponse
//	@Failure		500	{object}	schemas.InternalServerErrorResponse
//	@Router			/oauth/authorize/ [post]
func oauthLoginRoute(c *gin.Context) {
	// Get data from json
	jsonRequest := schemas.OauthLoginJsonRequest{}
	err := c.ShouldBindJSON(&jsonRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	// Get client_id
	_, err = repository.GetDetailOauth2SessionByClientId(models.DBConn, jsonRequest.ClientId)
	if err != nil {
		c.JSON(http.StatusForbidden, schemas.ForbiddenResponse{
			Message: "Forbidden",
		})
		return
	}

	// Get User
	user, err := repository.GetUserByUsername(models.DBConn, jsonRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: "invalid credentials",
		})
		return
	}

	// Check Password
	if !core.CheckPasswordHash(jsonRequest.Password, user.Password) {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: "invalid credentials",
		})
		return
	}

	code, err := repository.GenerateCodeForUser(models.DBConn, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, jsonRequest.RedirectUri+"?state="+jsonRequest.State+"&code="+code)
}
