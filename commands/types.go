package commands

import (
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"goprodukcji/config"
)

type CommandData struct {
	Command     func(ctx Context)
	Usage       string
	Description string
	Aliases     []string
}

type Context struct {
	Client  client.Client
	Message *discord.Message
	Args    []string
	Config  config.RunMode
}

func NewContext(c client.Client, m *discord.Message, a []string, cfg config.RunMode) Context {
	return Context{
		Client:  c,
		Message: m,
		Args:    a,
		Config:  cfg,
	}
}
