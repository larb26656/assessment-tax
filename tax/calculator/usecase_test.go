package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	// Test cases for CalculateTax function
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

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the CalculateTax function
			calculator := NewTaxCalculatorUseCase()
			result := calculator.CalculateTax(tc.netIncome)

			// Check if the result matches the expected tax
			assert.Equal(t, tc.expectedTax, result)
		})
	}
}
