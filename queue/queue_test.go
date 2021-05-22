package queue

import (
	"fmt"
	"log"
	"testing"

	"github.com/pranaySinghDev/goSAK/queue/config"
)

func TestRabbitMQFactory(t *testing.T) {
	queue, err := Build(&config.QueueConfig{
		Type: config.RabbitMQ,
		URL:  "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		log.Fatalf("Couldn't build rabbitMQ Factory: %v", err)
	}
	err = queue.Enqueue("my-queue", []byte("Q Q Q Q Q"))
	if err != nil {
		log.Fatalf("Couldn't enqueue rabbitMQ Factory: %v", err)
	}
	queue.Cleanup()
}

func TestRabbitMQFactory_Dequeue(t *testing.T) {
	queue, err := Build(&config.QueueConfig{
		Type: config.RabbitMQ,
		URL:  "amqp://guest:guest@localhost:5672/",
	})
	if err != nil {
		log.Fatalf("Couldn't build rabbitMQ Factory: %v", err)
	}
	msgs, err := queue.Dequeue("my-queue")
	if err != nil {
		log.Fatalf("Couldn't dequeue rabbitMQ Factory: %v", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved message: %s\n", d.Body)
		}
	}()
	fmt.Println("Ready to recieve messages")
	<-forever
	queue.Cleanup()
}
