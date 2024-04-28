package calculator

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	myValidator "github.com/larb26656/assessment-tax/validator"
	"github.com/stretchr/testify/assert"
)

type mockTaxCalculatorUsecase struct {
}

func (m *mockTaxCalculatorUsecase) CalculateAllowances(allowances []AllowanceReq, maxDonation float64, maxKReceipt float64) float64 {
	return 0.0
}

func (m *mockTaxCalculatorUsecase) CalculateTaxDeduction(personalDeduction, totalAllowances float64) float64 {
	return 0.0
}

func (m *mockTaxCalculatorUsecase) CalculateNetIncome(income, taxDeduction float64) float64 {
	return 0.0
}

func (m *mockTaxCalculatorUsecase) CalculateTax(netIncome float64, wht float64) (float64, float64, []TaxLevelRes) {
	return 0.0, 0.0, []TaxLevelRes{}
}

func (m *mockTaxCalculatorUsecase) Calculate(req TaxCalculatorReq) (TaxCalculatorRes, error) {
	return TaxCalculatorRes{
		Tax:       14000.0,
		TaxRefund: 0.0,
		TaxLevel: []TaxLevelRes{
			{
				Level: "0-150,000",
				Tax:   0.0,
			},
			{
				Level: "150,001-500,000",
				Tax:   14000.0,
			},
			{
				Level: "500,001-1,000,000",
				Tax:   0.0,
			},
			{
				Level: "1,000,001-2,000,000",
				Tax:   0.0,
			},
			{
				Level: "2,000,001 ขึ้นไป",
				Tax:   0.0,
			},
		},
	}, nil
}

func (m *mockTaxCalculatorUsecase) CalculateMultiRequest(reqs []TaxCalculatorReq) (TaxCalucalorMultipleRes, error) {
	return TaxCalucalorMultipleRes{}, nil
}

func mockCalculateTaxHttpReq(reqBody string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	e.Validator = myValidator.NewStructValidator(validator.New())

	req := httptest.NewRequest(http.MethodPost, "/tax/calculations", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return e, c, rec
}

// CalculateTax
func TestCalculateTaxHandler_ShouldGetBadRequest_WhenInvalidInput(t *testing.T) {
	// Arrange
	usecase := &mockTaxCalculatorUsecase{}
	handler := NewTaxCalculatorHttpHandler(
		usecase,
	)

	testCases := []struct {
		name    string
		reqBody string
	}{
		{
			"Test case 1",
			`{
				"totalIncome": -1,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 200000.0
					},
					{
						"allowanceType": "donation",
						"amount": 100000.0
					}
				]
			}`,
		},
		{
			"Test case 2",
			`{
				"totalIncome": 0.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt222",
						"amount": 200000.0
					},
					{
						"allowanceType": "donation",
						"amount": 100000.0
					}
				]
			}`,
		},
		{
			"Test case 3",
			`{
				"totalIncome": 0.0,
				"wht": 0.0,
				"allowances": [
					{
						"amount": 200000.0
					},
					{
						"allowanceType": "donation",
						"amount": 100000.0
					}
				]
			}`,
		},
		{
			"Test case 3",
			``,
		},
		{
			"Test case 3",
			`{
				"totalIncome": aasdasd,
				"wht": 0.0,
				"allowances": [
					{
						"amount": 200000.0
					},
					{
						"allowanceType": "donation",
						"amount": 100000.0
					}
				]
			}`,
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, c, _ := mockCalculateTaxHttpReq(tc.reqBody)
			err := handler.CalculateTax(c)

			// Assert
			assert.Error(t, err)
			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusBadRequest, he.Code)
		})
	}
}

type mockTaxCalculatorUsecaseCaseErrorOnCalculate struct {
}

func (m *mockTaxCalculatorUsecaseCaseErrorOnCalculate) CalculateAllowances(allowances []AllowanceReq, maxDonation float64, maxKReceipt float64) float64 {
	return 0.0
}

func (m *mockTaxCalculatorUsecaseCaseErrorOnCalculate) CalculateTaxDeduction(personalDeduction, totalAllowances float64) float64 {
	return 0.0
}

func (m *mockTaxCalculatorUsecaseCaseErrorOnCalculate) CalculateNetIncome(income, taxDeduction float64) float64 {
	return 0.0
}

func (m *mockTaxCalculatorUsecaseCaseErrorOnCalculate) CalculateTax(netIncome float64, wht float64) (float64, float64, []TaxLevelRes) {
	return 0.0, 0.0, []TaxLevelRes{}
}

func (m *mockTaxCalculatorUsecaseCaseErrorOnCalculate) Calculate(req TaxCalculatorReq) (TaxCalculatorRes, error) {
	return TaxCalculatorRes{}, errors.New("error on calculaate")
}

func (m *mockTaxCalculatorUsecaseCaseErrorOnCalculate) CalculateMultiRequest(reqs []TaxCalculatorReq) (TaxCalucalorMultipleRes, error) {
	return TaxCalucalorMultipleRes{}, nil
}

func TestCalculateTaxHandler_ShouldGetInternalServerError_WhenInvalidInput(t *testing.T) {
	// Arrange
	usecase := &mockTaxCalculatorUsecaseCaseErrorOnCalculate{}
	handler := NewTaxCalculatorHttpHandler(
		usecase,
	)

	testCases := []struct {
		name    string
		reqBody string
	}{
		{
			"Test case 1",
			`{
				"totalIncome": 500000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 200000.0
					},
					{
						"allowanceType": "donation",
						"amount": 100000.0
					}
				]
			}`,
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, c, _ := mockCalculateTaxHttpReq(tc.reqBody)
			err := handler.CalculateTax(c)

			// Assert
			assert.Error(t, err)
			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusInternalServerError, he.Code)
		})
	}
}

func TestCalculateTaxHandler_ShouldGetSuccess_WhenCorrectInput(t *testing.T) {
	// Arrange
	usecase := &mockTaxCalculatorUsecase{}
	handler := NewTaxCalculatorHttpHandler(
		usecase,
	)

	testCases := []struct {
		name             string
		reqBody          string
		expectedResponse string
	}{
		{
			"Test case 1",
			`{
				"totalIncome": 500000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 200000.0
					},
					{
						"allowanceType": "donation",
						"amount": 100000.0
					}
				]
			}`,
			`{
				"tax": 14000,
				"taxRefund": 0,
				"taxLevel": [
					{
						"level": "0-150,000",
						"tax": 0
					},
					{
						"level": "150,001-500,000",
						"tax": 14000
					},
					{
						"level": "500,001-1,000,000",
						"tax": 0
					},
					{
						"level": "1,000,001-2,000,000",
						"tax": 0
					},
					{
						"level": "2,000,001 ขึ้นไป",
						"tax": 0
					}
				]
			}`,
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, c, rec := mockCalculateTaxHttpReq(tc.reqBody)
			err := handler.CalculateTax(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
