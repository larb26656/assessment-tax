package personal

type PersonalDeductionsRepository interface {
	GetDeductions() (float64, error)
	UpdateDeductions(deductions float64) error
}

type personalDeductionsRepository struct {
}

func NewPersonalDeductionsRepository() PersonalDeductionsRepository {
	return &personalDeductionsRepository{}
}

func (p *personalDeductionsRepository) GetDeductions() (float64, error) {
	return 60000.0, nil
}

func (p *personalDeductionsRepository) UpdateDeductions(deductions float64) error {
	return nil
}
