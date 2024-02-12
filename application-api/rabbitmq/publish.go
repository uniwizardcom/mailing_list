package rabbitmq

import (
	"application-api/internal/commands"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func (r *RabbitMQ) PublishData(messType string, content []byte) error {
	cs := commands.CommandStructure{
		Type: messType,
		Data: content,
	}
	data, _ := json.Marshal(cs)

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain", // "application/json"
		Body:         data,
	}

	err := r.channel.PublishWithContext(context.Background(), r.conf.Exchange.Name, "", false, false, msg)
	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		log.Fatalf("basic.publish: %v", err)
	}

	return nil
}
