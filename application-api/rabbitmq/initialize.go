package rabbitmq

import (
	"application-api/cfg"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/url"
)

func (r *RabbitMQ) Init() error {
	r.conf = Config{}

	err := cfg.Read(&r.conf, "./configs/rabbitmq.yml")
	if err != nil {
		return fmt.Errorf("RabbitMQ Config error : %s\n", err)
	}

	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%d", r.conf.User, url.QueryEscape(r.conf.Password), r.conf.Host, r.conf.Port)
	r.conn, err = amqp.Dial(amqpUrl)
	if err != nil {
		return fmt.Errorf("RabbitMQ Dial [%s]\nerror : %s\n", amqpUrl, err)
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		return fmt.Errorf("channel.open: %s", err)
	}

	err = r.channel.ExchangeDeclare(r.conf.Exchange.Name, "direct", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("exchange.declare: %v", err)
	}

	err = r.channel.ExchangeBind(r.conf.Exchange.Name, "", r.conf.Exchange.Name, false, nil)
	if err != nil {
		return fmt.Errorf("exchange.bind: %v", err)
	}

	_, err = r.channel.QueueDeclare(r.conf.Queue.Name, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("queue.declare: %v", err)
	}

	err = r.channel.QueueBind(r.conf.Queue.Name, "", r.conf.Exchange.Name, false, nil)
	if err != nil {
		return fmt.Errorf("queue.bind: %v", err)
	}

	err = r.channel.Confirm(false)
	if err != nil {
		return fmt.Errorf("channel.Confirm: %v", err)
	}

	return nil
}
