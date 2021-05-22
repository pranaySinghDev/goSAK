package rabbitMQ

import (
	"github.com/pranaySinghDev/goSAK/queue/config"
	"github.com/pranaySinghDev/goSAK/queue/iface"
)

type RabbitQueueBuilder struct {
}

func (q *RabbitQueueBuilder) Build(config *config.QueueConfig) (iface.IQueue, error) {
	queue, err := SetConnection(config.URL)
	if err != nil {
		return nil, err
	}
	return queue, err
}
