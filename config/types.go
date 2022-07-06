package config

import "github.com/andersfylling/snowflake/v5"

type Config struct {
	Master RunMode `json:"master"`
	Slave  RunMode `json:"slave"`
}

type RunMode struct {
	DiscordToken            string              `json:"discord_token"`
	GhostToken              string              `json:"ghost_token"`
	DiscordNewsChannelGhost snowflake.Snowflake `json:"discord_news_channel_ghost"`
	Prefix                  string              `json:"prefix"`
	Shards                  uint16              `json:"shards"`
}
