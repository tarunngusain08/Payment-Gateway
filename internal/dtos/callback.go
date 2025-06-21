package dtos

// GatewayACallbackRequest represents the request structure for Gateway A callback

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
