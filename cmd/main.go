package main

import (
	"Payment-Gateway/pkg/logger"
	"os"

	"go.uber.org/zap"
)

func main() {
	if err := StartServer(); err != nil {
		logger.GetLogger().Fatal("Server exited with error", zap.Error(err))
	}
	logger.Sync()
	os.Exit(0)
}
