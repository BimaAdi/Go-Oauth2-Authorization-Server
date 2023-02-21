package tasks_test

import (
	"testing"

	"github.com/BimaAdi/Oauth2AuthorizationServer/tasks"
)

func TestMigrateDB(t *testing.T) {
	tasks.MigrateDB("../.env") // No Error
}
