package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// GetLogger returns a singleton zap.Logger instance.
func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			panic("failed to initialize zap logger: " + err.Error())
		}
	})
	return logger
}

// Sync flushes any buffered log entries.
func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}
