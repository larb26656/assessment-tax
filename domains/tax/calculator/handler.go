package calculator

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/larb26656/assessment-tax/constant/allowanceType"
)

type TaxCalculatorHttpHandler interface {
	CalculateTax(c echo.Context) error
	CalculateTaxWithCSV(c echo.Context) error
}

type taxCalculatorHttpHandler struct {
	taxCalculatorUseCase TaxCalculatorUseCase
}

func NewTaxCalculatorHttpHandler(taxCalculatorUseCase TaxCalculatorUseCase) TaxCalculatorHttpHandler {
	return &taxCalculatorHttpHandler{
		taxCalculatorUseCase: taxCalculatorUseCase,
	}
}

func (t *taxCalculatorHttpHandler) CalculateTax(c echo.Context) error {
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

func (t *taxCalculatorHttpHandler) CalculateTaxWithCSV(c echo.Context) error {
	// Read form file
	file, err := c.FormFile("taxFile")

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	src, err := file.Open()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	reader := csv.NewReader(src)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	var taxReqs []TaxCalculatorReq

	for i, row := range records {
		if i == 0 {
			continue
		}

		totalIncome, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			fmt.Println("Error converting TotalIncome:", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		}

		wht, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Println("Error converting WHT:", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		}

		donation, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Println("Error converting Donation:", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		}

		taxReq := TaxCalculatorReq{
			TotalIncome: totalIncome,
			WHT:         wht,
			Allowances: []AllowanceReq{
				{
					AllowanceType: allowanceType.Donation,
					Amount:        donation,
				},
			},
		}

		if err = c.Validate(taxReq); err != nil {
			fmt.Println("Error validate req:", err)
			return err
		}

		taxReqs = append(taxReqs, taxReq)
	}

	result, err := t.taxCalculatorUseCase.CalculateMultiRequest(taxReqs)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, result)
}
