package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"goprodukcji/commands"
	"goprodukcji/config"
	"goprodukcji/utils"
	"log"
)

func HandleReady(c state.IState, config config.RunMode) func(message gateway.ReadyEvent) {
	return func(message gateway.ReadyEvent) {
		commands.GenerateHelpOutput(config.Prefix)
		botUser, _ := c.CurrentUser()
		fmt.Printf("%v#%v is ready!\n", botUser.Username, botUser.Discriminator)
		c.SetPresence(gateway.StatusUpdate{Activities: []discord.Activity{{Name: fmt.Sprintf("NaProdukcji.xyz  |  %vhelp", config.Prefix)}}})
		rss, err := utils.RSS()
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = c.Channel(config.DiscordNewsChannelGhost).SendMessage(&discord.MessageCreateOptions{Content: rss})
		if err != nil {
			return
		}
	}
}
