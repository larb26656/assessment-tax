package calculator

import "github.com/larb26656/assessment-tax/domains/admin/deductions/personal"

type TaxCalculatorUseCase interface {
	CalculateAllowances(allowances []AllowanceReq, maxDonation float64) float64
	CalculateTaxDeduction(personalDeduction, totalAllowances float64) float64
	CalculateNetIncome(income, taxDeduction float64) float64
	CalculateTax(netIncome float64, wht float64) (float64, []TaxLevelRes)
	Calculate(req TaxCalculatorReq) (TaxCalculatorRes, error)
}

type taxCalculatorUseCase struct {
	personalDeductionsRepository personal.PersonalDeductionsRepository
}

func NewTaxCalculatorUseCase(personalDeductionsRepository personal.PersonalDeductionsRepository) TaxCalculatorUseCase {
	return &taxCalculatorUseCase{
		personalDeductionsRepository: personalDeductionsRepository,
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

func (t *taxCalculatorUseCase) CalculateTax(netIncome, wht float64) (float64, []TaxLevelRes) {
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

	taxLevels[lastTaxVisitIndex].Tax -= wht

	if taxLevels[lastTaxVisitIndex].Tax < 0 {
		taxLevels[lastTaxVisitIndex].Tax = 0
	}
	tax := taxLevels[lastTaxVisitIndex].Tax

	return tax, taxLevels
}

func (t *taxCalculatorUseCase) Calculate(req TaxCalculatorReq) (TaxCalculatorRes, error) {
	totalAllowances := t.CalculateAllowances(req.Allowances, 100000)
	selfTaxDeduction, err := t.personalDeductionsRepository.GetDeductions()

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

	tax, taxLevels := t.CalculateTax(netIncome, req.WHT)

	return TaxCalculatorRes{
		Tax:      tax,
		TaxLevel: taxLevels,
	}, nil
}
