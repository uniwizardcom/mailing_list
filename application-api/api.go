package main

import (
	"application-api/internal/commands"
	"application-api/internal/initializers"
	"application-api/internal/listeners"
	"application-api/mailbox"
	"application-api/postgresql"
	"application-api/rabbitmq"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Api struct {
	address    string
	routes     map[string]listeners.RouteListenerFunct
	rabbitmq   *rabbitmq.RabbitMQ
	postgresql *postgresql.PostgreSQL
	mailbox    *mailbox.MailBox
	commands   *commands.CommandHandlers
}

func main() {
	api := &Api{}
	api.commands = initializers.CommandHandlers()
	api.rabbitmq = initializers.RabbitMQ()
	api.postgresql = initializers.PostgreSQL()
	api.mailbox = initializers.MailBox()
	api.address, api.routes = initializers.Listeners(api.rabbitmq)
	api.start()
}

func (a *Api) start() {
	a.commands.Postgresql = a.postgresql
	a.commands.Mailbox = a.mailbox

	for path, funcListener := range a.routes {
		http.HandleFunc(path, funcListener)
	}

	a.rabbitmq.HandleMessage = func(content []byte) error {
		cs := commands.CommandStructure{}
		err := json.Unmarshal(content, &cs)
		if err != nil {
			return err
		}

		return a.commands.Execute(cs.Type, cs.Data)
	}

	err := http.ListenAndServe(a.address, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
