package mailing_messages

import (
	"application-api/internal/data"
	"encoding/json"
	"fmt"
)

func (c *MailingMessages) Create(body []byte) (uint64, error) {
	dbtblMessage := data.DBTableMessages{}
	err := json.Unmarshal(body, &dbtblMessage)
	if err != nil {
		return 0, fmt.Errorf("Body data have non supported data structure [%s]\njson.Unmarshal error [%s]\n", body, err)
	}

	var messageId uint64
	err = c.Postgresql.Create(fmt.Sprintf("INSERT INTO messages(messages_customer_id, messages_group_id, messages_title, messages_content, messages_createat)"+
		"VALUES(%d, %d, '%s', '%s', '%s') RETURNING (messages_id)",
		dbtblMessage.MessagesCustomerId,
		dbtblMessage.MessagesGroupId,
		dbtblMessage.MessagesTitle,
		dbtblMessage.MessagesContent,
		dbtblMessage.MessagesCreatedAt.Format("2006-01-02 15:04:05.999999999 -0700"),
	), &messageId)
	if err != nil {
		return 0, err
	}

	return messageId, nil
}
