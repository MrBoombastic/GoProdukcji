package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"goprodukcji/config"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var requestsCache = cache.New(5*time.Minute, 10*time.Minute)

func forceFetchArticles(options string) Articles {
	var articles Articles

	url := "https://naprodukcji.xyz/ghost/api/v3/content/posts/?key=" + config.GetConfig().GhostToken + options
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return Articles{}
	}
	err = json.NewDecoder(res.Body).Decode(&articles)
	if err != nil {
		log.Fatal(err)
		return Articles{}
	}
	requestsCache.Set("GetArticles", articles, cache.DefaultExpiration)
	return articles
}

func GetArticles(options string, caching bool) Articles {
	cacheRes, found := requestsCache.Get("GetArticles")

	if found && caching {
		cachedArticlees := cacheRes.(Articles)
		if len(cachedArticlees.Posts) > 0 {
			return cachedArticlees
		} else {
			return forceFetchArticles(options)
		}
	} else {
		return forceFetchArticles(options)
	}
}

func SearchArticle(query string) (Article, error) {
	articles := GetArticles("&limit=all&fields=id,title,url,primary_author,excerpt,published_at,feature_image&order=published_at%20desc&formats=plaintext&include=authors", true)
	for i := range articles.Posts {
		if strings.Contains(strings.ToLower(articles.Posts[i].Title), strings.ToLower(query)) {
			return articles.Posts[i], nil
		}
	}
	return Article{}, errors.New("artykuł nie został znaleziony")
}

func FormatBytes(b uint64) string {
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

func GetMemory() (x uint64) {
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
