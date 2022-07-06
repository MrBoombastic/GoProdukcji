package commands

import (
	"goprodukcji/utils"
	"strings"
)

var SearchCommand = CommandData{
	Command:     runSearch,
	Description: "przeszukuje artykuły po tytułach z Na Produkcji",
	Usage:       "search <fragment tytułu>",
	Aliases:     []string{"find"},
}

func runSearch(ctx Context) {
	message := ctx.Client.Channel(ctx.Message.ChannelID).SendMessage()
	if len(ctx.Args) == 0 {
		message.Content("Musisz podać tytuł artykułu do wyszukania!")
		_, err := message.Execute(ctx.Client)
		if err != nil {
			panic(err)
		}
	}
	foundArticle, err := utils.SearchArticle(strings.Join(ctx.Args, " "))
	if err != nil {
		message.Content("Błąd: " + err.Error())
		_, err := message.Execute(ctx.Client)
		if err != nil {
			panic(err)
		}
	}
	message.Embed(*embedArticle(foundArticle, strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)))
	if err != nil {
		panic(err)
	}
}
