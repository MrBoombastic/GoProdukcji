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
	message := ctx.Interaction.SendMessageReply()
	message.Embed(discord.MessageEmbed{
		Title:       "GoProdukcji Help",
		Description: helpOutput,
		Color:       colors.Orange,
		Thumbnail:   discord.EmbedMedia{Url: me.AvatarURL(cdn.ImageSize512, cdn.ImageFormatPNG, true)},
		Footer:      discord.EmbedFooter{Text: ctx.Interaction.User.Username, IconUrl: ctx.Interaction.User.AvatarURL(cdn.ImageSize512, cdn.ImageFormatPNG, true)},
	})
	err = message.Execute()
	if err != nil {
		panic(err)
	}
}
