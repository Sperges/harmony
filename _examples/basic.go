package main

import (
	"context"
	"fmt"
	"harmony"
)

type ChatMessage struct {
	Sender  string
	Message string
}

func HandleChatMessage(ctx context.Context, message ChatMessage) error {
	fmt.Printf("%s: %s", message.Sender, message.Message)
	return nil
}

type StatusUpdate struct {
	Status string
}

func HandleStatusUpdate(ctx context.Context, message StatusUpdate) error {
	fmt.Printf("The new status is: %s", message.Status)
	return nil
}

func main() {
	chatMsg, _ := harmony.NewBytes(ChatMessage{
		Sender:  "me",
		Message: "Hello, World!",
	})

	statusMsg, _ := harmony.NewBytes(StatusUpdate{
		Status: "Good!",
	})

	h := harmony.New()

	h.Register(harmony.NewHandler(HandleChatMessage))
	h.Register(harmony.NewHandler(HandleStatusUpdate))

	println(string(chatMsg))
	h.Handle(context.Background(), chatMsg)

	println(string(statusMsg))
	h.Handle(context.Background(), statusMsg)
}
