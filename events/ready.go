package events

import (
	"fmt"
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"goprodukcji/commands"
	"goprodukcji/config"
	"goprodukcji/utils"
	"log"
)

func HandleReady(c client.Client, config config.RunMode) {
	commands.GenerateHelpOutput()
	botUser, _ := c.CurrentUser()
	c.Presence().Set(discord.StatusOnline, discord.Activity{Name: "NaProdukcji.xyz  |  /help"})
	go func() {
		err := utils.RSS(c, config.DiscordChannel)
		if err != nil {
			log.Println(fmt.Sprintf("error: %v", err))
		}
	}()
	if config.DeployCommands {
		//commands.DeployCommandsGlobally(config.DiscordToken)
		commands.DeployCommandsLocally(config.DiscordToken, config.DiscordGuild)
	}
	c.Log().Info().Send(fmt.Sprintf("%v#%v is ready!", botUser.Username, botUser.Discriminator))
}
