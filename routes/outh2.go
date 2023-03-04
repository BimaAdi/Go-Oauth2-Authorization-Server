package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func oauthRoutes(rq *gin.RouterGroup) {
	auths := rq.Group("/oauth")

	auths.GET("/login", loginUiRoute)
}

func loginUiRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "login-ui.html", nil)
}
