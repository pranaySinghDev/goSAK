package logger

import (
	"github.com/pranaySinghDev/goSAK/logger/config"
	"github.com/pranaySinghDev/goSAK/logger/iface"
	"github.com/pranaySinghDev/goSAK/logger/logrus"
	"github.com/pranaySinghDev/goSAK/logger/zap"
)

type loggerFactory interface {
	Build(*config.LogConfig) (iface.Logger, error)
}

var loggerFactoryMap = map[config.LoggerType]loggerFactory{
	config.Logrus: &logrus.LogrusFactory{},
	config.Zap:    &zap.ZapFactory{},
}

func Build(config *config.LogConfig) (iface.Logger, error) {
	return loggerFactoryMap[config.Type].Build(config)
}
