package harmony

import "encoding/json"

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func NewBytes[T any](message T) ([]byte, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return []byte{}, err
	}
	b, err := json.Marshal(Message{
		Type: structNameAsJsonString(message),
		Data: data,
	})
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
