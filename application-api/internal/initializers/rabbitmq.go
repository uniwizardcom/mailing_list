package initializers

import (
	"application-api/rabbitmq"
	"os"
)

func RabbitMQ() *rabbitmq.RabbitMQ {
	rbmq := &rabbitmq.RabbitMQ{}

	err := rbmq.Init()
	if err != nil {
		os.Exit(1)
		return nil
	}

	err = rbmq.ConsumeMessage()
	if err != nil {
		os.Exit(1)
		return nil
	}

	return rbmq
}
