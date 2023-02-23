package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/BimaAdi/Oauth2AuthorizationServer/tasks"
	"github.com/urfave/cli/v2"
)

// @title									Go Oauth2 Authorization Server
// @version								1.0
// @description							Oauth2 Authorization server
//
// @securitydefinitions.oauth2.password	OAuth2Password
// @tokenurl								/auth/login
// @BasePath								/
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
				Subcommands: []*cli.Command{
					{
						Name:    "up",
						Aliases: []string{"u"},
						Usage:   "migrate up (create all tables)",
						Action: func(cCtx *cli.Context) error {
							tasks.MigrateUp(".env", "file://migrations/migrations_files/")
							return nil
						},
					},
					{
						Name:    "version",
						Aliases: []string{"v"},
						Usage:   "show current migrate version",
						Action: func(cCtx *cli.Context) error {
							version := tasks.GetMigrateVersion(".env", "file://migrations/migrations_files/")
							fmt.Println(version)
							return nil
						},
					},
					{
						Name:    "step",
						Aliases: []string{"s"},
						Usage:   "migrate to certain version, migrate step {version num}",
						Action: func(cCtx *cli.Context) error {
							stepInt, err := strconv.Atoi(cCtx.Args().First())
							if err != nil {
								panic(err.Error())
							}
							tasks.MigrateStep(".env", "file://migrations/migrations_files/", &stepInt)
							return nil
						},
					},
					{
						Name:    "down",
						Aliases: []string{"d"},
						Usage:   "migrate down (delete all tables)",
						Action: func(cCtx *cli.Context) error {
							tasks.MigrateDown(".env", "file://migrations/migrations_files/")
							return nil
						},
					},
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
