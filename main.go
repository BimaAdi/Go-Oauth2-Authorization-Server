package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BimaAdi/Oauth2AuthorizationServer/docs"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
	"github.com/BimaAdi/Oauth2AuthorizationServer/routes"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli/v2"
)

//	@title									Go Oauth2 Authorization Server
//	@version								1.0
//	@description							Oauth2 Authorization server
//
//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenurl								/auth/login
//	@BasePath								/
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "sayhello",
				Aliases: []string{"sh"},
				Usage:   "say hello to",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("hello ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "runserver",
				Aliases: []string{"rs"},
				Usage:   "run webserver",
				Action: func(cCtx *cli.Context) error {
					runServer()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runServer() {
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
		AllowHeaders:           []string{"Origin", "Content-Type", "authorization", "accept"},
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
