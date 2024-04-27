package personal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PersonalDeductionsHttpHandler interface {
	UpdateDeductions(c echo.Context) error
}

type personalDeductionsHttpHandler struct {
	personalDeductionsUsecase PersonalDeductionsUsecase
}

func NewPersonalDeductionsHttpHandler(personalDeductionsUsecase PersonalDeductionsUsecase) PersonalDeductionsHttpHandler {
	return &personalDeductionsHttpHandler{
		personalDeductionsUsecase: personalDeductionsUsecase,
	}
}

func (p *personalDeductionsHttpHandler) UpdateDeductions(c echo.Context) error {
	var req UpdatePersonalDeductionsReq

	err := c.Bind(&req)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	res, err := p.personalDeductionsUsecase.UpdateDeductions(req)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, res)
}
