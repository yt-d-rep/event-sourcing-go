package main

import (
	usercommand "yt-d-rep/github.com/event-sourcing-go/command/user_command"
)

func main() {
	userCommand := usercommand.Wire()

	err := userCommand.Create()
	if err != nil {
		panic(err)
	}
}
