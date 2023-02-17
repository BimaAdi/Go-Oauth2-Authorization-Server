package routes

import (
	"net/http"

	"github.com/BimaAdi/Oauth2AuthorizationServer/schemas"
	"github.com/gin-gonic/gin"
)

func authRoutes(rq *gin.RouterGroup) {
	auths := rq.Group("/auth")

	auths.POST("/login", authLoginRoute)
}

func authLoginRoute(c *gin.Context) {
	// TODO implement login route
	c.JSON(http.StatusNotImplemented, schemas.NotImplementedResponse{
		Error: "Not Implemented Yet",
	})
}
