package database

import (
	"database/sql"

	"github.com/larb26656/assessment-tax/config"
	_ "github.com/lib/pq"
)

func InitDatabase(appConfig *config.AppConfig) (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", appConfig.DatabaseUrl)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
