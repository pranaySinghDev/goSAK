package config

type LogConfig struct {
	Type     LoggerType
	Level    string
	Detailed bool
}

type LoggerType int

const (
	Logrus LoggerType = iota
	Zap
)

var (
	FieldKeyTime  = "ts"
	FieldKeyLevel = "level"
	FieldKeyMsg   = "msg"
	FieldKeyFunc  = "caller"
	FieldKeyFile  = "line"
)
