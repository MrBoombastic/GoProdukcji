package events

import (
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"goprodukcji/commands"
	"goprodukcji/config"
	"goprodukcji/utils"
	"strings"
)

func HandleMessageCreate(c client.Client, config config.RunMode, message discord.Message) {
	me, _ := c.CurrentUser()
	channel, _ := c.Channel(message.ChannelID).Get()
	// Todo: not implemented in bfcord yet
	// guild, _ := message.Guild().Get()

	//Repost all announcements/tweets
	if channel.Type == discord.ChannelTypeNews {
		err := c.API().Channel(channel.ID).Message(message.ID).CrossPost()
		if err != nil {
			panic(err)
		}
	}

	//React to pinging
	mentionedUsers := message.Mentions
	if len(mentionedUsers) > 0 {
		if mentionedUsers[0].ID == me.ID {
			embed := utils.MentionEmbed(config, "") //guild.IconURL(&discord.ImageOptions{})) //Todo: not implemented in bfcord yet
			message := c.Channel(message.ChannelID).SendMessage()
			message.Embed(embed)
			_, err := message.Execute(c)
			if err != nil {
				panic(err)
			}
		}
	}

	//Command handler section
	if !strings.HasPrefix(message.Content, config.Prefix) {
		return
	}

	args := strings.Fields(strings.TrimPrefix(message.Content, config.Prefix))
	command := args[0]
	args = args[1:]

	foundCommand, err := commands.FindCommand(command)
	if err != nil {
		return
	}
	handler := foundCommand.Command
	if handler != nil {
		handler(commands.NewContext(c, &message, args, config))
		return
	}
}
