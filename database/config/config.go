package config

type DBType int
type DBConfig struct {
	URL  string
	Type DBType
}

const (
	Mongodb DBType = iota
	PostgresSQL
)
