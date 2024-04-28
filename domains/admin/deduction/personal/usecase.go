package personal

import (
	"github.com/larb26656/assessment-tax/constant/deductionType"
	"github.com/larb26656/assessment-tax/domains/admin/deduction"
)

type PersonalDeductionUsecase interface {
	GetDeduction() (float64, error)
	UpdateDeduction(req UpdatePersonalDeductionReq) (UpdatePersonalDeductionRes, error)
}

type personalDeductionUsecase struct {
	deductionRepository deduction.DeductionRepository
}

func NewPersonalDeductionUsecase(deductionRepository deduction.DeductionRepository) PersonalDeductionUsecase {
	return &personalDeductionUsecase{
		deductionRepository: deductionRepository,
	}
}

func (p *personalDeductionUsecase) GetDeduction() (float64, error) {
	deduction, err := p.deductionRepository.GetDeduction(deductionType.Personal)

	if err != nil {
		return 0.0, err
	}

	return deduction, nil
}

func (p *personalDeductionUsecase) UpdateDeduction(req UpdatePersonalDeductionReq) (UpdatePersonalDeductionRes, error) {
	err := p.deductionRepository.UpdateDeduction(deductionType.Personal, req.Amount)

	if err != nil {
		return UpdatePersonalDeductionRes{}, err
	}

	return UpdatePersonalDeductionRes{
		PersonalDeduction: req.Amount,
	}, nil
}
