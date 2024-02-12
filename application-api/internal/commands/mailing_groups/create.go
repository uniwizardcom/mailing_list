package mailing_groups

import (
	"application-api/internal/data"
	"encoding/json"
	"fmt"
)

func (c *MailingGroup) Create(body []byte) (uint64, error) {
	dbtblGroup := data.DBTableGroups{}
	err := json.Unmarshal(body, &dbtblGroup)
	if err != nil {
		return 0, fmt.Errorf("Body data have non supported data structure [%s]\njson.Unmarshal error [%s]\n", body, err)
	}

	var groupId uint64
	err = c.Postgresql.Read(fmt.Sprintf("SELECT groups_id FROM groups WHERE groups_token='%s'", dbtblGroup.GroupsToken), &groupId)
	if err != nil {
		return 0, err
	}

	if groupId == 0 {
		err = c.Postgresql.Create(fmt.Sprintf("INSERT INTO groups(groups_token, groups_createat) VALUES('%s', now()) RETURNING (groups_id)", dbtblGroup.GroupsToken), &groupId)
		if err != nil {
			return 0, err
		}
	}

	return groupId, nil
}
