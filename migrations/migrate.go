package migrations

import (
	"errors"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up(file_path string, postgres_url string, step *int) {
	m, err := migrate.New(
		file_path,
		postgres_url)
	if err != nil {
		panic(err.Error())
	}

	if step == nil {
		m.Up()
	} else {
		m.Steps(*step)
	}
}

func Down(file_path string, postgres_url string, step *int) {
	m, err := migrate.New(
		file_path,
		postgres_url)
	if err != nil {
		panic(err.Error())
	}

	if step == nil {
		m.Down()
	} else {
		m.Steps(*step)
	}
}

func ShowCurrentVersion(file_path string, postgres_url string) string {
	m, err := migrate.New(
		file_path,
		postgres_url)
	if err != nil {
		panic(err.Error())
	}

	version, dirty, err := m.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			return "No migration applied"
		}
		panic(err.Error())
	}
	dirtyString := ""
	if dirty {
		dirtyString = "true"
	} else {
		dirtyString = "false"
	}
	return "version: " + strconv.FormatUint(uint64(version), 10) + " dirty: " + dirtyString
}
