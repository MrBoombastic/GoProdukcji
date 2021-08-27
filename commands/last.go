package commands

import (
	"github.com/BOOMfinity-Developers/bfcord/discord"
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
		_, err := ctx.Message.Reply(&discord.MessageCreateOptions{Content: "Błąd: " + err.Error()})
		if err != nil {
			return
		}
		return
	}
	foundArticle := articles.Posts[0]
	authorPicture := strings.Replace(foundArticle.PrimaryAuthor.ProfileImage, "//www.gravatar.com", "https://www.gravatar.com", 1)
	_, err = ctx.Message.Reply(&discord.MessageCreateOptions{Embed: embedArticle(foundArticle, authorPicture)})
	if err != nil {
		log.Fatal(err)
	}
}
