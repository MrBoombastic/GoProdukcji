package main

import (
	"encoding/json"
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord"
	"github.com/BOOMfinity-Developers/bfcord/cache"
	"github.com/BOOMfinity-Developers/bfcord/client"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"github.com/BOOMfinity-Developers/bfcord/gateway/intents"
	"github.com/BOOMfinity-Developers/bfcord/other"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getArticles() Articles {
	url := "https://naprodukcji.xyz/ghost/api/v3/content/posts/?key=" + os.Args[2]

	spaceClient := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "GoProdukcji v1")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	artciles := Articles{}
	jsonErr := json.Unmarshal(body, &artciles)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return artciles
}

func getLatestArticle() Post {
	return getArticles().Posts[0]
}

var prefix = "np!"

func main() {
	fmt.Println(os.Args)
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
	c.Gateway().MessageCreate(func(message gateway.MessageCreateEvent) {
		channel, _ := message.Channel().Get()
		//Repost all announcments/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				return
			}
		}
		//Ping command
		if message.Content == prefix+"ping" {
			_, err := message.Channel().SendMessage(&discord.MessageCreateOptions{Content: fmt.Sprintf("Gateway ping: %v", c.Manager().AveragePing())})
			if err != nil {
				return
			}
		}
	})
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
