package commands

import (
	"github.com/BOOMfinity/bfcord/api/cdn"
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
		Thumbnail:   discord.EmbedMedia{Url: me.AvatarURL(cdn.ImageSize512, cdn.ImageFormatPNG, true)},
		Footer:      discord.EmbedFooter{Text: ctx.Message.Author.Username, IconUrl: ctx.Message.Author.AvatarURL(cdn.ImageSize512, cdn.ImageFormatPNG, true)},
	})
	_, err = message.Execute(ctx.Client)
	if err != nil {
		panic(err)
	}
}
