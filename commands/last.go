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
		message := ctx.Interaction.SendMessageReply()
		message.Content("Błąd: " + err.Error())
		err := message.Execute()
		if err != nil {
			panic(err)
		}
	}
	foundArticle := articles.Posts[0]
	message := ctx.Interaction.SendMessageReply()
	message.Embed(*embedArticle(foundArticle, strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)))
	err = message.Execute()
	if err != nil {
		panic(err)
	}
}
