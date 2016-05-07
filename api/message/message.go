package message

import (
	"encoding/json"
	"net/http"
)

// Message Model
type Message struct {
	Message     string
	MessageType string
}

//PrintJSONMessage : print message in json format
func PrintJSONMessage(res http.ResponseWriter, msg string, msgtype string) {

	message := Message{
		Message:     msg,
		MessageType: msgtype,
	}
	json.NewEncoder(res).Encode(message)

}
