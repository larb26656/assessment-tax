package personal

type PersonalDeductionsUsecase interface {
	GetDeductions() float64
	UpdateDeductions(req UpdatePersonalDeductionsReq) (UpdatePersonalDeductionsRes, error)
}

type personalDeductionsUsecase struct {
	personalDeductionsRepository PersonalDeductionsRepository
}

func NewPersonalDeductionsUsecase(personalDeductionsRepository PersonalDeductionsRepository) PersonalDeductionsUsecase {
	return &personalDeductionsUsecase{
		personalDeductionsRepository: personalDeductionsRepository,
	}
}

func (p *personalDeductionsUsecase) GetDeductions() float64 {
	return 0.0
}

func (p *personalDeductionsUsecase) UpdateDeductions(req UpdatePersonalDeductionsReq) (UpdatePersonalDeductionsRes, error) {
	return UpdatePersonalDeductionsRes{}, nil
}
