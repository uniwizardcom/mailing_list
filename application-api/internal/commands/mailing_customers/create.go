package mailing_customers

import (
	"application-api/internal/commands/mailing_groups"
	"application-api/internal/commands/mailing_messages"
	"application-api/internal/data"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Create
// curl -X POST 172.16.10.10:8080/api/messages -d '{"email":"jan.kowalski@example.com","title":"Interview","content":"simple text","mailing_id":2, "insert_time": "2020-04-24T05:42:38.725412916Z"}'
func (c *MailingCustomer) Create(body []byte) error {
	ams := data.ApiMessagesCreate{}
	err := json.Unmarshal(body, &ams)
	if err != nil {
		return fmt.Errorf("Body data have non supported data structure [%s]\njson.Unmarshal error [%s]\n", body, err)
	}
	ams.GroupToken = strconv.Itoa(ams.MailingId)

	err = c.Postgresql.Read(fmt.Sprintf("SELECT customers_id FROM customers WHERE customers_email='%s'", ams.Email), &c.customerId)
	if err != nil {
		return err
	}

	if c.customerId == 0 {
		err = c.Postgresql.Create(fmt.Sprintf("INSERT INTO customers(customers_email, customers_createat) VALUES('%s', now()) RETURNING (customers_id)", ams.Email), &c.customerId)
		if err != nil {
			return err
		}
	}

	groupId, err := c.CreateGroup(ams.GroupToken)
	if err != nil {
		fmt.Printf("createGroup err: %s\n", err)
		return err
	}

	_, err = c.CreateMessage(groupId, ams.Title, ams.Content, ams.InsertTime)
	if err != nil {
		fmt.Printf("createMessage err: %s\n", err)
		return err
	}

	return nil
}

func (c *MailingCustomer) CreateGroup(groupsToken string) (uint64, error) {
	mg := &mailing_groups.MailingGroup{
		Postgresql: c.Postgresql,
	}
	dbtblGroup := data.DBTableGroups{
		GroupsToken: groupsToken,
	}
	dbtblGroupBytes, err := json.Marshal(dbtblGroup)
	if err != nil {
		return 0, err
	}

	return mg.Create(dbtblGroupBytes)
}

func (c *MailingCustomer) CreateMessage(groupId uint64, title string, content string, insertTime time.Time) (uint64, error) {
	mm := &mailing_messages.MailingMessages{
		Postgresql: c.Postgresql,
	}
	dbtblMessage := data.DBTableMessages{
		MessagesCustomerId: c.customerId,
		MessagesGroupId:    groupId,
		MessagesTitle:      title,
		MessagesContent:    content,
		MessagesCreatedAt:  insertTime,
	}
	dbtblMessageBytes, err := json.Marshal(dbtblMessage)
	if err != nil {
		return 0, err
	}

	return mm.Create(dbtblMessageBytes)
}
