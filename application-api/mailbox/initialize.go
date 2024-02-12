package mailbox

import (
	"application-api/cfg"
	"fmt"
)

func (p *MailBox) Init() error {
	p.conf = Config{}

	err := cfg.Read(&p.conf, "./configs/mailbox.yml")
	if err != nil {
		return fmt.Errorf("MailBox Config error : %s\n", err)
	}

	return err
}
