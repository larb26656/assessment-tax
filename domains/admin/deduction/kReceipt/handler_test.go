package kReceipt

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

type mockKReceiptDeductionUsecaseCaseSuccess struct {
}

func (m *mockKReceiptDeductionUsecaseCaseSuccess) GetDeduction() (float64, error) {
	return 50000.0, nil
}

func (m *mockKReceiptDeductionUsecaseCaseSuccess) UpdateDeduction(req UpdateKReceiptDeductionReq) (UpdateKReceiptDeductionRes, error) {
	return UpdateKReceiptDeductionRes{
		KReceipt: 50000.0,
	}, nil
}

func mockUpdateDeductionHttpReq(reqBody string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	e.Validator = myValidator.NewStructValidator(validator.New())

	req := httptest.NewRequest(http.MethodPost, "/admin/deductions/k-receipt", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return e, c, rec
}

func TestUpdateDeductionHandler_ShouldGetBadRequest_WhenWrongInput(t *testing.T) {
	// Arrange
	usecase := &mockKReceiptDeductionUsecaseCaseSuccess{}
	handler := NewKReceiptDeductionHttpHandler(
		usecase,
	)

	testCases := []struct {
		name    string
		reqBody string
	}{
		{
			"Test case 1",
			`{
				"amount": asdasd
			}`,
		},
		{
			"Test case 2",
			`{
				"amount": -1
			}`,
		},
		{
			"Test case 3",
			`{
				"amount": 100001.0
			}`,
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, c, _ := mockUpdateDeductionHttpReq(tc.reqBody)
			err := handler.UpdateDeduction(c)

			// Assert
			assert.Error(t, err)
			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusBadRequest, he.Code)
		})
	}
}

type mockKReceiptDeductionUsecaseCaseErrorOnUpdateDeduction struct {
}

func (m *mockKReceiptDeductionUsecaseCaseErrorOnUpdateDeduction) GetDeduction() (float64, error) {
	return 50000.0, nil
}

func (m *mockKReceiptDeductionUsecaseCaseErrorOnUpdateDeduction) UpdateDeduction(req UpdateKReceiptDeductionReq) (UpdateKReceiptDeductionRes, error) {
	return UpdateKReceiptDeductionRes{}, errors.New("error on update")
}

func TestUpdateDeductionHandler_ShouldGetInternalServerError_WhenErrorOnUpdateDeduction(t *testing.T) {
	// Arrange
	usecase := &mockKReceiptDeductionUsecaseCaseErrorOnUpdateDeduction{}
	handler := NewKReceiptDeductionHttpHandler(
		usecase,
	)

	reqBody := `{
		"amount": 50000.0
	}`

	_, c, _ := mockUpdateDeductionHttpReq(reqBody)

	// Act
	err := handler.UpdateDeduction(c)

	// Assert
	assert.Error(t, err)
	he, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
}

func TestUpdateDeductionHandler_ShouldGetSuccess_WhenCorrectInput(t *testing.T) {
	// Arrange
	usecase := &mockKReceiptDeductionUsecaseCaseSuccess{}
	handler := NewKReceiptDeductionHttpHandler(
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
				"amount": 50000.0
			}`,
			`{
				"kReceipt": 50000
			}`,
		},
		{
			"Test case 2",
			`{
				"amount": 70000.0
			}`,
			`{
				"kReceipt": 50000
			}`,
		},
		{
			"Test case 3",
			`{
				"amount": 10000.0
			}`,
			`{
				"kReceipt": 50000
			}`,
		},
		{
			"Test case 4",
			`{
				"amount": 100000.0
			}`,
			`{
				"kReceipt": 50000
			}`,
		},
		{
			"Test case 5",
			`{
				"amount": 0.0
			}`,
			`{
				"kReceipt": 50000
			}`,
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, c, rec := mockUpdateDeductionHttpReq(tc.reqBody)
			err := handler.UpdateDeduction(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
