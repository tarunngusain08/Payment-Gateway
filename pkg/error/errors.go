package error

import "errors"

var (
	ErrUnsupportedGateway      = errors.New("unsupported gateway")
	ErrInvalidRequest          = errors.New("invalid request payload")
	ErrProcessingFailed        = errors.New("gateway processing failed")
	ErrGatewayNotAvailable     = errors.New("gateway service not available")
	ErrGatewayTimeout          = errors.New("gateway request timed out")
	ErrInvalidGatewayConfig    = errors.New("invalid gateway configuration")
	ErrTransactionNotFound     = errors.New("transaction not found")
	ErrTransactionExists       = errors.New("transaction already exists")
	ErrInvalidTransactionData  = errors.New("invalid transaction data")
	ErrTransactionUpdateFailed = errors.New("failed to update transaction")
	ErrInvalidAmount           = errors.New("invalid transaction amount")
	ErrInvalidAccount          = errors.New("invalid account details")
	ErrCallbackInvalid         = errors.New("invalid callback data")
	ErrCallbackProcessing      = errors.New("callback processing failed")

	// Common Callback Validation Errors
	ErrMissingTransactionID  = errors.New("invalid callback: missing transaction ID")
	ErrMissingGatewayRef     = errors.New("invalid callback: missing gateway reference")
	ErrMissingAmount         = errors.New("invalid transaction: missing amount")
	ErrMissingCurrency       = errors.New("invalid transaction: missing currency")
	ErrMissingRequiredFields = errors.New("invalid callback: missing required fields")
)
