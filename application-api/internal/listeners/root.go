package listeners

import (
	"io"
	"log"
	"net/http"
)

func (a Listeners) Root(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "This is root!\n")
	if err != nil {
		log.Fatalf("getRoot error starting server: %s\n", err)
	}
}
