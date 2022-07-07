package main

import (
	"context"
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"github.com/BOOMfinity/bfcord/discord/interactions"
	"github.com/BOOMfinity/bfcord/gateway"
	"github.com/BOOMfinity/bfcord/gateway/intents"
	"goprodukcji/config"
	"goprodukcji/events"
)

var cfg = config.GetConfig()

func main() {
	// create client
	discordClient, _ := client.New(cfg.DiscordToken, client.WithIntents(intents.GuildMessages|intents.Guilds), client.WithShardCount(cfg.Shards))

	// on message event
	discordClient.Sub().Interaction(func(bot client.Client, shard *gateway.Shard, ev *interactions.Interaction) {
		defer discordClient.Log().Recover()
		events.HandleInteraction(bot, cfg, ev)
	})

	discordClient.Sub().MessageCreate(func(bot client.Client, shard *gateway.Shard, ev discord.Message) {
		defer discordClient.Log().Recover()
		events.HandleMessage(bot, ev)
	})

	discordClient.Sub().ShardReady(func(bot client.Client, shard *gateway.Shard, ev gateway.ReadyEvent) {
		defer discordClient.Log().Recover()
		events.HandleReady(discordClient, cfg)
	})
	err := discordClient.Start(context.Background())
	if err != nil {
		panic(err)
	}
	discordClient.Wait()
}
