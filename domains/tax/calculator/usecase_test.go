package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// CalculateAllowances
func TestCalculateAllowances_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase()
	testCases := []struct {
		name                    string
		allowances              []AllowanceReq
		expectedTotalAllowances float64
	}{
		{"Test case 1", []AllowanceReq{
			{
				AllowanceType: "donation",
				Amount:        2000,
			},
		}, 2000},
		{"Test case 2", []AllowanceReq{
			{
				AllowanceType: "donation",
				Amount:        2000,
			},
			{
				AllowanceType: "donation",
				Amount:        4000,
			},
			{
				AllowanceType: "donation",
				Amount:        4000,
			},
		}, 10000},
		{"Test case 3", []AllowanceReq{
			{
				AllowanceType: "donation",
				Amount:        2000,
			},
			{
				AllowanceType: "donation",
				Amount:        4000,
			},
			{
				AllowanceType: "donation",
				Amount:        200000,
			},
		}, 100000},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculator.CalculateAllowances(tc.allowances, 100000)

			// Assert
			assert.Equal(t, tc.expectedTotalAllowances, result)
		})
	}
}

// CalculateTaxDeduction
func TestCalculateTaxDeduction_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase()
	testCases := []struct {
		name                 string
		selfTaxDeduction     float64
		totalAllowances      float64
		expectedTaxDeduction float64
	}{
		{"Test case 1", 60000, 20, 60020},
		{"Test case 2", 70000, 30000, 100000}, // Expected tax is 5% of (200000 - 150000) Expected tax is 35% of (3000000 - 2000000) + 300000
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculator.CalculateTaxDeduction(tc.selfTaxDeduction, tc.totalAllowances)

			// Assert
			assert.Equal(t, tc.expectedTaxDeduction, result)
		})
	}
}

// CalculateNetIncome
func TestCalculateNetIncome_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase()
	testCases := []struct {
		name              string
		income            float64
		taxDeduction      float64
		expectedNetIncome float64
	}{
		{"Test case 1", 200000, 150000, 50000},
		{"Test case 2", 200000, 300000, -100000}, // Expected tax is 5% of (200000 - 150000) Expected tax is 35% of (3000000 - 2000000) + 300000
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculator.CalculateNetIncome(tc.income, tc.taxDeduction)

			// Assert
			assert.Equal(t, tc.expectedNetIncome, result)
		})
	}
}

// TestCalculateTax
func TestCalculateTax_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase()
	testCases := []struct {
		name        string
		netIncome   float64
		expectedTax float64
	}{
		{"Test case 1", 100000, 0},
		{"Test case 2", 200000, 5000},    // Expected tax is 5% of (200000 - 150000)
		{"Test case 3", 440000, 29000},   // Expected tax is 15% of (440000 - 500000) + 35000
		{"Test case 4", 600000, 50000},   // Expected tax is 15% of (600000 - 500000) + 35000
		{"Test case 5", 750000, 72500},   // Expected tax is 15% of (750000 - 500000) + 35000
		{"Test case 6", 1500000, 200000}, // Expected tax is 20% of (1500000 - 1000000) + 100000
		{"Test case 7", 3000000, 650000}, // Expected tax is 35% of (3000000 - 2000000) + 300000
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculator.CalculateTax(tc.netIncome)

			// Assert
			assert.Equal(t, tc.expectedTax, result)
		})
	}
}

// TestCalculate
func TestCalculate_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase()

	testCases := []struct {
		name        string
		req         TaxCalculatorReq
		expectedTax float64
	}{
		{"Test case 1", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         0.0,
			Allowances: []AllowanceReq{
				{AllowanceType: "donation", Amount: 0.0},
			},
		}, 29000},
		{"Test case 2", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         25000.0,
			Allowances: []AllowanceReq{
				{AllowanceType: "donation", Amount: 0.0},
			},
		}, 4000},
		{"Test case 3", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         0.0,
			Allowances: []AllowanceReq{
				{AllowanceType: "donation", Amount: 200000.0},
			},
		}, 19000.0},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := calculator.Calculate(tc.req)

			// Assert
			assert.Equal(t, tc.expectedTax, result.Tax)
		})
	}
}
