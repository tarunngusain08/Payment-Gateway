package handler

import "net/http"

type Handlers struct {
	DepositHandler    http.HandlerFunc
	WithdrawalHandler http.HandlerFunc
	GatewayACallback  GatewayACallbackHandler
	GatewayBCallback  GatewayBCallbackHandler
}
