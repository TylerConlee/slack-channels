package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Args[1])
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Printf("Name: %s\nID: %s\n\n", channel.Name, channel.ID)
		// channel is of type conversation & groupConversation
		// see all available methods in `conversation.go`
	}
}
