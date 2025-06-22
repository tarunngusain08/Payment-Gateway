package dtos

import errors "Payment-Gateway/pkg/error"

type HandleCallbackRequest struct {
	TransactionID string                 `json:"transaction_id" xml:"TransactionID"`
	Status        string                 `json:"status" xml:"Status"`
	Metadata      map[string]interface{} `json:"metadata" xml:"-"`
	GatewayRef    string                 `json:"gateway_ref" xml:"GatewayRef"`
	Amount        float64                `json:"amount" xml:"Amount"`
	Currency      string                 `json:"currency" xml:"Currency"`
	Timestamp     string                 `json:"timestamp" xml:"Timestamp"`
}

type HandleCallbackResponse struct {
	Status  string `json:"status" xml:"Status"`
	Message string `json:"message" xml:"Message"`
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
