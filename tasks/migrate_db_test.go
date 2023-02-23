package tasks_test

import (
	"testing"

	"github.com/BimaAdi/Oauth2AuthorizationServer/tasks"
	"github.com/stretchr/testify/assert"
)

func TestMigrateDB(t *testing.T) {
	assert.Equal(
		t, "No migration applied",
		tasks.GetMigrateVersion("../.env", "file://../migrations/migrations_files/"),
	)
	tasks.MigrateUp("../.env", "file://../migrations/migrations_files/")
	assert.NotEqual(
		t, "No migration applied",
		tasks.GetMigrateVersion("../.env", "file://../migrations/migrations_files/"),
	)
	tasks.MigrateDown("../.env", "file://../migrations/migrations_files/")
	assert.Equal(
		t, "No migration applied",
		tasks.GetMigrateVersion("../.env", "file://../migrations/migrations_files/"),
	)
}
