package listeners

import (
	"application-api/rabbitmq"
	"net/http"
)

type RouteListenerFunct func(w http.ResponseWriter, r *http.Request)

type Listeners struct {
	Rabbitmq *rabbitmq.RabbitMQ
}
