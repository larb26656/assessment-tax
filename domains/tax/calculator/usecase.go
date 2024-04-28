package calculator

import "github.com/larb26656/assessment-tax/domains/admin/deduction/personal"

type TaxCalculatorUseCase interface {
	CalculateAllowances(allowances []AllowanceReq, maxDonation float64) float64
	CalculateTaxDeduction(personalDeduction, totalAllowances float64) float64
	CalculateNetIncome(income, taxDeduction float64) float64
	CalculateTax(netIncome float64, wht float64) (float64, float64, []TaxLevelRes)
	Calculate(req TaxCalculatorReq) (TaxCalculatorRes, error)
}

type taxCalculatorUseCase struct {
	personalDeductionUsecase personal.PersonalDeductionUsecase
}

func NewTaxCalculatorUseCase(personalDeductionUsecase personal.PersonalDeductionUsecase) TaxCalculatorUseCase {
	return &taxCalculatorUseCase{
		personalDeductionUsecase: personalDeductionUsecase,
	}
}

func (t *taxCalculatorUseCase) CalculateAllowances(allowances []AllowanceReq, maxDonation float64) float64 {
	var totalAllowances float64 = 0
	var totalDonation float64 = 0

	for _, allowance := range allowances {
		if allowance.AllowanceType == "donation" {
			totalDonation += allowance.Amount
		}
	}

	if totalDonation > maxDonation {
		totalDonation = maxDonation
	}

	totalAllowances = totalDonation

	return totalAllowances
}

func (t *taxCalculatorUseCase) CalculateTaxDeduction(personalDeduction, totalAllowances float64) float64 {
	return totalAllowances + personalDeduction
}

func (t *taxCalculatorUseCase) CalculateNetIncome(income, taxDeduction float64) float64 {
	return income - taxDeduction
}

func (t *taxCalculatorUseCase) CalculateTax(netIncome, wht float64) (float64, float64, []TaxLevelRes) {
	taxLevels := []TaxLevelRes{
		{
			Level: "0-150,000",
			Tax:   0,
		},
		{
			Level: "150,001-500,000",
			Tax:   0,
		},
		{
			Level: "500,001-1,000,000",
			Tax:   0,
		},
		{
			Level: "1,000,001-2,000,000",
			Tax:   0,
		},
		{
			Level: "2,000,001 ขึ้นไป",
			Tax:   0,
		},
	}

	lastTaxVisitIndex := 0

	if netIncome >= 0 {
		taxLevels[0].Tax = 0
	}

	if netIncome > 150000 {
		taxInLevel := (netIncome - 150000) * 0.1

		if taxInLevel < 35000 {
			taxLevels[1].Tax = taxInLevel
		} else {
			taxLevels[1].Tax = 35000
		}

		lastTaxVisitIndex++
	}

	if netIncome > 500000 {
		taxInLevel := (netIncome-500000)*0.15 + 35000 // 35000 is the tax for the first bracket

		if taxInLevel < 100000 {
			taxLevels[2].Tax = taxInLevel
		} else {
			taxLevels[2].Tax = 100000
		}

		lastTaxVisitIndex++
	}

	if netIncome > 1000000 {
		taxInLevel := (netIncome-1000000)*0.2 + 100000 // 100000 is the tax for the second bracket

		if taxInLevel < 300000 {
			taxLevels[3].Tax = taxInLevel
		} else {
			taxLevels[3].Tax = 300000
		}

		lastTaxVisitIndex++
	}

	if netIncome > 2000000 {
		taxLevels[4].Tax = (netIncome-2000000)*0.35 + 300000 // 300000 is the tax for the third bracket
		lastTaxVisitIndex++
	}

	tax := taxLevels[lastTaxVisitIndex].Tax

	tax -= wht

	taxRefund := 0.0

	if tax < 0 {
		taxRefund = tax * -1
		tax = 0

	}

	return tax, taxRefund, taxLevels
}

func (t *taxCalculatorUseCase) Calculate(req TaxCalculatorReq) (TaxCalculatorRes, error) {
	totalAllowances := t.CalculateAllowances(req.Allowances, 100000)
	selfTaxDeduction, err := t.personalDeductionUsecase.GetDeduction()

	if err != nil {
		return TaxCalculatorRes{}, err
	}

	taxDeduction := t.CalculateTaxDeduction(
		selfTaxDeduction,
		totalAllowances,
	)
	netIncome := t.CalculateNetIncome(
		req.TotalIncome,
		taxDeduction,
	)

	tax, taxRefund, taxLevels := t.CalculateTax(netIncome, req.WHT)

	return TaxCalculatorRes{
		Tax:       tax,
		TaxRefund: taxRefund,
		TaxLevel:  taxLevels,
	}, nil
}
