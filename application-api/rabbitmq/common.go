package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Exchange struct {
	Name string `yaml:"name"`
}

type Queue struct {
	Name string `yaml:"name"`
}

type Config struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Exchange Exchange `yaml:"exchange"`
	Queue    Exchange `yaml:"queue"`
}

type RabbitMQ struct {
	conf    Config
	conn    *amqp.Connection
	channel *amqp.Channel

	HandleMessage func(body []byte) error
}
