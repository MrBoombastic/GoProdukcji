package events

import (
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"goprodukcji/commands"
	"goprodukcji/config"
	"goprodukcji/utils"
	"log"
	"strings"
)

func HandleMessageCreate(c state.IState, config config.RunMode) func(message gateway.MessageCreateEvent) {
	return func(message gateway.MessageCreateEvent) {
		me, _ := c.CurrentUser()
		guild, _ := message.Guild().Get()
		channel, _ := message.Channel().Get()

		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		//React to pinging
		mentionedUsers := message.Mentions().Users
		if len(mentionedUsers) > 0 {
			if mentionedUsers[0].ID == me.ID {
				embed := utils.MentionEmbed(config, guild.IconURL(&discord.ImageOptions{}))
				_, err := message.Reply(&discord.MessageCreateOptions{Embed: &embed})
				if err != nil {
					log.Fatal(err)
					return
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
			handler(commands.NewContext(c, message, args, config))
			return
		}
	}
}
