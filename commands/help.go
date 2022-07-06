package commands

import (
	"fmt"
	"github.com/BOOMfinity/bfcord/discord"
	"github.com/BOOMfinity/bfcord/discord/colors"
)

var HelpCommand = CommandData{
	Command:     runHelp,
	Description: "wyświetla niniejszą pomoc",
	Usage:       "",
	Aliases:     []string{"h"},
}

func runHelp(ctx Context) {
	me, err := ctx.Client.CurrentUser()
	if err != nil {
		panic(err)
	}
	message := ctx.Client.Channel(ctx.Message.ChannelID).SendMessage()
	message.Embed(discord.MessageEmbed{
		Title:       "GoProdukcji Help",
		Description: helpOutput,
		Color:       colors.Orange,
		Thumbnail:   discord.EmbedMedia{Url: fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.png", me.ID, me.Avatar)},
		Footer:      discord.EmbedFooter{Text: ctx.Message.Author.Username, IconUrl: fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%v.png", ctx.Message.Author.ID, ctx.Message.Author.Avatar)},
	})
	_, err = message.Execute(ctx.Client)
	if err != nil {
		panic(err)
	}
}
