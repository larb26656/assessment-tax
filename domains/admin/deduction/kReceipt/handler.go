package kReceipt

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KReceiptDeductionHttpHandler interface {
	UpdateDeduction(c echo.Context) error
}

type kReceiptDeductionHttpHandler struct {
	kReceiptDeductionUsecase KReceiptDeductionUsecase
}

func NewKReceiptDeductionHttpHandler(kReceiptDeductionUsecase KReceiptDeductionUsecase) KReceiptDeductionHttpHandler {
	return &kReceiptDeductionHttpHandler{
		kReceiptDeductionUsecase: kReceiptDeductionUsecase,
	}
}

func (p *kReceiptDeductionHttpHandler) UpdateDeduction(c echo.Context) error {
	var req UpdateKReceiptDeductionReq

	err := c.Bind(&req)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	res, err := p.kReceiptDeductionUsecase.UpdateDeduction(req)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, res)
}
