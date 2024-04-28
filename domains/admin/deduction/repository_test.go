package deduction

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/larb26656/assessment-tax/constant/deductionType"
	"github.com/stretchr/testify/assert"
)

// GetDeduction

func TestGetDeduction_ShouldReturnError_WhenErrorOnPrepare(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewDeductionsRepository(db)
	key := deductionType.Personal
	mock.ExpectPrepare(`SELECT value FROM tax_deduction_setting WHERE "key" = \$1`).WillReturnError(errors.New("error on prepare"))

	// Act
	_, err = repo.GetDeduction(key)

	// Assert
	assert.Error(t, err)
}

func TestGetDeduction_ShouldReturnError_WhenErrorOnScan(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewDeductionsRepository(db)
	key := deductionType.Personal
	mock.ExpectPrepare(`SELECT value FROM tax_deduction_setting WHERE "key" = \$1`).ExpectQuery().
		WithArgs(key).WillReturnError(errors.New("error on scan"))

	// Act
	_, err = repo.GetDeduction(key)

	// Assert
	assert.Error(t, err)
}

func TestGetDeduction_ShouldReturnDeduction_WhenCorrectInput(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewDeductionsRepository(db)
	key := deductionType.Personal
	rows := sqlmock.NewRows([]string{"value"}).AddRow(70000.0)
	mock.ExpectPrepare(`SELECT value FROM tax_deduction_setting WHERE "key" = \$1`).ExpectQuery().
		WithArgs(key).WillReturnRows(rows)

	// Act
	deduction, err := repo.GetDeduction(key)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 70000.0, deduction)
}

// UpdateDeduction

func TestUpdateDeduction_ShouldReturnErrorOnPrepare_WhenCorrectInput(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewDeductionsRepository(db)
	key := deductionType.Personal
	deduction := 20000.0
	mock.ExpectPrepare(`UPDATE tax_deduction_setting SET value=\$2 WHERE "key" = \$1`).WillReturnError(errors.New("error on prepare")) // 1 row affected

	// Act
	err = repo.UpdateDeduction(key, deduction)

	// Assert
	assert.Error(t, err)
}

func TestUpdateDeduction_ShouldReturnErrorOnExec_WhenCorrectInput(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewDeductionsRepository(db)
	key := deductionType.Personal
	deduction := 20000.0
	mock.ExpectPrepare(`UPDATE tax_deduction_setting SET value=\$2 WHERE "key" = \$1`).ExpectExec().WillReturnError(errors.New("error on prepare")) // 1 row affected

	// Act
	err = repo.UpdateDeduction(key, deduction)

	// Assert
	assert.Error(t, err)
}

func TestUpdateDeduction_ShouldSuccess_WhenCorrectInput(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewDeductionsRepository(db)
	key := deductionType.Personal
	deduction := 20000.0
	mock.ExpectPrepare(`UPDATE tax_deduction_setting SET value=\$2 WHERE "key" = \$1`).ExpectExec().
		WithArgs(key, deduction).WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Act
	err = repo.UpdateDeduction(key, deduction)

	// Assert
	assert.NoError(t, err)
}
