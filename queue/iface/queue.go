package iface

import "github.com/streadway/amqp"

// Queue represent common interface for queuing  function
type IQueue interface {
	Enqueue(name string, data []byte) error
	Dequeue(name string) (<-chan amqp.Delivery, error)
	Cleanup()
}
