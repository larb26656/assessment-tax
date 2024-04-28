package kReceipt

type UpdateKReceiptDeductionReq struct {
	Amount float64 `json:"amount" validate:"gte=0,lte=100000"`
}

type UpdateKReceiptDeductionRes struct {
	KReceipt float64 `json:"kReceipt"`
}
