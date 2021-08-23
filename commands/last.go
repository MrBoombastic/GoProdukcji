package commands

import (
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"goprodukcji/utils"
	"log"
	"strings"
	"time"
)

func LastHandler(ctx Context) {
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
	_, err = ctx.Message.Reply(&discord.MessageCreateOptions{
		Embed: &discord.MessageEmbed{
			Title:       foundArticle.Title,
			URL:         foundArticle.URL,
			Thumbnail:   discord.NewEmbedMedia(foundArticle.FeatureImage),
			Author:      discord.NewEmbedAuthor(foundArticle.PrimaryAuthor.Name, authorPicture, foundArticle.PrimaryAuthor.URL),
			Description: strings.ReplaceAll(foundArticle.Excerpt, "\n", " ") + " (...)",
			Footer: &discord.EmbedFooter{
				Text: foundArticle.PublishedAt.Format(time.RFC822),
			},
		}})
	if err != nil {
		log.Fatal(err)
	}
}
