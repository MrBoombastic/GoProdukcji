package commands

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"log"
)

func HelpHandler(ctx Context) {
	me, _ := ctx.Client.CurrentUser()
	myAvatar := me.GetAvatar(&discord.ImageOptions{})

	_, err := ctx.Message.Reply(&discord.MessageCreateOptions{
		Embed: &discord.MessageEmbed{
			Title: "GoProdukcji Help",
			Description: fmt.Sprintf(
				"`%vhelp` - wyświetla niniejszą pomoc\n"+
					"`%vstats` - wyświetla statystyki oraz ping bota\n"+
					"`%vsearch` - przeszukuje artykuły po tytułach z Na Produkcji\n"+
					"`%vlast` - wyświetla ostatni artykuł",
				ctx.Config.Prefix, ctx.Config.Prefix, ctx.Config.Prefix, ctx.Config.Prefix),
			Color:     colors.Orange,
			Thumbnail: &discord.EmbedMedia{Url: myAvatar},
			Footer:    &discord.EmbedFooter{Text: ctx.Message.Author.Username, IconURL: ctx.Message.Author.GetAvatar(&discord.ImageOptions{})},
		}})

	if err != nil {
		log.Fatal(err)
	}
}
