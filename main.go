package main

import (
	"github.com/BOOMfinity-Developers/bfcord"
	"github.com/BOOMfinity-Developers/bfcord/cache"
	"github.com/BOOMfinity-Developers/bfcord/client"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/gateway/intents"
	"github.com/BOOMfinity-Developers/bfcord/other"
	"goprodukcji/config"
	"goprodukcji/events"
)

var cfg = config.GetConfig()

func main() {
	// create client
	discordClient := state.New(client.Config{
		Logger:  other.NewDefaultLogger(false),
		Token:   cfg.DiscordToken,
		Intents: intents.GuildMessages | intents.Guilds,
		Cache: &cache.Config{
			Presences:   bfcord.Bool(false),
			Roles:       bfcord.Bool(false),
			MaxMessages: bfcord.Int(2),
		},
	})

	// on message event
	discordClient.Gateway().MessageCreate(events.HandleMessageCreate(discordClient, cfg))
	discordClient.Gateway().Ready(events.HandleReady(discordClient, cfg))

	discordClient.Start()
	discordClient.Wait()
}
