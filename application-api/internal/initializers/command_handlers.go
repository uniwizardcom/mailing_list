package initializers

import (
	"application-api/internal/commands"
	"application-api/internal/commands/mailing_customers"
	"application-api/internal/commands/mailing_groups"
	"application-api/internal/commands/mailing_messages"
)

func CommandHandlers() *commands.CommandHandlers {
	ch := commands.CommandHandlers{}
	ch.Init()
	ch.Add(&mailing_customers.MailingCustomer{})
	ch.Add(&mailing_groups.MailingGroup{})
	ch.Add(&mailing_messages.MailingMessages{})

	return &ch
}
