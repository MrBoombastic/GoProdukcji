package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"goprodukcji/commands"
	"goprodukcji/config"
	"log"
	"strings"
)

var cmds = map[string]commands.Handler{
	"stats":  commands.StatsHandler,
	"ping":   commands.StatsHandler,
	"help":   commands.HelpHandler,
	"search": commands.SearchHandler,
}

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
				_, err := message.Reply(&discord.MessageCreateOptions{Embed: &discord.MessageEmbed{
					Title: "Witaj!",
					Color: 0xff8000,
					Description: fmt.Sprintf("Jestem botem zaprojektowanym specjalnie dla serwera Na Produkcji!\n"+
						"Mój prefix to `%v`. Pomoc znajdziesz w `%vhelp`.\n"+
						"Kod źródłowy znajdziesz [tutaj](https://github.com/MrBoombastic/GoProdukcji).", config.Prefix, config.Prefix),
					Thumbnail: &discord.EmbedMedia{
						Url: "https://cdn.discordapp.com/icons/" + guild.ID.String() + "/" + guild.Icon + ".png?size=256",
					},
					URL: "https://naprodukcji.xyz",
					Image: &discord.EmbedMedia{
						Url: "https://naprodukcji.xyz/content/images/2021/06/comment_1622802543Quw49Z60cINC7fttv0aBcp.jpg",
					},
				}})
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
		handler := cmds[command]
		if handler != nil {
			handler(commands.NewContext(c, message, args, config))
			return
		}
	}
}
