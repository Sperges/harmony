package harmony

import (
	"context"
	"encoding/json"
	"fmt"
)

type Message[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type Harmony struct {
	handlers map[string]Handler
}

func New() *Harmony {
	return &Harmony{
		handlers: map[string]Handler{},
	}
}

func Register[T any](harmony *Harmony, handlerFunc HandlerFunc[T]) {
	handler := NewHandler(handlerFunc)
	harmony.Register(handler)
}

func (h *Harmony) Register(handler Handler) {
	handlerType := handler.Type()
	if _, ok := h.handlers[handlerType]; ok {
		// can this be turned into a compilation error?
		panic(fmt.Sprintf("attempt to register existing handler: %s", handlerType))
	}
	h.handlers[handlerType] = handler
}

func (h *Harmony) Handle(ctx context.Context, b []byte) error {
	var msg struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(b, &msg); err != nil {
		return fmt.Errorf("unable to unmarshal harmony message: %w", err)
	}
	handler, ok := h.handlers[msg.Type]
	if !ok {
		return fmt.Errorf("no registered handler for message type: %s", msg.Type)
	}
	return handler.Handle(ctx, []byte(msg.Data))
}

func NewBytes[T any](data T) ([]byte, error) {
	return json.Marshal(Message[T]{
		Type: structNameAsJsonString(data),
		Data: data,
	})
}
