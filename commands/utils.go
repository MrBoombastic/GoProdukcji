package commands

import (
	"errors"
	"fmt"
	"github.com/BOOMfinity/bfcord/discord"
	"github.com/BOOMfinity/bfcord/slash"
	"github.com/andersfylling/snowflake/v5"
	"goprodukcji/utils"
	"sort"
	"strings"
	"time"
)

var helpOutput string

var commandsList = map[string]CommandData{ //Map with all commands
	"stats":  StatsCommand,
	"help":   HelpCommand,
	"search": SearchCommand,
	"last":   LastCommand,
}

func getSortedCommands() []string { //Returns sorted keys from list
	keys := make([]string, 0, len(commandsList))
	for k := range commandsList {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func FindCommand(name string) (CommandData, error) { //Finds command by name or alias
	if commandsList[name].Command != nil {
		return commandsList[name], nil
	}
	return CommandData{}, errors.New("not found")
}

func embedArticle(article utils.Article, authorPicture string) *discord.MessageEmbed {
	return &discord.MessageEmbed{
		Title:       article.Title,
		Url:         article.URL,
		Thumbnail:   discord.EmbedMedia{Url: article.FeatureImage},
		Author:      discord.EmbedAuthor{Name: article.PrimaryAuthor.Name, IconUrl: authorPicture, Url: article.PrimaryAuthor.URL},
		Description: strings.ReplaceAll(article.Excerpt, "\n", " ") + " (...)",
		Footer: discord.EmbedFooter{
			Text: article.PublishedAt.Format(time.RFC822),
		},
	}
}

var GenerateHelpOutput = func() { //One-time help generator (fired on Ready)
	output := ""
	for _, com := range getSortedCommands() {
		output += fmt.Sprintf("- `%v%v` - %v\n", "/", com, commandsList[com].Description)
		if len(commandsList[com].Usage) > 0 {
			output += fmt.Sprintf("Użycie: `%v%v`\n", "/", commandsList[com].Usage)
		}
	}
	helpOutput = output
}

var DeployCommandsGlobally = func(token string) { //Deploys commands to all guilds
	for _, com := range getSortedCommands() {
		api := slash.NewClient(token)
		//todo
		err := api.Global().Create(com, commandsList[com].Description) //commandsList[com].Options), Type: 1})
		if err != nil {
			panic(err)
		}
	}
}

var DeployCommandsLocally = func(token string, guildID snowflake.Snowflake) { //Deploys commands to only one guild
	for _, com := range getSortedCommands() {
		api := slash.NewClient(token)
		//todo
		err := api.Guild(guildID).Create(com, commandsList[com].Description) //.Option(commandsList[com].Option)
		if err != nil {
			panic(err)
		}
	}
}
