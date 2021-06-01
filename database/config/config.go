package config

type DBType int
type DBConfig struct {
	URL     string
	Type    DBType
	MaxPool uint64
	MinPool uint64
}

const (
	Mongodb DBType = iota
	PostgresSQL
)
