package personal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDeductionRepositoryCaseDeductionNotFound struct {
}

func (p *mockDeductionRepositoryCaseDeductionNotFound) GetDeduction(key string) (float64, error) {
	return 0.0, errors.New("deduction not found")
}

func (p *mockDeductionRepositoryCaseDeductionNotFound) UpdateDeduction(key string, deductions float64) error {
	return nil
}

func TestGetDeduction_ShouldReturnErr_WhenDeductionNotFound(t *testing.T) {
	// Arrange
	usecase := NewPersonalDeductionUsecase(&mockDeductionRepositoryCaseDeductionNotFound{})

	// Act
	_, err := usecase.GetDeduction()

	// Assert
	assert.NotNil(t, err)
}

type mockDeductionRepositoryCaseDeductionFound struct {
}

func (p *mockDeductionRepositoryCaseDeductionFound) GetDeduction(key string) (float64, error) {
	return 60000.0, nil
}

func (p *mockDeductionRepositoryCaseDeductionFound) UpdateDeduction(key string, deductions float64) error {
	return nil
}

func TestGetDeduction_ShouldReturnDeduction_WhenDeductionFound(t *testing.T) {
	// Arrange
	usecase := NewPersonalDeductionUsecase(&mockDeductionRepositoryCaseDeductionFound{})

	// Act
	deduction, err := usecase.GetDeduction()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 60000.0, deduction)
}

type mockDeductionRepositoryCaseUpdateDeductionError struct {
	amount float64
}

func (p *mockDeductionRepositoryCaseUpdateDeductionError) GetDeduction(key string) (float64, error) {
	return 60000.0, nil
}

func (p *mockDeductionRepositoryCaseUpdateDeductionError) UpdateDeduction(key string, deduction float64) error {
	p.amount = deduction
	return errors.New("Update deduction error")
}

// UpdateDeduction
func TestUpdateDeduction_ShouldReturnError_WhenUpdateDeductionFail(t *testing.T) {
	// Arrange
	repo := &mockDeductionRepositoryCaseUpdateDeductionError{}
	usecase := NewPersonalDeductionUsecase(repo)
	req := UpdatePersonalDeductionReq{
		Amount: 70000.0,
	}

	// Act
	_, err := usecase.UpdateDeduction(req)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, req.Amount, repo.amount)
}

type mockDeductionRepositoryCaseUpdateSuccess struct {
	amount float64
}

func (p *mockDeductionRepositoryCaseUpdateSuccess) GetDeduction(key string) (float64, error) {
	return 60000.0, nil
}

func (p *mockDeductionRepositoryCaseUpdateSuccess) UpdateDeduction(key string, deduction float64) error {
	p.amount = deduction
	return nil
}

func TestUpdateDeduction_ShouldSuccess_WhenCorrectInput(t *testing.T) {
	// Arrange
	repo := &mockDeductionRepositoryCaseUpdateSuccess{}
	usecase := NewPersonalDeductionUsecase(repo)
	req := UpdatePersonalDeductionReq{
		Amount: 70000.0,
	}

	// Act
	result, err := usecase.UpdateDeduction(req)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, req.Amount, repo.amount)
	assert.Equal(t, req.Amount, result.PersonalDeduction)
}
