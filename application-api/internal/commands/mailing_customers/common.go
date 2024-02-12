package mailing_customers

import "application-api/postgresql"

type MailingCustomer struct {
	Postgresql *postgresql.PostgreSQL

	customerId uint64
}
