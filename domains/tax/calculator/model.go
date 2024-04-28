package calculator

type AllowanceReq struct {
	AllowanceType string  `json:"allowanceType" validate:"required,oneof='donation' 'k-receipt'"`
	Amount        float64 `json:"amount" validate:"gte=0"`
}

type TaxCalculatorReq struct {
	TotalIncome float64        `json:"totalIncome" validate:"gte=0"`
	WHT         float64        `json:"wht" validate:"gte=0"`
	Allowances  []AllowanceReq `json:"allowances" validate:"required,dive"`
}

type TaxLevelRes struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}
type TaxCalculatorRes struct {
	Tax       float64       `json:"tax"`
	TaxRefund float64       `json:"taxRefund"`
	TaxLevel  []TaxLevelRes `json:"taxLevel"`
}

type TaxCalucalorMultipleDetailRes struct {
	TotalIncome float64 `json:"totalIncome"`
	Tax         float64 `json:"tax"`
	TaxRefund   float64 `json:"taxRefund"`
}

type TaxCalucalorMultipleRes struct {
	Taxes []TaxCalucalorMultipleDetailRes `json:"taxes"`
}
