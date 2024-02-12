package initializers

import (
	"application-api/mailbox"
	"os"
)

func MailBox() *mailbox.MailBox {
	mb := &mailbox.MailBox{}

	err := mb.Init()
	if err != nil {
		os.Exit(1)
		return nil
	}

	return mb
}
