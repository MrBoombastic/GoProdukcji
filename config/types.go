package config

import "github.com/andersfylling/snowflake/v5"

type Config struct {
	Master RunMode `json:"master"`
	Slave  RunMode `json:"slave"`
}

type RunMode struct {
	DiscordToken   string              `json:"discord_token"`
	GhostToken     string              `json:"ghost_token"`
	DiscordChannel snowflake.Snowflake `json:"discord_channel"`
	DiscordGuild   snowflake.Snowflake `json:"discord_guild"`
	Shards         uint16              `json:"shards"`
	DeployCommands bool                `json:"deploy_commands"`
}
