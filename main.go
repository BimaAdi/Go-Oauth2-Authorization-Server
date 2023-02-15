package main

import (
	"github.com/BimaAdi/Oauth2AuthorizationServer/docs"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/routes"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Oauth2 Authorization Server
// @version 1.0
// @description Oauth2 Authorization server
func main() {
	// Initialize environtment variable
	settings.InitiateSettings(".env")

	// Initiate Database connection
	models.Initiate()
	models.AutoMigrate() // Run Migration

	// Initialize gin route
	routes := routes.GetRoutes()

	// setup swagger
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = settings.SERVER_HOST + ":" + settings.SERVER_PORT
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// run gin server
	routes.Run(settings.SERVER_HOST + ":" + settings.SERVER_PORT)
}
