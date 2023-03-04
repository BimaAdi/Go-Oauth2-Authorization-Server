package routes

import (
	"github.com/gin-gonic/gin"
)

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRoutes(router *gin.Engine) *gin.Engine {
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")
	noPrefixRoutes := router.Group("/")
	authRoutes(noPrefixRoutes)
	userRoutes(noPrefixRoutes)
	oauthRoutes(noPrefixRoutes)
	return router
}
