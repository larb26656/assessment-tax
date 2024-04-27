package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/larb26656/assessment-tax/config"
	myValidator "github.com/larb26656/assessment-tax/validator"
)

func InitServer(appConfig *config.AppConfig) *echo.Echo {
	e := echo.New()

	e.Validator = myValidator.NewStructValidator(validator.New())

	RegisterRoute(appConfig, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appConfig.Port)))

	return e
}
