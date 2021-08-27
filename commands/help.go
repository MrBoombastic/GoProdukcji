package commands

import (
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"log"
)

var HelpCommand = CommandData{
	Command:     runHelp,
	Description: "wyświetla niniejszą pomoc",
	Usage:       "",
	Aliases:     []string{"h"},
}

func runHelp(ctx Context) {
	me, _ := ctx.Client.CurrentUser()
	myAvatar := me.GetAvatar(&discord.ImageOptions{})
	_, err := ctx.Message.Reply(&discord.MessageCreateOptions{
		Embed: &discord.MessageEmbed{
			Title:       "GoProdukcji Help",
			Description: helpOutput,
			Color:       colors.Orange,
			Thumbnail:   &discord.EmbedMedia{Url: myAvatar},
			Footer:      &discord.EmbedFooter{Text: ctx.Message.Author.Username, IconURL: ctx.Message.Author.GetAvatar(&discord.ImageOptions{})},
		}})

	if err != nil {
		log.Fatal(err)
	}
}
