package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/larb26656/assessment-tax/config"
	"github.com/larb26656/assessment-tax/domains/admin"
	"github.com/larb26656/assessment-tax/domains/admin/deduction"
	"github.com/larb26656/assessment-tax/domains/admin/deduction/kReceipt"
	"github.com/larb26656/assessment-tax/domains/admin/deduction/personal"
	"github.com/larb26656/assessment-tax/domains/tax/calculator"
)

func RegisterRoute(appConfig *config.AppConfig, db *sql.DB, e *echo.Echo) {

	// deduction
	deductionRepository := deduction.NewDeductionsRepository(db)

	// personal deduction
	personalDeductionsUsecase := personal.NewPersonalDeductionUsecase(deductionRepository)
	personalDeductionsHttpHandler := personal.NewPersonalDeductionHttpHandler(personalDeductionsUsecase)

	// k-Receipt deduction
	kReceiptDeductionsUsecase := kReceipt.NewKReceiptDeductionUsecase(deductionRepository)
	kReceiptDeductionsHttpHandler := kReceipt.NewKReceiptDeductionHttpHandler(kReceiptDeductionsUsecase)

	// admin
	adminRepository := admin.NewAdminRepository(appConfig)
	adminUsecase := admin.NewAdminUsecase(adminRepository)

	adminGroup := e.Group("/admin")

	adminGroup.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if !adminUsecase.Authenticate(username, password) {
			return false, nil
		}

		return true, nil
	}))

	adminGroup.POST("/deductions/personal", personalDeductionsHttpHandler.UpdateDeduction)
	adminGroup.POST("/deductions/k-receipt", kReceiptDeductionsHttpHandler.UpdateDeduction)

	// tax
	taxCalculatorUsecase := calculator.NewTaxCalculatorUseCase(personalDeductionsUsecase, kReceiptDeductionsUsecase)
	taxCalculatorHttpHandler := calculator.NewTaxCalculatorHttpHandler(taxCalculatorUsecase)

	e.POST("/tax/calculations", taxCalculatorHttpHandler.CalculateTax)
	e.POST("/tax/calculations/upload-csv", taxCalculatorHttpHandler.CalculateTaxWithCSV)
}
