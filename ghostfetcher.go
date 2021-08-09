package main

import (
	"encoding/json"
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
