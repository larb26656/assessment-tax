package calculator

type TaxCalculatorUseCase interface {
	CalculateAllowances(allowances []AllowanceReq) float64
	CalculateTaxDeduction(selfTaxDeduction, totalAllowances float64) float64
	CalculateNetIncome(income, taxDeduction float64) float64
	CalculateTax(netIncome float64) float64
	Calculate(req TaxCalculatorReq) TaxCalculatorRes
}

type taxCalculatorUseCase struct {
}

func NewTaxCalculatorUseCase() TaxCalculatorUseCase {
	return &taxCalculatorUseCase{}
}

func (t taxCalculatorUseCase) CalculateAllowances(allowances []AllowanceReq) float64 {
	var totalAllowances float64 = 0

	for _, allowance := range allowances {
		totalAllowances += allowance.Amount
	}

	return totalAllowances
}

func (t taxCalculatorUseCase) CalculateTaxDeduction(selfTaxDeduction, totalAllowances float64) float64 {
	return totalAllowances + selfTaxDeduction
}

func (t taxCalculatorUseCase) CalculateNetIncome(income, taxDeduction float64) float64 {
	return income - taxDeduction
}

func (t taxCalculatorUseCase) CalculateTax(netIncome float64) float64 {
	var tax float64 = 0

	if netIncome <= 150000 {
		tax = 0
	} else if netIncome <= 500000 {
		tax = (netIncome - 150000) * 0.1
	} else if netIncome <= 1000000 {
		tax = (netIncome-500000)*0.15 + 35000 // 35000 is the tax for the first bracket
	} else if netIncome <= 2000000 {
		tax = (netIncome-1000000)*0.2 + 100000 // 100000 is the tax for the second bracket
	} else {
		tax = (netIncome-2000000)*0.35 + 300000 // 300000 is the tax for the third bracket
	}

	return tax
}

func (t taxCalculatorUseCase) Calculate(req TaxCalculatorReq) TaxCalculatorRes {
	totalAllowances := t.CalculateAllowances(req.Allowances)
	selfTaxDeduction := 60000.0
	taxDeduction := t.CalculateTaxDeduction(
		selfTaxDeduction,
		totalAllowances,
	)
	netIncome := t.CalculateNetIncome(
		req.TotalIncome,
		taxDeduction,
	)

	tax := t.CalculateTax(netIncome)
	netTax := tax - req.WHT

	return TaxCalculatorRes{
		Tax: netTax,
	}
}
