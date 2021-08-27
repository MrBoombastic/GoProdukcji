package commands

import (
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"goprodukcji/utils"
	"log"
	"strings"
)

var SearchCommand = CommandData{
	Command:     runSearch,
	Description: "przeszukuje artykuły po tytułach z Na Produkcji",
	Usage:       "search <fragment tytułu>",
	Aliases:     []string{"find"},
}

func runSearch(ctx Context) {
	if len(ctx.Args) == 0 {
		_, err := ctx.Message.Reply(&discord.MessageCreateOptions{Content: "Musisz podać tytuł artykułu do wyszukania!"})
		if err != nil {
			return
		}
	}

	foundArticle, err := utils.SearchArticle(strings.Join(ctx.Args, " "))
	if err != nil {
		_, _ = ctx.Message.Reply(&discord.MessageCreateOptions{Content: "Błąd: " + err.Error()})
	}
	authorPicture := strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)
	_, err = ctx.Message.Reply(&discord.MessageCreateOptions{Embed: embedArticle(foundArticle, authorPicture)})
	if err != nil {
		log.Fatal(err)
	}
}
