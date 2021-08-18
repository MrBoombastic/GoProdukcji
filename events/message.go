package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"goprodukcji/commands"
	"goprodukcji/config"
	"goprodukcji/utils"
	"log"
	"strings"
	"time"
)

func HandleMessageCreate(c state.IState, config config.RunMode) func(message gateway.MessageCreateEvent) {
	return func(message gateway.MessageCreateEvent) {
		if !strings.HasPrefix(message.Content, config.Prefix) {
			return
		}
		channel, _ := message.Channel().Get()
		args := strings.Fields(strings.TrimPrefix(message.Content, config.Prefix))
		command := args[0]
		args = args[1:]

		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		type CommandHandler func(c state.IState, message gateway.MessageCreateEvent)

		var commands = map[string]CommandHandler{
			"ping": CommandHandler(commands.StatsHandler),
		}
		handler := commands[command]
		if handler == nil {
			// Komenda nie istnieje
		}

		if command == "help" {
			_, err := message.Reply(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title: "GoProdukcji Help",
					Description: fmt.Sprintf("`%vhelp` - wyświetla niniejszą pomoc\n"+
						"`%vstats` - wyświetla statystyki oraz ping bota",
						config.Prefix, config.Prefix),
					Color: colors.Orange,
					Footer: &discord.EmbedFooter{
						Text: "Bot w głównej mierze ma pomagać w automatyce niektórych rzeczy, stąd mało komend",
					},
				}})

			if err != nil {
				log.Fatal(err)
			}
		}
		if command == "search" {
			if len(args) == 0 {
				_, err := message.Reply(&discord.MessageCreateOptions{Content: "Musisz podać tytuł artykułu do wyszukania!"})
				if err != nil {
					return
				}
				return
			}

			foundArticle, err := utils.SearchArticle(strings.Join(args, " "))
			if err != nil {
				_, err := message.Reply(&discord.MessageCreateOptions{Content: "Błąd: " + err.Error() + "!"})
				if err != nil {
					return
				}
				return
			}
			_, err = message.Reply(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title:       foundArticle.Title,
					URL:         foundArticle.URL,
					Thumbnail:   discord.NewEmbedMedia(foundArticle.FeatureImage),
					Author:      discord.NewEmbedAuthor(foundArticle.PrimaryAuthor.Name, &foundArticle.PrimaryAuthor.ProfileImage, &foundArticle.PrimaryAuthor.URL),
					Description: strings.ReplaceAll(foundArticle.Excerpt, "\n", " ") + " (...)",
					Footer: &discord.EmbedFooter{
						Text: foundArticle.PublishedAt.Format(time.RFC822),
					},
				}})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
