package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"io/ioutil"
	"log"
	"time"
)

func HandleReady(c state.IState) func(message gateway.ReadyEvent) {
	return func(message gateway.ReadyEvent) {
		botUser, _ := c.CurrentUser()
		fmt.Printf("%v#%v is ready!\n", botUser.Username, botUser.Discriminator)

		go func() {
			for {
				time.Sleep(2 * time.Minute)
				latestSavedArticleID := "0"
				fetchLatestSavedArticleID, err := ioutil.ReadFile("lastArticle")
				if err != nil {
					log.Fatal(err)
				} else {
					latestSavedArticleID = string(fetchLatestSavedArticleID)
				}
				latestArticle := GetArticles().Posts[0]
				if latestSavedArticleID != latestArticle.ID {
					_, err := c.Channel(873917016585666560).SendMessage(&discord.MessageCreateOptions{Content: latestArticle.URL})
					if err != nil {
						log.Fatal(err)
					}
					save := ioutil.WriteFile("lastArticle", []byte(latestArticle.ID), 0777)
					if save != nil {
						log.Fatal(save)
					}
				}
			}
		}()
	}
}
