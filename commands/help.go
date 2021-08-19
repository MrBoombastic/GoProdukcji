package commands

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"log"
)

func HelpHandler(ctx Context) {
	_, err := ctx.Message.Reply(&discord.MessageCreateOptions{
		Embed: &discord.MessageEmbed{
			Title: "GoProdukcji Help",
			Description: fmt.Sprintf(
				"`%vhelp` - wyświetla niniejszą pomoc\n"+
					"`%vstats` - wyświetla statystyki oraz ping bota\n"+
					"`%vsearch` - przeszukuje artykuły po tytułach z Na Produkcji",
				ctx.Config.Prefix, ctx.Config.Prefix, ctx.Config.Prefix),
			Color: colors.Orange,
			Footer: &discord.EmbedFooter{
				Text: "Bot w głównej mierze ma pomagać w automatyce niektórych rzeczy, stąd mało komend",
			},
		}})

	if err != nil {
		log.Fatal(err)
	}
}
