package harmony

import (
	"context"
	"encoding/json"
	"fmt"
)

type IHandler interface {
	Type() string
	Handle(context.Context, any) error
}

type Handler[T any] struct {
	Handler     func(context.Context, T) error
	handlerType string
}

func NewHandler[T any](handler func(context.Context, T) error) IHandler {
	h := Handler[T]{
		Handler: handler,
	}
	h.handlerType = structNameAsJsonString(*new(T))
	return h
}

func (h Handler[T]) Type() string {
	return h.handlerType
}

func (h Handler[T]) Handle(ctx context.Context, message any) error {
	var obj T
	// TODO: casting `any` here seems sussy
	if err := json.Unmarshal(message.([]byte), &obj); err != nil {
		return fmt.Errorf("unable to unmarshal message: %w", err)
	}
	return h.Handler(ctx, obj)
}
