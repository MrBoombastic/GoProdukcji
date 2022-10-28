package commands

import (
	"goprodukcji/utils"
	"log"
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
		message.Content("Error: " + err.Error())
		err := message.Execute()
		if err != nil {
			log.Println(err)
		}
		return
	}
	foundArticle := articles.Posts[0]
	message := ctx.Interaction.SendMessageReply()
	message.Embed(*embedArticle(foundArticle, strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)))
	err = message.Execute()
	if err != nil {
		panic(err)
	}
}
