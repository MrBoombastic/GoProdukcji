package commands

import (
	"github.com/BOOMfinity/bfcord/slash"
	"goprodukcji/utils"
	"strings"
)

var SearchCommand = CommandData{
	Command:     runSearch,
	Description: "przeszukuje artykuły po tytułach z Na Produkcji",
	Usage:       "search <fragment tytułu>",
	Aliases:     []string{"find"},
	Options:     []slash.Option{{Name: "fragment", Description: "fragment tytułu do wyszukania", Required: true, Type: 3}},
}

func runSearch(ctx Context) {
	err := ctx.Interaction.Defer(true, false)
	if err != nil {
		panic(err)
	}
	message := ctx.Interaction.EditOriginalReply()
	foundArticle, err := utils.SearchArticle(ctx.Interaction.Data.Options.Get("fragment").Value.(string))
	if err != nil {
		message.Content("Błąd: " + err.Error())
		err := message.Execute()
		if err != nil {
			panic(err)
		}
	}
	message.Embed(*embedArticle(foundArticle, strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)))
	err = message.Execute()
	if err != nil {
		panic(err)
	}
}
