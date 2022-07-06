package commands

import (
	"goprodukcji/utils"
	"strings"
)

var LastCommand = CommandData{
	Command:     runLast,
	Description: "wyświetla ostatni artykuł",
	Usage:       "",
}

func runLast(ctx Context) {
	articles, err := utils.GetArticles("all", true)
	if err != nil {
		message := ctx.Client.Channel(ctx.Message.ChannelID).SendMessage()
		message.Content("Błąd: " + err.Error())
		_, err := message.Execute(ctx.Client)
		if err != nil {
			panic(err)
		}
	}
	foundArticle := articles.Posts[0]
	message := ctx.Client.Channel(ctx.Message.ChannelID).SendMessage()
	message.Embed(*embedArticle(foundArticle, strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)))
	_, err = message.Execute(ctx.Client)
	if err != nil {
		panic(err)
	}
}
