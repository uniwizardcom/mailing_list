package mailing_groups

import (
	"application-api/internal/data"
	"encoding/json"
	"fmt"
	"strconv"
)

func (c *MailingGroup) Send(body []byte) error {
	ams := data.ApiMessagesSend{}
	err := json.Unmarshal(body, &ams)
	if err != nil {
		return fmt.Errorf("Body data have non supported data structure [%s]\njson.Unmarshal error [%s]\n", body, err)
	}
	ams.GroupToken = strconv.Itoa(ams.MailingId)

	var groupId uint64
	err = c.Postgresql.Read(fmt.Sprintf("SELECT groups_id FROM groups WHERE groups_token='%s'", ams.GroupToken), &groupId)
	if err != nil {
		fmt.Printf("mailing_groups Send err: %s\n", err)
		return err
	}

	messages := make([]map[string]interface{}, 0)
	err = c.Postgresql.ReadAll(fmt.Sprintf("SELECT messages_customer_id, messages_title, messages_content FROM messages WHERE messages_group_id=%d", groupId), &messages)
	if err != nil {
		fmt.Printf("mailing_groups Send err: %s\n", err)
		return err
	}

	for _, message := range messages {
		customer, err := c.getCustomerData(message["messages_customer_id"].(int64))
		if err != nil {
			fmt.Printf("messages getCustomerData err: %s\n", err)
			return err
		}

		err = c.Mailbox.Send(
			[]string{customer["customers_email"].(string)},
			message["messages_title"].(string),
			message["messages_content"].(string),
		)
		if err != nil {
			fmt.Printf("Mailbox.Send error: %s\n", err)
			return err
		}
	}

	return nil
}

func (c *MailingGroup) getCustomerData(customerId int64) (map[string]interface{}, error) {
	customerData := make(map[string]interface{})
	err := c.Postgresql.ReadRecord(fmt.Sprintf("SELECT customers_email FROM customers WHERE customers_id=%d", customerId), &customerData)
	if err != nil {
		if err.Error() != "Record not exist" {
			fmt.Printf("getCustomerData err: %s\n", err)
			return nil, err
		}
	}

	return customerData, nil
}
