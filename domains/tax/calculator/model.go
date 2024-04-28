package calculator

type AllowanceReq struct {
	AllowanceType string  `json:"allowanceType" validate:"required"`
	Amount        float64 `json:"amount" validate:"gte=0"`
}

type TaxCalculatorReq struct {
	TotalIncome float64        `json:"totalIncome" validate:"gte=0"`
	WHT         float64        `json:"wht" validate:"gte=0"`
	Allowances  []AllowanceReq `json:"allowances" validate:"required,dive"`
}

type TaxLevelRes struct {
	Level string  `json:"leveL"`
	Tax   float64 `json:"tax"`
}
type TaxCalculatorRes struct {
	Tax       float64       `json:"tax"`
	TaxRefund float64       `json:"taxRefund"`
	TaxLevel  []TaxLevelRes `json:"taxLevel"`
}
