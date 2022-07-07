package commands

import (
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord/interactions"
	"github.com/BOOMfinity/bfcord/slash"
	"goprodukcji/config"
)

type CommandData struct {
	Command     func(ctx Context)
	Usage       string
	Description string
	Aliases     []string
	Options     []slash.Option
}

type Context struct {
	Client      client.Client
	Interaction *interactions.Interaction
	Config      config.RunMode
}

func NewContext(c client.Client, i *interactions.Interaction, cfg config.RunMode) Context {
	return Context{
		Client:      c,
		Interaction: i,
		Config:      cfg,
	}
}
