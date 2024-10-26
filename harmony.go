package harmony

import (
	"context"
	"encoding/json"
	"fmt"
)

type Harmony struct {
	handlers map[string]IHandler
}

func New() *Harmony {
	return &Harmony{
		handlers: map[string]IHandler{},
	}
}

func (h *Harmony) Register(handler IHandler) {
	c := handler.Type()
	if _, ok := h.handlers[c]; ok {
		// can this be turned into a compilation error?
		panic(fmt.Sprintf("attempt to register existing handler: %s", c))
	}
	h.handlers[c] = handler
}

func (h *Harmony) Handle(ctx context.Context, b []byte) error {
	var msg Message
	if err := json.Unmarshal(b, &msg); err != nil {
		return fmt.Errorf("unable to unmarshal harmony message: %w", err)
	}
	handler, ok := h.handlers[msg.Type]
	if !ok {
		return fmt.Errorf("no registered handler for message type: %s", msg.Type)
	}
	return handler.Handle(ctx, []byte(msg.Data))
}
