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
			Guilds:      bfcord.Bool(true),
			Members:     bfcord.Bool(true),
			Users:       bfcord.Bool(false),
			Channels:    bfcord.Bool(true),
			Presences:   bfcord.Bool(false),
			Messages:    bfcord.Bool(true),
			Roles:       bfcord.Bool(false),
			MaxMessages: bfcord.Int(10),
		},
	})

	// on message event
	discordClient.Gateway().MessageCreate(events.HandleMessageCreate(discordClient, cfg))
	discordClient.Gateway().Ready(events.HandleReady(discordClient, cfg))

	discordClient.Start()
	discordClient.Wait()
}
