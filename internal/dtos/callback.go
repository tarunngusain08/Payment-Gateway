package dtos

import errors "Payment-Gateway/pkg/error"

type HandleCallbackRequest struct {
	TransactionID string                 `json:"transaction_id"`
	Status        string                 `json:"status"`
	Metadata      map[string]interface{} `json:"metadata"`
	GatewayRef    string                 `json:"gateway_ref"`
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	Timestamp     string                 `json:"timestamp"`
}

type HandleCallbackResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (r *HandleCallbackRequest) Validate() error {
	if r.TransactionID == "" || r.GatewayRef == "" {
		return errors.ErrMissingRequiredFields
	}
	if r.Amount <= 0 {
		return errors.ErrMissingAmount
	}
	if r.Currency == "" {
		return errors.ErrMissingCurrency
	}
	return nil
}
