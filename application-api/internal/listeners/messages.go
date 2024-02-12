package listeners

import (
	"application-api/internal/data"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func (a Listeners) Messages(w http.ResponseWriter, r *http.Request) {
	/*_, err := io.WriteString(w, "This is messages!\n")
	if err != nil {
		log.Fatalf("Messages error starting server: %s\n", err)
		return
	}*/

	if r.Method == "POST" {
		err := a.messagesPost(w, r)
		if err != nil {
			log.Print(err)
		}

		return
	}

	if r.Method == "DELETE" {
		err := a.messagesDelete(w, r)
		if err != nil {
			log.Print(err)
		}

		return
	}

	log.Printf("This request doesn't support  with this method [%s]", r.Method)
	return
}

// path "/api/messages" for "mailing\customers\create" by POST method
// Subscribe new messages for mailing
// e.q.: curl -X POST 172.16.10.10:8080/api/messages -d '{"email":"jan.kowalski@example.com","title":"Interview","content":"simple text","mailing_id":2, "insert_time": "2020-04-24T05:42:38.725412916Z"}'
func (a Listeners) messagesPost(w http.ResponseWriter, r *http.Request) error {
	if r.ContentLength > 0 {
		dataBody, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("Body.Read: [%s]\n", err)
		}

		amc := data.ApiMessagesCreate{}
		err = json.Unmarshal(dataBody, &amc)
		if err != nil {
			return fmt.Errorf("Body data have non supported data structure [%s]\njson.Unmarshal error [%s]\n", dataBody, err)
		}

		err = a.Rabbitmq.PublishData("mailing\\customers\\create", dataBody)
		if err != nil {
			return fmt.Errorf("Rabbitmq.SendMessage error: %s\n", err)
		}
	} else {
		return fmt.Errorf("Body.Read: no data")
	}

	return nil
}

// path "/api/messages" for "mailing\groups\delete" by DELETE method
// Deleting subcribe from mailing
// e.q.: curl -X POST 172.16.10.10:8080/api/messages/{id}
func (a Listeners) messagesDelete(w http.ResponseWriter, r *http.Request) error {
	reg := regexp.MustCompile(`/api/messages/(.*)`)
	matches := reg.FindStringSubmatch(r.RequestURI)

	val, err := strconv.Atoi(matches[1])
	if err != nil {
		return fmt.Errorf("Body.Delete error: [%s]", r.RequestURI)
	}

	err = a.Rabbitmq.PublishData("mailing\\groups\\delete", []byte(strconv.Itoa(val)))
	if err != nil {
		return fmt.Errorf("Rabbitmq.SendMessage error: %s\n", err)
	}

	return nil
}
