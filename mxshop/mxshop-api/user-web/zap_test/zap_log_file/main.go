package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stdout",
		"stderr",
		"./myproject.log",
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		return
	}
	sugar := logger.Sugar()
	defer sugar.Sync()
	url := "http://www.baidu.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
