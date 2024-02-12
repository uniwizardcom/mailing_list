package listeners

import (
	"application-api/internal/data"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (a Listeners) MessagesSend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := a.messagesSendPost(w, r)
		if err != nil {
			log.Print(err)
		}

		return
	}

	log.Printf("This request doesn't support  with this method [%s]", r.Method)
	return
}

// path "/api/messages" for "mailing\groups\send" by POST method
// The sending action to subscribe to all (mailing_id = -1) or only from specific group (mailing_id > 0)
// e.q.: curl -X POST 172.16.10.10:8080/api/messages/send -d '{"mailing_id":1}'
func (a Listeners) messagesSendPost(w http.ResponseWriter, r *http.Request) error {
	if r.ContentLength > 0 {
		dataBody, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("Body.Read: [%s]\n", err)
		}

		ams := data.ApiMessagesSend{}
		err = json.Unmarshal(dataBody, &ams)
		if err != nil {
			return fmt.Errorf("Body data have non supported data structure [%s]\njson.Unmarshal error [%s]\n", dataBody, err)
		}

		err = a.Rabbitmq.PublishData("mailing\\groups\\send", dataBody)
		if err != nil {
			return fmt.Errorf("Rabbitmq.SendMessage error: %s\n", err)
		}
	} else {
		return fmt.Errorf("Body.Read: no data")
	}

	return nil
}
