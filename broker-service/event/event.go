package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_declare", // name
		"topic",        // type
		true,           // durable yn
		false,          // auto delete yn
		false,          // internal yn
		false,          // no-wait yn
		nil,            // arguments yn

	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name yn
		false, // durable yn
		false, // auto delete yn
		true,  // exclusive yn
		false, // no-wait yn
		nil,   // arguments yn
	)
}
