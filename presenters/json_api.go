package presenters

import (
	"banco_de_espana/entities"
	"encoding/json"
)

type JSONPayload struct {
	Data interface{} `json:"data"`
}
type ClientJSONPayload struct {
	Data entities.Client `json:"data"`
}

func SerializeJSON(payload interface{}) ([]byte, error) {
	return json.Marshal(payload)
}
