package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"log"
	"os/exec"
)

var prefix = "np!"

func HandleMessageCreate(c state.IState) func(message gateway.MessageCreateEvent) {
	return func(message gateway.MessageCreateEvent) {
		channel, _ := message.Channel().Get()
		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		//Stats command
		if message.Content == prefix+"stats" {
			version, versionErr := exec.Command("git", "rev-parse", "HEAD").Output()
			if versionErr != nil {
				log.Fatal(versionErr)
			}
			_, sendErr := message.Channel().SendMessage(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title: "GoProdukcji Stats",
					Description: fmt.Sprintf("Gateway ping: %v\nVersion: [%v](https://github.com/MrBoombastic/GoProdukcji/commit/%v)",
						c.Manager().AveragePing(), string(version), string(version)),
					Color: colors.Orange,
				}})

			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}
	}
}
