package error

import "errors"

var (
	ErrUnsupportedGateway = errors.New("unsupported gateway")
	ErrInvalidRequest     = errors.New("invalid request payload")
	ErrProcessingFailed   = errors.New("gateway processing failed")
)
