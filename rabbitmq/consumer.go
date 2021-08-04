package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func GetTask() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err!=nil{
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err!=nil{
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"vkTasks", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err!=nil{
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err!=nil{
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
