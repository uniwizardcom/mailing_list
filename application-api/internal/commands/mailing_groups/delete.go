package mailing_groups

import (
	"fmt"
)

func (c *MailingGroup) Delete(body []byte) error {
	groupToken := string(body)
	var groupId uint64
	err := c.Postgresql.Read(fmt.Sprintf("SELECT groups_id FROM groups WHERE groups_token='%s'", groupToken), &groupId)
	if err != nil {
		return err
	}

	err = c.Postgresql.Delete(fmt.Sprintf("DELETE FROM messages WHERE messages_group_id=%d", groupId))
	if err != nil {
		return err
	}

	err = c.Postgresql.Delete(fmt.Sprintf("DELETE FROM groups WHERE groups_id=%d", groupId))
	if err != nil {
		return err
	}

	return nil
}
