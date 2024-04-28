package server

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/larb26656/assessment-tax/config"
	myValidator "github.com/larb26656/assessment-tax/validator"
)

func InitServer(appConfig *config.AppConfig, db *sql.DB) *echo.Echo {
	e := echo.New()

	e.Validator = myValidator.NewStructValidator(validator.New())

	RegisterRoute(appConfig, db, e)

	return e
}
