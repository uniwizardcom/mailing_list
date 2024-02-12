package rabbitmq

import (
	"fmt"
	"log"
)

func (r *RabbitMQ) ConsumeMessage() error {
	eventQueue, err := r.channel.Consume(r.conf.Queue.Name, "", false, false, false, false, nil)
	if err != nil {
		// TODO: re-send message to deadletter queue
		return fmt.Errorf("consume: %v", err)
	}

	go func() {
		for a := range eventQueue {
			err = r.channel.Ack(a.DeliveryTag, false)
			if err != nil {
				// TODO: re-send message to deadletter queue
				log.Fatalf("consume: %v", err)
				log.Fatalf("Re-send message to deadletter queue event %s\n", string(a.Body))
			} else {
				err = r.HandleMessage(a.Body)
				if err != nil {
					// TODO: re-send message to deadletter queue
					log.Fatalf("Re-send message to deadletter queue event %s\n", string(a.Body))
				}
			}
		}
	}()

	return nil
}
