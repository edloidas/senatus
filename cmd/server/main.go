package main

import (
	"log"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger := createLogger()
	defer logger.Sync()

	sugar := logger.Sugar()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ave!"))
	})

	serverPort := "8080"
	serverAddr := ":" + serverPort

	sugar.Infow("Starting server", "port", serverPort)
	sugar.Infof("Starting server on %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		sugar.Fatalf("Failed to start server: %v", err)
	}
}

func createLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.DebugLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.Encoding = "console"

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	return logger
}
