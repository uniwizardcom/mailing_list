package data

import "time"

type ApiMessagesCreate struct {
	Email      string    `json:"email"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	MailingId  int       `json:"mailing_id"`
	InsertTime time.Time `json:"insert_time"`
	GroupToken string
}

type ApiMessagesSend struct {
	MailingId  int `json:"mailing_id"`
	GroupToken string
}
