package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"goprodukcji/config"
	"goprodukcji/utils"
	"io/ioutil"
	"log"
	"time"
)

func HandleReady(c state.IState, config config.RunMode) func(message gateway.ReadyEvent) {
	return func(message gateway.ReadyEvent) {
		botUser, _ := c.CurrentUser()
		fmt.Printf("%v#%v is ready!\n", botUser.Username, botUser.Discriminator)
		c.SetPresence(gateway.StatusUpdate{Activities: []discord.Activity{{Name: fmt.Sprintf("NaProdukcji.xyz  |  %vhelp", config.Prefix)}}})
		go func() {
			for {
				time.Sleep(2 * time.Minute)
				latestSavedArticleID := "0"
				fetchLatestSavedArticleID, err := ioutil.ReadFile("./lastArticle")
				if err != nil {
					err := ioutil.WriteFile("./lastArticle", []byte(latestSavedArticleID), 0777)
					if err != nil {
						return
					}
				} else {
					latestSavedArticleID = string(fetchLatestSavedArticleID)
				}
				articles, err := utils.GetArticles("id", false)
				if err != nil {
					fmt.Println(err)
					return
				}
				latestArticle := articles.Posts[0]
				if latestSavedArticleID != latestArticle.ID {
					_, err := c.Channel(config.DiscordNewsChannelGhost).SendMessage(&discord.MessageCreateOptions{Content: latestArticle.URL})
					if err != nil {
						log.Fatal(err)
					}
					save := ioutil.WriteFile("./lastArticle", []byte(latestArticle.ID), 0777)
					if save != nil {
						log.Fatal(save)
					}
				}
			}
		}()
	}
}
