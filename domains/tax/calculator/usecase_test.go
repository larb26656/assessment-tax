package calculator

import (
	"errors"
	"testing"

	"github.com/larb26656/assessment-tax/constant/allowanceType"
	"github.com/larb26656/assessment-tax/domains/admin/deduction/kReceipt"
	"github.com/larb26656/assessment-tax/domains/admin/deduction/personal"
	"github.com/stretchr/testify/assert"
)

type mockPersonalDeductionUsecase struct {
}

func (p *mockPersonalDeductionUsecase) GetDeduction() (float64, error) {
	return 60000.0, nil
}

func (p *mockPersonalDeductionUsecase) UpdateDeduction(req personal.UpdatePersonalDeductionReq) (personal.UpdatePersonalDeductionRes, error) {
	return personal.UpdatePersonalDeductionRes{}, nil
}

type mockKReceiptDeductionUsecase struct {
}

func (p *mockKReceiptDeductionUsecase) GetDeduction() (float64, error) {
	return 50000.0, nil
}

func (p *mockKReceiptDeductionUsecase) UpdateDeduction(req kReceipt.UpdateKReceiptDeductionReq) (kReceipt.UpdateKReceiptDeductionRes, error) {
	return kReceipt.UpdateKReceiptDeductionRes{}, nil
}

// CalculateAllowances
func TestCalculateAllowances_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecase{},
	)
	testCases := []struct {
		name                    string
		allowances              []AllowanceReq
		expectedTotalAllowances float64
	}{
		{"Test case 1", []AllowanceReq{
			{
				AllowanceType: allowanceType.Donation,
				Amount:        2000.0,
			},
		}, 2000.0},
		{"Test case 2", []AllowanceReq{
			{
				AllowanceType: allowanceType.Donation,
				Amount:        2000,
			},
			{
				AllowanceType: allowanceType.Donation,
				Amount:        4000,
			},
			{
				AllowanceType: allowanceType.Donation,
				Amount:        4000,
			},
		}, 10000},
		{"Test case 3", []AllowanceReq{
			{
				AllowanceType: allowanceType.Donation,
				Amount:        2000.0,
			},
			{
				AllowanceType: allowanceType.Donation,
				Amount:        4000.0,
			},
			{
				AllowanceType: allowanceType.Donation,
				Amount:        200000.0,
			},
		}, 100000.0},
		{"Test case 4", []AllowanceReq{
			{
				AllowanceType: allowanceType.Donation,
				Amount:        2000.0,
			},
			{
				AllowanceType: allowanceType.KReceipt,
				Amount:        4000.0,
			},
		}, 6000.0},
		{"Test case 5", []AllowanceReq{
			{
				AllowanceType: allowanceType.Donation,
				Amount:        200000.0,
			},
			{
				AllowanceType: allowanceType.KReceipt,
				Amount:        200000.0,
			},
		}, 150000.0},
		{"Test case 6", []AllowanceReq{
			{
				AllowanceType: allowanceType.KReceipt,
				Amount:        5000.0,
			},
		}, 5000.0},
		{"Test case 7", []AllowanceReq{
			{
				AllowanceType: allowanceType.KReceipt,
				Amount:        200000.0,
			},
		}, 50000.0},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculator.CalculateAllowances(tc.allowances, 100000.0, 50000.0)

			// Assert
			assert.Equal(t, tc.expectedTotalAllowances, result)
		})
	}
}

// CalculateTaxDeduction
func TestCalculateTaxDeduction_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecase{},
	)
	testCases := []struct {
		name                 string
		selfTaxDeduction     float64
		totalAllowances      float64
		expectedTaxDeduction float64
	}{
		{"Test case 1", 60000.0, 20.0, 60020.0},
		{"Test case 2", 70000.0, 30000.0, 100000.0}, // Expected tax is 5% of (200000 - 150000) Expected tax is 35% of (3000000 - 2000000) + 300000
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
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecase{},
	)
	testCases := []struct {
		name              string
		income            float64
		taxDeduction      float64
		expectedNetIncome float64
	}{
		{"Test case 1", 200000, 150000, 50000},
		{"Test case 2", 200000, 300000, 0},
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
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecase{},
	)
	testCases := []struct {
		name              string
		netIncome         float64
		wht               float64
		expectedTax       float64
		expectedTaxRefund float64
		expectedTaxLevel  []TaxLevelRes
	}{
		{
			"Test case 1",
			100000.0,
			0.0,
			0.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   0.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 2",
			200000.0,
			0.0,
			5000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   5000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		}, // Expected tax is 5% of (200000 - 150000)
		{"Test case 3",
			440000.0,
			0.0,
			29000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   29000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		}, // Expected tax is 15% of (440000 - 500000) + 35000
		{"Test case 4",
			440000.0,
			25000.0,
			4000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   29000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 5",
			440000.0,
			29000.0,
			0.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   29000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{
			"Test case 6",
			440000,
			39000,
			0,
			10000,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   29000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 7",
			600000.0,
			0,
			50000.0,
			0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   35000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   50000.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		}, // Expected tax is 15% of (600000 - 500000) + 35000
		{"Test case 8",
			750000.0,
			0.0,
			72500.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   35000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   72500.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		}, // Expected tax is 15% of (750000 - 500000) + 35000
		{"Test case 9",
			1500000.0,
			0.0,
			200000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   35000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   100000.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   200000.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		}, // Expected tax is 20% of (1500000 - 1000000) + 100000
		{"Test case 10",
			3000000.0,
			0.0,
			650000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   35000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   100000.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   300000.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   650000.0,
				},
			},
		}, // Expected tax is 35% of (3000000 - 2000000) + 300000
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tax, taxRefund, taxLevel := calculator.CalculateTax(tc.netIncome, tc.wht)

			// Assert
			assert.Equal(t, tc.expectedTax, tax)
			assert.Equal(t, tc.expectedTaxRefund, taxRefund)
			assert.Equal(t, tc.expectedTaxLevel, taxLevel)
		})
	}
}

// TestCalculate

type mockPersonalDeductionUsecaseGetDeductionNotFound struct {
}

func (p *mockPersonalDeductionUsecaseGetDeductionNotFound) GetDeduction() (float64, error) {
	return 0.0, errors.New("Not found")
}

func (p *mockPersonalDeductionUsecaseGetDeductionNotFound) UpdateDeduction(req personal.UpdatePersonalDeductionReq) (personal.UpdatePersonalDeductionRes, error) {
	return personal.UpdatePersonalDeductionRes{}, nil
}

func TestCalculate_ShouldReturnErr_WhenGetPersonalDeductionNotFound(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecaseGetDeductionNotFound{},
		&mockKReceiptDeductionUsecase{},
	)

	req := TaxCalculatorReq{
		TotalIncome: 500000.0,
		WHT:         0.0,
		Allowances: []AllowanceReq{
			{AllowanceType: allowanceType.Donation, Amount: 0.0},
		},
	}

	// Act
	result, err := calculator.Calculate(req)

	// Assert
	assert.Equal(t, 0.0, result.Tax)
	assert.Equal(t, 0.0, result.TaxRefund)
	assert.NotNil(t, err)
}

type mockKReceiptDeductionUsecaseGetDeductionNotFound struct {
}

func (p *mockKReceiptDeductionUsecaseGetDeductionNotFound) GetDeduction() (float64, error) {
	return 0.0, errors.New("Not found")
}

func (p *mockKReceiptDeductionUsecaseGetDeductionNotFound) UpdateDeduction(req kReceipt.UpdateKReceiptDeductionReq) (kReceipt.UpdateKReceiptDeductionRes, error) {
	return kReceipt.UpdateKReceiptDeductionRes{}, nil
}

func TestCalculate_ShouldReturnErr_WhenGetKReceiptDeductionNotFound(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecaseGetDeductionNotFound{},
	)

	req := TaxCalculatorReq{
		TotalIncome: 500000.0,
		WHT:         0.0,
		Allowances: []AllowanceReq{
			{AllowanceType: allowanceType.Donation, Amount: 0.0},
		},
	}

	// Act
	result, err := calculator.Calculate(req)

	// Assert
	assert.Equal(t, 0.0, result.Tax)
	assert.Equal(t, 0.0, result.TaxRefund)
	assert.NotNil(t, err)
}

func TestCalculate_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecase{},
	)

	testCases := []struct {
		name              string
		req               TaxCalculatorReq
		expectedTax       float64
		expectedTaxRefund float64
		expectedTaxLevel  []TaxLevelRes
	}{
		{"Test case 1", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         0.0,
			Allowances: []AllowanceReq{
				{AllowanceType: allowanceType.Donation, Amount: 0.0},
			},
		},
			29000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   29000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 2", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         25000.0,
			Allowances: []AllowanceReq{
				{AllowanceType: allowanceType.Donation, Amount: 0.0},
			},
		},
			4000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   29000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 3", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         0.0,
			Allowances: []AllowanceReq{
				{AllowanceType: allowanceType.Donation, Amount: 200000.0},
			},
		},
			19000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   19000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 4", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         29000.0,
			Allowances: []AllowanceReq{
				{AllowanceType: allowanceType.Donation, Amount: 200000.0},
			},
		},
			0.0,
			10000.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   19000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   100000.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   200000.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
		{"Test case 5", TaxCalculatorReq{
			TotalIncome: 500000.0,
			WHT:         0.0,
			Allowances: []AllowanceReq{
				{AllowanceType: allowanceType.KReceipt, Amount: 200000.0},
				{AllowanceType: allowanceType.Donation, Amount: 100000.0},
			},
		},
			14000.0,
			0.0,
			[]TaxLevelRes{
				{
					Level: "0-150,000",
					Tax:   0.0,
				},
				{
					Level: "150,001-500,000",
					Tax:   14000.0,
				},
				{
					Level: "500,001-1,000,000",
					Tax:   0.0,
				},
				{
					Level: "1,000,001-2,000,000",
					Tax:   0.0,
				},
				{
					Level: "2,000,001 ขึ้นไป",
					Tax:   0.0,
				},
			},
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := calculator.Calculate(tc.req)

			// Assert
			assert.Equal(t, tc.expectedTax, result.Tax)
			assert.Equal(t, tc.expectedTaxRefund, result.TaxRefund)
		})
	}
}

// CalculateMultiRequest

func TestCalculateTaxWithCSV_ShouldReturnErr_WhenGetDeductionNotFound(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecaseGetDeductionNotFound{},
		&mockKReceiptDeductionUsecase{},
	)

	reqs := []TaxCalculatorReq{
		{
			TotalIncome: 500000.0,
			WHT:         0.0,
			Allowances: []AllowanceReq{
				{AllowanceType: allowanceType.Donation, Amount: 0.0},
			},
		},
	}

	// Act
	_, err := calculator.CalculateMultiRequest(reqs)

	// Assert
	assert.NotNil(t, err)

}

func TestCalculateMultiRequest_ShouldCalculateCorrect_WhenCorrectInput(t *testing.T) {
	// Arrange
	calculator := NewTaxCalculatorUseCase(
		&mockPersonalDeductionUsecase{},
		&mockKReceiptDeductionUsecase{},
	)

	testCases := []struct {
		name        string
		reqs        []TaxCalculatorReq
		expectedRes TaxCalucalorMultipleRes
	}{
		{"Test case 1",
			[]TaxCalculatorReq{
				{
					TotalIncome: 500000.0,
					WHT:         0.0,
					Allowances: []AllowanceReq{
						{AllowanceType: allowanceType.Donation, Amount: 0.0},
					},
				},
				{
					TotalIncome: 600000.0,
					WHT:         40000.0,
					Allowances: []AllowanceReq{
						{AllowanceType: allowanceType.Donation, Amount: 20000.0},
					},
				},
				{
					TotalIncome: 750000.0,
					WHT:         50000.0,
					Allowances: []AllowanceReq{
						{AllowanceType: allowanceType.Donation, Amount: 15000.0},
					},
				},
			},
			TaxCalucalorMultipleRes{
				Taxes: []TaxCalucalorMultipleDetailRes{
					{
						TotalIncome: 500000.0,
						Tax:         29000.0,
						TaxRefund:   0.0,
					},
					{
						TotalIncome: 600000.0,
						Tax:         0.0,
						TaxRefund:   2000.0,
					},
					{
						TotalIncome: 750000.0,
						Tax:         11250.0,
						TaxRefund:   0.0,
					},
				},
			},
		},
	}

	// Act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := calculator.CalculateMultiRequest(tc.reqs)

			// Assert
			assert.Equal(t, tc.expectedRes, result)
		})
	}
}
