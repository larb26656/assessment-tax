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

type TaxCalculatorRes struct {
	Tax float64 `json:"tax"`
}
