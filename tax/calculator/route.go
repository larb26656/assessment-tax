package calculator

import (
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	taxCalculatorUseCase := NewTaxCalculatorUseCase()
	taxCalculatorHandler := NewTaxCalculatorHandler(taxCalculatorUseCase)

	e.POST("/tax/calculations", taxCalculatorHandler.CalculateTax)
}
