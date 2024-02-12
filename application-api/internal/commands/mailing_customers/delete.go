package mailing_customers

import (
	"fmt"
)

func (c *MailingCustomer) Delete(body []byte) error {
	fmt.Printf("mailing_customers Delete: %s\n", body)

	return nil
}
