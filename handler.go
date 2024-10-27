package harmony

import (
	"context"
	"encoding/json"
	"fmt"
)

type Handler interface {
	Type() string
	Handle(context.Context, any) error
}

type GenericHandler[T any] struct {
	Handler     func(context.Context, T) error
	handlerType string
}

func NewHandler[T any](handler func(context.Context, T) error) Handler {
	h := GenericHandler[T]{
		Handler: handler,
	}
	h.handlerType = structNameAsJsonString(*new(T))
	return h
}

func (h GenericHandler[T]) Type() string {
	return h.handlerType
}

func (h GenericHandler[T]) Handle(ctx context.Context, message any) error {
	var obj T
	// TODO: casting `any` here seems sussy
	if err := json.Unmarshal(message.([]byte), &obj); err != nil {
		return fmt.Errorf("unable to unmarshal message: %w", err)
	}
	return h.Handler(ctx, obj)
}
