package main

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord"
	"github.com/BOOMfinity-Developers/bfcord/cache"
	"github.com/BOOMfinity-Developers/bfcord/client"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"github.com/BOOMfinity-Developers/bfcord/gateway/intents"
	"github.com/BOOMfinity-Developers/bfcord/other"
	"io/ioutil"
	"os"
)

func getLatestArticle() Post {
	return getArticles().Posts[0]
}

func main() {
	// create client
	c := state.New(client.Config{
		Logger:  other.NewDefaultLogger(false),
		Token:   os.Args[1],
		Intents: intents.GuildMessages | intents.Guilds,
		Cache: &cache.Config{
			Guilds:      bfcord.Bool(true),
			Members:     bfcord.Bool(true),
			Users:       bfcord.Bool(false),
			Channels:    bfcord.Bool(true),
			Presences:   bfcord.Bool(false),
			Messages:    bfcord.Bool(true),
			Roles:       bfcord.Bool(false),
			MaxMessages: bfcord.Int(10),
		},
	})

	// on message event
	c.Gateway().MessageCreate(handleMessageCreate)
	c.Gateway().Ready(func(_ gateway.ReadyEvent) {
		fmt.Println("GoProdukcji is ready!")

		var latestSavedArticleID = "0"
		fetchLatestSavedArticleID, err := ioutil.ReadFile("lastArticle")
		if err != nil {
			fmt.Println(err)
		} else {
			latestSavedArticleID = string(fetchLatestSavedArticleID)
		}
		var latestArticle = getLatestArticle()
		if latestSavedArticleID != latestArticle.ID {
			_, err := c.Channel(873917016585666560).SendMessage(&discord.MessageCreateOptions{Content: latestArticle.URL})
			if err != nil {
				return
			}
			savererr := ioutil.WriteFile("lastArticle", []byte(latestArticle.ID), 0777)
			if savererr != nil {
				fmt.Println(savererr)
			}
		}
	})

	c.Start()
	c.Wait()
}
