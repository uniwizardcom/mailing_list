package mailbox

import (
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

func (p *MailBox) Send(destinyAccounts []string, subject string, body string) error {
	// fmt.Printf("MailBox conf: %v\n", p.conf)

	msg := "From: " + p.conf.Name + "\n" +
		"To: " + strings.Join(destinyAccounts, ";") + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(p.conf.Smtp.Host+":"+strconv.Itoa(p.conf.Smtp.Port),
		smtp.PlainAuth(p.conf.Auth.Identity, p.conf.Auth.Account, p.conf.Auth.Pass, p.conf.Smtp.Host),
		p.conf.Name, destinyAccounts, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}
