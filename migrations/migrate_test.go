package migrations_test

import (
	"testing"

	"github.com/BimaAdi/Oauth2AuthorizationServer/migrations"
	"github.com/BimaAdi/Oauth2AuthorizationServer/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MigrateMigrateTestSuite struct {
	suite.Suite
}

func (suite *MigrateMigrateTestSuite) SetupSuite() {
	settings.InitiateSettings("../.env")
}

// ==========================================

func (suite *MigrateMigrateTestSuite) TestMigration() {
	// Given
	migrations_files := "file://migrations_files/"
	postgresUser := settings.POSTGRESQL_USER
	postgresPassword := settings.POSTGRESQL_PASSWORD
	postgresHost := settings.POSTGRESQL_HOST
	postgresPort := settings.POSTGRESQL_PORT
	postgresDatabase := settings.POSTGRESQL_DATABASE
	postgresSslMode := settings.POSTGRESQL_SSL_MODE
	postgers_url := "postgres://" + postgresUser + ":" + postgresPassword + "@" +
		postgresHost + ":" + postgresPort + "/" + postgresDatabase + "?sslmode=" + postgresSslMode

	// When
	assert.Equal(
		suite.T(), "No migration applied",
		migrations.ShowCurrentVersion(migrations_files, postgers_url),
	)
	migrations.Up(migrations_files, postgers_url, nil)
	assert.NotEqual(
		suite.T(), "No migration applied",
		migrations.ShowCurrentVersion(migrations_files, postgers_url),
	)
	migrations.Down(migrations_files, postgers_url, nil)
	assert.Equal(
		suite.T(), "No migration applied",
		migrations.ShowCurrentVersion(migrations_files, postgers_url),
	)

	// Expect
	// No Error
}

// ==========================================

func TestMigrateMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateMigrateTestSuite))
}
