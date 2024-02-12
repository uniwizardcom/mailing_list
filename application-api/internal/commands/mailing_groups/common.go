package mailing_groups

import (
	"application-api/mailbox"
	"application-api/postgresql"
)

type MailingGroup struct {
	Postgresql *postgresql.PostgreSQL
	Mailbox    *mailbox.MailBox
}
