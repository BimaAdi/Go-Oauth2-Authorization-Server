package tasks

import (
	"github.com/BimaAdi/Oauth2AuthorizationServer/migrations"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
)

func MigrateUp(envPath string, migrations_file_path string) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)
	postgresUser := settings.POSTGRESQL_USER
	postgresPassword := settings.POSTGRESQL_PASSWORD
	postgresHost := settings.POSTGRESQL_HOST
	postgresPort := settings.POSTGRESQL_PORT
	postgresDatabase := settings.POSTGRESQL_DATABASE
	postgresSslMode := settings.POSTGRESQL_SSL_MODE

	// Initiate Database connection
	postgres_url := "postgres://" + postgresUser + ":" + postgresPassword + "@" +
		postgresHost + ":" + postgresPort + "/" + postgresDatabase + "?sslmode=" + postgresSslMode

	// Migrate
	migrations.Up(migrations_file_path, postgres_url, nil)
}

func MigrateStep(envPath string, migrations_file_path string, step *int) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)
	postgresUser := settings.POSTGRESQL_USER
	postgresPassword := settings.POSTGRESQL_PASSWORD
	postgresHost := settings.POSTGRESQL_HOST
	postgresPort := settings.POSTGRESQL_PORT
	postgresDatabase := settings.POSTGRESQL_DATABASE
	postgresSslMode := settings.POSTGRESQL_SSL_MODE

	// Initiate Database connection
	postgres_url := "postgres://" + postgresUser + ":" + postgresPassword + "@" +
		postgresHost + ":" + postgresPort + "/" + postgresDatabase + "?sslmode=" + postgresSslMode

	// Migrate
	migrations.Up(migrations_file_path, postgres_url, step)
}

func MigrateDown(envPath string, migrations_file_path string) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)
	postgresUser := settings.POSTGRESQL_USER
	postgresPassword := settings.POSTGRESQL_PASSWORD
	postgresHost := settings.POSTGRESQL_HOST
	postgresPort := settings.POSTGRESQL_PORT
	postgresDatabase := settings.POSTGRESQL_DATABASE
	postgresSslMode := settings.POSTGRESQL_SSL_MODE

	// Initiate Database connection
	postgres_url := "postgres://" + postgresUser + ":" + postgresPassword + "@" +
		postgresHost + ":" + postgresPort + "/" + postgresDatabase + "?sslmode=" + postgresSslMode

	// Migrate
	migrations.Down(migrations_file_path, postgres_url, nil)
}

func GetMigrateVersion(envPath string, migrations_file_path string) string {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)
	postgresUser := settings.POSTGRESQL_USER
	postgresPassword := settings.POSTGRESQL_PASSWORD
	postgresHost := settings.POSTGRESQL_HOST
	postgresPort := settings.POSTGRESQL_PORT
	postgresDatabase := settings.POSTGRESQL_DATABASE
	postgresSslMode := settings.POSTGRESQL_SSL_MODE

	// Initiate Database connection
	postgres_url := "postgres://" + postgresUser + ":" + postgresPassword + "@" +
		postgresHost + ":" + postgresPort + "/" + postgresDatabase + "?sslmode=" + postgresSslMode

	return migrations.ShowCurrentVersion(migrations_file_path, postgres_url)
}
