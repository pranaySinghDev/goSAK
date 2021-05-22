package main

import (
	"fmt"
	"log"

	"github.com/pranaySinghDev/goSAK/queue"
	"github.com/pranaySinghDev/goSAK/queue/config"
)

func main() {
	queue, err := queue.Build(&config.QueueConfig{
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
