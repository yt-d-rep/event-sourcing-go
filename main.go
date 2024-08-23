package main

import (
	usercommand "yt-d-rep/github.com/event-sourcing-go/command/user_command"
)

func main() {
	userCommand := usercommand.Wire()

	userCommand.Create()
}
