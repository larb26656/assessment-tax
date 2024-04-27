package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/larb26656/assessment-tax/config"
	"github.com/larb26656/assessment-tax/domains/admin"
	"github.com/larb26656/assessment-tax/domains/admin/deductions/personal"
	"github.com/larb26656/assessment-tax/domains/tax/calculator"
)

func RegisterRoute(appConfig *config.AppConfig, e *echo.Echo) {

	// personal deductions
	personalDeductionsRepository := personal.NewPersonalDeductionsRepository()
	personalDeductionsUsecase := personal.NewPersonalDeductionsUsecase(personalDeductionsRepository)
	personalDeductionsHttpHandler := personal.NewPersonalDeductionsHttpHandler(personalDeductionsUsecase)
	// tax
	taxCalculatorUsecase := calculator.NewTaxCalculatorUseCase(personalDeductionsRepository)
	taxCalculatorHttpHandler := calculator.NewTaxCalculatorHttpHandler(taxCalculatorUsecase)

	e.POST("/tax/calculations", taxCalculatorHttpHandler.CalculateTax)

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

	adminGroup.POST("/deductions/personal", personalDeductionsHttpHandler.UpdateDeductions)
}
