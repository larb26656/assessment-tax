package personal

type UpdatePersonalDeductionReq struct {
	Amount float64 `json:"amount" validate:"gte=10000,lte=100000"`
}

type UpdatePersonalDeductionRes struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}
