package config

type QueueConfig struct {
	Type     QueueType
	URL      string
	Detailed bool
}

type QueueType int

const (
	RabbitMQ QueueType = iota
	CloudTask
)
