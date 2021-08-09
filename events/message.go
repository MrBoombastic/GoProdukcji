package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
)

var prefix = "np!"

func handleMessageCreate(message gateway.MessageCreateEvent) discord.Message {
	channel, _ := message.Channel().Get()
	//Repost all announcments/tweets
	if channel.Type == discord.GuildNewsChannel {
		err := message.Crosspost()
		if err != nil {
		}
	}
	//Ping command
	if message.Content == prefix+"ping" {
		_, err := message.Channel().SendMessage(&discord.MessageCreateOptions{Content: fmt.Sprintf("Gateway ping: %v", c.Manager().AveragePing())})
		if err != nil {
		}
	}
}
