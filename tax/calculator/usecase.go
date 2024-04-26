package calculator

import "fmt"

type TaxCalculatorUseCase interface {
	Calculate(req TaxCalculatorReq) TaxCalculatorRes
	CalculateAllowances(req TaxCalculatorReq) float64
	CalculateNetIncome(income, selfTaxDeduction, wht, totalAllowances float64) float64
	CalculateTax(netIncome float64) float64
}

type taxCalculatorUseCase struct {
}

func NewTaxCalculatorUseCase() TaxCalculatorUseCase {
	return &taxCalculatorUseCase{}
}

func (t taxCalculatorUseCase) Calculate(req TaxCalculatorReq) TaxCalculatorRes {
	totalAllowances := t.CalculateAllowances(req)
	selfTaxDeduction := 60000.0
	totalIncome := t.CalculateNetIncome(
		req.TotalIncome,
		selfTaxDeduction,
		req.WHT,
		totalAllowances,
	)

	tax := t.CalculateTax(totalIncome)

	fmt.Println(totalAllowances)
	fmt.Println(totalIncome)
	fmt.Println(tax)

	return TaxCalculatorRes{
		Tax: tax,
	}
}

func (*taxCalculatorUseCase) CalculateAllowances(req TaxCalculatorReq) float64 {
	var totalAllowances float64 = 0

	for _, allowance := range req.Allowances {
		totalAllowances += allowance.Amount
	}

	return totalAllowances
}

func (*taxCalculatorUseCase) CalculateNetIncome(income, wht, selfTaxDeduction, totalAllowances float64) float64 {
	taxDeduction := (wht + totalAllowances + selfTaxDeduction)
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
