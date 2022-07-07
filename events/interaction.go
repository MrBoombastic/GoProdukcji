package events

import (
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord/interactions"
	"goprodukcji/commands"
	"goprodukcji/config"
)

func HandleInteraction(c client.Client, config config.RunMode, interaction *interactions.Interaction) {
	if interaction.IsCommand() {
		foundCommand, err := commands.FindCommand(interaction.Data.Name)
		if err != nil {
			panic(err)
		}
		handler := foundCommand.Command
		if handler != nil {
			handler(commands.NewContext(c, interaction, config))
			return
		}
	}
}
