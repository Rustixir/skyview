package main

import (
	"github.com/Rustixir/skyview/component"
	"github.com/Rustixir/skyview/example/chat"
)

func main() {
	hub := chat.NewHub()
	builder := component.NewBuilder("secret", "skyview")
	builder.AddComponent("channel", chat.NewChannel(hub), []string{"new_message"})
	builder.Start(":7070")
}
