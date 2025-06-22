package handler

type Handlers struct {
	TransactionHandler TransactionHandler
	GatewayACallback   GatewayACallbackHandler
	GatewayBCallback   GatewayBCallbackHandler
}
