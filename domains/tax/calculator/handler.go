package calculator

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaxCalculatorHttpHandler interface {
	CalculateTax(c echo.Context) error
}

type taxCalculatorHttpHandler struct {
	taxCalculatorUseCase TaxCalculatorUseCase
}

func NewTaxCalculatorHttpHandler(taxCalculatorUseCase TaxCalculatorUseCase) TaxCalculatorHttpHandler {
	return &taxCalculatorHttpHandler{
		taxCalculatorUseCase: taxCalculatorUseCase,
	}
}

func (t taxCalculatorHttpHandler) CalculateTax(c echo.Context) error {
	var req TaxCalculatorReq

	err := c.Bind(&req)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	res, err := t.taxCalculatorUseCase.Calculate(req)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, res)
}
