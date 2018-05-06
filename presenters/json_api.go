package presenters

import "encoding/json"

type JSONPayload struct {
	Data interface{} `json:"data"`
}

func SerializeJSON(payload interface{}) ([]byte, error) {
	return json.Marshal(payload)
}
