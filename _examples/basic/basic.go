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
	fmt.Printf("%s: %s\n", message.Sender, message.Message)
	return nil
}

type StatusUpdate struct {
	Status string
}

func HandleStatusUpdate(ctx context.Context, message StatusUpdate) error {
	fmt.Printf("The new status is: %s\n", message.Status)
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

	// which looks better?

	// h.Register(harmony.NewHandler(HandleChatMessage))
	// h.Register(harmony.NewHandler(HandleStatusUpdate))

	harmony.Register(h, HandleChatMessage)
	harmony.Register(h, HandleStatusUpdate)

	println(string(chatMsg))
	h.Handle(context.Background(), chatMsg)

	println(string(statusMsg))
	h.Handle(context.Background(), statusMsg)
}
