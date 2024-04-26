package main

import (
	"github.com/labstack/echo/v4"
	"github.com/larb26656/assessment-tax/tax/calculator"
)

func main() {
	e := echo.New()

	// register
	calculator.RegisterRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
