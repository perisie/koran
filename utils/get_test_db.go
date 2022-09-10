package utils

import (
	"database/sql"

	"github.com/arikama/go-arctic-tern/arctictern"
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func GetTestDb() (*sql.DB, error) {
	container, err := mysqltestcontainer.Create("test")
	if err != nil {
		return nil, err
	}

	db := container.GetDb()

	err = arctictern.Migrate(db, "./../migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}
