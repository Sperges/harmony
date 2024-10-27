package harmony

import (
	"context"
	"encoding/json"
	"fmt"
)

type HandlerFunc[T any] func(context.Context, T) error

type Handler interface {
	Type() string
	Handle(context.Context, []byte) error
}

type GenericHandler[T any] struct {
	Handler     HandlerFunc[T]
	handlerType string
}

func NewHandler[T any](handler HandlerFunc[T]) Handler {
	h := GenericHandler[T]{
		Handler: handler,
	}
	h.handlerType = structNameAsJsonString(*new(T))
	return h
}

func (h GenericHandler[T]) Type() string {
	return h.handlerType
}

func (h GenericHandler[T]) Handle(ctx context.Context, b []byte) error {
	var obj T
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("unable to unmarshal message: %w", err)
	}
	return h.Handler(ctx, obj)
}
