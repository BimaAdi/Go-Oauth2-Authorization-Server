package main

import (
	"github.com/BimaAdi/Oauth2AuthorizationServer/docs"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/routes"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title									Go Oauth2 Authorization Server
//	@version								1.0
//	@description							Oauth2 Authorization server
//
//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenurl								/auth/login
//	@BasePath								/
func main() {
	// Initialize environtment variable
	settings.InitiateSettings(".env")

	// Initiate Database connection
	models.Initiate()
	models.AutoMigrate() // Run Migration

	// Cors Middleware
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           []string{},
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowHeaders:           []string{"Origin", "Content-Type"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 0,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	// Initialize gin route
	routes := routes.GetRoutes(router)

	// setup swagger
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = settings.SERVER_HOST + ":" + settings.SERVER_PORT
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// run gin server
	routes.Run(settings.SERVER_HOST + ":" + settings.SERVER_PORT)
}
