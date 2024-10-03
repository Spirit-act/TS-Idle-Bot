package main

import (
	"doc/ts3bot/pkg/teamspeak3"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"time"
)

func main() {
	// create an logging for default logging -> stdout
	// var stdout = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// stdout log with log level Debug
	var stdout = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// create logging for errors -> stderr
	var stderr = slog.New(slog.NewJSONHandler(os.Stderr, nil))

	// api := teamspeak3.NewAPIClientFromEnv()

	api := teamspeak3.NewAPIClient(
		"http://127.0.0.1:10080",
		"1",
		// "BABhzVy1pSv_d_gBgr8xKCmkrmfGJp70uAJjjSC",
		"BAB6ZaEnX4IDmktN67TqvS6ug3hXaT03ORXJbbm",
	)

	// map to track the channels of the moved clients
	client_tracker := map[string]string{}

	// get AFK Chanell ID from Environment variable
	afk_channel_id, exists := os.LookupEnv("TSIDLE_AFK_CHANNEL_ID")

	if exists {
		stderr.Error("the AFK_CHANNEL_ID env variable needs to be set")
	}

	// get the channel from the ID
	afk_channel, _ := api.GetChannelById(afk_channel_id)

	// List of excluded chennels by ID
	excluded_channels := []string{
		"4",
	}

	for {
		// get list of all clients
		clients, err := api.ListClients()

		if err != nil {
			// if we are unable to get the client list, log an error and panic
			stderr.Error(err.Error(), "details", "unable to retrieve client list")
			// if can not get any clients, what is the point. PANIC
			panic(err.Error())
		}

		for _, client := range clients {
			if client.Type == teamspeak3.CLIENT_QUERY {
				// if the client is a query client
				// leave it, most do not see it
				continue
			}

			// check if the client id is in the tracker, also
			// if the id is in the tracker get the channel id
			cid, exists := client_tracker[client.Id]

			if exists {
				// user was moved by so we also need to move him back
				if client.RecentActive(60000) {
					// if client was in the last 1min active
					// delete it from the tracker
					delete(client_tracker, client.Id)

					if client.Channel_id != afk_channel.Id {
						// user left the afk channel himself  do NOT move him

						stdout.Debug(fmt.Sprintf(
							"Client %v has already left the configured Channel",
							client.Nickname,
						))

						continue
					}

					if _, err := api.MoveClient(client, cid); err != nil {
						// try to move the client
						// if it does not work, log an error
						stderr.Error(fmt.Sprintf("Unable to move %v to %v", client.Nickname, cid))
						continue
					}

					stdout.Info(fmt.Sprintf("%v moved back to channel %v", client.Nickname, cid))
				}

				continue
			}

			if client.ToIdle(120000) {
				if slices.Contains(excluded_channels, client.Channel_id) {
					// if a client is in an excluded channel, leave it allown
					continue
				}

				// store the current channel with in the tracker
				// bevor moving it
				client_tracker[client.Id] = client.Channel_id

				// MOVE IT
				if _, err := api.MoveClient(client, afk_channel.Id); err != nil {
					// try to move the client
					// if it does not work, log an error
					stderr.Error(fmt.Sprintf(
						"Unable to move %v to %v",
						client.Nickname,
						afk_channel.Id,
					))

					continue
				}
				stdout.Info(fmt.Sprintf("%v moved to channel %v", client.Nickname, afk_channel.Name))
			}
		}

		// wait some time bevor rechecking
		time.Sleep(10 * time.Second)
	}
}
