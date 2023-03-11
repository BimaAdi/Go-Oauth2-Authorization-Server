package routes

import (
	"net/http"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/repository"
	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func oauthRoutes(rq *gin.RouterGroup) {
	oauth := rq.Group("/oauth")

	oauth.GET("/authorize/", loginUiRoute)
	oauth.POST("/authorize/", oauthLoginRoute)
	oauth.POST("/token/", oauthTokenRoute)
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

// Oauth Token
//
//	@Summary		Oauth Token
//	@Description	Get Authorization Token
//	@Tags			Oauth
//	@Produce		json
//	@Accept			json
//	@Param			payload	body		schemas.Oauth2TokenJsonRequest	true	"code data"
//	@Success		200		{object}	schemas.LoginResponse
//	@Failure		400		{object}	schemas.BadRequestResponse
//	@Failure		403		{object}	schemas.ForbiddenResponse
//	@Failure		500		{object}	schemas.InternalServerErrorResponse
//	@Router			/oauth/token/ [post]
func oauthTokenRoute(c *gin.Context) {
	// Get data from json
	jsonRequest := schemas.Oauth2TokenJsonRequest{}
	err := c.ShouldBindJSON(&jsonRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	var jwt_token string
	err = models.DBConn.Transaction(func(tx *gorm.DB) error {
		// Validate Client Id and Client Secret
		_, err = repository.GetDetailOauth2SessionByClientIdAndClientSecret(tx, jsonRequest.ClientId, jsonRequest.ClientSecret)
		if err != nil {
			c.JSON(http.StatusForbidden, schemas.ForbiddenResponse{
				Message: "Forbidden",
			})
			return err
		}

		// Validate Code
		oauthToken, err := repository.GetOauthTokenByCode(
			tx, jsonRequest.Code,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, schemas.ForbiddenResponse{
				Message: "Invalid Code",
			})
			return err
		}

		// Generate JWT
		tx.Preload("User").Find(&oauthToken)
		jwt_token, err = core.GenerateJWTTokenFromUser(tx, oauthToken.User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return err
		}

		// Remove Code
		err = repository.DeleteOauthToken(tx, &oauthToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
				Error: err.Error(),
			})
			return err
		}

		return nil
	})

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, schemas.LoginResponse{
		AccessToken: jwt_token,
		TokenType:   "Bearer",
	})
}
