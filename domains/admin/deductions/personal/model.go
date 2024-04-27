package personal

type UpdatePersonalDeductionsReq struct {
	Amount float64 `json:"amount" validate:"gte=0,lte=100000"`
}

type UpdatePersonalDeductionsRes struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}
