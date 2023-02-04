package main

import (
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/routes"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
)

func main() {
	// Initialize environtment variable
	settings.InitiateSettings(".env")

	// Initiate Database connection
	models.Initiate()
	models.AutoMigrate() // Run Migration

	// run gin server
	routes := routes.GetRoutes()
	routes.Run(settings.SERVER_HOST + ":" + settings.SERVER_PORT)
}
