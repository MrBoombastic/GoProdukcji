package events

import (
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
		if !strings.HasPrefix(message.Content, config.Prefix) {
			return
		}
		channel, _ := message.Channel().Get()
		args := strings.Fields(strings.TrimPrefix(message.Content, config.Prefix))
		command := args[0]
		args = args[1:]

		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		handler := cmds[command]
		if handler != nil {
			handler(commands.NewContext(c, message, args, config))
			return
		}
	}
}
