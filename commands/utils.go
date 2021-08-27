package commands

import (
	"errors"
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"goprodukcji/utils"
	"sort"
	"strings"
	"time"
)

var list = map[string]CommandData{ //Map with all commands
	"stats":  StatsCommand,
	"help":   HelpCommand,
	"search": SearchCommand,
	"last":   LastCommand,
}

func getSortedCommands() []string { //Returns sorted keys from list
	keys := make([]string, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func FindCommand(name string) (CommandData, error) { //Finds command by name or alias
	if list[name].Command != nil {
		return list[name], nil
	} else {
		for com := range list { //Commands loop
			for _, alias := range list[com].Aliases { //Aliases of command loop
				if alias == name {
					return list[com], nil
				}
			}
		}
	}
	return CommandData{}, errors.New("not found")
}

func EmbedArticle(article utils.Article, authorPicture string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Title:       article.Title,
		URL:         article.URL,
		Thumbnail:   discord.NewEmbedMedia(article.FeatureImage),
		Author:      discord.NewEmbedAuthor(article.PrimaryAuthor.Name, authorPicture, article.PrimaryAuthor.URL),
		Description: strings.ReplaceAll(article.Excerpt, "\n", " ") + " (...)",
		Footer: &discord.EmbedFooter{
			Text: article.PublishedAt.Format(time.RFC822),
		},
	}
}

var GenerateHelpOutput = func(prefix string) { //One-time help generator (fired on Ready)
	output := ""
	for _, com := range getSortedCommands() {
		output += fmt.Sprintf("- `%v%v` - %v\n", prefix, com, list[com].Description)
		if len(list[com].Usage) > 0 {
			output += fmt.Sprintf("Użycie: `%v%v`\n", prefix, list[com].Usage)
		}
	}
	helpOutput = output
}

var helpOutput string
