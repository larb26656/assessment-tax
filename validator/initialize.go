package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewStructValidator(validator *validator.Validate) *StructValidator {
	return &StructValidator{
		validator: validator,
	}
}

func (cv *StructValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
