package commands

import (
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"goprodukcji/config"
)

type CommandData struct {
	Command     func(ctx Context)
	Usage       string
	Description string
	Aliases     []string
}

type Context struct {
	Client  state.IState
	Message *discord.Message
	Args    []string
	Config  config.RunMode
}

func NewContext(c state.IState, m *discord.Message, a []string, cfg config.RunMode) Context {
	return Context{
		Client:  c,
		Message: m,
		Args:    a,
		Config:  cfg,
	}
}
