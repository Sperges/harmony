package harmony

import "encoding/json"

type Message[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

func NewBytes[T any](data T) ([]byte, error) {
	b, err := json.Marshal(Message[T]{
		Type: structNameAsJsonString(data),
		Data: data,
	})
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
