package kReceipt

import (
	"github.com/larb26656/assessment-tax/constant/allowanceType"
	"github.com/larb26656/assessment-tax/domains/admin/deduction"
)

type KReceiptDeductionUsecase interface {
	GetDeduction() (float64, error)
	UpdateDeduction(req UpdateKReceiptDeductionReq) (UpdateKReceiptDeductionRes, error)
}

type kReceiptDeductionUsecase struct {
	deductionRepository deduction.DeductionRepository
}

func NewKReceiptDeductionUsecase(deductionRepository deduction.DeductionRepository) KReceiptDeductionUsecase {
	return &kReceiptDeductionUsecase{
		deductionRepository: deductionRepository,
	}
}

func (p *kReceiptDeductionUsecase) GetDeduction() (float64, error) {
	deduction, err := p.deductionRepository.GetDeduction(allowanceType.KReceipt)

	if err != nil {
		return 0.0, err
	}

	return deduction, nil
}

func (p *kReceiptDeductionUsecase) UpdateDeduction(req UpdateKReceiptDeductionReq) (UpdateKReceiptDeductionRes, error) {
	err := p.deductionRepository.UpdateDeduction(allowanceType.KReceipt, req.Amount)

	if err != nil {
		return UpdateKReceiptDeductionRes{}, err
	}

	return UpdateKReceiptDeductionRes{
		KReceipt: req.Amount,
	}, nil
}
