package dtos

type TransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type TransactionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
