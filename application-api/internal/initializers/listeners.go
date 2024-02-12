package initializers

import (
	"application-api/internal/listeners"
	"application-api/rabbitmq"
)

func Listeners(rabbitmq *rabbitmq.RabbitMQ) (string, map[string]listeners.RouteListenerFunct) {
	apiListeners := &listeners.Listeners{
		Rabbitmq: rabbitmq,
	}
	address := ":8080"
	routes := map[string]listeners.RouteListenerFunct{
		"/":                  apiListeners.Root,
		"/api/messages":      apiListeners.Messages,
		"/api/messages/":     apiListeners.Messages,
		"/api/messages/send": apiListeners.MessagesSend,
	}

	return address, routes
}
