package personal

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

type mockPersonalDeductionUsecaseCaseSuccess struct {
}

func (m *mockPersonalDeductionUsecaseCaseSuccess) GetDeduction() (float64, error) {
	return 60000.0, nil
}

func (m *mockPersonalDeductionUsecaseCaseSuccess) UpdateDeduction(req UpdatePersonalDeductionReq) (UpdatePersonalDeductionRes, error) {
	return UpdatePersonalDeductionRes{
		PersonalDeduction: 60000.0,
	}, nil
}

func mockUpdateDeductionHttpReq(reqBody string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	e.Validator = myValidator.NewStructValidator(validator.New())

	req := httptest.NewRequest(http.MethodPost, "/admin/deductions/personal", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return e, c, rec
}

func TestUpdateDeductionHandler_ShouldGetBadRequest_WhenWrongInput(t *testing.T) {
	// Arrange
	usecase := &mockPersonalDeductionUsecaseCaseSuccess{}
	handler := NewPersonalDeductionHttpHandler(
		usecase,
	)

	testCases := []struct {
		name    string
		reqBody string
	}{
		{
			"Test case 1",
			``,
		},
		{
			"Test case 2",
			`{}`,
		},
		{
			"Test case 3",
			`{
				"amount": asdasd
			}`,
		},
		{
			"Test case 4",
			`{
				"amount": 99.0
			}`,
		},
		{
			"Test case 5",
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

type mockPersonalDeductionUsecaseCaseErrorOnUpdateDeduction struct {
}

func (m *mockPersonalDeductionUsecaseCaseErrorOnUpdateDeduction) GetDeduction() (float64, error) {
	return 60000.0, nil
}

func (m *mockPersonalDeductionUsecaseCaseErrorOnUpdateDeduction) UpdateDeduction(req UpdatePersonalDeductionReq) (UpdatePersonalDeductionRes, error) {
	return UpdatePersonalDeductionRes{}, errors.New("error on update")
}

func TestUpdateDeductionHandler_ShouldGetInternalServerError_WhenErrorOnUpdateDeduction(t *testing.T) {
	// Arrange
	usecase := &mockPersonalDeductionUsecaseCaseErrorOnUpdateDeduction{}
	handler := NewPersonalDeductionHttpHandler(
		usecase,
	)

	reqBody := `{
		"amount": 60000.0
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
	usecase := &mockPersonalDeductionUsecaseCaseSuccess{}
	handler := NewPersonalDeductionHttpHandler(
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
				"amount": 60000.0
			}`,
			`{
				"personalDeduction": 60000
			}`,
		},
		{
			"Test case 2",
			`{
				"amount": 70000.0
			}`,
			`{
				"personalDeduction": 60000
			}`,
		},
		{
			"Test case 3",
			`{
				"amount": 10000.0
			}`,
			`{
				"personalDeduction": 60000
			}`,
		},
		{
			"Test case 4",
			`{
				"amount": 100000.0
			}`,
			`{
				"personalDeduction": 60000
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

	reqBody := `{
		"amount": 60000.0
	}`

	_, c, rec := mockUpdateDeductionHttpReq(reqBody)

	expectedResponse := `{
		"personalDeduction": 60000
	}`

	// Act
	err := handler.UpdateDeduction(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}
