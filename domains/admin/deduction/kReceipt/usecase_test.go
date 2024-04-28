package kReceipt

import (
	"errors"
	"testing"

	"github.com/larb26656/assessment-tax/constant/allowanceType"
	"github.com/stretchr/testify/assert"
)

type mockDeductionRepositoryCaseDeductionNotFound struct {
	key string
}

func (p *mockDeductionRepositoryCaseDeductionNotFound) GetDeduction(key string) (float64, error) {
	p.key = key
	return 0.0, errors.New("deduction not found")
}

func (p *mockDeductionRepositoryCaseDeductionNotFound) UpdateDeduction(key string, deductions float64) error {
	return nil
}

func TestGetDeduction_ShouldReturnErr_WhenDeductionNotFound(t *testing.T) {
	// Arrange
	repo := &mockDeductionRepositoryCaseDeductionNotFound{}
	usecase := NewKReceiptDeductionUsecase(repo)

	// Act
	_, err := usecase.GetDeduction()

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, allowanceType.KReceipt, repo.key)
}

type mockDeductionRepositoryCaseDeductionFound struct {
	key string
}

func (p *mockDeductionRepositoryCaseDeductionFound) GetDeduction(key string) (float64, error) {
	p.key = key
	return 50000.0, nil
}

func (p *mockDeductionRepositoryCaseDeductionFound) UpdateDeduction(key string, deductions float64) error {
	return nil
}

func TestGetDeduction_ShouldReturnDeduction_WhenDeductionFound(t *testing.T) {
	// Arrange
	repo := &mockDeductionRepositoryCaseDeductionFound{}
	usecase := NewKReceiptDeductionUsecase(repo)

	// Act
	deduction, err := usecase.GetDeduction()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, allowanceType.KReceipt, repo.key)
	assert.Equal(t, 50000.0, deduction)
}

type mockDeductionRepositoryCaseUpdateDeductionError struct {
	key    string
	amount float64
}

func (p *mockDeductionRepositoryCaseUpdateDeductionError) GetDeduction(key string) (float64, error) {
	return 50000.0, nil
}

func (p *mockDeductionRepositoryCaseUpdateDeductionError) UpdateDeduction(key string, deduction float64) error {
	p.key = key
	p.amount = deduction
	return errors.New("Update deduction error")
}

// UpdateDeduction
func TestUpdateDeduction_ShouldReturnError_WhenUpdateDeductionFail(t *testing.T) {
	// Arrange
	repo := &mockDeductionRepositoryCaseUpdateDeductionError{}
	usecase := NewKReceiptDeductionUsecase(repo)
	req := UpdateKReceiptDeductionReq{
		Amount: 50000.0,
	}

	// Act
	_, err := usecase.UpdateDeduction(req)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, allowanceType.KReceipt, repo.key)
	assert.Equal(t, req.Amount, repo.amount)
}

type mockDeductionRepositoryCaseUpdateSuccess struct {
	key    string
	amount float64
}

func (p *mockDeductionRepositoryCaseUpdateSuccess) GetDeduction(key string) (float64, error) {
	return 60000.0, nil
}

func (p *mockDeductionRepositoryCaseUpdateSuccess) UpdateDeduction(key string, deduction float64) error {
	p.key = key
	p.amount = deduction
	return nil
}

func TestUpdateDeduction_ShouldSuccess_WhenCorrectInput(t *testing.T) {
	// Arrange
	repo := &mockDeductionRepositoryCaseUpdateSuccess{}
	usecase := NewKReceiptDeductionUsecase(repo)
	req := UpdateKReceiptDeductionReq{
		Amount: 70000.0,
	}

	// Act
	result, err := usecase.UpdateDeduction(req)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, allowanceType.KReceipt, repo.key)
	assert.Equal(t, req.Amount, repo.amount)
	assert.Equal(t, req.Amount, result.KReceipt)
}
