package service

import (
	"Payment-Gateway/internal/dtos"
)

type Callback interface {
	HandleCallback(req dtos.HandleCallbackRequest) (dtos.HandleCallbackResponse, error)
}
