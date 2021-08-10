package events

import (
	"encoding/json"
	"goprodukcji/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetArticles() Articles {
	url := "https://naprodukcji.xyz/ghost/api/v3/content/posts/?key=" + config.GetConfig().GhostToken

	spaceClient := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 seconds
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
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
				log.Fatal(err)
			}
		}(res.Body)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	articles := Articles{}
	jsonErr := json.Unmarshal(body, &articles)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return articles
}
