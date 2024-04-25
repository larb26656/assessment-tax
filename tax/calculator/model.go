package calculator

type AllowanceReq struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type TaxCalculatorReq struct {
	TotalIncome float64        `json:"totalIncome"`
	WHT         float64        `json:"wht"`
	Allowances  []AllowanceReq `json:"allowances"`
}

type TaxCalculatorRes struct {
	Tax float64 `json:"tax"`
}
