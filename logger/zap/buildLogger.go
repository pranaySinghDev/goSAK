package zap

import (
	"time"

	cfg "github.com/pranaySinghDev/goSAK/logger/config"
	"github.com/pranaySinghDev/goSAK/logger/iface"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapFactory struct{}

// build zap logger
func (f *ZapFactory) Build(config *cfg.LogConfig) (iface.Logger, error) {
	//standard configuration
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:    cfg.FieldKeyMsg,
		LevelKey:      cfg.FieldKeyLevel,
		TimeKey:       cfg.FieldKeyTime,
		NameKey:       "",
		CallerKey:     cfg.FieldKeyFile,
		FunctionKey:   cfg.FieldKeyFunc,
		StacktraceKey: "",
		LineEnding:    "",
		EncodeLevel: func(zapcore.Level, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeTime: SyslogTimeEncoder,
		EncodeDuration: func(time.Duration, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeCaller: func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeName: func(string, zapcore.PrimitiveArrayEncoder) {
		},
		ConsoleSeparator: "",
	}
	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	// flushes buffer, if any
	defer zapLogger.Sync()
	return zapLogger.Sugar(), nil
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339Nano))
}
