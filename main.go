package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/larb26656/assessment-tax/structValidator"
	"github.com/larb26656/assessment-tax/tax/calculator"
)

func main() {
	e := echo.New()

	// register
	calculator.RegisterRouter(e)

	e.Validator = structValidator.NewStructValidator(validator.New())

	e.Logger.Fatal(e.Start(":8080"))
}
