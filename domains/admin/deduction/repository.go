package deduction

import (
	"database/sql"
)

type DeductionRepository interface {
	GetDeduction(key string) (float64, error)
	UpdateDeduction(key string, deduction float64) error
}

type deductionRepository struct {
	db *sql.DB
}

func NewDeductionsRepository(db *sql.DB) DeductionRepository {
	return &deductionRepository{
		db: db,
	}
}

func (p *deductionRepository) GetDeduction(key string) (float64, error) {
	stmt, err := p.db.Prepare(`SELECT value FROM tax_deduction_setting WHERE "key" = $1`)

	if err != nil {
		return 0.0, err
	}

	row := stmt.QueryRow(key)
	deductions := 0.0

	err = row.Scan(&deductions)

	if err != nil {
		return 0.0, err
	}

	return deductions, nil
}

func (p *deductionRepository) UpdateDeduction(key string, deduction float64) error {
	stmt, err := p.db.Prepare(`UPDATE tax_deduction_setting SET value=$2 WHERE "key" = $1`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(key, deduction)

	if err != nil {
		return err
	}

	return nil
}
