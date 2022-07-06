package events

import (
	"fmt"
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"goprodukcji/commands"
	"goprodukcji/config"
	"goprodukcji/utils"
)

func HandleReady(c client.Client, config config.RunMode) {
	commands.GenerateHelpOutput(config.Prefix)
	botUser, _ := c.CurrentUser()
	c.Presence().Set(discord.StatusOnline, discord.Activity{Name: fmt.Sprintf("NaProdukcji.xyz  |  %vhelp", config.Prefix)})
	go func() {
		err := utils.RSS(c, config.DiscordNewsChannelGhost)
		if err != nil {
			panic(err)
		}
	}()
	c.Log().Info().Send(fmt.Sprintf("%v#%v is ready!", botUser.Username, botUser.Discriminator))
}
