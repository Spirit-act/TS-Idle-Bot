package main

import (
	"doc/ts3bot/pkg/teamspeak3"
	"fmt"
)

func main() {
	api := teamspeak3.NewAPIClient(
		"http://127.0.0.1:10080",
		"1",
		"BADwny9ZGk28fadawg3N_fX-TlQM_rx2299fvsY",
	)

	channels, err := api.ListChannels()

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Printf("%-20s%-12s\n", "Channel Name", "Channel ID")
	fmt.Println("--------------------------------")
	for _, channel := range channels {
		fmt.Printf("%-20s%s\n", channel.Name, channel.Id)
	}
}
