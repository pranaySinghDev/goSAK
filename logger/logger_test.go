package logger

import (
	"log"
	"testing"

	"github.com/pranaySinghDev/goSAK/logger/config"
)

func TestLogrusLoggerFactory(t *testing.T) {
	cfg := &config.LogConfig{
		Type:     config.Logrus,
		Level:    "DEBUG",
		Detailed: true,
	}
	myLogger, err := Build(cfg)
	if err != nil {
		log.Fatalf("Init Logrus failed %v", err)
	}
	myLogger.Infof("info logrus")
}

func TestZapLoggerFactory(t *testing.T) {
	cfg := &config.LogConfig{
		Type:     config.Zap,
		Level:    "DEBUG",
		Detailed: true,
	}
	myLogger, err := Build(cfg)
	if err != nil {
		log.Fatalf("Init zap failed %v", err)
	}
	myLogger.Infof("info zap")
}
