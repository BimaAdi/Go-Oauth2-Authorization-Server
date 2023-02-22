package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func up(file_path string, postgres_url string) {
	m, err := migrate.New(
		file_path,
		postgres_url)
	if err != nil {
		panic(err.Error())
	}
	m.Up()
}
