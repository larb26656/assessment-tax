package personal

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PersonalDeductionHttpHandler interface {
	UpdateDeduction(c echo.Context) error
}

type personalDeductionHttpHandler struct {
	personalDeductionUsecase PersonalDeductionUsecase
}

func NewPersonalDeductionHttpHandler(personalDeductionUsecase PersonalDeductionUsecase) PersonalDeductionHttpHandler {
	return &personalDeductionHttpHandler{
		personalDeductionUsecase: personalDeductionUsecase,
	}
}

func (p *personalDeductionHttpHandler) UpdateDeduction(c echo.Context) error {
	var req UpdatePersonalDeductionReq

	err := c.Bind(&req)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	res, err := p.personalDeductionUsecase.UpdateDeduction(req)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, res)
}
