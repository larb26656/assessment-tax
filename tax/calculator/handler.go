package calculator

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaxCalculatorHttpHandler interface {
	CalculateTax(c echo.Context) error
}

type taxCalculatorHttpHandler struct {
	taxCalculatorUseCase TaxCalculatorUseCase
}

func NewTaxCalculatorHandler(taxCalculatorUseCase TaxCalculatorUseCase) TaxCalculatorHttpHandler {
	return &taxCalculatorHttpHandler{
		taxCalculatorUseCase: taxCalculatorUseCase,
	}
}

func (t taxCalculatorHttpHandler) CalculateTax(c echo.Context) error {
	// TODO validate req payload
	var req TaxCalculatorReq

	err := c.Bind(&req)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	res := t.taxCalculatorUseCase.Calculate(req)

	return c.JSON(http.StatusOK, res)
}
