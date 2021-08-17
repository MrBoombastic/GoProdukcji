package events

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"goprodukcji/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetArticles(config config.RunMode, options string) Articles {
	url := "https://naprodukcji.xyz/ghost/api/v3/content/posts/?key=" + config.GhostToken + options

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

func SearchArticle(query string) (Post, error) {
	articles := GetArticles(config.GetConfig(), "&limit=all&fields=id,title,url,excerpt,published_at,feature_image&order=published_at%20desc&formats=plaintext")
	for i := range articles.Posts {
		if strings.Contains(strings.ToLower(articles.Posts[i].Title), strings.ToLower(query)) {
			return articles.Posts[i], nil
		}
	}
	return Post{}, errors.New("artykuł nie został znaleziony")
}

func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func getMemory() (x uint64) {
	data, err := os.ReadFile("/proc/self/statm")
	if err != nil {
		return 0
	}
	d := bytes.Split(data, []byte(" "))

	r, _ := strconv.Atoi(string(d[1]))
	x += uint64(r)
	r, _ = strconv.Atoi(string(d[2]))
	x += uint64(r)
	x = x * 1024
	return
}
