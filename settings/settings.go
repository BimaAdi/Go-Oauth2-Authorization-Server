package settings

import (
	"os"

	"github.com/joho/godotenv"
)

// Server connection
var ENVIRONTMENT string
var SERVER_HOST string
var SERVER_PORT string

// Postgresql connection
var POSTGRESQL_HOST string
var POSTGRESQL_USER string
var POSTGRESQL_PASSWORD string
var POSTGRESQL_DATABASE string
var POSTGRESQL_PORT string
var POSTGRESQL_SSL_MODE string
var POSTGRESQL_TIMEZONE string

func InitiateSettings(pathToEnvFile string) {
	if os.Getenv("ENVIRONTMENT") != "PROD" {
		err := godotenv.Load(pathToEnvFile)
		if err != nil {
			panic(err)
		}
	}

	ENVIRONTMENT = os.Getenv("ENVIRONTMENT")
	SERVER_HOST = os.Getenv("SERVER_HOST")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	POSTGRESQL_HOST = os.Getenv("POSTGRESQL_HOST")
	POSTGRESQL_USER = os.Getenv("POSTGRESQL_USER")
	POSTGRESQL_PASSWORD = os.Getenv("POSTGRESQL_PASSWORD")
	POSTGRESQL_DATABASE = os.Getenv("POSTGRESQL_DATABASE")
	POSTGRESQL_PORT = os.Getenv("POSTGRESQL_PORT")
	POSTGRESQL_SSL_MODE = os.Getenv("POSTGRESQL_SSL_MODE")
	POSTGRESQL_TIMEZONE = os.Getenv("POSTGRESQL_TIMEZONE")
}
