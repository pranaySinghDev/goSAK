package rabbitMQ

import (
	"github.com/streadway/amqp"
)

type RabbitQueue struct {
	conn *amqp.Connection
}

func SetConnection(URL string) (*RabbitQueue, error) {
	conn, err := amqp.Dial(URL)
	if err != nil {
		return nil, err
	}
	rQ := &RabbitQueue{conn}
	return rQ, nil
}
func (q *RabbitQueue) Cleanup() {
	q.conn.Close()
}

func (q *RabbitQueue) Enqueue(name string, data []byte) error {
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(name, false, false, false, false, nil)
	err = ch.Publish("", name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	if err != nil {
		return err
	}
	defer ch.Close()
	return nil
}

func (q *RabbitQueue) Dequeue(name string) (<-chan amqp.Delivery, error) {
	ch, err := q.conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch.Consume(name, "", true, false, false, false, nil)
}
