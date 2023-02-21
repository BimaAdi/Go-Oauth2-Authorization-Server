package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BimaAdi/Oauth2AuthorizationServer/tasks"
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
					tasks.RunServer(".env")
					return nil
				},
			},
			{
				Name:    "migrate-db",
				Aliases: []string{"md"},
				Usage:   "migrate database",
				Action: func(cCtx *cli.Context) error {
					tasks.MigrateDB(".env")
					return nil
				},
			},
			{
				Name:    "init-superuser",
				Aliases: []string{"is"},
				Usage:   "create initial superuser",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "username",
						Value: "admin",
						Usage: "superuser username (default: admin)",
					},
					&cli.StringFlag{
						Name:  "email",
						Value: "admin@local.com",
						Usage: "superuser email (default: admin@local.com)",
					},
					&cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "superuser password (required)",
					},
				},
				Action: func(cCtx *cli.Context) error {
					if cCtx.String("password") == "" {
						panic("--password not defined, --password is required, see init-superuser --help")
					}
					tasks.CreateSuperUser(".env", cCtx.String("email"), cCtx.String("username"), cCtx.String("password"))
					fmt.Println("init superuser")
					fmt.Println("email: " + cCtx.String("email"))
					fmt.Println("username: " + cCtx.String("username"))
					fmt.Println("password: " + cCtx.String("password"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
