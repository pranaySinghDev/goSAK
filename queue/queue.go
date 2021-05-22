package queue

import (
	"github.com/pranaySinghDev/goSAK/queue/config"
	"github.com/pranaySinghDev/goSAK/queue/iface"
	"github.com/pranaySinghDev/goSAK/queue/rabbitMQ"
)

type queueFactory interface {
	Build(*config.QueueConfig) (iface.IQueue, error)
}

var loggerFactoryMap = map[config.QueueType]queueFactory{
	config.RabbitMQ:  &rabbitMQ.RabbitQueueBuilder{},
	config.CloudTask: nil,
}

func Build(config *config.QueueConfig) (iface.IQueue, error) {
	return loggerFactoryMap[config.Type].Build(config)
}
