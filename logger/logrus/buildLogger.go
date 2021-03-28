package logrus

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	cfg "github.com/pranaySinghDev/goSAK/logger/config"
	"github.com/pranaySinghDev/goSAK/logger/iface"
	"github.com/sirupsen/logrus"
)

type LogrusFactory struct{}

// build logrus logger
func (f *LogrusFactory) Build(config *cfg.LogConfig) (iface.Logger, error) {
	//standard configuration
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  time.RFC3339Nano,
		CallerPrettyfier: setFileFormatter,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  cfg.FieldKeyTime,
			logrus.FieldKeyLevel: cfg.FieldKeyLevel,
			logrus.FieldKeyMsg:   cfg.FieldKeyMsg,
			logrus.FieldKeyFunc:  cfg.FieldKeyFunc,
			logrus.FieldKeyFile:  cfg.FieldKeyFile,
		},
	})
	log.SetReportCaller(config.Detailed)
	log.SetOutput(os.Stdout)
	logLevel := &log.Level
	err := logLevel.UnmarshalText([]byte(config.Level))
	if err != nil {
		return nil, errors.New(fmt.Sprint(err))
	}
	log.SetLevel(*logLevel)
	return log, nil
}

func setFileFormatter(frame *runtime.Frame) (function, file string) {
	sp := strings.Split(frame.File, "/")
	file = sp[len(sp)-1] + "/" + sp[len(sp)-2] + ":" + strconv.Itoa(frame.Line)
	return frame.Function, file
}
